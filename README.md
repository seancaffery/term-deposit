# Term deposit balance calculator

A simple term deposit calculator written in Go.

## Design overview / tradeoffs

- Uses [Cobra](https://github.com/spf13/cobra) CLI application to define commands and manage user interaction (flag parsing, help output, etc). Using Cobra rather than [flags](https://pkg.go.dev/flag) directly has a few advantages:
    - a well known (or at least opinionated) application layout
    - error handling and display to the user
    - help message and usage display
    - simpler flag enforcement
- Input validation is via a `validate() error` interface and joining all encountered errors. This enables composing more complex validations out of smaller building blocks. An extension of the current validations could include something like a `StartingBalanceValidator` which implement `Validator` and internally reference `PositiveValueValidator` and a `MinValueValidator`.
    - Note: there are some input validations - required argument checks and basic paramater validation - but it is likely that there are inputs that will produce inaccurate results or break the application.
- The `TermDeposit` struct handles most operations and is the public interface for the total balance calculation. This is on the path to something that I would be happy with but it's not quite there. The idea to is to protect the user from inaccurate / broken results by forcing validation to run first. The current implementation has an awkward interface that binds the validation, calculation and output all together. Returning validated arguments or a `TermDeposit` from the validator could be a way to keep the guarantee but with a better interface. Injecting a presenter or output handler would also improve the situation.

### Assumptions

- All proceeds are reinvested for the term of the deposit duration
- A simple interest calculation is used when the interest paid valid is 'maturity'
- A compound interest calculation is used for all other interest paid values

## Running the application

There are a couple of options for running the application.

Run with `-h` to see command help.

### Using Go

Run directly with `go run`

```
go run main.go --startingBalance 10000 --termYears 3 --interestRate 1.1 --interestPaid maturity
```

 Or build and run

```
go build 

./term-deposit --startingBalance 10000 --termYears 3 --interestRate 1.1 --interestPaid maturity
```

### Using Docker

```
docker build . -t termdeposit:1
docker run termdeposit:1 --startingBalance 10000 --termYears 3 --interestRate 1.1 --interestPaid maturity
```

## Running tests

```
go test -v ./...
```

## Potential changes given more time

To keep to the time limit, I've left some things as they were written, not necessarily where/how they should be. Here are some of the things that would improve it.

* Standalone validators for each input argument and a better way to construct them e.g. something that implements `ArgumentValidator(name string, validations []Validator) Validator`
* More comprehensive tests. The existsing tests are ok to validate that common cases work but extension via property based testing or at least range tests for inputs would be better.
* Clear up the output responsibility. There could be a way to supply a presenter to handle formatting and output. Using `fmt.Printf` is fine for a command line but useless if there was a webserver component.
* Clean up variable naming. It's inconsistent and potentially confusing.
* Improve behaviour locations. Because validation, calculation and presentation are all coupled, the behaviour locations are awkward.
* Improved method visibilty / module public interface. `TotalBalance` should at least be module private or moved to `TermDeposit`
* Normalise balance rounding or, probably better, make it a presentation concern rather than a calculation one.
* More realistic input validation for parameters e.g. banks are unlikely to offer unbounded interest rates on any balance
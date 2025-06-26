/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/seancaffery/term-deposit/term_deposit"
	"github.com/spf13/cobra"
)

var interestPaid string
var termYears int
var interestRate float64
var startingBalance float64

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "term-deposit",
	Short: "A term deposit total balance calulator",
	Long: `Calculates the resulting balance for a term deposit given the following:
- starting balance (e.g. 10000)
- interest rate (e.g. 1.10)
- investment term in years (e.g. 3)
- interest payment frequency (monthly, quarterly, annually, at maturity)
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(term_deposit.TotalBalance(startingBalance, interestRate, termYears, interestPaid))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.term-deposit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().Float64Var(&startingBalance, "startingBalance", 0.0, "The initial balance of the term deposit")
	rootCmd.MarkFlagRequired("startingBalance")

	rootCmd.Flags().Float64Var(&interestRate, "interestRate", 1.1, "The interest rate of the term deposit")
	rootCmd.MarkFlagRequired("interestRate")

	rootCmd.Flags().IntVar(&termYears, "termYears", 3, "The investment term in years")
	rootCmd.MarkFlagRequired("termYears")

	rootCmd.Flags().StringVar(&interestPaid, "interestPaid", "maturity", "Interest payment frequency. Available values: monthly, quarterly, annually, maturity")
	rootCmd.MarkFlagRequired("interestPaid")

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}

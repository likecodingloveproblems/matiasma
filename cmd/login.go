package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/likecodingloveproblems/matiasma/internal/auth"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login <phone-number>",
	Short: "Login to Telegram account",
	Long: `Authenticate with Telegram using your phone number.
You will receive a verification code via Telegram that you'll need to input.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		phoneNumber := args[0]
		if err := auth.Login(ctx, phoneNumber); err != nil {
			fmt.Printf("Login failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Successfully logged in!")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

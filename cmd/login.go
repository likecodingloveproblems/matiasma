package cmd

import (
	"context"
	"fmt"
	"github.com/gotd/td/examples"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	_ "github.com/lib/pq"
	"github.com/likecodingloveproblems/matiasma/internal/session"
	"go.uber.org/zap"
	"time"

	"github.com/joho/godotenv"
	my_auth "github.com/likecodingloveproblems/matiasma/internal/auth"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login <phone-number>",
	Short: "Login to Telegram account",
	Long: `Authenticate with Telegram using your phone number.
You will receive a verification code via Telegram that you'll need to input.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewProduction()
		defer func(logger *zap.Logger) {
			err := logger.Sync()
			if err != nil {
				fmt.Println(err.Error())
			}
		}(logger)
		err := godotenv.Load()
		if err != nil {
			panic(err.Error())
		}
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()

		phoneNumber := args[0]
		sessionStorage := session.New(ctx, phoneNumber, logger)
		defer func(sessionStorage *session.PostgresSessionStorage) {
			err := sessionStorage.Close()
			if err != nil {
				logger.Error(fmt.Sprintf("Can not close session storage connection: %s", err.Error()))
			}
		}(sessionStorage)
		client, err := telegram.ClientFromEnvironment(
			telegram.Options{
				Logger:         logger,
				SessionStorage: sessionStorage,
			},
		)
		if err != nil {
			panic(err.Error())
		}
		flow := auth.NewFlow(examples.Terminal{PhoneNumber: phoneNumber}, auth.SendCodeOptions{})
		err = client.Run(ctx, my_auth.AuthenticateIfNecessary(client, flow, phoneNumber, logger))
		if err != nil {
			panic(err.Error())
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
	loginCmd.PersistentFlags().String("phone_number", "", "A help for phone number")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

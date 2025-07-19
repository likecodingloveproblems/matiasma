package cmd

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gotd/td/telegram"
	_ "github.com/lib/pq"
	"github.com/likecodingloveproblems/matiasma/internal/session"
	"go.uber.org/zap"
	"os"
	"strings"
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

		codeChannel := make(chan string)
		passwordChannel := make(chan string)
		notifPasswordRequiredChannel := make(chan struct{})

		codeChannelWriter := func() {
			fmt.Print("Enter code: ")
			code, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				panic(err.Error())
			}
			codeChannel <- strings.TrimSpace(code)
		}
		passwordChannelWriter := func() {
			select {
			case <-notifPasswordRequiredChannel:
				fmt.Print("Enter password: ")
				code, err := bufio.NewReader(os.Stdin).ReadString('\n')
				if err != nil {
					panic(err.Error())
				}
				passwordChannel <- strings.TrimSpace(code)
			case <-ctx.Done():
				logger.Info("Context canceled while waiting for password requirement")
				return
			}
		}
		channelAuthenticator := my_auth.New(phoneNumber, codeChannel, passwordChannel, notifPasswordRequiredChannel)

		err = client.Run(ctx, my_auth.AuthenticateIfNecessary(client, channelAuthenticator, logger, codeChannelWriter, passwordChannelWriter))
		if err != nil {
			panic(err.Error())
		}
		cancel()
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

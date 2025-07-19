package auth

import (
	"context"
	"fmt"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"go.uber.org/zap"
	_ "time"

	"github.com/gotd/td/tg"
	"golang.org/x/crypto/ssh/terminal"
)

type AuthFlow struct {
	PhoneNumber string
}

//func (a AuthFlow) SignUp(ctx context.Context, client auth.SignUpClient) (*tg.AuthAuthorization, error) {
//	return nil, fmt.Errorf("signup not implemented")
//}

func (a AuthFlow) AcceptTermsOfService(ctx context.Context, tos tg.HelpTermsOfService) error {
	return nil
}

func (a AuthFlow) Code(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
	fmt.Print("Enter code: ")
	code, err := terminal.ReadPassword(0)
	if err != nil {
		return "", err
	}
	return string(code), nil
}

func (a AuthFlow) Password(ctx context.Context, passwordHint string) (string, error) {
	fmt.Print("Enter password: ")
	pass, err := terminal.ReadPassword(0)
	if err != nil {
		return "", err
	}
	return string(pass), nil
}

func (a AuthFlow) Phone(ctx context.Context) (string, error) {
	return a.PhoneNumber, nil
}

func Login(ctx context.Context, client *telegram.Client, apiID int, apiHash string, phoneNumber string) error {
	//flow := AuthFlow{PhoneNumber: phoneNumber}
	//
	//// Setup Telegram client with QR code login support
	//qr := qrlogin.NewQR(client, apiID, apiHash, qrlogin.Options{})
	//client, err := auth.NewClient(qr, flow, apiID, apiHash)
	//if err != nil {
	//	return fmt.Errorf("create client: %w", err)
	//}
	//
	//if err := qr.Run(ctx); err != nil {
	//	return fmt.Errorf("qr login: %w", err)
	//}
	//
	//// Start authentication flow
	//if _, err := client.Auth(ctx, flow.PhoneNumber, func(ctx context.Context, client auth.Client) error {
	//	return client.AuthFlow(ctx, flow)
	//}); err != nil {
	//	return fmt.Errorf("auth flow: %w", err)
	//}

	return nil
}

func AuthenticateIfNecessary(client *telegram.Client, flow auth.Flow, phoneNumber string, logger *zap.Logger) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		// Checking auth status.
		status, err := client.Auth().Status(ctx)
		if err != nil {
			return err
		}
		// Can be already authenticated if we have valid session in
		// session storage.
		if !status.Authorized {
			// Otherwise, perform bot authentication.
			if err := client.Auth().IfNecessary(ctx, flow); err != nil {
				return err
			}
		}

		// All good, manually authenticated.
		logger.Info("Done")
		return nil
	}
}

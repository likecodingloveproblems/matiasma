package auth

import (
	"context"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"go.uber.org/zap"
	_ "time"
)

func AuthenticateIfNecessary(client *telegram.Client, authenticator auth.UserAuthenticator, logger *zap.Logger, codeChannelWrite func(), passwordChannelWriter func()) func(ctx context.Context) error {
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
			go codeChannelWrite()
			go passwordChannelWriter()
			flow := auth.NewFlow(authenticator, auth.SendCodeOptions{})
			if err := client.Auth().IfNecessary(ctx, flow); err != nil {
				return err
			}
		}

		// All good, manually authenticated.
		logger.Info("Done")
		return nil
	}
}

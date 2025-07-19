package auth

import (
	"context"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
)

type ChannelAuthFlow struct {
	phoneNumber                  string
	codeChannel                  <-chan string
	passwordChannel              <-chan string
	notifPasswordRequiredChannel chan<- struct{}
}

func New(phoneNumber string, codeChannel, passwordChannel <-chan string, notifPasswordRequiredChannel chan<- struct{}) *ChannelAuthFlow {
	return &ChannelAuthFlow{
		phoneNumber:                  phoneNumber,
		codeChannel:                  codeChannel,
		passwordChannel:              passwordChannel,
		notifPasswordRequiredChannel: notifPasswordRequiredChannel,
	}
}

func (c ChannelAuthFlow) Phone(ctx context.Context) (string, error) {
	return c.phoneNumber, nil
}

func (c ChannelAuthFlow) Password(ctx context.Context) (string, error) {
	c.notifPasswordRequiredChannel <- struct{}{}
	return <-c.passwordChannel, nil
}

func (c ChannelAuthFlow) AcceptTermsOfService(ctx context.Context, tos tg.HelpTermsOfService) error {
	panic("it's not implemented")
}

func (c ChannelAuthFlow) SignUp(ctx context.Context) (auth.UserInfo, error) {
	panic("its not implemented")
}

func (c ChannelAuthFlow) Code(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
	return <-c.codeChannel, nil
}

var _ auth.UserAuthenticator = &ChannelAuthFlow{}

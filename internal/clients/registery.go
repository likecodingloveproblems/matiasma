package clients

import (
	"github.com/gotd/td/telegram"
	"sync"
)

type Registry interface {
	AddNewClient(phoneNumber string, client *telegram.Client) error
	DoesClientExists(phoneNumber string) bool
}

type ClientRegistry struct {
	lock    sync.RWMutex
	clients map[string]*telegram.Client
}

func New() ClientRegistry {
	return ClientRegistry{}
}

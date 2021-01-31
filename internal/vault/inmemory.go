package vault

import (
	"fmt"
)

// InMemoryValues stores key/values pairs
var InMemoryValues map[string]string = map[string]string{}

// InMemoryClient implements the interface
// vaultInterface to a in memory storage.
// Warning: Don't use it in production!
type InMemoryClient struct{}

// Save encrypt and stores the value inside the map using the
// given key
func (c InMemoryClient) Save(key string, value string) error {
	InMemoryValues[key] = value

	return nil
}

// Retrieve given the key, retrive it from map, decrypt, and
// return it's value.
func (c InMemoryClient) Retrieve(key string) (string, error) {
	if value, hasValue := InMemoryValues[key]; hasValue {
		return value, nil
	}

	return "", fmt.Errorf("The key %s doesn't exists", key)
}

func newInMemoryClient(backendType BackendType) (Vault, error) {
	return Vault{
		Type:   backendType,
		Client: InMemoryClient{},
	}, nil
}

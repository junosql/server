package vault

import "github.com/hashicorp/vault/api"

// github.com/hashicorp/vault/api

// HashiCorpClient implements the
// interface vaultInterface with
// HashiCorp Vault
type HashiCorpClient struct {
	client api.Client
}

// Save encrypts and save a given value under
// a key in HashiCorp Vault
func (c HashiCorpClient) Save(key string, value string) error {
	return nil
}

// Retrieve gets a decrypted value for a given key
func (c HashiCorpClient) Retrieve(key string) (string, error) {
	return "", nil
}

func newHashiCorpClient(backendType BackendType) (Vault, error) {
	hashiCorpClient, err := api.NewClient(&api.Config{})

	thisClient := HashiCorpClient{
		client: *hashiCorpClient,
	}

	thisVault := Vault{
		Type:   backendType,
		Client: thisClient,
	}

	if err != nil {
		return thisVault, err
	}

	return thisVault, nil
}

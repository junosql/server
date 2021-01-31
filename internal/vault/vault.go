// Package vault provides functions
// to interact with secrets vault used
// by JunoSQL and do operations like
// secrets retrieve or save new secret.
package vault

// Vault represents a Vault client
type Vault struct {
	// Type indicates what is the backend used
	Type BackendType

	// Client is a struct implementing the Interface
	// for integrate with a specific backend
	Client Interface
}

// Interface is the default interface
// of integration with Vault backends.
type Interface interface {
	Save(string, string) error
	Retrieve(string) (string, error)
}

// BackendType is the type of backend
type BackendType int

const (
	// InMemoryBackend indicates to JunoSQL Vault
	// that the secrets must be stored in memory.
	// Not suitable for production.
	InMemoryBackend BackendType = iota + 1

	// HashiCorpBackend indicates to JunoSQL Vault
	// that the secrets must be stores in HashiCorp
	// Vault. Is the default option.
	HashiCorpBackend
)

// NewVault creates a new
func NewVault(backendType BackendType) (Vault, error) {
	switch backendType {

	case InMemoryBackend:
		return newInMemoryClient(backendType)

	case HashiCorpBackend:
		return newHashiCorpClient(backendType)

	default:
		return newHashiCorpClient(backendType)
	}
}

// Save receives a value referenced by a key,
// encrypt it's value, and save on the Vault backend
func (v Vault) Save(key string, value string) error {
	return v.Client.Save(key, value)
}

// Retrieve gets the value from storage, decrypt
// and return it.
func (v Vault) Retrieve(key string) (string, error) {
	return v.Client.Retrieve(key)
}

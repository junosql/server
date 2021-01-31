// Package vault provides functions
// to interact with secrets vault used
// by JunoSQL and do operations like
// secrets retrieve or save new secret.
package vault

// Vault represents a Vault client
type Vault struct {
	// Type indicates what is the backend used
	Type BackendType

	Conn interface{}
}

type vaultInterface interface {
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
func NewVault(backendType BackendType) Vault {
	switch backendType {

	case InMemoryBackend:
		return newInMemoryVault(backendType)

	case HashiCorpBackend:
		return newHashiCorpVault(backendType)

	default:
		return newHashiCorpVault(backendType)
	}
}

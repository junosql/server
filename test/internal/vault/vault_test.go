package vault__test

import (
	"testing"

	"github.com/junosql/server/internal/vault"
)

func TestNewVault(t *testing.T) {
	t.Run("when type is vault.InMemoryBackend", func(t *testing.T) {
		thisVault, _ := vault.NewVault(vault.InMemoryBackend)

		t.Run("should set Type according backendType", func(t *testing.T) {
			if thisVault.Type != vault.InMemoryBackend {
				t.Errorf("expected thisVault.Type go be %+v but it was %+v", vault.HashiCorpBackend, thisVault.Type)
			}
		})

		t.Run("should set Client according backendType", func(t *testing.T) {
			switch thisVault.Client.(type) {
			case vault.InMemoryClient:
				return

			default:
				t.Errorf("expected thisVault.Client go be %T but it was %T", vault.InMemoryClient{}, thisVault.Client)
			}
		})
	})
}

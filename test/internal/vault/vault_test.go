package vault__test

import (
	"testing"

	"github.com/junosql/server/internal/vault"
)

func TestVaultPackage(t *testing.T) {
	thisVault := vault.NewVault(vault.HashiCorpBackend)

	t.Run("should set type according argument backendType", func(t *testing.T) {
		if thisVault.Type != vault.HashiCorpBackend {
			t.Errorf("expected thisVault.Type go be %T but it was %T", vault.HashiCorpBackend, thisVault.Type)
		}
	})

}

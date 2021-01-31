package vault__test

import (
	"fmt"
	"testing"

	"github.com/junosql/server/internal/vault"

	uuid "github.com/satori/go.uuid"
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

func TestSave(t *testing.T) {
	t.Run("when type is vault.InMemoryBackend", func(t *testing.T) {
		thisVault, _ := vault.NewVault(vault.InMemoryBackend)

		t.Run("shouldn't return an error", func(t *testing.T) {
			if err := thisVault.Save("key", "value"); err != nil {
				t.Errorf("expected no error, but got %s", err)
			}
		})

		t.Run("should store the value inside global map", func(t *testing.T) {
			key := fmt.Sprintf("%s", uuid.NewV4())
			val := fmt.Sprintf("%s", uuid.NewV4())

			thisVault.Save(key, val)

			if vault.InMemoryValues[key] != val {
				t.Errorf("expected '%s', but got '%s'", val, vault.InMemoryValues[key])
			}
		})
	})
}

func TestRetrieve(t *testing.T) {
	t.Run("when type is vault.InMemoryBackend", func(t *testing.T) {
		thisVault, _ := vault.NewVault(vault.InMemoryBackend)

		t.Run("when key exists should return it's value with no errors", func(t *testing.T) {
			key := fmt.Sprintf("%s", uuid.NewV4())
			val := fmt.Sprintf("%s", uuid.NewV4())

			vault.InMemoryValues[key] = val

			actualVal, actualErr := thisVault.Retrieve(key)

			if actualErr != nil {
				t.Errorf("expected no error, but got '%s'", actualErr)
			}

			if actualVal != val {
				t.Errorf("expected '%s', but got '%s'", val, actualVal)
			}
		})

		t.Run("when key doesn't exists should return error", func(t *testing.T) {
			inexistentKey := "inexistentKey"

			expectVal := ""
			expectedErr := fmt.Errorf("The key %s doesn't exists", inexistentKey)

			actualVal, actualErr := thisVault.Retrieve(inexistentKey)

			if fmt.Sprintf("%s", actualErr) != fmt.Sprintf("%s", expectedErr) {
				t.Errorf("expected error '%s' but got '%s'", expectedErr, actualErr)
			}

			if actualVal != expectVal {
				t.Errorf("expected empty value but got '%s'", actualVal)
			}
		})
	})
}

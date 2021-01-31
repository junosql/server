package vault

func newInMemoryVault(backendType BackendType) Vault {
	return Vault{Type: backendType}
}

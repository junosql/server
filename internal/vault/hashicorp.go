package vault

func newHashiCorpVault(backendType BackendType) Vault {
	return Vault{Type: backendType}
}

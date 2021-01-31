package vault

func newInMemoryVault() Vault {
	return Vault{Type: "INMEMORY"}
}

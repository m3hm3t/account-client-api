package deleter

type AccountDeleter interface {
	DeleteAccount(accountID string, version string) (int, error)
}

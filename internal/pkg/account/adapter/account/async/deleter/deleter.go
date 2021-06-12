package deleter

type AccountDeleterAsync interface {
	DeleteAll(accountIDs []string, versions []string) ([]int, error)
}

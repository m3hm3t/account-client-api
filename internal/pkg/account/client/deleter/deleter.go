package deleter

type RestDeleter interface {
	MakeDeleteRequest(url string) (int, error)
}

package delete

type RestDeleter interface {
	MakeDeleteRequest(url string) (int, error)
}

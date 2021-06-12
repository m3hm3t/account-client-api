package get

type RestGetter interface {
	MakeGetRequest(url string) ([]byte, int, error)
}

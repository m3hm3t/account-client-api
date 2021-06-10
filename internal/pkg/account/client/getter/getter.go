package getter

type RestGetter interface {
	MakeGetRequest(url string) ([]byte, int, error)
}

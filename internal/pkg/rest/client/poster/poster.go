package poster

type RestPoster interface {
	MakePostRequest(url string, input interface{}) ([]byte, int, error)
}

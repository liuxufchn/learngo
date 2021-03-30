package mock

type Retriever struct {
	Content string
}

func (r *Retriever) Post(url string, param map[string]string) string {
	return r.Content
}

func (r Retriever) Get(url string) string {
	return r.Content
}

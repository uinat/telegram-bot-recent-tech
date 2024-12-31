package hackernews

type Story struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Url   string `json:"url"`
}

package msql

type Thing struct {
	Writer  string `json:"writer"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

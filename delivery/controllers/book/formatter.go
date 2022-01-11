package book

type CreateBookRequest struct {
	Title    string `json:"title"`
	AuthorID int    `json:"author"`
}

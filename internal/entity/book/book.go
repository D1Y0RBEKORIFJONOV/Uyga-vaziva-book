package bookEntity

import "time"

type (
	Book struct {
		ID           string    `json:"id"`
		Title        string    `json:"title"`
		Author       string    `json:"author"`
		PublisherDay time.Time `json:"publisherDay"`
	}
	CreateBookRequest struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	UpdateBookRequest struct {
		Title        string `json:"title"`
		Author       string `json:"author"`
		PublisherDay string `json:"publisherDay"`
		ID           string `json:"id"`
	}
	FieldValueReq struct {
		Field string `json:"field"`
		Value string `json:"value"`
	}
	GetAllBookReq struct {
		Field  string `json:"field"`
		Value  string `json:"value"`
		Limit  int    `json:"limit"`
		Offset int    `json:"offset"`
	}
)

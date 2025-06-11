package response

type Campaign struct {
	Metadata
	Title       string `json:"title"`
	Description string `json:"description"`
}

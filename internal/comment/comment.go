package comment

// Comment - a representation of a comment
type Comment struct {
	ID     string
	Slug   string `json:"slug"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

// Comments - a representation of a slice of comments
type Comments []Comment

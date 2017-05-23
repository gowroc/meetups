package query

// UserQueries .
type UserQueries struct {
	SelectAll        string
	Insert           string
	DeleteAll        string
	GetByID          string
	GetUserWithPosts string
}

// User queries.
var User = UserQueries{}

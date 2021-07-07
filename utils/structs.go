package utils

type Comment struct {
	Author   User                   `json:"author"`
	Body     string                 `json:"body"`
	ID       int                    `json:"id"`
	Post     map[string]interface{} `json:"post"`
	Replying map[string]interface{} `json:"replying"`
	Self     string                 `json:"self"`
}

type User struct {
	AboutMe     interface{} `json:"about_me"`
	Confirmed   bool        `json:"confirmed"`
	ID          int         `json:"id"`
	LastSeen    string      `json:"last_seen"`
	Location    interface{} `json:"location"`
	MemberSince string      `json:"member_since"`
	Name        string      `json:"name"`
	Self        string      `json:"self"`
	Username    string      `json:"username"`
}

type Post struct {
	Author   User          `json:"author"`
	Columns  []interface{} `json:"columns"`
	Comments []interface{} `json:"comments"`
	Content  string        `json:"content"`
	ID       int           `json:"id"`
	Private  bool          `json:"private"`
	Self     string        `json:"self"`
	Title    string        `json:"title"`
}

type Column struct {
	Author User   `json:"author"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Posts  []Post `json:"posts"`
	URL    string `json:"self"`
}

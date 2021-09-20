package utils

import (
	"time"
)

type Timestamp struct {
	time.Time
}

func (t *Timestamp) String() string {
	return t.Time.Format(time.RFC1123Z)
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	s := string(b)[1 : len(b)-1]
	parsedTime, err := time.Parse(time.RFC3339, s+"Z")
	if err != nil {
		return err
	}
	t.Time = parsedTime.Local()
	return nil
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type Comment struct {
	Author   User                   `json:"author"`
	Body     string                 `json:"body"`
	ID       int                    `json:"id"`
	Post     map[string]interface{} `json:"post"`
	Replying map[string]interface{} `json:"replying"`
	URL      string                 `json:"self"`
}

type User struct {
	AboutMe     interface{} `json:"about_me"`
	Confirmed   bool        `json:"confirmed"`
	ID          int         `json:"id"`
	LastSeen    Timestamp   `json:"last_seen"`
	Location    interface{} `json:"location"`
	MemberSince Timestamp   `json:"member_since"`
	Name        string      `json:"name"`
	Self        string      `json:"self"`
	Username    string      `json:"username"`
	Coins       float32     `json:"coins"`
	Experience  int         `json:"experience"`
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
	Coins    int           `json:"coins"`
}

type Column struct {
	Author User   `json:"author"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Posts  []Post `json:"posts"`
	URL    string `json:"self"`
}

type Notification struct {
	Id      int       `json:"id"`
	Message string    `json:"message"`
	Time    Timestamp `json:"timestamp"`
}

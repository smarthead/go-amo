package users

type User struct {
	ID     int         `json:"id"`
	Name   string      `json:"name"`
	Email  string      `json:"email"`
	Lang   string      `json:"lang"`
	Rights interface{} `json:"rights"`
	Links  interface{} `json:"_links"`
}

// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User          *UserInformation `json:"user,omitempty"`
	Authenticated *bool            `json:"authenticated,omitempty"`
	Message       *string          `json:"message,omitempty"`
}

type SignUpRequest struct {
	Username  *string `json:"username,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Email     *string `json:"email,omitempty"`
	Password  *string `json:"password,omitempty"`
}

type SignUpResponse struct {
	Message    *string `json:"message,omitempty"`
	StatusCode *int    `json:"statusCode,omitempty"`
}

type UserInformation struct {
	UserID    *string `json:"userId,omitempty"`
	Username  *string `json:"username,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Jwt       *string `json:"JWT,omitempty"`
}

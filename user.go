package testrail

import (
	"fmt"
	"net/url"
)

// User represents a User
type User struct {
	Email    string `json:"email"`
	ID       int    `json:"id"`
	IsActive bool   `json:"is_active"`
	Name     string `json:"name"`
}

// GetUser returns the user userID
func (c *Client) GetUser(userID int) (user User, err error) {
	err = c.sendRequest("GET", fmt.Sprintf("get_user/%d", userID), nil, &user)
	return
}

// GetUserByEmail returns the user corresponding to email email
func (c *Client) GetUserByEmail(email string) (user User, err error) {
	vals := url.Values{"email": []string{email}}
	err = c.sendRequest("GET", fmt.Sprintf("get_user_by_email?%s", vals.Encode()), nil, &user)
	return
}

// GetUsers returns the list of users
func (c *Client) GetUsers() (users []User, err error) {
	err = c.sendRequest("GET", "get_users", nil, &users)
	return
}

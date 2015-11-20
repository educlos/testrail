package testrail

import "strconv"

type User struct {
	Email    string `json:"email"`
	ID       int    `json:"id"`
	IsActive bool   `json:"is_active"`
	Name     string `json:"name"`
}

// Returns the existing user userID
func (c *Client) GetUser(userID int) (User, error) {
	returnUser := User{}
	err := c.sendRequest("GET", "get_user/"+strconv.Itoa(userID), nil, &returnUser)
	return returnUser, err
}

// Returns the existing user with email email
func (c *Client) GetUserByEmail(email string) (User, error) {
	returnUser := User{}
	err := c.sendRequest("GET", "get_user_by_email&email="+email, nil, &returnUser)
	return returnUser, err
}

// Returns the list of user
func (c *Client) GetUsers() ([]User, error) {
	returnUser := []User{}
	err := c.sendRequest("GET", "get_users", nil, &returnUser)
	return returnUser, err
}

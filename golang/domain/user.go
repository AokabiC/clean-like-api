package domain

type User struct {
	ID       UserID
	Username Username
}

func NewUser(id UserID, username Username) (*User, error) {
	return &User{
		ID:       id,
		Username: username,
	}, nil
}

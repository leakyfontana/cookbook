package repositories

type User struct {
	ID    int
	Name  string
	Email string
}

func GetUserByID(id int) (*User, error) {
	// Placeholder for retrieving user from database
	return &User{ID: id, Name: "John Doe", Email: "john@example.com"}, nil
}
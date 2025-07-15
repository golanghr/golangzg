package users

import "github.com/org/my-lib/books"

type User struct {
	ID    int
	Name  string
	Email string
}

func (u *User) GetRented() []books.Book {
	return []books.Book{}
}

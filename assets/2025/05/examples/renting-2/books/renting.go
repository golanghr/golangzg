package books

import (
	"time"

	"github.com/org/my-lib/users"
)

type Renting struct {
	ID     int
	BookID int
	UserID int
	Start  time.Time
	End    time.Time
}

type Library struct {
	Orders []Renting
}

var library Library

func Borrow(book Book, user users.User) error {
	library.Orders = append(library.Orders, Renting{
		ID:     len(library.Orders) + 1,
		BookID: book.ID,
		UserID: user.ID,
		Start:  time.Now(),
	})
	return nil
}

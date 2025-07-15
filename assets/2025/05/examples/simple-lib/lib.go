package main

import (
    "fmt"

    "github.com/org/my-lib/books"
    "github.com/org/my-lib/users"
)

func main() {
    book := books.Book{
        ID:     1,
        Title:  "Go Programming",
        Author: "Alice Gopher",
    }
    user := users.User{
        ID:    1,
        Name:  "Bob Smith",
        Email: "bsmith@domain.com",
    }
    fmt.Println("Book:", book.Title, "by", book.Author)
    fmt.Println("User:", user.Name+",", user.Email)
}

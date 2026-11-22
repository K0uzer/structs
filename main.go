package main

import "fmt"

type Book struct {
	title       string
	authtor     string
	year        int
	ISBN        string
	extradition bool
}

type Member struct {
	name      string
	email     string
	booksList []int
}

type Library struct {
	title   string
	address string
	books   []Book
	readers []Member
}

// type Example struct {
//     a bool // 1
// 	// 7
//     b int64 // 8
//     c int32 // 4
//     d bool // 1
// 	// 3
// }
// 8 + 16 = 24

// type Example struct {
//     b int64 // 8
//     c int32 // 4
//     a bool // 1
//     d bool // 1
// 	// 2
// }
// 16

func NewBook(title string, authtor string, year int, ISBN string, extradition bool) (*Book, error) {
	if title == "" || authtor == "" {
		return nil, fmt.Errorf("Пустое название или автор")
	}
	if year <= 0 {
		return nil, fmt.Errorf("Год меньше или равен нулю")
	}
	return &Book{
		title:       title,
		authtor:     authtor,
		year:        year,
		ISBN:        ISBN,
		extradition: extradition,
	}, nil
}

func NewMember(name string, email string, booksList []int) (*Member, error) {
	if name == "" {
		return nil, fmt.Errorf("Пустое имя")
	}
	if email == "" {
		return nil, fmt.Errorf("Пустой email")
	}
	return &Member{
		name:      name,
		email:     email,
		booksList: booksList,
	}, nil
}

func NewLibrary(title, address string) *Library {
	return &Library{
		title:   title,
		address: address,
		books:   []Book{},
		readers: []Member{},
	}
}

func main() {

}

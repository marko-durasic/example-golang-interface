package literature

import (
	"fmt"
	"time"
)

type book struct {
	Title  string
	Author string
	Pages  int
}

func NewBook(title, author string, pages int) *book {
	return &book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

func (b *book) Read() bool {
	for p := 1; p <= b.Pages; p++ {
		fmt.Println("Reading page: ", p, " from ", b.Title)
		time.Sleep(time.Second)
	}
	return true
}

func (b *book) GetAuthor() string {
	return b.Author
}

func (b *book) GetTitle() string {
	return b.Title
}

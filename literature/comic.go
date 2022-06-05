package literature

import (
	"fmt"
	"time"
)

type comic struct {
	Title  string
	Author string
	Pages  int
}

func NewComic(title, author string, pages int) *comic {
	return &comic{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

func (c *comic) Read() bool {
	for p := 1; p <= c.Pages; p++ {
		fmt.Println("Reading page: ", p, " from ", c.Title)
		time.Sleep(time.Second)
	}
	return true
}

func (c *comic) GetAuthor() string {
	return c.Author
}

func (c *comic) GetTitle() string {
	return c.Title
}

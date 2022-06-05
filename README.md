# Example code Golang (Go) with interface #

This is an example of how to use interfaces in golang and their benefits

## File structure ##

    .
    ├── literature           # Custom package
    │   ├── book.go          # Custom book struct
    │   ├── comic.go         # Custom comic struct
    │   └── literature.go    # Interface 
    ├── go.mod               # Go.mod for dependency management (generated by go mod init)
    └── main.go              # Main executable of our module

## Code ##

### literature/book.go ###

```go
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
```

### literature/comic.go ###

```go
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
```

### literature/literature.go ###

```go
package literature

type Literature interface {
 Read() bool
 GetAuthor() string
 GetTitle() string
}
```

### main.go ###

```go
package main

//do go mod init and go mod tidy

import (
 "fmt"
 "mine/literature"
 "sync"
)

func main() {
 //initialize different literatures
 b := literature.NewBook("Data Structures", "Marko", 100)
 c := literature.NewComic("Lady Bug Man", "Charles", 50)

 //run functions concurrently
 var wg sync.WaitGroup
 wg.Add(2)
 go func() { readAll(b); wg.Done() }()
 go func() { readAll(c); wg.Done() }()
 wg.Wait()
}

//custom function that takes any kind of literature
func readAll(l literature.Literature) {
 fmt.Println("Starting to read: "+l.GetTitle()+
  " by: ", l.GetAuthor())
 l.Read()
}
```

Run the following command in the root of your project:

```shell
go mod init mine 
```

Here the wording `mine` represents the module name that we chose. It can be anything.
This will result in creation of the following file ( `go.mod` ) in the root of your project:

```go
module mine

go 1.16
```

As we can see in the code above.
We have import path of the module at the top:

```go
module mine
```

and version of go that we are using:

```go
1.16
```

If we had any external dependencies they would be represented here. But in this example we don't have them.

## Interfaces ##

Interfaces in go are implemented implicitly.
Basically if a struct implements methods that are named in the interface, then it is said that that struct implements that interface.

So in our case, as long as the structs `book` and `comic` implement methods :

```go
 Read() bool
 GetAuthor() string
 GetTitle() string

```

they are considered to implement interface `Literature`.

Then we can use the function `readAll` and pass `Literature` as the input type:

```go
func readAll(l literature.Literature){}
```

This way we have achieved polymorphism in the `readAll` function.

One of the benefits of interfaces in Go is that we can define interfaces at a later time (after the structs are written) and then use them in our code.


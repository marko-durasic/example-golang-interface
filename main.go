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

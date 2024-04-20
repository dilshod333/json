package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type BookData struct {
	Books []Book `json:"books"`
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
}

func main() {

	f, err := os.Open("file.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var booksData BookData
	err = json.NewDecoder(f).Decode(&booksData)
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	for _, b := range booksData.Books {
		fmt.Println("Title:", b.Title)
		count++
		fmt.Println("Author:", b.Author)
		fmt.Println("Genre:", b.Genre)
		fmt.Println("Year:", b.Year)
		fmt.Println()
	}
	fmt.Println("count of the books are/is ", count)
	fmt.Println()
	fmt.Println("-----------------Before published 1950--------------------")
	fmt.Println()
	for _, i := range booksData.Books {
		if i.Year < 1950 {
			fmt.Println("Title:", i.Title)
			fmt.Println("Author:", i.Author)
			fmt.Println("Genre:", i.Genre)
			fmt.Println("Year:", i.Year)
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Println("--------------Added new book----------------")
	fmt.Println()

	new_book := Book{
		Title:  "atomic habits",
		Author: "David",
		Genre:  "self-study",
		Year:   2017,
	}

	booksData.Books = append(booksData.Books, new_book)
	for _, i := range booksData.Books {
			fmt.Println("Title:", i.Title)
			fmt.Println("Author:", i.Author)
			fmt.Println("Genre:", i.Genre)
			fmt.Println("Year:", i.Year)
			fmt.Println()
	}
}

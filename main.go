package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type bmksMap map[string]string

func createBookmark(bookmarks *bmksMap) {
	var title string
	var url string
	fmt.Println("-- Создание закладки -- ")
	fmt.Print("Введите название сайта: ")
	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Введите URL адрес: ")
	reader1 := bufio.NewReader(os.Stdin)
	url, err = reader1.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	(*bookmarks)[strings.TrimSpace(title)] = strings.TrimSpace(url)
}

func printBookmark(bookmark *bmksMap) {
	var title string
	fmt.Print("Введите необходимое название ссылки: ")
	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%8s | %12s\n", strings.TrimSpace(title), (*bookmark)[strings.TrimSpace(title)])

}

func printBookmarks(bookmarks *bmksMap) {
	for key, value := range *bookmarks {
		fmt.Printf("%8s | %12s\n", key, value)

	}
}

func delBookmark(bookmakr *bmksMap) {
	var title string
	fmt.Print("Введите необходимое название ссылки: ")
	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	delete(*bookmakr, strings.TrimSpace(title))
}

func menu() {
	fmt.Println(`
1. Создать закладку
2. Вывести все закладки
3. Вывести закладку по ее названию
4. Удалить закладку
5. Выход	
	`)
}

func choisePart(bookmart *bmksMap) error {
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		createBookmark(bookmart)
	case 2:
		printBookmarks(bookmart)
	case 3:
		printBookmark(bookmart)
	case 4:
		delBookmark(bookmart)
	case 5:
		return errors.New("выход")
	}
	return nil

}

func main() {
	bookMarks := make(bmksMap)

	for {
		menu()
		err := choisePart(&bookMarks)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// type BookMarkHistory = dictionary{"title":"url"}
type bmh map[string]string

func textAlignCenter(text string, maxWidth int, separator string) string {
	lenText := utf8.RuneCountInString(text)
	spaceCount := (maxWidth - lenText) / 2
	spaceString := strings.Repeat(separator, spaceCount)
	return fmt.Sprintf("%s%s%s", spaceString, text, spaceString)
}

func menu(maxWidth int) {
	fmt.Println(textAlignCenter("<<<<--Меню-->>>>", maxWidth+3, "-"))
	fmt.Printf("%2s %s\n", "(1)", textAlignCenter("Создать закладку", maxWidth, " "))
	fmt.Printf("%2s %s\n", "(2)", textAlignCenter("Вывести все закладки", maxWidth, " "))
	fmt.Printf("%2s %s\n", "(3)", textAlignCenter("Вывести закладку по имени", maxWidth, " "))
	fmt.Printf("%2s %s\n", "(4)", textAlignCenter("Удалить закладку", maxWidth, " "))
	fmt.Printf("%2s %s\n", "(5)", textAlignCenter("Выйти из программы", maxWidth, " "))
	fmt.Println(textAlignCenter("<<<-->>>", maxWidth+3, "-"))
}

func choiceMenuPart(bookmarks *bmh) error {
	var choice int
	fmt.Println()
	fmt.Print("Введите номер варианта: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		err := setBoormark(bookmarks)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			fmt.Println()
			break
		}
	case 2:
		getBookmarks(bookmarks)
	case 3:
		bm, err := getBookmark(bookmarks)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			fmt.Println()
			break
		} else {
			fmt.Println()
			fmt.Println(bm)
			fmt.Println()
			break
		}
	case 4:
		err := deleteBokmark(bookmarks)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			fmt.Println()
			break
		}
	case 5:
		return errors.New("EXIT")
	default:
		fmt.Println()
		fmt.Println("Неизвестная команда")
		fmt.Println()
	}
	return nil
}

func setBoormark(bookmarks *bmh) error {
	var title string
	var url string
	fmt.Println(textAlignCenter("Создание закладки", 28, " "))
	fmt.Print("Введите название закладки: ")
	reader := bufio.NewReader(os.Stdin)
	title, _ = reader.ReadString('\n')
	clearTitle := strings.TrimSpace(title)
	if utf8.RuneCountInString(clearTitle) == 0 {
		return errors.New("задано пустое название закладки")
	}
	fmt.Print("Введите URL адрес: ")
	reader = bufio.NewReader(os.Stdin)
	url, _ = reader.ReadString('\n')
	clearURL := strings.TrimSpace(url)
	if utf8.RuneCountInString(clearURL) == 0 {
		return errors.New("задан пустой URL адрес")
	}
	(*bookmarks)[clearTitle] = clearURL
	fmt.Println()
	fmt.Println("Закладка успешно записана")
	fmt.Println()
	return nil
}

func deleteBokmark(bookmarks *bmh) error {
	if len(*bookmarks) == 0 {
		return errors.New("у вас еще нет никаких закладок")
	}
	var title string
	fmt.Print("Введите назание желаемой закладки: ")
	reader := bufio.NewReader(os.Stdin)
	title, _ = reader.ReadString('\n')
	clearTitle := strings.TrimSpace(title)
	if utf8.RuneCountInString(clearTitle) == 0 {
		return errors.New("вы не ввели название закладки")
	}
	_, flag := (*bookmarks)[clearTitle]
	if flag {
		delete(*bookmarks, clearTitle)
	} else {
		return errors.New("такой закладки нет")
	}
	fmt.Println()
	fmt.Println("Ваша закладка успешно удалена")
	fmt.Println()

	return nil
}

func getBookmark(bookmarks *bmh) (string, error) {
	var title string
	fmt.Print("Введите название закладки: ")
	reader := bufio.NewReader(os.Stdin)
	title, _ = reader.ReadString('\n')
	clearTitle := strings.TrimSpace(title)
	if utf8.RuneCountInString(clearTitle) == 0 {
		return "", errors.New("вы не ввели название закладки")
	}
	_, flag := (*bookmarks)[clearTitle]
	if flag {
		bmText := fmt.Sprintf("%s : %s", clearTitle, (*bookmarks)[clearTitle])
		bm := textAlignCenter(bmText, utf8.RuneCountInString(bmText)*2, " ")
		return bm, nil
	} else {
		return "", errors.New("такой закладки нет")
	}
}

func getBookmarks(bookmarks *bmh) {
	if len(*bookmarks) == 0 {
		fmt.Println()
		fmt.Println("У вас еще нет закладок")
		fmt.Println()
	}
	for key, value := range *bookmarks {
		fmt.Println(key, value)
	}
}

func main() {
	bookmarks := make(bmh)
	for {

		menu(28)
		err := choiceMenuPart(&bookmarks)
		if err != nil {
			break
		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display() error
// }

type outputtable interface {
	saver
	Display()
	DoSomething(int) string
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println()
		return
	}

	printSomething(todo)

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	// todo.Display()
	// err = saveData(todo)
	err = outputData(todo)

	if err != nil {
		fmt.Println("Saving the todo failed")
		return
	}

	// userNote.Display()
	// err = saveData(userNote)
	outputData(userNote)

	// if err != nil {
	// 	return
	// }
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Value is an integer:", intVal)
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Value is an float:", floatVal)
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("Value is an string:", stringVal)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float: ", value)
	// case string:
	// 	fmt.Println("String: ", value)
	// default:
	// 	fmt.Println("Unknown type: ", value)
	// }
	// fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)
	// var value string

	// fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

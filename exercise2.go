/*
Exercise 2: Writing files
Write a go program:
- Where a text file is created, e.g sample.txt.
- If the file already exists then truncate the contents of the file.
- Write some random data to the file (“Lorem ispum”) and close the file.
- Then read the same file and display its contents on the console.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func createFile(filename string) {
	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v file created successfully\n", filename)

	defer f.Close()
}

func writeFile(filename, content string) {
	f, err := os.OpenFile(filename, os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString(content)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v file written successfully\n", filename)

	defer f.Close()
}

func readFile(filename string) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v file read successfully\n", filename)
	fmt.Printf("Content: %v\n", string(content))

	defer f.Close()
}

func main() {
	createFile("sample.txt")
	writeFile("sample.txt", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")
	readFile("sample.txt")
}
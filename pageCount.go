package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func getCount() string {
	// read the whole file at once
	file, err := os.OpenFile("count.txt", os.O_RDWR, os.ModeAppend) // For read access.
	if err != nil {
		log.Println(err)
		return "gagal 1"
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	co, err := strconv.Atoi(string(data[:count]))
	if err != nil {
		fmt.Println(err)
	}
	co = co + 1

	cos := strconv.Itoa(co)
	writeFile(cos)
	return cos
}

func writeFile(count string) {
	file, err := os.OpenFile("count.txt", os.O_WRONLY, os.ModeAppend) // For read access.
	if err != nil {
		log.Println(err)
	}
	_, err = file.WriteString(count)
	if err != nil {
		log.Println(err)
	}
}

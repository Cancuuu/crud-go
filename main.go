package main

import (
	"encoding/json"
	"io"
	"os"

	tasks "github.com/cancuuu/go-api/tasks"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR, 0666)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []tasks.Task

	info, err := file.Stat()
	errChecking(err)

	fileIsNotEmpty := info.Size() != 0

	if fileIsNotEmpty {
		// read file
		fileBytes, err := io.ReadAll(file)
		errChecking(err)

		// parse bytes to json and save into the tasks array
		jsonErr := json.Unmarshal(fileBytes, &tasks)
		errChecking(jsonErr)
	} else {

	}

}

func errChecking(err error) {
	if err != nil {
		panic(err)
	}
}

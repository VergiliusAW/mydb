package helpers

import (
	"github.com/atotto/clipboard"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const FilePath = "d:/Ashcheulov/dat/mydb.txt"

// CopyToClipboard Скопировать в буфер обмена
func CopyToClipboard(name string) {
	dat, _ := ioutil.ReadFile(FilePath)
	if strings.Contains(string(dat), name) {
		lines := strings.Split(string(dat), "\n")
		for i := range lines {
			line := strings.Split(lines[i], " ")
			if line[0] == name {
				clipboard.WriteAll(line[1])
			}
		}
	}
}

// Get получить value переменной
func Get(name string) string {
	dat, _ := ioutil.ReadFile(FilePath)
	if strings.Contains(string(dat), name) {
		lines := strings.Split(string(dat), "\n")
		for i := range lines {
			line := strings.Split(lines[i], " ")
			if line[0] == name {
				return line[1]
			}
		}
		return "error"
	}
	return "no variable found"
}

//добавить или изменить переменную
func Set(name string, value string) {
	dat, _ := ioutil.ReadFile(FilePath)
	if strings.Contains(string(dat), name) {

		return
	}
	f, err := os.OpenFile(FilePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(name + " " + value + "\n"); err != nil {
		log.Println(err)
	}
}

package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type File struct {
	filename *string
	grep *string
}

func New(filename *string, grep *string) *File {
	return &File{
		filename: filename,
		grep:     grep,
	}
}

func (f *File) IsExists() bool {
	_, err := os.Open(*f.filename)
	return err != nil
}

func (f *File) Find() {
	file, err := os.Open(*f.filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	buf := bufio.NewScanner(file)
	line := 1

	for buf.Scan() {
		colorReset := "\033[0m"
		colorRed := "\x1b[31m"
		text := strings.ToLower(buf.Text())
		grep := strings.ToLower(*f.grep)
		if strings.Contains(text, grep) {
			current_idx := strings.Index(text, grep)
			last_idx := (strings.Index(text, grep)-1) + len(grep)
			string_out := []string{}
			for _, data := range text {
				string_out = append(string_out, string(data))
			}
			for i := current_idx; i <= last_idx; i++ {
				string_out[i] = "";
			}
			string_out[current_idx] = colorRed + grep
			string_out[last_idx] = colorReset + string_out[last_idx]
			output := strings.Join(string_out, "") + fmt.Sprintf(" Line: %d Chars: %d", line, current_idx-1)
			fmt.Println(output)
		}
		line++
	}
}
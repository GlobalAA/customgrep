package main

import (
	"flag"
	"log"
	"main/internal/app"
	"strings"
)

func logHelpFatal() {
	log.Fatalln("Check that all arguments are specified!")
}

func main() {
	grep := flag.String("grep", "", "What to look for")
	filename := flag.String("path", "", "What file open")
	flag.Parse()
	
	if len(strings.Trim(*grep, " ")) <= 0 || len(strings.Trim(*filename, " ")) <= 0 {
		logHelpFatal()
	}
	
	file := app.New(filename, grep)
	if file.IsExists() {
		log.Fatalln("File not found!")
	}
	file.Find()
}
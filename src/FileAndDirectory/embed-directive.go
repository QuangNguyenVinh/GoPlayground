package main

import (
	"embed"
)

//"go:embed" is a compiler directive that allows programs to include arbitrary files and folders in the Go binary at build time

//go:embed text.txt
var fileString string

//go:embed text.txt
var fileByte []byte

//go:embed text.txt
//go:embed EmbedDirective/*.txt
var folder embed.FS

func main() {

	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("EmbedDirective/text1.txt")
	print(string(content1))

	content2, _ := folder.ReadFile("EmbedDirective/text2.txt")
	print(string(content2))
}

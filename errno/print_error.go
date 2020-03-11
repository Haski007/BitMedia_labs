package errno

import (
	"github.com/gookit/color"
	"log"
)

// PrintError - print your error with red sign ERROR!
func PrintError(err error) {
	color.Error.Printf("Error: ")
	log.Println(err)
}
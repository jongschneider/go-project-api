package main

import (
	"log"
	"os"
	"strconv"

	"github.com/jongschneider/go-project/server"
)

// default port number
var port = 8000

func init() {
	rawPort := os.Getenv("PORT")
	if len(rawPort) > 0 {
		var err error
		if port, err = strconv.Atoi(rawPort); err != nil {
			log.Fatalln(err)
		}
	}
}
func main() {
	s := server.New(port)
	s.Start()
	log.Println("Application has run!")
}

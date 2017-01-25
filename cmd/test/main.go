package main

import (
	"os"

	"github.com/jncornett/lg"
)

func main() {
	logger := lg.Wrap{lg.New(os.Stdin)}
	logger.Infof("hello %v", "world")
}

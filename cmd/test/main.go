package main

import "github.com/jncornett/lg"

func main() {
	logger := lg.New(nil, lg.ColorFormatter, 0)
	for i := 0; i < 10; i++ {
		logger.Infof("hello %v", "world")
	}
}

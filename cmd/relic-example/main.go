package main

import (
	"github.com/aliocode/relic-example/internal/app"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	c, err := app.New()
	if err != nil {
		return err
	}
	defer c.Close()
	return c.Run()
}

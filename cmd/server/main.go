package main

import "refty-backend/internal/app"

func main() {
	a, err := app.New()
	if err != nil {
		panic(err)
	}

	if err := a.Run(); err != nil {
		panic(err)
	}
}

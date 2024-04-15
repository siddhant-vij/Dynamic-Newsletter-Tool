package main

import (
	"github.com/siddhant-vij/Dynamic-Newsletter-Tool/internal"
)

func main() {
	dynamicContent, err := internal.ParseJson("newsletter.json")
	if err != nil {
		panic(err)
	}

	err = internal.GenerateHtml(dynamicContent)
	if err != nil {
		panic(err)
	}

	err = internal.StartHttpServer()
	if err != nil {
		panic(err)
	}
}

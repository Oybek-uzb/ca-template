package main

import (
	"ca-library-app/internal/composites"
	"context"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// entrypoint
	//bookStorage := book2.NewStorage()
	//bookService := book.NewService(bookStorage)
	//bookHandler := book3.NewHandler(bookService)

	router := httprouter.New()

	mongodbComposite, m := composites.NewMongoDBComposite(context.Background(), "", "", "", "", "", "")
	if m != nil {
		return
	}

	bookComposite, err := composites.NewBookComposite(mongodbComposite)
	if err != nil {
		return
	}

	//authorHandler := author3.NewHandler(authorService)
	//authorService := author.NewService(authorStorage)
	//authorStorage := author2.NewStorage()
	authorComposite, err := composites.NewAuthorComposite(mongodbComposite)
	if err != nil {
		return
	}

	bookComposite.Handler(router)
	authorComposite.Handler(router)

	//start
}

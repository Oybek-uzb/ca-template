package composites

import (
	"ca-library-app/internal/adapters/api"
	"ca-library-app/internal/adapters/api/book"
	book3 "ca-library-app/internal/adapters/db/book"
	book2 "ca-library-app/internal/domain/book"
)

type BookComposite struct {
	Service book.Service
	Storage book2.Storage
	Handler api.Handler
}

func NewBookComposite(mongodbComposite *MongoDBComposite) (*BookComposite, error) {
	storage := book3.NewStorage(mongodbComposite.db)
	service := book2.NewService(storage)
	handler := book.NewHandler(service)

	return &BookComposite{
		Service: service,
		Storage: storage,
		Handler: handler,
	}, nil
}

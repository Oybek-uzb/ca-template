package composites

import (
	"ca-library-app/internal/adapters/api"
	"ca-library-app/internal/adapters/api/author"
	author3 "ca-library-app/internal/adapters/db/author"
	author2 "ca-library-app/internal/domain/author"
)

type AuthorComposite struct {
	Service author.Service
	Storage author2.Storage
	Handler api.Handler
}

func NewAuthorComposite(mongodbComposite *MongoDBComposite) (*AuthorComposite, error) {
	storage := author3.NewStorage(mongodbComposite.db)
	service := author2.NewService(storage)
	handler := author.NewHandler(service)
	return &AuthorComposite{
		Service: service,
		Storage: storage,
		Handler: handler,
	}, nil
}

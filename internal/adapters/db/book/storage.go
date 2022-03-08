package book

import "ca-library-app/internal/domain/book"

type bookStorage struct {
}

func NewStorage() book.Storage {
	return &bookStorage{}
}

func (bs *bookStorage) GetOne(uuid string) *book.Book {
	return nil
}
func (bs *bookStorage) GetAll(limit, offset int) []*book.Book {
	return nil
}
func (bs *bookStorage) Create(book *book.Book) *book.Book {
	return nil
}
func (bs *bookStorage) Delete(uuid string) error {
	return nil
}

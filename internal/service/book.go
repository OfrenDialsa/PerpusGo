package service

import (
	"PerpusGo/domain"
	"PerpusGo/dto"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type bookService struct {
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
}

func NewBook(bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository) domain.BookService {
	return &bookService{
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
	}
}

// Index implements domain.BookService.
func (b *bookService) Index(ctx context.Context) ([]dto.BookData, error) {
	result, err := b.bookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var bookData []dto.BookData
	for _, c := range result {
		bookData = append(bookData, dto.BookData{
			Id:          c.Id,
			Isbn:        c.Isbn,
			Title:       c.Title,
			Description: c.Description,
		})
	}
	return bookData, nil
}

// Create implements domain.BookService.
func (b *bookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
	book := domain.Book{
		Id:          uuid.NewString(),
		Isbn:        req.Isbn,
		Title:       req.Title,
		Description: req.Description,
		Created_at:  sql.NullTime{Time: time.Now(), Valid: true},
	}
	return b.bookRepository.Save(ctx, &book)
}

// Update implements domain.BookService.
func (b *bookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
	persisted, err := b.bookRepository.FindById(ctx, req.Id)
	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return errors.New("book's data not found")
	}

	persisted.Isbn = req.Isbn
	persisted.Title = req.Title
	persisted.Description = req.Description
	persisted.Updated_at = sql.NullTime{Time: time.Now(), Valid: true}
	return b.bookRepository.Update(ctx, &persisted)
}

// Delete implements domain.BookService.
func (b *bookService) Delete(ctx context.Context, id string) error {
	persisted, err := b.bookRepository.FindById(ctx, id)

	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return errors.New("book's data not found")
	}
	err = b.bookRepository.Delete(ctx, persisted.Id)
	if err != nil {
		return err
	}
	return b.bookStockRepository.DeleteByBookId(ctx, persisted.Id)
}

// Show implements domain.BookService.
func (b *bookService) Show(ctx context.Context, id string) (dto.BookData, error) {
	persisted, err := b.bookRepository.FindById(ctx, id)

	if err != nil {
		return dto.BookData{}, err
	}

	if persisted.Id == "" {
		return dto.BookData{}, errors.New("book's data not found")
	}

	return dto.BookData{
		Id:    persisted.Id,
		Isbn:  persisted.Isbn,
		Title: persisted.Title,
	}, nil
}

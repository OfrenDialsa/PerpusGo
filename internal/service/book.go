package service

import (
	"PerpusGo/domain"
	"PerpusGo/dto"
	"PerpusGo/internal/config"
	"context"
	"database/sql"
	"errors"
	"path"
	"time"

	"github.com/google/uuid"
)

type bookService struct {
	cnf                 *config.Config
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
	mediaRepository     domain.MediaRepository
}

func NewBook(cnf *config.Config, bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository, mediaRepository domain.MediaRepository) domain.BookService {
	return &bookService{
		cnf:                 cnf,
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
		mediaRepository:     mediaRepository,
	}
}

// Index implements domain.BookService.
func (b *bookService) Index(ctx context.Context) ([]dto.BookData, error) {
	result, err := b.bookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	coverId := make([]string, 0)
	for _, v := range result {
		if v.CoverId.Valid {
			coverId = append(coverId, v.CoverId.String)
		}
	}
	covers := make(map[string]string)
	if len(coverId) > 0 {
		coversDb, _ := b.mediaRepository.FindByIds(ctx, coverId)
		for _, v := range coversDb {
			covers[v.Id] = path.Join(b.cnf.Server.Asset, v.Path)
		}
	}

	var bookData []dto.BookData
	for _, c := range result {

		var coverUrl string
		if v2, e := covers[c.CoverId.String]; e {
			coverUrl = v2
		}

		bookData = append(bookData, dto.BookData{
			Id:          c.Id,
			Isbn:        c.Isbn,
			Title:       c.Title,
			CoverUrl:    coverUrl,
			Description: c.Description,
		})
	}
	return bookData, nil
}

// Create implements domain.BookService.
func (b *bookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
	coverId := sql.NullString{Valid: false, String: req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}

	book := domain.Book{
		Id:          uuid.NewString(),
		Isbn:        req.Isbn,
		Title:       req.Title,
		Description: req.Description,
		CoverId:     coverId,
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

	coverId := sql.NullString{Valid: false, String: req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}

	persisted.Isbn = req.Isbn
	persisted.Title = req.Title
	persisted.Description = req.Description
	persisted.Updated_at = sql.NullTime{Time: time.Now(), Valid: true}
	persisted.CoverId = coverId
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
func (b *bookService) Show(ctx context.Context, id string) (dto.BookShowData, error) {
	persisted, err := b.bookRepository.FindById(ctx, id)

	if err != nil {
		return dto.BookShowData{}, err
	}

	if persisted.Id == "" {
		return dto.BookShowData{}, domain.ErrBookNotFound
	}

	stocks, err := b.bookStockRepository.FindByBookId(ctx, persisted.Id)
	if err != nil {
		return dto.BookShowData{}, err
	}
	stocksData := make([]dto.BookStockData, 0)
	for _, v := range stocks {
		stocksData = append(stocksData, dto.BookStockData{
			Code:   v.Code,
			Status: v.Status,
		})
	}
	var coverUrl string
	if persisted.CoverId.Valid {
		cover, _ := b.mediaRepository.FindById(ctx, persisted.CoverId.String)
		if cover.Path != "" {
			coverUrl = path.Join(b.cnf.Server.Asset, cover.Path)
		}

	}

	return dto.BookShowData{
		BookData: dto.BookData{
			Id:          persisted.Id,
			Isbn:        persisted.Isbn,
			Title:       persisted.Title,
			CoverUrl:    coverUrl,
			Description: persisted.Description,
		},
		Stocks: stocksData,
	}, nil
}

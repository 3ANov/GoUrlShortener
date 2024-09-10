package interfaces

import (
	"GoUrlShortener/internal/models"
	"context"
)

type ShortURLDao interface {
	Create(ctx context.Context, shortURL *models.ShortURL) error
	GetByID(ctx context.Context, id int) (*models.ShortURL, error)
	GetByShortCode(ctx context.Context, shortCode string) (*models.ShortURL, error)
	ShortCodeExists(ctx context.Context, shortCode string) (bool, error)
	IncrementUsage(ctx context.Context, shortCode string) error
	GetAll(ctx context.Context) ([]models.ShortURL, error)
}

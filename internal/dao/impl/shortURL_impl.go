package impl

import (
	"GoUrlShortener/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type ShortURLImpl struct {
	dbPool *pgxpool.Pool
}

func NewShortURLImpl(dbPool *pgxpool.Pool) *ShortURLImpl {
	return &ShortURLImpl{dbPool: dbPool}
}

func (dao *ShortURLImpl) Create(ctx context.Context, shortURL *models.ShortURL) error {
	var id int

	query := "INSERT INTO short_urls (original_url, short_code, expires_at) VALUES ($1, $2, $3) RETURNING id"
	err := dao.dbPool.QueryRow(ctx, query, shortURL.OriginalUrl, shortURL.ShortCode, shortURL.ExpiresAt).Scan(&id)

	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
		return err
	}

	shortURL.Id = id
	return nil
}

func (dao *ShortURLImpl) GetByID(ctx context.Context, id int) (*models.ShortURL, error) {
	query := "SELECT id, original_url, short_code, created_at, expires_at, usage_count FROM short_urls WHERE id=$1"
	var shortURL models.ShortURL
	err := dao.dbPool.QueryRow(ctx, query, id).Scan(
		&shortURL.Id,
		&shortURL.OriginalUrl,
		&shortURL.ShortCode,
		&shortURL.CreatedAt,
		&shortURL.ExpiresAt,
		&shortURL.UsageCount,
	)
	if err != nil {
		return nil, err
	}
	return &shortURL, nil
}

func (dao *ShortURLImpl) GetByShortCode(ctx context.Context, shortCode string) (*models.ShortURL, error) {
	query := "SELECT id, original_url, short_code, created_at, expires_at, usage_count FROM short_urls WHERE short_code=$1"
	var shortURL models.ShortURL
	err := dao.dbPool.QueryRow(ctx, query, shortCode).Scan(
		&shortURL.Id,
		&shortURL.OriginalUrl,
		&shortURL.ShortCode,
		&shortURL.CreatedAt,
		&shortURL.ExpiresAt,
		&shortURL.UsageCount,
	)
	if err != nil {
		return nil, err
	}
	return &shortURL, nil
}

func (dao *ShortURLImpl) ShortCodeExists(ctx context.Context, shortCode string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM short_urls WHERE short_code = $1)"

	err := dao.dbPool.QueryRow(ctx, query, shortCode).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (dao *ShortURLImpl) IncrementUsage(ctx context.Context, shortCode string) error {
	query := `UPDATE short_urls SET usage_count = usage_count + 1 WHERE short_code = $1`
	_, err := dao.dbPool.Exec(ctx, query, shortCode)
	return err
}

func (dao *ShortURLImpl) GetAll(ctx context.Context) ([]models.ShortURL, error) {
	query := "SELECT id, original_url, short_code, created_at, expires_at, usage_count FROM short_urls"
	shortURLS := make([]models.ShortURL, 0)

	rows, err := dao.dbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close() // Закрыть rows после использования

	var row models.ShortURL
	_, err = pgx.ForEachRow(rows, []any{&row.Id, &row.OriginalUrl, &row.ShortCode, &row.CreatedAt, &row.ExpiresAt, &row.UsageCount}, func() error {
		shortURLS = append(shortURLS, row)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return shortURLS, nil
}

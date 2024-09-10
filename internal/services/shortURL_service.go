package services

import (
	"GoUrlShortener/internal/dao/interfaces"
	"GoUrlShortener/internal/models"
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"
)

type ShortURLService struct {
	shortURLDao interfaces.ShortURLDao
}

func NewShortURLService(shortURLDao interfaces.ShortURLDao) *ShortURLService {
	return &ShortURLService{shortURLDao: shortURLDao}
}

func (s *ShortURLService) GetURLByShortCode(ctx context.Context, shortCode string) (*models.ShortURL, error) {
	shortURL, err := s.shortURLDao.GetByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	if time.Now().After(shortURL.ExpiresAt) {
		return nil, errors.New("short URL has expired")
	}

	if err := s.shortURLDao.IncrementUsage(ctx, shortCode); err != nil {
		return nil, err
	}

	return shortURL, nil
}

func (s *ShortURLService) GetAllUrls(ctx context.Context) ([]models.ShortURL, error) {
	return s.shortURLDao.GetAll(ctx)
}

func (s *ShortURLService) CreateShortURL(ctx context.Context, originalURL string, expiresAt time.Time) (*models.ShortURL, error) {
	shortCode, err := generateShortCode()
	if err != nil {
		return nil, err
	}

	exists, err := s.shortURLDao.ShortCodeExists(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	for exists {
		shortCode, err = generateShortCode()
		if err != nil {
			return nil, err
		}
		exists, err = s.shortURLDao.ShortCodeExists(ctx, shortCode)
		if err != nil {
			return nil, err
		}
	}

	shortURL := &models.ShortURL{
		OriginalUrl: originalURL,
		ShortCode:   shortCode,
		CreatedAt:   time.Now(),
		ExpiresAt:   expiresAt,
		UsageCount:  0,
	}

	err = s.shortURLDao.Create(ctx, shortURL)
	if err != nil {
		return nil, err
	}

	return shortURL, nil
}

func generateShortCode() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

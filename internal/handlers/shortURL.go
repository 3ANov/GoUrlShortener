package handlers

import (
	"GoUrlShortener/internal/models"
	"GoUrlShortener/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShortURLHandler struct {
	shortURLService *services.ShortURLService
}

func NewShortURLHandler(shortURLService *services.ShortURLService) *ShortURLHandler {
	return &ShortURLHandler{shortURLService: shortURLService}
}

func (h *ShortURLHandler) CreateShortURL(c *gin.Context) {
	var request models.ShortURLRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL, err := h.shortURLService.CreateShortURL(c.Request.Context(), request.OriginalUrl, request.ExpiresAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create short URL"})
		return
	}

	response := models.ShortURLResponse{
		ShortCode: shortURL.ShortCode,
		ExpiresAt: shortURL.ExpiresAt,
	}
	c.JSON(http.StatusOK, response)
}

func (h *ShortURLHandler) RedirectToOriginalURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	ctx := c.Request.Context()

	shortURL, err := h.shortURLService.GetURLByShortCode(ctx, shortCode)
	if err != nil {
		if err.Error() == "short URL has expired" {
			c.JSON(http.StatusGone, gin.H{"error": "Short URL has expired"})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}
	c.Redirect(http.StatusFound, shortURL.OriginalUrl)
}

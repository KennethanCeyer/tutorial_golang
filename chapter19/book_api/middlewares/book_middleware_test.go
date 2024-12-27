package middlewares


import (
	"book_api/models"
	"book_api/repositories"
	"net/http"
	"net/http/httptest"
	"testing"


	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBookLoader(t *testing.T) {
	mockBooks := []models.Book{
		{Model: gorm.Model{ID: 1}, Title: "테스트 책", Author: "테스트 저자", Year: 2023},
	}
	mockRepo := &repositories.MockBookRepository{MockBooks: mockBooks}

	r := gin.Default()
	r.GET("/books/:id", BookLoader(mockRepo), func(c *gin.Context) {
		book, exists := c.Get("book")
		assert.True(t, exists, "Context에 'book' 데이터가 저장되어야 합니다.")

		bookModel, ok := book.(models.Book)
		assert.True(t, ok, "Context의 'book' 데이터가 models.Book 타입이어야 합니다.")

		c.JSON(http.StatusOK, bookModel)
	})

	req := httptest.NewRequest("GET", "/books/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "테스트 책")
}


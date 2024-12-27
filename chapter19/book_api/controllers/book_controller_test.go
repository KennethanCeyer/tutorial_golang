package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"book_api/models"
	"book_api/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetBooks(t *testing.T) {
	// 1. Mock 데이터 정의
	mockBooks := []models.Book{
		{
			Model:  gorm.Model{ID: 1},
			Title:  "테스트 책 1",
			Author: "테스트 저자 1",
			Year:   2023,
		},
		{
			Model:  gorm.Model{ID: 2},
			Title:  "테스트 책 2",
			Author: "테스트 저자 2",
			Year:   2022,
		},
	}

	// 2. Mock Repository와 Controller 초기화
	mockRepo := &repositories.MockBookRepository{MockBooks: mockBooks}
	controller := &BookController{Repository: mockRepo}

	// 3. Gin Context Mocking
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// 4. 핸들러 실행
	controller.GetBooks(c)

	// 5. 결과 검증
	assert.Equal(t, http.StatusOK, w.Code)

	// 응답 데이터 비교
	var response []models.Book
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, len(mockBooks), len(response)) // 길이 비교

	for i, mockBook := range mockBooks {
		assert.Equal(t, mockBook.ID, response[i].ID, "ID가 일치하지 않습니다.")
		assert.Equal(t, mockBook.Title, response[i].Title, "Title이 일치하지 않습니다.")
		assert.Equal(t, mockBook.Author, response[i].Author, "Author가 일치하지 않습니다.")
		assert.Equal(t, mockBook.Year, response[i].Year, "Year가 일치하지 않습니다.")
	}
}

func TestDeleteBook(t *testing.T) {
    // Mock 데이터 및 Repository 생성
    mockBooks := []models.Book{
        {Model: gorm.Model{ID: 1}, Title: "테스트 책", Author: "테스트 저자", Year: 2023},
    }
    mockRepo := &repositories.MockBookRepository{MockBooks: mockBooks}
    bookController := &BookController{Repository: mockRepo}

    // Mock Gin Context 생성
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Mock Context에 필요한 데이터 설정
    c.Set("book", mockBooks[0]) // BookLoader 미들웨어의 동작을 대신 수행

    // 핸들러 실행
    bookController.DeleteBook(c)

    // 결과 검증
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "삭제되었습니다")
    assert.Equal(t, 0, len(mockRepo.MockBooks)) // MockBooks 리스트가 비었는지 확인
}

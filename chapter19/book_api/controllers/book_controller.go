package controllers


import (
	"book_api/models"
	"book_api/repositories"
	"net/http"


	"github.com/gin-gonic/gin"
)


type BookController struct {
	Repository repositories.BookRepository
}

func (bc *BookController) ShowIndexPage(c *gin.Context) {
	books, err := bc.Repository.FetchBooks()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "데이터를 불러올 수 없습니다."})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"Books": books})
}

func (bc *BookController) GetBooks(c *gin.Context) {
	books, err := bc.Repository.FetchBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "데이터를 가져오는 데 실패했습니다"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (bc *BookController) GetBookByID(c *gin.Context) {
	book, exists := c.Get("book")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "책 데이터를 로드하지 못했습니다"})
		return
	}

	bookModel, ok := book.(models.Book)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "잘못된 책 데이터 형식"})
		return
	}

	c.JSON(http.StatusOK, bookModel)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	book, exists := c.Get("book")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "책 데이터를 로드하지 못했습니다"})
		return
	}

	bookModel, ok := book.(models.Book)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "잘못된 책 데이터 형식"})
		return
	}

	if err := bc.Repository.DeleteBook(bookModel.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "책 삭제에 실패했습니다"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "책이 삭제되었습니다."})
}


package handlers

import (
	stdErrors "errors"
	"github.com/andrey-lawyer/go-gin-todo-app/errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleError(c *gin.Context, err error) {
	var validationError *errors.ValidationError
	var authError *errors.AuthError

	// Ошибка валидации
	if stdErrors.As(err, &validationError) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ошибка авторизации
	if stdErrors.As(err, &authError) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Ошибка MongoDB (или любая другая внутренняя)
	if err != nil && (errors.IsMongoError(err)) {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// Любая другая ошибка
	log.Printf("Internal error: %v", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}

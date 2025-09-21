package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "github.com/Mamvriyskiy/lab1-template/person/model"
	logger "github.com/Mamvriyskiy/lab1-template/person/logger"
)

func (h *Handler) GetInfoPerson(c *gin.Context) {
	person, err := h.GetInfoPerson()
	if err != nil {
		//TODO: logger, просмотреть код ошибки

		return
	}


	c.JSON(http.StatusOK, person)

	return
}

func (s *Handler) GetInfoPersons(c *gin.Context) {
	person, err := s.GetInfoPersons()

	return
}

func (s *Handler) CreateNewRecordPerson(c *gin.Context) {
	person, err := s.CreateNewRecordPerson()

	return
}

func (s *Handler) UpdateRecordPerson(c *gin.Context) {
	person, err := s.UpdateRecordPerson()

	return 
}

func (s *Handler) DeleteRecordPerson(c *gin.Context) {
	person, err := s.DeleteRecordPerson()

	return 
}

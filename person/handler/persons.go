package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/lab1-template/person/model"
	// logger "github.com/Mamvriyskiy/lab1-template/person/logger"
)

func getPersonID(c *gin.Context) (int, error) {
	id, ok := c.Get("personId")
	if !ok {
		//TODO: создать ошибку
		return 0, nil
	}

	personID, ok := id.(int)
	if !ok {
		//TODO: создать ошибку
		return 0, nil
	}

	return personID, nil
}

func (h *Handler) GetInfoPerson(c *gin.Context) {
	personID, err := getPersonID(c)
	if err != nil {
		//TODO: logger, просмотреть код ошибки
		// logger.Log("Warning", "Get", "userID is not a string", nil, "userID")
		return
	}

	person, err := h.services.GetInfoPerson(personID)
	if err != nil {
		//TODO: logger, просмотреть код ошибки

		return
	}


	c.JSON(http.StatusOK, person)

	return
}

func (s *Handler) GetInfoPersons(c *gin.Context) {
	persons, err := s.services.GetInfoPersons()
	if err != nil {
		//TODO: logger, просмотреть код ошибки
		return
	}

	c.JSON(http.StatusOK, persons)

	return
}

func (s *Handler) CreateNewRecordPerson(c *gin.Context) {
	var newPerson model.Person
	if err := c.BindJSON(&newPerson); err != nil {
		//TODO: logger, просмотреть код ошибки
		return
	}

	createPerson, err := s.services.CreateNewRecordPerson(newPerson)
	if err != nil {
		//TODO: logger, просмотреть код ошибки
		return
	}

	c.JSON(http.StatusOK, createPerson)

	return
}

func (s *Handler) UpdateRecordPerson(c *gin.Context) {
	var person model.Person
	if err := c.BindJSON(&person); err != nil {
		//TODO: logger, просмотреть код ошибки
		return
	}

	updatePerson, err := s.services.UpdateRecordPerson(person)
	if err != nil {
		//TODO: logger, просмотреть код ошибки
		return
	}

	c.JSON(http.StatusOK, updatePerson)

	return 
}

func (s *Handler) DeleteRecordPerson(c *gin.Context) {
	personID, err := getPersonID(c)
	if err != nil {
		//TODO: logger, просмотреть код ошибки
		// logger.Log("Warning", "Get", "userID is not a string", nil, "userID")
		return
	}

	err = s.services.DeleteRecordPerson(personID)
	if err != nil {
		//TODO: logger, просмотреть код ошибки
		return
	}

	return 
}

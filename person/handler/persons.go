package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	logger "github.com/Mamvriyskiy/lab1-template/person/logger"
	"github.com/Mamvriyskiy/lab1-template/person/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetPersonID(c *gin.Context) (int, error) {
	personId := c.Param("personId")
    if personId == "" {
        return 0, errors.New("personId отсутствует в параметрах пути")
    }
    
    id, err := strconv.Atoi(personId)
    if err != nil {
        return 0, errors.New("значение ключа personId имеет некорректный тип, ожидался int")
    }
    
    return id, nil
}

func (h *Handler) GetInfoPerson(c *gin.Context) {
	personID, err := GetPersonID(c)
	if err != nil {
		logger.Error("", zap.Error(err))
		c.Status(http.StatusBadRequest)
		return
	}

	person, err := h.services.GetInfoPerson(personID)
	if err != nil {
		c.Status(http.StatusNotFound)
		logger.Error("", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, person)
}

func (s *Handler) GetInfoPersons(c *gin.Context) {
	persons, err := s.services.GetInfoPersons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, persons)
}

func (s *Handler) CreateNewRecordPerson(c *gin.Context) {
	var newPerson model.Person
	if err := c.BindJSON(&newPerson); err != nil {
		logger.Error("не удалось распарсить тело запроса в структуру Person",
        	zap.Error(err),
    	)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createPerson, err := s.services.CreateNewRecordPerson(newPerson)
	if err != nil {
		logger.Error("", zap.Error(err))
		return
	}

	c.Header("Location", fmt.Sprintf("/api/v1/persons/%d", createPerson.PersonID))
	c.JSON(http.StatusCreated, createPerson)
}

func (s *Handler) UpdateRecordPerson(c *gin.Context) {
	personID, err := GetPersonID(c)
	if err != nil {
		logger.Error("", zap.Error(err))
		c.Status(http.StatusBadRequest)
		return
	}

	var person model.Person
	if err := c.BindJSON(&person); err != nil {
		logger.Error("не удалось распарсить тело запроса в структуру Person",
        	zap.Error(err),
    	)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person.PersonID = personID

	updatePerson, err := s.services.UpdateRecordPerson(person)
	if err != nil {
		logger.Error("", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, updatePerson)
}

func (s *Handler) DeleteRecordPerson(c *gin.Context) {
	personID, err := GetPersonID(c)
	if err != nil {
		logger.Error("", zap.Error(err))
		c.Status(http.StatusBadRequest)
		return
	}

	err = s.services.DeleteRecordPerson(personID)
	if err != nil {
		logger.Error("", zap.Error(err))
		return
	}

	c.Status(http.StatusNoContent)
}

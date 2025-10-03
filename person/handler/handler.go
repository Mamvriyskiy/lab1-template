package handler

import (
	"github.com/Mamvriyskiy/lab1-template/person/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	persons := router.Group("api/v1/persons")

	// Инофрмация о человека
	persons.GET("/:personId", h.GetInfoPerson)

	// Информация по всем людям
	persons.GET("", h.GetInfoPersons)

	// Создание новой записи о человеке
	persons.POST("", h.CreateNewRecordPerson)

	// Обновление существующей записи о человеке
	persons.PATCH("/:personId", h.UpdateRecordPerson)

	// Удаление записи о человеке
	persons.DELETE("/:personId", h.DeleteRecordPerson)

	return router
}
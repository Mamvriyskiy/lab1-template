package handler

import (
	"github.com/Mamvriyskiy/database_course/main/pkg/service"
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

	persons := router.Group("persons")

	// Инофрмация о человека
	persons.Get("/:personId", h.GetInfoPerson)

	// Информация по всем людям
	persons.Get("", h.GetInfoPersons)

	// Создание новой записи о человеке
	persons.Post("", h.CreateNewRecordPerson)

	// Обновление существующей записи о человеке
	persons.PATCH("/:personId", h.UpdateRecordPerson)

	// Удаление записи о человеке
	persons.DELETE("/:personId", h.DeleteRecordPerson)

	return router
}
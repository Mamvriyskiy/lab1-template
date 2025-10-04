package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPersonID_NotFound(t *testing.T) {
	c := &gin.Context{}

	c.Params = gin.Params{}

	id, err := GetPersonID(c)

	assert.Equal(t, 0, id)
	assert.EqualError(t, err, "personId отсутствует в параметрах пути")
}

func TestGetPersonID_WrongType(t *testing.T) {
	c := &gin.Context{}

	c.Params = gin.Params{gin.Param{Key: "personId", Value: "not-an-int"}}

	id, err := GetPersonID(c)

	assert.Equal(t, 0, id)
	assert.EqualError(t, err, "значение ключа personId имеет некорректный тип, ожидался int")
}

func TestGetPersonID_Success(t *testing.T) {
	c := &gin.Context{}
	
	c.Params = gin.Params{gin.Param{Key: "personId", Value: "42"}}

	id, err := GetPersonID(c)

	assert.Equal(t, 42, id, "id должен совпадать с сохранённым")
	assert.NoError(t, err)
}

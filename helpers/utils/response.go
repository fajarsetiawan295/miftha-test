package util

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON returns a well formated response with a status code

type StructData struct {
	PerPage       int         `json:"perPage"`
	Total         int         `json:"total"`
	TotalFiltered int         `json:"totalFiltered"`
	List          interface{} `json:"list"`
}

func JSON(c *gin.Context, message string, data interface{}) {
	var hasil, dataList interface{}
	hasil = data

	Str := StructData{}
	iniData, _ := json.Marshal(data)
	json.Unmarshal(iniData, &Str)

	if Str.PerPage != 0 {
		dataList = Str.List
		if Str.List == nil {
			dataList = []string{}
		}
		hasil = StructData{
			PerPage:       Str.PerPage,
			Total:         Str.Total,
			TotalFiltered: Str.TotalFiltered,
			List:          dataList,
		}
	}

	c.JSON(http.StatusOK, struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  true,
		Message: message,
		Data:    hasil,
	})
	return
}

// ERROR returns a jsonified error response along with a status code.
func ERROR(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Status  bool        `json:"status"`
			Message string      `json:"message"`
			Data    interface{} `json:"data"`
		}{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusBadRequest, nil)
}
func ERRORDATA(c *gin.Context, err error, data interface{}) {

	var hasil interface{}
	hasil = data
	if data == nil {
		hasil = []string{}
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Status  bool        `json:"status"`
			Message string      `json:"message"`
			Data    interface{} `json:"data"`
		}{
			Status:  false,
			Message: err.Error(),
			Data:    hasil,
		})
		return
	}
	c.JSON(http.StatusBadRequest, nil)
}

func FAILED(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  false,
		Message: message,
		Data:    nil,
	})
	return
}

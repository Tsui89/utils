package response

import (
	"reflect"

	"gitee.com/k2tf/utils/page"
	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewBaseResponse() *BaseResponse {
	b := BaseResponse{
		0,
		"success",
	}
	return &b
}
func (br *BaseResponse) Set(code int, message string) {
	br.Code = code
	br.Message = message
}

type ListResponse struct {
	BaseResponse
	PageInfo page.PageInfo `json:"page_info"`
	Data     interface{}   `json:"data"`
}

type ListNoneResponse struct {
	BaseResponse
	PageInfo page.PageInfo `json:"page_info"`
	Data     []interface{} `json:"data"`
}

type ListNoneWithoutPageResponse struct {
	BaseResponse
	Data     []interface{} `json:"data"`
}

type ListWithoutPageResponse struct {
	BaseResponse
	Data     interface{} `json:"data"`
}


type DataResponse struct {
	BaseResponse
	Data interface{} `json:"data"`
}

func ResponseList(c *gin.Context, data interface{}, info page.PageInfo, br BaseResponse) {

	if isNotNull(data) {
		c.JSON(0, ListResponse{
			br,
			info,
			data,
		})
	} else {
		c.JSON(0, ListNoneResponse{
			br,
			info,
			[]interface{}{},
		})
	}
}

func ResponseListWithotPage(c *gin.Context, data interface{}, br BaseResponse) {

	if isNotNull(data) {
		c.JSON(0, ListWithoutPageResponse{
			br,
			data,
		})
	} else {
		c.JSON(0, ListNoneWithoutPageResponse{
			br,
			[]interface{}{},
		})
	}
}

func ResponseData(c *gin.Context, data interface{}, br BaseResponse) {

	if data == nil {
		data = map[string]interface{}{}
	}
	if isNotNull(data) {
		c.JSON(0, DataResponse{
			br,
			data,
		})
	} else {
		c.JSON(0, DataResponse{
			br,
			map[string]interface{}{},
		})
	}
}

func isNotNull(i interface{}) bool {
	if i == nil {
		return false
	}
	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.Slice:
		if v.Len() > 0 {
			return true
		}
	case reflect.Map:
		if len(v.MapKeys()) > 0 {
			return true
		}
	default:
		return true
	}
	return false
}

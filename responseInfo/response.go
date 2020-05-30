package responseInfo

import (
	"reflect"

	page "github.com/Tsui89/utils/pageInfo"
	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	HttpCode int    `json:"-"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Info     string `json:"info"`
}

func NewBaseResponse() *BaseResponse {
	b := BaseResponse{
		200,
		0,
		"success",
		"成功",
	}
	return &b
}
func (br *BaseResponse) Set(code int, message, info string) {
	br.Code = code
	br.Message = message
	br.Info = info
}

type ListResponse struct {
	BaseResponse
	PageInfo page.PageInfo `json:"pageInfo"`
	Data     interface{}   `json:"data"`
}

type ListNoneResponse struct {
	BaseResponse
	PageInfo page.PageInfo `json:"pageInfo"`
	Data     []interface{} `json:"data"`
}

type ListNoneWithoutPageResponse struct {
	BaseResponse
	Data []interface{} `json:"data"`
}

type ListWithoutPageResponse struct {
	BaseResponse
	Data interface{} `json:"data"`
}

type DataResponse struct {
	BaseResponse
	Data interface{} `json:"data"`
}

func ResponseList(c *gin.Context, data interface{}, info page.PageInfo, br BaseResponse) {

	if isNotNull(data) {
		c.JSON(200, ListResponse{
			br,
			info,
			data,
		})
	} else {
		c.JSON(200, ListNoneResponse{
			br,
			info,
			[]interface{}{},
		})
	}
}

func ResponseListWithotPage(c *gin.Context, data interface{}, br BaseResponse) {

	if isNotNull(data) {
		c.JSON(200, ListWithoutPageResponse{
			br,
			data,
		})
	} else {
		c.JSON(200, ListNoneWithoutPageResponse{
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
		c.JSON(200, DataResponse{
			br,
			data,
		})
	} else {
		c.JSON(200, DataResponse{
			br,
			map[string]interface{}{},
		})
	}
}
func ResponseWithoutData(c *gin.Context, br BaseResponse) {
	c.JSON(200, br)
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

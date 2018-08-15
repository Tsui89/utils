package pageInfo

import (
	"errors"
	"math"
	"strconv"

	"net/http"
)

type PageInfo struct {
	Size     int `json:"size"`     //当前页大小
	Total    int `json:"total"`    //总大小
	Pages    int `json:"pages"`    //总页数
	PageNum  int `json:"pageNum"`  //当前页
	PageSize int `json:"pageSize"` //每页多少
}

func NewPageInfo(req *http.Request) *PageInfo {
	perPage, err := strconv.Atoi(req.FormValue("pageSize"))
	if err != nil || perPage == 0 {
		perPage = 20
	} else if perPage < 0 {
		perPage = -1
	}
	page, err := strconv.Atoi(req.FormValue("pageNum"))
	if err != nil || page <= 0 {
		page = 1
	}
	pageInfo := new(PageInfo)
	pageInfo.SetPage(page)
	pageInfo.SetPerPage(perPage)
	return pageInfo
}

func (p *PageInfo) SetTotalPage() {
	p.Pages = int(math.Ceil(float64(p.Total) / float64(p.PageSize)))
	if p.Pages <= 0 {
		p.Pages = 1
	}
}

func (p *PageInfo) PageCheck() error {
	if p.Total == 0 || p.PageNum > p.Pages {
		return errors.New("has no data")
	}
	return nil
}

func (p *PageInfo) SetTotalSize(totalSize int) {
	p.Total = totalSize
}

func (p *PageInfo) SetSize(size int) {
	p.Size = size
}

func (p *PageInfo) SetPage(page int) {
	p.PageNum = page
}

func (p *PageInfo) SetPerPage(perPage int) {
	p.PageSize = perPage
}

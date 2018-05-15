package page

import (
	"errors"
	"math"
	"strconv"

	"net/http"
)

type PageInfo struct {
	Size      int `json:"size"`       //当前页大小
	TotalSize int `json:"total_size"` //总大小
	TotalPage int `json:"total_page"` //总页数
	Page      int `json:"page"`       //当前页
	PerPage   int `json:"per_page"`   //每页多少
}

func NewPageInfo(req *http.Request) *PageInfo {
	perPage, err := strconv.Atoi(req.FormValue("per_page"))
	if err != nil || perPage == 0 {
		perPage = 20
	} else if perPage < 0 {
		perPage = -1
	}
	page, err := strconv.Atoi(req.FormValue("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	pageInfo := new(PageInfo)
	pageInfo.SetPage(page)
	pageInfo.SetPerPage(perPage)
	return pageInfo
}

func (p *PageInfo) SetTotalPage() {
	p.TotalPage = int(math.Ceil(float64(p.TotalSize) / float64(p.PerPage)))
	if p.TotalPage <=0{
		p.TotalPage = 1
	}
}

func (p *PageInfo) PageCheck() error {
	if p.TotalSize == 0 || p.Page > p.TotalPage {
		return errors.New("has no data")
	}
	return nil
}

func (p *PageInfo) SetTotalSize(totalSize int) {
	p.TotalSize = totalSize
}

func (p *PageInfo) SetSize(size int) {
	p.Size = size
}

func (p *PageInfo) SetPage(page int) {
	p.Page = page
}

func (p *PageInfo) SetPerPage(perPage int) {
	p.PerPage = perPage
}

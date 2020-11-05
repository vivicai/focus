package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Index = new(indexApi)

type indexApi struct{}

// @summary 展示站点首页
// @tags    首页
// @produce html
// @router  / [GET]
// @success 200 {string} html "页面HTML"
func (a *indexApi) Index(r *ghttp.Request) {
	r.Response.WriteTpl("web/layout/layout.html", g.Map{
		"mainTpl": "web/index.html",
		"title":   "gf bbs - 首页",
	})
}

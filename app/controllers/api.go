package controllers

import (
	"altria/app/models"
	"strconv"

	"github.com/revel/revel"
)

type Api struct {
	*revel.Controller
}

// 分页规矩：o：offset，第几页；p：perpage，每页返回个数
func (c Api) Favorites() revel.Result {
	o, _ := strconv.Atoi(c.Params.Query.Get("o"))
	p, _ := strconv.Atoi(c.Params.Query.Get("p"))

	if p == 0 || p > 100 {
		p = 20
	}

	results := models.FindList(o, p)
	// 读取数据库中的列表，将所有资料混在一起
	// 要分页
	// fmt.Println(o, p)
	res := map[string]interface{}{"data": results, "count": len(results)}
	return c.RenderJSON(res)
}

func (c Api) FavoritesOnChannel() revel.Result {
	o, _ := strconv.Atoi(c.Params.Query.Get("o"))
	p, _ := strconv.Atoi(c.Params.Query.Get("p"))
	channel := c.Params.Route.Get("channel")

	if p == 0 || p > 100 {
		p = 20
	}

	// 读取特定分区的热门视频，得出列表
	// 要分页
	// fmt.Println(o, p, channel)
	results := models.FindListWithChannel(o, p, channel)
	res := map[string]interface{}{"data": results, "count": len(results)}

	return c.RenderJSON(res)
}

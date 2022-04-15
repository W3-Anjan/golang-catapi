package routers

import (
	"golang-catapi/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:BreedSearch")
	beego.Router("/id", &controllers.BreedListController{}, "post:FavBreedById")
	beego.Router("/list", &controllers.BreedListController{}, "get:GetFavBreeds")
}

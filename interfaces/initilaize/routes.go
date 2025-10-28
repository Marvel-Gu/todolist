package initilaize

import (
	"github.com/gin-gonic/gin"

	middleware "github.com/CocaineCong/todolist-ddd/interfaces/midddleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1/")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
	}
	return r
}

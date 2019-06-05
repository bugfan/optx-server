package option

import (
	"errors"
	"net/http"
	"optx-server/models"
	"time"

	"github.com/bugfan/to"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	opt := models.Options{}
	err := ctx.Bind(&opt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	if opt.Question == "" {
		ctx.JSON(http.StatusBadRequest, errors.New("问题为空"))
		return
	}
	opt.Created = time.Now()
	if _, err := models.GetEngine().Insert(&opt); err != nil {
		ctx.JSON(http.StatusBadRequest, opt)
		return
	}
	ctx.JSON(http.StatusCreated, opt.Question)
}

func All(ctx *gin.Context) {
	all := make([]*models.Options, 0, 1)
	models.All(&all)
	ctx.JSON(http.StatusOK, all)
}

func Delete(ctx *gin.Context) {
	id := to.Int64(ctx.Param("id"))
	if id < 1 {
		ctx.JSON(http.StatusBadRequest, errors.New("ID不正确"))
		return
	}
	var o models.Options
	_, err := models.GetEngine().Id(id).Delete(&o)
	ctx.JSON(http.StatusOK, err)
}

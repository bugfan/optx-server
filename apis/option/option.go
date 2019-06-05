package option

import (
	"errors"
	"net/http"
	"optx-server/models"

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
	if _, err := models.GetEngine().Insert(&opt); err != nil {
		ctx.JSON(http.StatusBadRequest, opt)
		return
	}
	ctx.JSON(http.StatusOK, opt.Question)
}

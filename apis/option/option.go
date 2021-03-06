package option

import (
	"errors"
	"net/http"
	"optx-server/models"
	"strconv"
	"time"

	"github.com/bugfan/to"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	kind := ctx.DefaultQuery("kind", "0")
	opt := models.Options{
		Kind: to.Int64(kind),
	}
	err := ctx.Bind(&opt)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if opt.Question == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("问题为空"))
		return
	}
	if opt.Answer < 1 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("答案格式不正确"))
		return
	}
	opt.Created = time.Now()
	if _, err := models.GetEngine().Insert(&opt); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, opt.ID)
}

func All(ctx *gin.Context) {
	kind := ctx.DefaultQuery("kind", "0")
	all := make([]*models.Options, 0, 0)
	models.GetEngine().Where("kind = ?", to.Int64(kind)).Find(&all)
	ctx.JSON(http.StatusOK, all)
}

func Delete(ctx *gin.Context) {
	id := to.Int64(ctx.Param("id"))
	if id < 1 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("ID不正确"))
		return
	}
	var o models.Options
	_, err := models.GetEngine().Id(id).Delete(&o)
	ctx.JSON(http.StatusOK, err)
}

func Update(ctx *gin.Context) {
	id := to.Int64(ctx.Param("id"))
	if id < 1 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("ID不正确"))
		return
	}
	opt := models.Options{}
	err := ctx.Bind(&opt)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if opt.Question == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("问题为空"))
		return
	}
	if opt.Answer < 1 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("答案格式不正确"))
		return
	}
	_, err = models.GetEngine().Id(id).AllCols().Update(&opt)
	ctx.JSON(http.StatusOK, id)
}

// 批量导入
func CreateLot(ctx *gin.Context) {
	kind := ctx.DefaultQuery("kind", "0")
	opts := make([]*models.Options, 0, 0)
	err := ctx.Bind(&opts)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	for _, v := range opts {
		v.Kind = to.Int64(kind)
	}
	_, err = models.GetEngine().Insert(opts)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, err)
}

/// 获取一些题目

func GetSMS(ctx *gin.Context) {
	if ctx.Query("au") != "XNlcm5hbW" {
		ctx.AbortWithError(403, errors.New(""))
		return
	}

	c, _ := strconv.Atoi(ctx.Query("count"))
	if c < 1 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("数据太少"))
		return
	}
	list, err := models.GetCountOptions(c)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func GetSMSTemp(ctx *gin.Context) {
	if ctx.Query("au") != "XNlcm5hbW" {
		ctx.AbortWithError(403, errors.New(""))
		return
	}

	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if limit < 1 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("数据太少"))
		return
	}
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	list, err := models.GetOptionsPages(offset, limit)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, list)
}

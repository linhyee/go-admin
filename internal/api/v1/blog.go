package v1

import (
	"blog/global"
	"blog/internal/model"
	"blog/pkg/app"
	"blog/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type BlogApi struct{}

func NewBlogApi() BlogApi {
	return BlogApi{}
}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "文章标题" maxlength(100)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Blog "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Router /api/v1/blog [get]
func (b BlogApi) List(c *gin.Context) {
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	}{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error : %v", errs)
		response.ToErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(gin.H{})
}

func (b BlogApi) Get(c *gin.Context) {
	var blog model.Blog
	global.DB().Where("id=?", 1).First(&blog)
	app.NewResponse(c).ToResponse(&blog)
}

func (b BlogApi) Create(c *gin.Context) {

}

func (b BlogApi) Update(c *gin.Context) {

}

func (b BlogApi) Delete(c *gin.Context) {

}

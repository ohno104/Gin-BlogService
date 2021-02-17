package v1

import (
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/app"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summery 獲取單個文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.Success)
	return
}

func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}

package v1

import (
	"felix.bs.com/felix/BeStrongerInGO/02_GinBlog/pkg/app"
	"felix.bs.com/felix/BeStrongerInGO/02_GinBlog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.Success)
	return
}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}

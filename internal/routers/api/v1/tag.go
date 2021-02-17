package v1

import (
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/global"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/internal/service"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/app"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/convert"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary 取得多個標籤
// @Produce json
// @Param name query string false "標籤名稱" maxlength(100)
// @Param state query int false "狀態" Enums(0, 1) default(1)
// @Param page query int false "頁碼"
// @Param page_size query int false "每頁數量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}

	totalRows, err := svc.CountTag(&service.CountTagRequest{
		Name:  param.Name,
		State: param.State,
	})
	if err != nil {
		global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
	return
}

// @Summary 新增標籤
// @Produce json
// @Param name body string true "標籤名稱" minlength(3) maxlength(100)
// @Param state body int false "狀態" Enums(0, 1) default(1)
// @Param create_by body string true "建立者" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新標籤
// @Produce json
// @Param id path int true "標籤ID"
// @Param name body string false "標籤名稱" minlength(3) maxlength(100)
// @Param state body int false "狀態" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 刪除標籤
// @Produce json
// @Param id path int true "標籤ID"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags/{id} [delete]]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

package app

import (
	"github.com/DATOULIN/dtservice/internal/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

type ResultList struct {
	Pager Pager       `json:"pager"`
	List  interface{} `json:"list"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// ToErrorResponse 异常响应，状态码根据自己的code去返回
func (r *Response) ToErrorResponse(err *errno.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg(), "result": nil}
	details := err.Details()
	if len(details) > 0 {
		response["msg"] = details
	}
	//logger.LogInfo("details", details)
	r.Ctx.JSON(err.StatusCode(), response)
}

// ToSuccessResponse 成功的响应，只返回message
func (r *Response) ToSuccessResponse(err *errno.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg(), "result": nil}
	r.Ctx.JSON(http.StatusOK, response)
}

// ToResponseList 成功的响应返回分页信息
func (r *Response) ToResponseList(list interface{}, totalRows int, err *errno.Error) {
	pager := Pager{
		Page:      GetPage(r.Ctx),
		PageSize:  GetPageSize(r.Ctx),
		TotalRows: totalRows,
	}
	r.Ctx.JSON(http.StatusOK, gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
		"result": ResultList{
			List:  list,
			Pager: pager,
		},
	})
}

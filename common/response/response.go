package response

import (
	"Gopan/common/errorx"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

type Body struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// http返回
func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		//成功返回
		r := &Body{
			Code:    0,
			Message: "success",
			Data:    resp,
		}
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := uint32(500)
		errmsg := "服务器错误"

		causeErr := errors.Cause(err)                  // err类型
		if e, ok := causeErr.(*errorx.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.Code
			errmsg = e.Msg
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				errcode = grpcCode
				errmsg = gstatus.Message()
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		httpx.WriteJson(w, http.StatusBadRequest, &Body{
			Code:    errcode,
			Message: errmsg,
			Data:    nil,
		})
	}
}

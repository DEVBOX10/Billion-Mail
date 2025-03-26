package middlewares

import (
	"billionmail-core/internal/service/public"
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/util/gvalid"
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
)

func HandleApiResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// If the request has been exited, do nothing.
	if r.Response.BufferLength() > 0 {
		return
	}

	// Handle the error if it exists.
	if r.GetError() != nil {
		// Clear the response buffer.
		r.Response.ClearBuffer()

		// Catch validation errors and respond with 412.
		if v, ok := r.GetError().(gvalid.Error); ok {
			errorStr := v.Error()
			errorLen := len(errorStr)
			// If the error message is in the format Lang{xxx}, convert it to the corresponding language.
			if strings.Contains(errorStr, "Lang{") && errorStr[errorLen-1] == '}' {
				errorStr = errorStr[5 : errorLen-1]
				errorStr = public.LangCtx(r.Context(), errorStr)
			}

			r.Response.WriteJson(api_v1.StandardRes{
				Success: false,
				Code:    412,
				Msg:     errorStr,
			})

			return
		}

		// Respond with 500 error.
		r.Response.WriteJson(api_v1.StandardRes{
			Success: false,
			Code:    500,
			Msg:     r.GetError().Error(),
		})

		return
	}

	// Get the response data.
	resp := r.GetHandlerResponse()

	// Respond with the data if it exists.
	if resp != nil {
		r.Response.WriteJson(resp)
	}
}

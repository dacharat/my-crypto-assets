package testutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func NewDefaultContext() (res *httptest.ResponseRecorder, context *gin.Context) {
	return NewWithRequestContext(http.MethodGet, "/", JSON(nil))
}

func NewWithRequestContext(method, url string, body io.Reader, headers ...map[string]string) (res *httptest.ResponseRecorder, context *gin.Context) {
	r := httptest.NewRequest(method, url, body)

	if len(headers) > 0 {
		for k, v := range headers[0] {
			r.Header.Set(k, v)
		}
	}

	res = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(res)
	context.Request = r
	return res, context
}

func JSON(v interface{}) io.Reader {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		panic(fmt.Errorf("JSON encode error: %v", err))
	}
	return &buf
}

package golanghttproutes

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type LogMidleWare struct {
	http.Handler
}

func (midleware *LogMidleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Receive request"))
	midleware.Handler.ServeHTTP(w, r)
}

func TestMidlewareRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()
	midlerware := LogMidleWare{
		Handler: router,
	}

	midlerware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello World", string(body))
}

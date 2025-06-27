package handler_test

import (
	"backend/core"
	"backend/presentation/middleware"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
)

// createTestEcho はテスト用のEchoインスタンスを作成します
func createTestEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Error(core.NewLogger("dev")))
	return e
}

// createTestRequest はテスト用のHTTPリクエストとEchoコンテキストを作成します
func createTestRequest(method, path, body string) (*httptest.ResponseRecorder, echo.Context) {
	e := createTestEcho()

	var req *http.Request
	if body != "" {
		req = httptest.NewRequest(method, path, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	} else {
		req = httptest.NewRequest(method, path, nil)
	}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return rec, c
}

// createTestPostRequest はPOSTリクエスト用のヘルパー関数です
func createTestPostRequest(path, body string) (*httptest.ResponseRecorder, echo.Context) {
	return createTestRequest(http.MethodPost, path, body)
}

// createTestGetRequest はGETリクエスト用のヘルパー関数です
func createTestGetRequest(path string) (*httptest.ResponseRecorder, echo.Context) {
	return createTestRequest(http.MethodGet, path, "")
}

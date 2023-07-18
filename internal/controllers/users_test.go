package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAddUser(t *testing.T) {
	t.Run("Add User Successfully", func(t *testing.T) {
		// Arrange
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)

		context.Request = &http.Request{
			Header: make(http.Header),
		}

		context.Request.Method = "POST"
		context.Request.Header.Set("Content-Type", "application/json")

		content := map[string]interface{}{"Name": "User", "Role": "Admin", "Description": "Master Admin User"}

		jsonbytes, err := json.Marshal(content)
		if err != nil {
			panic(err)
		}

		context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	})
}

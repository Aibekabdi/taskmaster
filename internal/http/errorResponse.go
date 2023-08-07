package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type errorResponseStruct struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func errorResponse(c *gin.Context, status int, msg string) {
	fmt.Printf("%s %s [%s]\t%s%s - %d - %s\n", time.Now().Format("2006/01/02 15:04:05"), c.Request.Proto, c.Request.Method, c.Request.Host, c.Request.RequestURI, status, http.StatusText(status))
	fmt.Println(msg)

	c.AbortWithStatusJSON(status, errorResponseStruct{Status: status, Msg: msg})
}

package main

import (
	"errors"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	e := gin.Default()
	e.POST("/", convert)

	e.Run(":2247")
}

type request struct {
	Data string `json:"data"`
	Type string `json:"type"` // bplustree
	Out  string `json:"out"`  // svg, dot
}

func convert(c *gin.Context) {
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	out, err := dot(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if req.Out == "dot" {
		c.JSON(http.StatusOK, gin.H{
			"data": string(out),
		})
	} else {
		c.Data(http.StatusOK, "image/svg+xml", out)
	}
}

func dot(req *request) ([]byte, error) {
	var w strings.Builder
	g := G[req.Type](&w)
	g.Generate(req.Data)
	if req.Out == "dot" {
		return []byte(w.String()), nil
	}
	dotfile, err := os.CreateTemp("", "tmpfile-*.dot")
	if err != nil {
		return nil, errors.New("create tmp file error")
	}
	defer dotfile.Close()
	defer os.Remove(dotfile.Name())

	dotfile.WriteString(w.String())

	cmd := exec.Command("dot", "-Tsvg", dotfile.Name())
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return out, nil
}

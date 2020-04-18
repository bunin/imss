//go:generate esc -o data.go -private -pkg=ui public

package ui

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	prefix    = "/public"
	indexFile = "/index.html"
	apiPrefix = "/api"
)

func Handle(ctx *gin.Context) {
	if strings.HasPrefix(ctx.Request.URL.Path, apiPrefix) {
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New("not found"))
		return
	}
	if ctx.Request.Method != http.MethodGet && ctx.Request.Method != http.MethodHead {
		ctx.Header("Allow", http.MethodGet+", "+http.MethodHead)
		http.Error(ctx.Writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	r := ctx.Request.URL.Path
	if r == "/" || r == "" {
		r = indexFile
	}
	if filepath.Ext(r) == "" {
		r = indexFile
	}
	file, err := _escStatic.Open(prefix + r)
	if err == os.ErrNotExist || file == nil {
		http.NotFound(ctx.Writer, ctx.Request)
		ctx.Abort()
		return
	}
	fi, err := file.Stat()
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		ctx.Abort()
		return
	}
	modTime := fi.ModTime()
	if r == indexFile {
		modTime = time.Now()
	}
	http.ServeContent(ctx.Writer, ctx.Request, fi.Name(), modTime, file)
}

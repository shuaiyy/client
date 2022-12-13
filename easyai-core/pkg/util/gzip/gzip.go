package gzip

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

/**
forked from gin-contrib/gzip
*/

const (
	// BestCompression ...
	BestCompression = gzip.BestCompression
	// BestSpeed ...
	BestSpeed = gzip.BestSpeed
	// DefaultCompression ...
	DefaultCompression = gzip.DefaultCompression
	// NoCompression ...
	NoCompression = gzip.NoCompression
)

// Gzip  ...
func Gzip(level int, options ...Option) gin.HandlerFunc {
	return newGzipHandler(level, options...).Handle
}

type gzipWriter struct {
	gin.ResponseWriter
	writer *gzip.Writer
}

// WriteString  ...
func (g *gzipWriter) WriteString(s string) (int, error) {
	return g.writer.Write([]byte(s))
}

// Write ...
func (g *gzipWriter) Write(data []byte) (int, error) {
	return g.writer.Write(data)
}

// WriteHeader Fix: https://github.com/mholt/caddy/issues/38
func (g *gzipWriter) WriteHeader(code int) {
	g.Header().Del("Content-Length")
	g.ResponseWriter.WriteHeader(code)
}

// handler
type gzipHandler struct {
	*Options
	gzPool *sync.Pool
}

func newGzipHandler(level int, options ...Option) *gzipHandler {
	var gzPool sync.Pool
	gzPool.New = func() interface{} {
		gz, err := gzip.NewWriterLevel(io.Discard, level)
		if err != nil {
			panic(err)
		}
		return gz
	}
	handler := &gzipHandler{
		Options: DefaultOptions,
		gzPool:  &gzPool,
	}
	for _, setter := range options {
		setter(handler.Options)
	}
	return handler
}

// Handle  ...
func (g *gzipHandler) Handle(c *gin.Context) {
	if fn := g.DecompressFn; fn != nil && c.Request.Header.Get("Content-Encoding") == "gzip" {
		fn(c)
	}

	if !g.shouldCompress(c.Request) {
		return
	}

	gz := g.gzPool.Get().(*gzip.Writer)
	defer g.gzPool.Put(gz)
	defer gz.Reset(io.Discard)
	gz.Reset(c.Writer)

	c.Header("Content-Encoding", "gzip")
	c.Header("Vary", "Accept-Encoding")
	c.Writer = &gzipWriter{c.Writer, gz}
	defer func() {
		gz.Close()
		c.Header("Content-Length", fmt.Sprint(c.Writer.Size()))
	}()
	c.Next()
}

func (g *gzipHandler) shouldCompress(req *http.Request) bool {
	if !strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") ||
		strings.Contains(req.Header.Get("Connection"), "Upgrade") ||
		strings.Contains(req.Header.Get("Content-Type"), "text/event-stream") {

		return false
	}

	extension := filepath.Ext(req.URL.Path)
	if g.ExcludedExtensions.Contains(extension) {
		return false
	}

	return !g.ExcludedPaths.Contains(req.URL.Path)
}

// option

var (
	// DefaultExcludedExtentions  ...
	DefaultExcludedExtentions = NewExcludedExtensions([]string{
		".png", ".gif", ".jpeg", ".jpg",
	})
	// DefaultOptions  ...
	DefaultOptions = &Options{
		ExcludedExtensions: DefaultExcludedExtentions,
	}
)

// Options  ...
type Options struct {
	ExcludedExtensions ExcludedExtensions
	ExcludedPaths      ExcludedPaths
	DecompressFn       func(c *gin.Context)
}

// Option  ...
type Option func(*Options)

// WithExcludedExtensions  ...
func WithExcludedExtensions(args []string) Option {
	return func(o *Options) {
		o.ExcludedExtensions = NewExcludedExtensions(args)
	}
}

// WithExcludedPaths  ...
func WithExcludedPaths(args []string) Option {
	return func(o *Options) {
		o.ExcludedPaths = NewExcludedPaths(args)
	}
}

// WithDecompressFn ...
func WithDecompressFn(decompressFn func(c *gin.Context)) Option {
	return func(o *Options) {
		o.DecompressFn = decompressFn
	}
}

// ExcludedExtensions Using map for better lookup performance
type ExcludedExtensions map[string]bool

// NewExcludedExtensions  ...
func NewExcludedExtensions(extensions []string) ExcludedExtensions {
	res := make(ExcludedExtensions)
	for _, e := range extensions {
		res[e] = true
	}
	return res
}

// Contains  ...
func (e ExcludedExtensions) Contains(target string) bool {
	_, ok := e[target]
	return ok
}

// ExcludedPaths  ...
type ExcludedPaths []string

// NewExcludedPaths  ...
func NewExcludedPaths(paths []string) ExcludedPaths {
	return ExcludedPaths(paths)
}

// Contains  ...
func (e ExcludedPaths) Contains(requestURI string) bool {
	for _, path := range e {
		if strings.HasPrefix(requestURI, path) {
			return true
		}
	}
	return false
}

// DefaultDecompressHandle  ...
func DefaultDecompressHandle(c *gin.Context) {
	if c.Request.Body == nil {
		return
	}
	r, err := gzip.NewReader(c.Request.Body)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Request.Header.Del("Content-Encoding")
	c.Request.Header.Del("Content-Length")
	c.Request.Body = r
}

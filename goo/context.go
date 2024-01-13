package goo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// Context is the encapsulation of http.Request and http.ResponseWriter
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
}

// newContext is the constructor of Context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm is the method of Context to get POST form
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query is the method of Context to get query string
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Param is the method of Context to get router param
func (c *Context) Param(key string) string {
	value := c.Params[key]
	return value
}

// Status is the method of Context to write status code
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader is the method of Context to write response header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String is the method of Context to write string response
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON is the method of Context to write json response
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data is the method of Context to write data response
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML is the method of Context to write html response
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

package goo

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestAddRoute(t *testing.T) {
	r := newTestRouter()
	fun := func(method string, path string, panicOK bool) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				if !panicOK {
					t.Fatal("the router panic")
				}
				return
			}
			if panicOK {
				t.Fatal("the router should panic")
			}
		}()
		r.addRoute(method, path, nil)
	}
	fun("GET", "/hello/:name/*assets", false)
	fun("GET", "/hello/:age", true)
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/goo")
	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}
	if ps["name"] != "goo" {
		t.Fatal("name should be equal to 'goo'")
	}
	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])

	n, ps = r.getRoute("GET", "/assets/file1.txt")
	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if n.pattern != "/assets/*filepath" {
		t.Fatal("should match /assets/*filepath")
	}
	if ps["filepath"] != "file1.txt" {
		t.Fatal("name should be equal to 'file1.txt'")
	}
	fmt.Printf("matched path: %s, params['filepath']: %s\n", n.pattern, ps["filepath"])
}

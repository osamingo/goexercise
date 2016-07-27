package goexercise

import (
	"fmt"
	"net/http"

	"github.com/celrenheit/lion"
	"golang.org/x/net/context"
)

func init() {
	l := lion.Classic()
	l.GetFunc("/", func(c context.Context, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "TEST")
	})
	l.Run()
}

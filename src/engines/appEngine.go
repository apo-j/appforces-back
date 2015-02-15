package engines

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

type AppEngine struct {
	DataEngine     DataEngine
	Request        *http.Request
	Renderer       render.Render
	MartiniContext martini.Context
}


package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"engines"
	"log"
	"net/http"
	"controllers"
	"json"
	"models"
)

var m *martini.Martini

func main() {
	if err := http.ListenAndServe(":8080", m); err != nil {
		log.Fatal(err)
	}
	/*go func() {
		// Listen on http: to raise an error and indicate that https: is required.
		//
		// This could also be achieved by passing the same `m` martini instance as
		// used by the https server, and by using a middleware that checks for https
		// and returns an error if it is not a secure connection. This would have the benefit
		// of handling only the defined routes. However, it is common practice to define
		// APIs on separate web servers from the web (html) pages, for maintenance and
		// scalability purposes, so it's not like it will block otherwise valid routes.
		//
		// It is also common practice to use a different subdomain so that cookies are
		// not transfered with every API request.
		// So with that in mind, it seems reasonable to refuse each and every request
		// on the non-https server, regardless of the route. This could of course be done
		// on a reverse-proxy in front of this web server.
		//
		if err := http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "https scheme is required", http.StatusBadRequest)
			})); err != nil {
			log.Fatal(err)
		}
	}()

	// Listen on https: with the preconfigured martini instance. The certificate files
	// can be created using this command in this repository's root directory:
	//
	// go run /path/to/goroot/src/pkg/crypto/tls/generate_cert.go --host="localhost"
	//
	if err := http.ListenAndServeTLS(":8001", "cert.pem", "key.pem", m); err != nil {
		log.Fatal(err)
	}*/
}

func init(){
	ctrls := controllers.InitControllers()

	m = martini.New()
	//set middleware
	m.Use(render.Renderer())
	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(PopulateAppContext)
	m.Use(CloseDatabase)


	//set routes
	r := martini.NewRouter()
	//Ex: /api/config/iluxe-privee.com
	r.Get("/api/config/:url", ctrls.ConfigController.Get)
	//Ex: /api/data/1/article/2

	r.Get("/api/data/:appId/:dataName/:id", ctrls.AppDataController.Get)
	r.Post("/api/data/:appId/:dataName", binding.Bind(models.AppDataSearchCriteria{}), ctrls.AppDataController.Search)
	r.Get("/api/pages/:appId/:pageId", ctrls.PageController.Get)

	//r.Get("/api/test", binding.Bind(models.AppDataSearchCriteria{}), ctrls.TestController.Get)
	r.Post("/api/test/new",binding.Bind(json.AppDataFieldsJSON{}), ctrls.TestController.New)

	// Add the router action
	m.Action(r.Handle)
}

func CloseDatabase(martiniContext martini.Context, appEngine *engines.AppEngine) {
	martiniContext.Next()
	appEngine.DataEngine.Dispose()
}

func PopulateAppContext(martiniContext martini.Context, w http.ResponseWriter, request *http.Request, renderer render.Render) {
	dataEngine := engines.CreateDataEngine()
	appEngine := &engines.AppEngine{Request: request, Renderer: renderer, MartiniContext: martiniContext, DataEngine: dataEngine}

	martiniContext.Map(appEngine)
}

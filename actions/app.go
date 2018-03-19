package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/cdacontrol/cda/models"
	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// T is a Translator instance for accessing the localization resources
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_cda_session",
		})

		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		// Setup and use translations:
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
			app.Stop(err)
		}
		app.Use(T.Middleware())

		app.GET("/", HomeHandler)

		app.ServeFiles("/assets", assetsBox)

		app.Resource("/users", UsersResource{&buffalo.BaseResource{}})
		app.Resource("/media", MediaResource{&buffalo.BaseResource{}})

		app.POST("/media/upload", MediaUploadHandler)
		app.Resource("/scenes", ScenesResource{&buffalo.BaseResource{}})
		app.Resource("/devices", DevicesResource{&buffalo.BaseResource{}})
		app.Resource("/events", EventsResource{&buffalo.BaseResource{}})
		app.Resource("/props", PropsResource{&buffalo.BaseResource{}})

		app.GET("/scene/editor/{scene}", SceneEditorHandler)
		// TODO: Add SceneData resource
		app.GET("/scene/data/{scene}", SceneGetHandler)
		app.POST("/scene/data", SceneCreateHandler)
		app.PUT("/scene/data/{scene}", SceneUpdateHandler)
		app.DELETE("/scene/data/{scene}", SceneDeleteHandler)
		// Device end-points
		app.POST("/scene/data/{scene}/device", SceneDeviceCreateHandler)
		app.PUT("/scene/data/{scene}/device/{device}", SceneDeviceUpdateHandler)
		app.DELETE("/scene/data/{scene}/device/{device}", SceneDeviceDeleteHandler)
		// Event end-points
		app.POST("/scene/data/{scene}/event", SceneEventCreateHandler)
		app.PUT("/scene/data/{scene}/event/{event}", SceneEventUpdateHandler)
		app.DELETE("/scene/data/{scene}/event/{event}", SceneEventDeleteHandler)
	}

	return app
}

package game

import (
	"github.com/tomyhero/submarine-game/utils"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

func NewMux(webApp utils.WebApp) *web.Mux {
	mux := web.New()
	mux.Use(middleware.SubRouter)

	root := Root{WebApp: webApp}
	mux.Get("/", root.Index)

	return mux
}

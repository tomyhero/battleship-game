package game

import (
	"github.com/tomyhero/battleship-game/utils"
	"github.com/zenazn/goji/web"
	"net/http"
)

type Root struct {
	utils.WebApp
}

func (self *Root) Index(c web.C, w http.ResponseWriter, r *http.Request) {
	stash := map[string]interface{}{"matching_endpoint": self.Config.GameServer.MatchingEndpoint}
	self.RenderHTML(w, "game/index", stash)
}

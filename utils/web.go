package utils

import (
	//"database/sql"
	"fmt"
	"net/http"
	"text/template"
)

// WebApp struct hold global data for this web application and also has utility methods.
// global data is not thread safe so please do not modify.
type WebApp struct {
	//DB     *sql.DB
	Config *Config
}

func (self *WebApp) RenderHTML(w http.ResponseWriter, path string, stash map[string]interface{}) {
	tmpl, err := template.ParseFiles(self.Config.TemplatePath(path))

	if err != nil {
		// render ERROR
		fmt.Println(err)
		return
	}

	err = tmpl.Execute(w, stash)

	if err != nil {
		// render ERROR
		fmt.Println(err)
		return
	}
}

// RenderJSON render stash into browser with JSON format.
func (self *WebApp) RenderJSON(w http.ResponseWriter, stash map[string]interface{}) {
	RenderJSON(w, stash)
}

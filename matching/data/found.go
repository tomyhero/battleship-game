package data

import ()

type Found struct {
	GameID string
}

func (self Found) Load(d map[string]interface{}) {
	//	self.GameID = d["GameID"].(string)
	// no body data
}

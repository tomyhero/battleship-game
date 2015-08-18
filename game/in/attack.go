package in

import ()

type Attack struct {
	X          int
	Y          int
	MatchingID string
}

func (self *Attack) Load(d map[string]interface{}) {
	self.X = int(d["x"].(float64))
	self.Y = int(d["y"].(float64))
	self.MatchingID = d["matching_id"].(string)
}

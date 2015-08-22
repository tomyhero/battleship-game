package in

type Resume struct {
	MatchingID string
	UserID     string
}

func (self *Resume) Load(d map[string]interface{}) {
	self.MatchingID = d["matching_id"].(string)
	self.UserID = d["user_id"].(string)
}

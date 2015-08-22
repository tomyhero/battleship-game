package in

type Login struct {
	MatchingID string
	UserID     string
}

func (self *Login) Load(d map[string]interface{}) {
	self.MatchingID = d["matching_id"].(string)
	self.UserID = d["user_id"].(string)
}

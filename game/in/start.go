package in

type Start struct {
	MatchingID string
	UserID     string
}

func (self *Start) Load(d map[string]interface{}) {
	self.MatchingID = d["matching_id"].(string)
	self.UserID = d["user_id"].(string)
}

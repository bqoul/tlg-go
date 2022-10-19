package upd

type Responce struct {
	Ok     bool                     `json:"ok"`
	Result []map[string]interface{} `json:"result"`
}

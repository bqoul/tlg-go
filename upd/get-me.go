package upd

type GetMe struct {
	Ok     bool `json:"ok"`
	Result User `json:"result"`
}

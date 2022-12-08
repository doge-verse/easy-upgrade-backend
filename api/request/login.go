package request

type Login struct {
	Address   string `json:"address" form:"address"`
	Signature string `json:"signature" form:"signature"`
	SignData  string `json:"signData" form:"signData"`
}

package entities

type Shipment struct {
	Id       int    `json:"id"`
	Address  string `json:"address"`
	S_status string `json:"s_status"`
}

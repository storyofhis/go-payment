package views

type Payment struct {
	ReferenceId string `json:"reference_id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

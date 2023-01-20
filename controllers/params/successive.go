package params

type SuccessiveReq struct {
	OrderID     string `json:"order_id"`
	ReferenceId string `json:"reference_id"`
	Type        string `json:"type"`
	Amount      int64  `json:"amount"`
}

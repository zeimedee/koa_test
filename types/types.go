package types

type Request struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Target   string  `json:"target"`
}

type Res struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

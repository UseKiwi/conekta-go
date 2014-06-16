package conekta

type ConektaEvent struct {
	Id        string     `json:"id"`
	Type      string     `json:"type,omitempty"`
	CreatedAt *timestamp `json:"created_at,omitempty"`
	Data      EventData  `json:"data"`
}

type EventData struct {
	Object EventObject `json:"object"`
}

type EventObject struct {
	Id             string        `json:"id"`
	Object         string        `json:"object,omitempty"`
	CreatedAt      *timestamp    `json:"created_at,omitempty"`
	Status         string        `json:"status,omitempty"`
	ReferenceId    string        `json:"reference_id"`
	FailureCode    string        `json:"failure_code,omitempty"`
	FailureMessage string        `json:"failure_message,omitempty"`
	Amount         int64         `json:"amount,omitempty"`
	PaidAt         *timestamp    `json:"paid_at,omitempty"`
	Fee            int64         `json:"fee,omitempty"`
	PaymentMethod  PaymentMethod `json:"payment_method"`
}

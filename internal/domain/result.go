package domain

type EventResult struct {
	Index   int    `json:"index"`
	EventID string `json:"event_id,omitempty"`
	Status  string `json:"status"`
	Code    string `json:"error_code,omitempty"`
}

type BatchResult struct {
	Accepted     int           `json:"accepted"`
	Published    int           `json:"published,omitempty"`
	Queued       int           `json:"queued,omitempty"`
	Failed       int           `json:"failed,omitempty"`
	BatchMode    string        `json:"batch_mode"`
	DeliveryMode string        `json:"delivery_mode"`
	Results      []EventResult `json:"results"`
}

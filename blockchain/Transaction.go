package blockchain

type Transaction struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Value    int    `json:"value"`
	Metadata string `json:"metadata"`
}

func NewTransaction(_from string, _to string, _value int, _metadata string) *Transaction {
	return &Transaction{
		From:     _from,
		To:       _to,
		Value:    _value,
		Metadata: _metadata,
	}
}

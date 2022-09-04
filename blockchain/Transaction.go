package blockchain

type TransactionList []*Transaction

type Transaction struct {
	from     string
	to       string
	value    int
	metadata string
}

func NewTransaction(_from string, _to string, _value int, _metadata string) *Transaction {
	return &Transaction{
		from:     _from,
		to:       _to,
		value:    _value,
		metadata: _metadata,
	}
}

//
//func loadJSON(path string) (transaction, error) {
//	content, err := os.ReadFile(path)
//	if err != nil {
//		return transaction{}, err
//	}
//
//	var loadedJSON transaction
//	err = json.Unmarshal(content, &loadedJSON)
//	if err != nil {
//		return transaction{}, err
//	}
//
//	return loadedJSON, nil
//}

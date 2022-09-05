package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Nonce     int            `json:"nonce"`
	Timestamp string         `json:"timestamp"`
	Payload   []*Transaction `json:"payload"`
	Hash      string         `json:"hash"`
	PrevHash  string         `json:"prevHash"`
}

func NewBlock(_payload []*Transaction, _prevHash string) *Block {
	return &Block{
		Timestamp: time.Now().String(),
		Payload:   _payload,
		PrevHash:  _prevHash,
	}
}

func (b *Block) HashBlock(_nonce int) {
	b.Nonce = _nonce

	// Payload to []byte
	encodedPayload := new(bytes.Buffer)
	json.NewEncoder(encodedPayload).Encode(b.Payload)

	// PrevHash to []byte
	encodedPrevHash := []byte(b.PrevHash)

	// Append nonce to encodedPayload
	binary.LittleEndian.AppendUint64(encodedPayload.Bytes(), uint64(b.Nonce))

	// Calculate hash with sha256
	hash := sha256.Sum256(append(encodedPrevHash, encodedPayload.Bytes()...))

	fmt.Println(string(hash[:]))

	b.Hash = string(hash[:])
}

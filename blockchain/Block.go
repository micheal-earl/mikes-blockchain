package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"time"
)

type Block struct {
	nonce     uint32
	timestamp time.Time
	payload   TransactionList
	hash      []byte
	prevHash  []byte
}

func NewBlock(_payload TransactionList, _prevHash []byte) *Block {
	return &Block{
		timestamp: time.Now(),
		payload:   _payload,
		prevHash:  _prevHash,
	}
}

func (b *Block) HashBlock(_nonce uint32) {
	b.nonce = _nonce
	encodedPayload := new(bytes.Buffer)
	json.NewEncoder(encodedPayload).Encode(b.payload)
	binary.LittleEndian.AppendUint32(encodedPayload.Bytes(), b.nonce)
	hash := sha256.Sum256(append(b.prevHash, encodedPayload.Bytes()...))
	b.hash = hash[:]
}

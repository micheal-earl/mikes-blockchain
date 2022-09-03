package blockchain

import (
	"crypto/sha256"
	"encoding/binary"
	"time"
)

type Block struct {
	nonce     uint32
	timestamp time.Time
	payload   []byte
	hash      []byte
	prevHash  []byte
}

func NewBlock(_payload, _prevHash []byte) *Block {
	return &Block{
		timestamp: time.Now(),
		payload:   _payload,
		prevHash:  _prevHash,
	}
}

func (b *Block) HashBlock(_nonce uint32) {
	b.nonce = _nonce
	binary.LittleEndian.AppendUint32(b.payload, b.nonce)
	hash := sha256.Sum256(append(b.payload, b.prevHash...))
	b.hash = hash[:]
}

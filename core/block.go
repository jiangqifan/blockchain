package core

import (
	"time"
	"crypto/sha256"
)

type Block struct {
	Head head
	Body string
}

type head struct {
	PreviousHash [32]byte
	Hash [32]byte
	Timestamp time.Time
	Nonce int
	BodyHash [32]byte	
}

func CreateBlock(previous Block) Block {
	return Block{
		Head: head{
			PreviousHash: previous.Head.Hash,
			Timestamp: time.Now(),
			Nonce: 0,
		},
	}
}

func createGodBlock(content string) Block{
	return Block{
		Head: head{
			Timestamp: time.Now(),
			Nonce: 0,
			Hash: sha256.Sum256([]byte(content)),
		},
		Body: content,
	}
}
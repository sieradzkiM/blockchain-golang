package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

const subsidy = 100

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

type TxOutput struct {
	Amount       int
	ScriptPubKey string
}

type TxInput struct {
	ID          []byte
	OutputIndex int
	ScriptSig   string
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func CoinbaseTx(recipient, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to %s", recipient)
	}

	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{subsidy, recipient}

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	tx.SetID()

	return &tx
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].OutputIndex == -1
}

func (in *TxInput) CanUnlockOutputWith(data string) bool {
	return in.ScriptSig == data
}

func (out *TxOutput) CanBeUnlockedWith(data string) bool {
	return out.ScriptPubKey == data
}

func NewTransaction(sender, recipient string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := chain.FindSpendableOutputs(sender, amount)

	if acc < amount {
		log.Panic("Error: not enough funds")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxInput{txID, out, sender}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{amount, recipient})

	if acc > amount {
		outputs = append(outputs, TxOutput{acc - amount, sender})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}
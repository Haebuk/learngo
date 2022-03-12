package blockchain

import (
	"errors"
	"time"

	"github.com/learngo/nomadcoin/utils"
)

const (
	minerReward int = 50 
)

type mempool struct {
	Txs []*Tx
}

var Mempool *mempool = &mempool{}

type Tx struct {
	Id string `json:"id"`
	Timestamp int `json:"timestamp"`
	TxIns []*TxIn `json:"txIns"`
	TxOuts []*TxOut `json:"txOut"`
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

type TxIn struct {
	TxID string `json:"txID"`
	Index int `json:"index"`
	Owner string `json:"owner"`
}

type TxOut struct {
	Owner string `json:"owner"`
	Amount int `json:"amount"`
}

type UTxOut struct {
	TxID string
	Index int
	Amount int
}

func isOnMempool(uTxOut *UTxOut) bool {
	exists := false
	for _, tx := range Mempool.Txs {
		for _, input := range tx.TxIns {
			exists = input.TxID == uTxOut.TxID && input.Index == uTxOut.Index
		}
	}
	return exists 
}

func makeCoinbaseTx(address string) *Tx {
	 txIns := []*TxIn{
		 {"", -1, "COINBASE"},
	 }
	 txOuts := []*TxOut{
		 {address, minerReward},
	 }
	 tx := Tx {
		 Id: "",
		 Timestamp: int(time.Now().Unix()),	
		 TxIns: txIns,
		TxOuts: txOuts,
	 }
	 tx.getId()
	 return &tx
} 

func makeTx(from, to string, amount int) (*Tx, error) {
	if Blockchain().BalanceByAddress(from) < amount {
		return nil, errors.New("not enough money")
	}
	var txOuts []*TxOut
	var txIns []*TxIn
	total := 0
	uTxOuts := Blockchain().UTxOutsByAddress(from)
	for _, uTxOut := range uTxOuts {
		if total >= amount {
			break
		}
		txIn := &TxIn{uTxOut.TxID, uTxOut.Index, from}
		txIns = append(txIns, txIn)
		total += uTxOut.Amount
	}
	if change := total - amount; change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}
	txOut := &TxOut{to, amount}
	txOuts = append(txOuts, txOut)
	tx := &Tx{
		Id: "",
		Timestamp: int(time.Now().Unix()),
		TxIns: txIns,
		TxOuts: txOuts,
	}
	tx.getId()
	return tx, nil
}

func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("nico", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("nico")
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil
	return txs
}
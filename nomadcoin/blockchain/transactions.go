package blockchain

import (
	"errors"
	"sync"
	"time"

	"github.com/learngo/nomadcoin/utils"
	"github.com/learngo/nomadcoin/wallet"
)

const (
	minerReward int = 50 
)

type mempool struct {
	Txs []*Tx
	m sync.Mutex
}	 

var m *mempool
var memOnce sync.Once

func Mempool() *mempool {
	memOnce.Do(func() {
		m = &mempool{}
	})
	return m
}

type Tx struct {
	Id string `json:"id"`
	Timestamp int `json:"timestamp"`
	TxIns []*TxIn `json:"txIns"`
	TxOuts []*TxOut `json:"txOut"`
}

type TxIn struct {
	TxID string `json:"txID"`
	Index int `json:"index"`
	Signature string `json:"signature"`
}

type TxOut struct {
	Address string `json:"address"`
	Amount int `json:"amount"`
}

type UTxOut struct {
	TxID string
	Index int
	Amount int
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

func (t *Tx) sign() {
	for _, txIn := range t.TxIns {
		txIn.Signature = wallet.Sign(t.Id, wallet.Wallet())
	}
}

func validate(tx *Tx) bool {
	valid := true
	for _, txIn := range tx.TxIns {
		prevTx := FindTx(Blockchain(), txIn.TxID)
		if prevTx == nil {
			valid = false
			break
		}
		address := prevTx.TxOuts[txIn.Index].Address
		valid = wallet.Verify(txIn.Signature, tx.Id, address)
		if !valid {
			break
		} 
	}
	return valid
}

func isOnMempool(uTxOut *UTxOut) bool {
	exists := false
	Outer:
	for _, tx := range Mempool().Txs {
		for _, input := range tx.TxIns {
			if input.TxID == uTxOut.TxID && input.Index  == uTxOut.Index {
				exists = true
				break Outer
			}
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

var ErrorNoMoney = errors.New("not enough money")
var ErrorNotValid = errors.New("Tx Invalid")

func makeTx(from, to string, amount int) (*Tx, error) {
	if BalanceByAddress(from, Blockchain()) < amount {
		return nil, ErrorNoMoney
	}
	var txOuts []*TxOut
	var txIns []*TxIn
	total := 0
	uTxOuts := UTxOutsByAddress(from, Blockchain())
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
	tx.sign()
	valid := validate(tx)
	if !valid {
		return nil, ErrorNotValid 
	}
	return tx, nil
}

func (m *mempool) AddTx(to string, amount int) (*Tx, error) {
	tx, err := makeTx(wallet.Wallet().Address, to, amount)
	if err != nil {
		return nil, err
	}
	m.Txs = append(m.Txs, tx)
	return tx, nil
}

func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx(wallet.Wallet().Address)
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil
	return txs
}

func (m *mempool) AddPeerTx(tx *Tx) {
	m.m.Lock()
	defer m.m.Unlock()
	
	m.Txs = append(m.Txs, tx)
}
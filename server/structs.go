package server

import "github.com/lyulka/trivial-ledger/structs"

type ProposeTransactionRequest structs.ProposedTransaction

type ProposeTransactionResponse struct {
	BlockNum int `json:"blockNum"`
	TxNumber int `json:"txNumber"`
}

type GetTransactionRequestParameter struct {
	BlockNum int `json:"block_num"`
	TxNumber int `json:"tx_number"`
}

type GetTransactionRequest struct {
	BlockNum int `json:"blockNum"`
	TxNumber int `json:"txNumber"`
}

type GetTransactionResponse structs.Transaction

type GetBlockRequest struct {
	BlockNum int `json:"blockNum"`
}

type GetBlockRequestParameter struct {
	BlockNum int `json:"block_num"`
}


type GetBlockResponse structs.Block

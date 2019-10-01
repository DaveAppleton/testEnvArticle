package main

import (
	"math/big"

	"github.com/DaveAppleton/memorykeys"

	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var hushSend bool

func sendEth(ec bind.ContractTransactor, sender string, recipient string, amount *big.Int, gasLimit uint64) (*types.Transaction, error) {
	senderKey, _ := memorykeys.GetPrivateKey(sender)
	senderAddress, _ := memorykeys.GetAddress(sender)
	recipientAddress, _ := memorykeys.GetAddress(recipient)
	nonce, err := ec.PendingNonceAt(context.Background(), *senderAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := ec.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	s := types.HomesteadSigner{}
	data := common.FromHex("0x")
	t := types.NewTransaction(nonce, *recipientAddress, amount, gasLimit, gasPrice, data)
	tx, err := types.SignTx(t, s, senderKey)
	if err != nil {
		return nil, err
	}
	err = ec.SendTransaction(context.Background(), tx)
	return tx, err
}

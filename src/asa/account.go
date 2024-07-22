package main

import (
	"context"

	"github.com/algorand/go-algorand-sdk/v2/crypto"
	"github.com/algorand/go-algorand-sdk/v2/mnemonic"
)

func generateAccount() {
	account := crypto.GenerateAccount()
	mn, err := mnemonic.FromPrivateKey(account.PrivateKey)
	writeMnemonic(&addressMnemonic{Mnemonic: mn, Address: account.Address.String(), err: err, Err: ""})
}

func microToAlgo(microAlgos uint64) uint64 {
	return microAlgos / 1_000_000
}

func sandboxAccounts() {
	algodClient := getAlgodClient()
	accts, err := getSandboxAccounts()
	accounts := make([]*accountResp, 0)
	if err != nil {
		acc := &accountResp{}
		acc.err = err
		accounts = append(accounts, acc)
	}

	for _, j := range accts {
		info := algodClient.AccountInformation(j.Address.String())
		resp, err := info.Do(context.Background())
		acc := &accountResp{Address: j.Address.String(), Amount: microToAlgo(resp.Amount), err: err}
		accounts = append(accounts, acc)
	}
	writeAccounts(accounts...)
}

func importAccount(userMnemonic string) (crypto.Account, error) {
	var acc crypto.Account
	key, err := mnemonic.ToPrivateKey(userMnemonic)
	if err != nil {
		return acc, err
	}
	return crypto.AccountFromPrivateKey(key)
}

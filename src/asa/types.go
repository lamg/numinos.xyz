package main

import (
	"fmt"
)

type addressMnemonic struct {
	Mnemonic string `json:"mnemonic"`
	Address  string `json:"address"`
	Err      string `json:"err"`
	err      error
}

type assetResp struct {
	AssetId uint64 `json:"assetId"`
	Round   uint64 `json:"round"`
	TxId    string `json:"txId"`
	err     error
	Err     string `json:"err"`
	Request string `json:"request"`
}

func newMintAssetResp(r *assetResp) *assetResp {
	r.Request = "mint"
	return r
}

func newConfigureAssetResp(r *assetResp) *assetResp {
	r.Request = "configure"
	return r
}

func newAssetOptInResp(r *assetResp) *assetResp {
	r.Request = "optIn"
	return r
}

func newAssetOptOutResp(r *assetResp) *assetResp {
	r.Request = "optOut"
	return r
}

func newXferAssetResp(r *assetResp) *assetResp {
	r.Request = "xfer"
	return r
}

func newClawbackAssetResp(r *assetResp) *assetResp {
	r.Request = "clawback"
	return r
}

func newDestroyAssetResp(r *assetResp) *assetResp {
	r.Request = "destroy"
	return r
}

func newFreezeAssetResp(r *assetResp) *assetResp {
	r.Request = "freeze"
	return r
}

type accountResp struct {
	Address string `json:"address"`
	Amount  uint64 `json:"amount"` // in Algos
	Err     string `json:"err"`
	err     error
}

type suggestedParamsErr struct {
	err error
}

func (e *suggestedParamsErr) Error() string {
	return fmt.Sprintf("error getting suggested tx params: %s", e.err)
}

func newSuggestedParamsErr(err error) (e *suggestedParamsErr) {
	e = &suggestedParamsErr{err: err}
	return
}

type txErr struct {
	err error
}

func (e *txErr) Error() string {
	return fmt.Sprintf("failed to make transaction: %s", e.err)
}

func newTransactionErr(err error) (e *txErr) {
	e = &txErr{err: err}
	return
}

type signTxErr struct{ err error }

func (e *signTxErr) Error() string {
	return fmt.Sprintf("failed to sign transaction: %s", e.err)
}

func newSignTxErr(err error) (e *signTxErr) {
	e = &signTxErr{err: err}
	return
}

type sendTxErr struct{ err error }

func (e *sendTxErr) Error() string {
	return fmt.Sprintf("failed to send transaction: %s", e.err)
}

func newSendTxErr(err error) (e *sendTxErr) {
	e = &sendTxErr{err: err}
	return
}

type waitConfirmationErr struct{ err error }

func (e *waitConfirmationErr) Error() string {
	return fmt.Sprintf("error waiting for confirmation:  %s", e.err)
}

func newWaitConfirmationErr(err error) (e *waitConfirmationErr) {
	e = &waitConfirmationErr{err: err}
	return
}

package main

import (
	"context"
	"fmt"

	"github.com/algorand/go-algorand-sdk/v2/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/v2/client/v2/common/models"
	"github.com/algorand/go-algorand-sdk/v2/crypto"
	"github.com/algorand/go-algorand-sdk/v2/transaction"
	"github.com/algorand/go-algorand-sdk/v2/types"
)

type assetRequest struct {
	name     string
	unitName string
	url      string
}

func signSendWait(client *algod.Client, account crypto.Account, txn types.Transaction, err error) *assetResp {
	var (
		txId         string
		stx          []byte
		confirmedTxn models.PendingTransactionInfoResponse
		assetId      uint64
		round        uint64
	)
	if err != nil {
		err = newTransactionErr(err)
		goto end
	}
	// sign the transaction
	txId, stx, err = crypto.SignTransaction(account.PrivateKey, txn)
	if err != nil {
		err = newSignTxErr(err)
		goto end
	}

	// Broadcast the transaction to the network
	_, err = client.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		err = newSendTxErr(err)
		goto end
	}

	// Wait for confirmation
	confirmedTxn, err = transaction.WaitForConfirmation(client, txId, 4, context.Background())
	if err != nil {
		err = newWaitConfirmationErr(err)
		goto end
	}
	assetId = confirmedTxn.AssetIndex
	round = confirmedTxn.ConfirmedRound
end:

	return &assetResp{AssetId: assetId, Round: round, TxId: txId, Request: "", Err: "", err: err}
}

func createAsset(algodClient *algod.Client, creator crypto.Account, n *assetRequest) *assetResp {
	// example: ASSET_CREATE
	// Configure parameters for asset creation
	var (
		creatorAddr       = creator.Address.String()
		assetName         = n.name
		unitName          = n.unitName
		assetURL          = n.url
		assetMetadataHash = "" // optional
		defaultFrozen     = false
		decimals          = uint32(0)
		totalIssuance     = uint64(1)

		manager  = creatorAddr
		reserve  = creatorAddr
		freeze   = creatorAddr
		clawback = creatorAddr

		note []byte
	)

	var (
		txParams types.SuggestedParams
		err      error
		txn      types.Transaction
		resp     = &assetResp{}
	)

	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		resp.err = newSuggestedParamsErr(err)
		goto end
	}

	// Construct the transaction
	txn, err = transaction.MakeAssetCreateTxn(
		creatorAddr, note, txParams, totalIssuance, decimals,
		defaultFrozen, manager, reserve, freeze, clawback,
		unitName, assetName, assetURL, assetMetadataHash,
	)

	resp = signSendWait(algodClient, creator, txn, err)
	// example: ASSET_CREATE
end:
	return newMintAssetResp(resp)
}

func configureAsset(client *algod.Client, assetID uint64, creator crypto.Account) *assetResp {
	// example: ASSET_CONFIG
	creatorAddr := creator.Address.String()
	var (
		newManager  = creatorAddr
		newFreeze   = creatorAddr
		newClawback = creatorAddr
		newReserve  = ""

		strictAddrCheck = false
		note            []byte
	)
	var (
		sp   types.SuggestedParams
		err  error
		txn  types.Transaction
		resp = &assetResp{}
	)

	// Get network-related transaction parameters and assign
	sp, err = client.SuggestedParams().Do(context.Background())
	if err != nil {
		resp.err = newSuggestedParamsErr(err)
		goto end
	}

	txn, err = transaction.MakeAssetConfigTxn(creatorAddr, note, sp, assetID, newManager, newReserve, newFreeze, newClawback, strictAddrCheck)
	if err != nil {
		resp.err = newTransactionErr(err)
		goto end
	}
	resp = signSendWait(client, creator, txn, err)

end:
	// example: ASSET_CONFIG
	return newConfigureAssetResp(resp)
}

func optInAsset(client *algod.Client, assetID uint64, user crypto.Account) *assetResp {
	// example: ASSET_OPTIN
	userAddr := user.Address.String()
	var (
		sp   types.SuggestedParams
		err  error
		txn  types.Transaction
		resp = &assetResp{}
	)

	sp, err = client.SuggestedParams().Do(context.Background())
	if err != nil {
		resp.err = newSuggestedParamsErr(err)
		goto end
	}

	txn, err = transaction.MakeAssetAcceptanceTxn(userAddr, nil, sp, assetID)
	if err != nil {
		resp.err = newTransactionErr(err)
		goto end
	}
	resp = signSendWait(client, user, txn, err)

end:
	return newAssetOptInResp(resp)
}

func optOutAsset(client *algod.Client, assetID uint64, creator, user crypto.Account) *assetResp {
	// example: ASSET_OPT_OUT
	userAddr := user.Address.String()
	var (
		sp   types.SuggestedParams
		err  error
		txn  types.Transaction
		resp = &assetResp{}
	)

	sp, err = client.SuggestedParams().Do(context.Background())
	if err != nil {
		resp.err = newSuggestedParamsErr(err)
		goto end
	}

	txn, err = transaction.MakeAssetTransferTxn(userAddr, creator.Address.String(), 0, nil, sp, creator.Address.String(), assetID)
	if err != nil {
		resp.err = newTransactionErr(err)
		goto end
	}
	resp = signSendWait(client, user, txn, err)

end:
	// example: ASSET_OPT_OUT
	return newAssetOptOutResp(resp)
}

func xferAsset(client *algod.Client, assetID uint64, owner crypto.Account, userAddress string) *assetResp {
	// example: ASSET_XFER
	var (
		ownerAddress = owner.Address.String()
	)
	var (
		sp   types.SuggestedParams
		err  error
		txn  types.Transaction
		resp = &assetResp{}
	)

	sp, err = client.SuggestedParams().Do(context.Background())
	if err != nil {
		resp.err = newSuggestedParamsErr(err)
		goto end
	}

	txn, err = transaction.MakeAssetTransferTxn(ownerAddress, userAddress, 1, nil, sp, "", assetID)
	if err != nil {
		resp.err = newTransactionErr(err)
		goto end
	}
	resp = signSendWait(client, owner, txn, err)

end:
	// example: ASSET_XFER
	return newXferAssetResp(resp)
}

func freezeAsset(client *algod.Client, assetID uint64, creator crypto.Account, userAddress string) *assetResp {
	// example: ASSET_FREEZE
	var (
		creatorAddr = creator.Address.String()

		resp = &assetResp{}
		txn  types.Transaction
	)

	sp, err := client.SuggestedParams().Do(context.Background())
	if err != nil {
		resp.err = newSuggestedParamsErr(err)
		goto end
	}

	// Create a freeze asset transaction with the target of the user address
	// and the new freeze setting of `true`
	txn, err = transaction.MakeAssetFreezeTxn(creatorAddr, nil, sp, assetID, userAddress, true)
	if err != nil {
		resp.err = newTransactionErr(err)
		goto end
	}
	resp = signSendWait(client, creator, txn, err)
end:
	// example: ASSET_FREEZE
	return newFreezeAssetResp(resp)
}

func clawbackAsset(client *algod.Client, assetID uint64, creator crypto.Account, destAddress string) *assetResp {
	// example: ASSET_CLAWBACK
	var (
		creatorAddr = creator.Address.String()

		resp = &assetResp{}
		txn  types.Transaction
		err  error
	)

	sp, err := client.SuggestedParams().Do(context.Background())
	if err != nil {
		resp.err = newSuggestedParamsErr(err)
		goto end
	}

	// Create a new clawback transaction with the target of the user address and the recipient as the creator
	// address, being sent from the address marked as `clawback` on the asset, in this case the same as creator
	txn, err = transaction.MakeAssetRevocationTxn(creatorAddr, destAddress, 1, creatorAddr, nil, sp, assetID)
	if err != nil {
		resp.err = newTransactionErr(err)
		goto end
	}
	resp = signSendWait(client, creator, txn, err)
end:
	// example: ASSET_CLAWBACK
	return newClawbackAssetResp(resp)
}

func destroyAsset(client *algod.Client, assetID uint64, creator crypto.Account) *assetResp {
	// example: ASSET_DELETE
	var (
		creatorAddr = creator.Address.String()
		resp        = &assetResp{}
		txn         types.Transaction
	)

	sp, err := client.SuggestedParams().Do(context.Background())
	if err != nil {
		resp.err = newSuggestedParamsErr(err)
		goto end
	}

	// Create a new clawback transaction with the target of the user address and the recipient as the creator
	// address, being sent from the address marked as `clawback` on the asset, in this case the same as creator
	txn, err = transaction.MakeAssetDestroyTxn(creatorAddr, nil, sp, assetID)
	if err != nil {
		resp.err = newTransactionErr(err)
		goto end
	}
	resp = signSendWait(client, creator, txn, err)

end:
	return newDestroyAssetResp(resp)
}

func infoAsset(client *algod.Client, address string, assetId uint64) {
	x, err := client.AccountAssetInformation(address, assetId).Do(context.Background())
	if err == nil {
		fmt.Printf("asset: %v\n", x.CreatedAsset)
	} else {
		fmt.Printf("error %s\n", err.Error())
	}
}

package main

import (
	"flag"
	"fmt"
)

func deleteAssetCli(mnemonic string, assetId uint64) {
	client := getAlgodClient()
	acc, err := importAccount(mnemonic)
	if err != nil {
		fmt.Printf(`{"err": "failed mnemonic account import with %s"}\n`, err)
		return
	}
	resp := destroyAsset(client, assetId, acc)
	writeResp(resp)
}

func deleteAssetSandbox(sandboxAccIndex uint64, assetId uint64) {
	client := getAlgodClient()
	accs, err := getSandboxAccounts()
	if err != nil {
		fmt.Printf(`{"err": "error getting sandbox accounts %s"}\n`, err)
	}
	if sandboxAccIndex < length(accs) {
		acc := accs[sandboxAccIndex]
		resp := destroyAsset(client, assetId, acc)
		writeResp(resp)
	}
}

func createAssetCli(sandboxAccIndex uint64, n *assetRequest) {
	client := getAlgodClient()
	accs, err := getSandboxAccounts()
	if err != nil {
		fmt.Printf(`{"err": "error getting sandbox accounts %s"}\n`, err)
	}
	xs := make([]*assetResp, 0)
	if sandboxAccIndex < length(accs) {
		acc := accs[sandboxAccIndex]
		resp0 := createAsset(client, acc, n)
		xs = append(xs, resp0)
		if resp0.err == nil {
			resp1 := configureAsset(client, resp0.AssetId, acc)
			xs = append(xs, resp1)
			if resp1.err == nil {
				resp2 := optInAsset(client, resp0.AssetId, acc)
				xs = append(xs, resp2)
			}
		}
	}
	writeResp(xs...)
}

func freezeAssetSandbox(sandboxAccIndex uint64, assetId uint64, destAddress string) {
	client := getAlgodClient()
	accs, err := getSandboxAccounts()
	if err != nil {
		fmt.Printf(`[{"err": "error getting sandbox accounts %s"}]`, err)
		return
	}
	xs := make([]*assetResp, 0)
	if sandboxAccIndex < length(accs) {
		acc := accs[sandboxAccIndex]
		xs = append(xs, freezeAsset(client, assetId, acc, destAddress))
	}
	writeResp(xs...)
}

func freezeAssetCli(mnemonic string, assetId uint64, destAddress string) {
	client := getAlgodClient()
	acc, err := importAccount(mnemonic)
	if err != nil {
		fmt.Printf(`[{"err": "failed mnemonic account import with %s"}]`, err)
		return
	}
	resp := freezeAsset(client, assetId, acc, destAddress)
	writeResp(resp)
}

func transferAssetSandbox(sandboxAccIndex uint64, assetId uint64, destAddress string) {
	client := getAlgodClient()
	accs, err := getSandboxAccounts()
	if err != nil {
		fmt.Printf(`[{"err": "error getting sandbox accounts %s"}]`, err)
		return
	}
	xs := make([]*assetResp, 0)
	if sandboxAccIndex < length(accs) {
		acc := accs[sandboxAccIndex]
		xs = append(xs, xferAsset(client, assetId, acc, destAddress))
	}
	writeResp(xs...)
}

func transferAssetCli(mnemonic string, assetId uint64, destAddress string) {
	client := getAlgodClient()
	acc, err := importAccount(mnemonic)
	if err != nil {
		fmt.Printf(`{"err": "failed mnemonic account import with %s"}\n`, err)
		return
	}
	resp := xferAsset(client, assetId, acc, destAddress)
	writeResp(resp)
}

func clawbackAssetSandbox(sandboxAccIndex uint64, assetId uint64, destAddress string) {
	client := getAlgodClient()
	accs, err := getSandboxAccounts()
	if err != nil {
		fmt.Printf(`[{"err": "error getting sandbox accounts %s"}]`, err)
		return
	}
	xs := make([]*assetResp, 0)
	if sandboxAccIndex < length(accs) {
		acc := accs[sandboxAccIndex]
		xs = append(xs, clawbackAsset(client, assetId, acc, destAddress))
	}
	writeResp(xs...)
}

func clawbackAssetCli(mnemonic string, assetId uint64, destAddress string) {
	client := getAlgodClient()
	acc, err := importAccount(mnemonic)
	if err != nil {
		fmt.Printf(`{"err": "failed mnemonic account import with %s"}\n`, err)
		return
	}
	resp := clawbackAsset(client, assetId, acc, destAddress)
	writeResp(resp)
}

func infoAssetCli(sandboxAccIndex, assetId uint64) {
	client := getAlgodClient()
	accs, err := getSandboxAccounts()
	if err != nil {
		fmt.Printf(`[{"err": "error getting sandbox accounts %s"}]`, err)
		return
	}
	if sandboxAccIndex < length(accs) {
		acc := accs[sandboxAccIndex]
		infoAsset(client, acc.Address.String(), assetId)
	}
}

func main() {
	var (
		// accounts
		list         bool
		genAccount   bool
		mnemonic     string
		sandboxIndex uint64

		// asset destruction

		destroy bool
		assetId uint64

		// asset creation
		create    bool
		assetName string
		unitName  string
		url       string

		// other asset operations
		clawback    bool
		freeze      bool
		transfer    bool
		destAddress string
		info        bool
	)

	// accounts
	flag.BoolVar(&list, "l", false, "get sandbox accounts")
	flag.BoolVar(&genAccount, "genAccount", false, "generate account")
	flag.StringVar(&mnemonic, "mnemonic", "", "mnemonic")
	flag.Uint64Var(&sandboxIndex, "sandboxIndex", 0, "sandbox account index to use in the operation")

	// target selection
	flag.Uint64Var(&assetId, "assetId", 0, "asset ID")
	flag.StringVar(&destAddress, "destAddress", "", "destination address")

	// asset creation
	flag.BoolVar(&create, "c", false, "create an asset")
	flag.StringVar(&assetName, "assetName", "", "asset name")
	flag.StringVar(&unitName, "unitName", "", "unit name")
	flag.StringVar(&url, "url", "", "asset url")

	// asset management
	flag.BoolVar(&clawback, "clawback", false, "clawback")
	flag.BoolVar(&freeze, "freeze", false, "freeze")
	flag.BoolVar(&transfer, "transfer", false, "transfer")
	flag.BoolVar(&destroy, "destroy", false, "delete asset")
	flag.BoolVar(&info, "info", false, "asset information")

	flag.Parse()

	if list {
		sandboxAccounts()
		return
	}
	if genAccount {
		generateAccount()
		return
	}

	if destroy {
		if mnemonic == "" {
			deleteAssetSandbox(sandboxIndex, assetId)
		} else {
			deleteAssetCli(mnemonic, assetId)
		}
		return
	}

	if create {
		req := &assetRequest{name: assetName, unitName: unitName, url: url}
		createAssetCli(sandboxIndex, req)
		return
	}

	if freeze {
		if mnemonic == "" {
			freezeAssetSandbox(sandboxIndex, assetId, destAddress)

		} else {
			freezeAssetCli(mnemonic, assetId, destAddress)
		}
		return
	}

	if transfer {
		if mnemonic == "" {
			transferAssetSandbox(sandboxIndex, assetId, destAddress)
		} else {
			transferAssetCli(mnemonic, assetId, destAddress)
		}
		return
	}

	if clawback {
		if mnemonic == "" {
			clawbackAssetSandbox(sandboxIndex, assetId, destAddress)
		} else {
			clawbackAssetCli(mnemonic, assetId, destAddress)
		}
		return
	}

	if info {
		infoAssetCli(sandboxIndex, assetId)
		return
	}
	flag.Usage()
}

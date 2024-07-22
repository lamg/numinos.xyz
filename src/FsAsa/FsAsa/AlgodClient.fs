module FsAsa.AlgodClient

open System
open System.Net.Http
open FsHttp


(*
crypto.SignTransaction
client.SendRawTransaction
transaction.WaitForConfirmation
algodClient.SuggestedParams
transaction.MakeAssetCreateTxn
transaction.MakeAssetConfigTxn
transaction.MakeAssetAcceptanceTxn
transaction.MakeAssetTransferTxn
transaction.MakeAssetTransferTxn
transaction.MakeAssetFreezeTxn
transaction.MakeAssetRevocationTxn
transaction.MakeAssetDestroyTxn
client.AccountAssetInformation
*)

[<Measure>]
type MicroAlgo

[<Measure>]
type Round

type TransactionParametersResponse =
  { fee: uint64<MicroAlgo>
    ``genesis-id``: string
    ``genesis-hash``: byte array
    ``last-round``: uint64<Round>
    ``min-fee``: uint64<MicroAlgo>
    ``consensus-version``: string }

type SuggestedParams =
  { fee: uint64<MicroAlgo>
    genesisID: string
    genesisHash: byte array
    firstRoundValid: uint64<Round>
    lastRoundValid: uint64<Round>
    consensusVersion: string
    minFee: uint64<MicroAlgo> }

[<Literal>]
let localAlgod = "https://localhost:4001"

let defaultAuthToken = String.replicate 64 "a"

[<Literal>]
let algodAuthHeader = "X-Algo-API-Token"

type AlgodClient(baseUrl: string, authToken: string) =
  let client =
    http {
      config_useBaseUrl baseUrl
      header algodAuthHeader authToken
    }

  member _.suggestedParams() : Async<SuggestedParams> =
    async {
      let! r = client { GET "/v2/transactions/params" } |> Request.sendAsync
      let! p = Response.deserializeJsonAsync<TransactionParametersResponse> r

      return
        { fee = p.fee
          genesisID = p.``genesis-id``
          genesisHash = p.``genesis-hash``
          firstRoundValid = p.``last-round``
          lastRoundValid = p.``last-round`` + 1000UL<Round>
          minFee = p.``min-fee``
          consensusVersion = p.``consensus-version`` }
    }

  member _.sendRawTransaction() = ()
  member _.accountAssetInformation() = ()

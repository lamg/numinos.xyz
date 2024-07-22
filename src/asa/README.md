# Command line utility for performing relevant domain operations in the Algorand network

In order to use install first [algokit][algokit]

Functionality:

- [x] Mint NFT

```sh
asa -c -assetName coco -unitName coco -url https://numinos.xyz/coco
```

Follow the following ARCs
- [ARC3][ARC3]: conventions for fungible/non-fungible tokens
- [ARC16][ARC16]: Convention for declaring traits of an NFT's
- [ARC18][ARC18]: Royalty Enforcement Specification
- [ARC19][ARC19]: Templating of NFT ASA URLs for mutability
- [ARC36][ARC36]: Convention for declaring filters of an NFT
- [ARC72][ARC72]: Algorand Smart Contract NFT Specification
- [ARC74][ARC74]: NFT Indexer API
- [ARC69][ARC69]: ASA Parameters Conventions, Digital Media
- [ARC × Platform support matrix](https://arc.algorand.foundation/nfts)

- [ ] Create invoice
- [ ] Check if payment was received for an invoice
- [ ] Mint an NFT and send it to the payer for accounts that have that configuration
- [x] Destroy asset

```sh
asa -assetId 1019 -destroy
```

- [x] List sandbox accounts

```sh
asa -l
```

- [x] Inspect asset

```sh
asa -info -assetId 1041
```

[algokit]: https://developer.algorand.org/docs/get-started/algokit/#install-algokit
[ARC3]: https://arc.algorand.foundation/ARCs/arc-0003
[ARC16]: https://arc.algorand.foundation/ARCs/arc-0016
[ARC19]: https://arc.algorand.foundation/ARCs/arc-0019
[ARC36]: https://arc.algorand.foundation/ARCs/arc-0036
[ARC69]: https://arc.algorand.foundation/ARCs/arc-0069
[ARC72]: https://arc.algorand.foundation/ARCs/arc-0072
[ARC74]: https://arc.algorand.foundation/ARCs/arc-0074

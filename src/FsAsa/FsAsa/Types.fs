module FsAsa.Types

// reference https://github.com/algorand/go-algorand-sdk/blob/develop/types/basics.go
// https://github.com/algorand/go-algorand-sdk/blob/develop/types/transaction.go

type Address = byte array // [hashLenBytes]byte

type AssetParams =
  {
    // Total specifies the total number of units of this asset
    // created.
    Total: uint64 //`codec:"t"`
    // Decimals specifies the number of digits to display after the decimal
    // place when displaying this asset. A value of 0 represents an asset
    // that is not divisible, a value of 1 represents an asset divisible
    // into tenths, and so on. This value must be between 0 and 19
    // (inclusive).
    Decimals: uint32 //`codec:"dc"`
    // DefaultFrozen specifies whether slots for this asset
    // in user accounts are frozen by default or not.
    DefaultFrozen: bool //`codec:"df"`
    // UnitName specifies a hint for the name of a unit of
    // this asset.
    UnitName: string //`codec:"un,allocbound=config.MaxAssetUnitNameBytes"`
    // AssetName specifies a hint for the name of the asset.
    AssetName: string //`codec:"an,allocbound=config.MaxAssetNameBytes"`
    // URL specifies a URL where more information about the asset can be
    // retrieved
    URL: string //`codec:"au,allocbound=config.MaxAssetURLBytes"`
    // MetadataHash specifies a commitment to some unspecified asset
    // metadata. The format of this metadata is up to the application.
    MetadataHash: byte array // [AssetMetadataHashLen]byte //`codec:"am"`
    // Manager specifies an account that is allowed to change the
    // non-zero addresses in this AssetParams.
    Manager: Address //`codec:"m"`
    // Reserve specifies an account whose holdings of this asset
    // should be reported as "not minted".
    Reserve: Address //`codec:"r"`
    // Freeze specifies an account that is allowed to change the
    // frozen state of holdings of this asset.
    Freeze: Address //`codec:"f"`
    // Clawback specifies an account that is allowed to take units
    // of this asset from any account.
    Clawback: Address } //`codec:"c"`


type TxType = string

type Header = string

type VotePK = byte array // type VotePK [ed25519.PublicKeySize]byte

type VRFPK = byte array

type MerkleVerifier = byte array // [KeyStoreRootSize]byte

type AssetIndex = uint64

[<Measure>]
type MicroAlgo

[<Measure>]
type Round

type KeyregTxnFields =
  { VotePK: VotePK // `codec:"votekey"`
    SelectionPK: VRFPK // `codec:"selkey"`
    VoteFirst: uint64<Round> // `codec:"votefst"`
    VoteLast: uint64<Round> // `codec:"votelst"`
    VoteKeyDilution: uint64 // `codec:"votekd"`
    Nonparticipation: bool // `codec:"nonpart"`
    StateProofPK: MerkleVerifier } // `codec:"sprfkey"`

type PaymentTxnFields =
  { Receiver: Address // `codec:"rcv"`
    Amount: uint64<MicroAlgo> // `codec:"amt"`
    CloseRemainderTo: Address } //codec:"close"`

type AssetConfigTxnFields =
  {
    // ConfigAsset is the asset being configured or destroyed.
    // A zero value means allocation.
    ConfigAsset: AssetIndex // `codec:"caid"`
    // AssetParams are the parameters for the asset being
    // created or re-configured.  A zero value means destruction.
    AssetParams: AssetParams } // `codec:"apar"`

// AssetTransferTxnFields captures the fields used for asset transfers.
type AssetTransferTxnFields =
  { XferAsset: AssetIndex //`codec:"xaid"`
    // AssetAmount is the amount of asset to transfer.
    // A zero amount transferred to self allocates that asset
    // in the account's Assets map.
    AssetAmount: uint64 //`codec:"aamt"`
    // AssetSender is the sender of the transfer.  If this is not
    // a zero value, the real transaction sender must be the Clawback
    // address from the AssetParams.  If this is the zero value,
    // the asset is sent from the transaction's Sender.
    AssetSender: Address //`codec:"asnd"`
    // AssetReceiver is the recipient of the transfer.
    AssetReceiver: Address //`codec:"arcv"`
    // AssetCloseTo indicates that the asset should be removed
    // from the account's Assets map, and specifies where the remaining
    // asset holdings should be transferred.  It's always valid to transfer
    // remaining asset holdings to the creator account.
    AssetCloseTo: Address } //`codec:"aclose"`

// AssetFreezeTxnFields captures the fields used for freezing asset slots.
type AssetFreezeTxnFields =
  {
    // FreezeAccount is the address of the account whose asset
    // slot is being frozen or un-frozen.
    FreezeAccount: Address //`codec:"fadd"`
    // FreezeAsset is the asset ID being frozen or un-frozen.
    FreezeAsset: AssetIndex //`codec:"faid"`
    // AssetFrozen is the new frozen value.
    AssetFrozen: bool } //`codec:"afrz"`

type AppIndex = uint64
type OnCompletion = uint64

// StateSchema sets maximums on the number of each type that may be stored
type StateSchema =
  { NumUint: uint64 //`codec:"nui"`
    NumByteSlice: uint64 } //`codec:"nbs"`


// BoxReference names a box by the index in the foreign app array
type BoxReference =
  {
    // The index of the app in the foreign app array.
    ForeignAppIdx: uint64 // `codec:"i"`
    // The name of the box unique to the app it belongs to
    Name: byte array } //`codec:"n"`

// ApplicationCallTxnFields captures the transaction fields used for all
// interactions with applications
type ApplicationCallTxnFields =
  { ApplicationID: AppIndex //`codec:"apid"`
    OnCompletion: OnCompletion //`codec:"apan"`
    ApplicationArgs: byte array array //`codec:"apaa,allocbound=encodedMaxApplicationArgs,maxtotalbytes=config.MaxAppTotalArgLen"`
    Accounts: Address array //`codec:"apat,allocbound=encodedMaxAccounts"`
    ForeignApps: AppIndex array //`codec:"apfa,allocbound=encodedMaxForeignApps"`
    ForeignAssets: AssetIndex array //`codec:"apas,allocbound=encodedMaxForeignAssets"`
    BoxReferences: BoxReference array //`codec:"apbx,allocbound=encodedMaxBoxReferences"`

    LocalStateSchema: StateSchema //`codec:"apls"`
    GlobalStateSchema: StateSchema //`codec:"apgs"`
    ApprovalProgram: byte array //`codec:"apap"`
    ClearStateProgram: byte array //`codec:"apsu"`
    ExtraProgramPages: uint32 } //`codec:"apep"`

type ApplicationFields = ApplicationCallTxnFields


// MerkleSignature represents a Falcon signature in a compressed-form
//
//msgp:allocbound MerkleSignature FalconMaxSignatureSize
type MerkleSignature = byte array

// GenericDigest is a digest that implements CustomSizeDigest, and can be used as hash output.
//
//msgp:allocbound GenericDigest MaxHashDigestSize
type GenericDigest = byte array

// HashType represents different hash functions
type HashType = uint16

// HashFactory is responsible for generating new hashes accordingly to the type it stores.
//
//msgp:postunmarshalcheck HashFactory Validate
type HashFactory = { HashType: HashType } // `codec:"t"`

// Proof is used to convince a verifier about membership of leaves: h0,h1...hn
// at indexes i0,i1...in on a tree. The verifier has a trusted value of the tree
// root hash.
type Proof =
  {
    // Path is bounded by MaxNumLeavesOnEncodedTree since there could be multiple reveals, and
    // given the distribution of the elt positions and the depth of the tree,
    // the path length can increase up to 2^MaxEncodedTreeDepth / 2
    Path: GenericDigest array //`codec:"pth,allocbound=MaxNumLeavesOnEncodedTree/2"`
    HashFactory: HashFactory //`codec:"hsh"`
    // TreeDepth represents the depth of the tree that is being proven.
    // It is the number of edges from the root to a leaf.
    TreeDepth: uint8 } //`codec:"td"`

// SingleLeafProof is used to convince a verifier about membership of a specific
// leaf h at index i on a tree. The verifier has a trusted value of the tree
// root hash. it corresponds to merkle verification path.
type SingleLeafProof = Proof


// FalconSignatureStruct represents a signature in the merkle signature scheme using falcon signatures as an underlying crypto scheme.
// It consists of an ephemeral public key, a signature, a merkle verification path and an index.
// The merkle signature considered valid only if the Signature is verified under the ephemeral public key and
// the Merkle verification path verifies that the ephemeral public key is located at the given index of the tree
// (for the root given in the long-term public key).
// More details can be found on Algorand's spec
type FalconSignatureStruct =
  { Signature: MerkleSignature //`codec:"sig"`
    VectorCommitmentIndex: uint64 //`codec:"idx"`
    Proof: SingleLeafProof //`codec:"prf"`
    VerifyingKey: FalconVerifier } //`codec:"vkey"`

// A sigslotCommit is a single slot in the sigs array that forms the state proof.
type SigslotCommit =
  {
    // Sig is a signature by the participant on the expected message.
    Sig: FalconSignatureStruct //`codec:"s"`

    // L is the total weight of signatures in lower-numbered slots.
    // This is initialized once the builder has collected a sufficient
    // number of signatures.
    L: uint64 } //`codec:"l"`

// Commitment represents the root of the vector commitment tree built upon the MSS keys.
type Commitment = byte array // [MerkleSignatureSchemeRootSize]byte

// Verifier is used to verify a merklesignature.Signature produced by merklesignature.Secrets.
type Verifier =
  { Commitment: Commitment //`codec:"cmt"`
    KeyLifetime: uint64 } //`codec:"lf"`

// A Participant corresponds to an account whose AccountData.Status
// is Online, and for which the expected sigRound satisfies
// AccountData.VoteFirstValid <= sigRound <= AccountData.VoteLastValid.
//
// In the Algorand ledger, it is possible for multiple accounts to have
// the same PK.  Thus, the PK is not necessarily unique among Participants.
// However, each account will produce a unique Participant struct, to avoid
// potential DoS attacks where one account claims to have the same VoteID PK
// as another account.
type Participant =
  {
    // PK is the identifier used to verify the signature for a specific participant
    PK: Verifier //`codec:"p"`
    // Weight is AccountData.MicroAlgos.
    Weight: uint64 } //`codec:"w"`

// Reveal is a single array position revealed as part of a state
// proof.  It reveals an element of the signature array and
// the corresponding element of the participants array.
type Reveal =
  { SigSlot: SigslotCommit // `codec:"s"`
    Part: Participant } // `codec:"p"`


// StateProof represents a proof on Algorand's state.
type StateProof =
  { SigCommit: GenericDigest //`codec:"c"`
    SignedWeight: uint64 //`codec:"w"`
    SigProofs: Proof //`codec:"S"`
    PartProofs: Proof //`codec:"P"`
    MerkleSignatureSaltVersion: byte //`codec:"v"`
    // Reveals is a sparse map from the position being revealed
    // to the corresponding elements from the sigs and participants
    // arrays.
    Reveals: Map<uint64, Reveal> // `codec:"r,allocbound=MaxReveals"`
    PositionsToReveal: uint64 array } // `codec:"pr,allocbound=MaxReveals"`


// Message represents the message that the state proofs are attesting to. This message can be
// used by lightweight client and gives it the ability to verify proofs on the Algorand's state.
// In addition to that proof, this message also contains fields that
// are needed in order to verify the next state proofs (VotersCommitment and LnProvenWeight).
type Message =
  {
    // BlockHeadersCommitment contains a commitment on all light block headers within a state proof interval.
    BlockHeadersCommitment: byte array //`codec:"b,allocbound=Sha256Size"`
    VotersCommitment: byte array //`codec:"v,allocbound=SumhashDigestSize"`
    LnProvenWeight: uint64 //`codec:"P"`
    FirstAttestedRound: uint64 //`codec:"f"`
    LastAttestedRound: uint64 } //`codec:"l"`

// StateProofTxnFields captures the fields used for stateproof transactions.
type StateProofTxnFields =
  { StateProofType: StateProofType // `codec:"sptype"`
    StateProof: StateProof // `codec:"sp"`
    Message: Message } // `codec:"spmsg"`

type Transaction =
  { txType: TxType
    header: Header
    keyRegTxnFields: KeyregTxnFields
    paymentTxnFields: PaymentTxnFields
    assetConfigTxnFields: AssetConfigTxnFields
    assetTransferTxnFields: AssetTransferTxnFields
    assetFreezeTxnFields: AssetFreezeTxnFields
    applicationFields: ApplicationFields
    stateProofTxnFields: StateProofTxnFields }

// type Transaction = {
// 	_struct struct{} `codec:",omitempty,omitemptyarray"`

// 	// Type of transaction
// 	Type TxType `codec:"type"`

// 	// Common fields for all types of transactions
// 	Header

// 	// Fields for different types of transactions
// 	KeyregTxnFields
// 	PaymentTxnFields
// 	AssetConfigTxnFields
// 	AssetTransferTxnFields
// 	AssetFreezeTxnFields
// 	ApplicationFields
// 	StateProofTxnFields
// }

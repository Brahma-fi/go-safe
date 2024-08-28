# go-safe
go safe is a utility package written in go for [safe](https://github.com/safe-global/safe-smart-account) which has following features

### Gas estimation
gas-estimate package contains the utility functions for safe related gas estimations
- `EstimateSafeGas`: for a given safe transaction it finds the gas limit which can be used as safeTxGas to limit the amount of gas used during actual execution, which avoids transactions from using all the gas available in safe, it is compatible with safe version >= 1.3.0
- `EstimateSafeTransactionGasLimit`:  estimates the gas limit for a safe transaction. It uses a binary search algorithm to find the optimal gas limit for the transaction.  And finds the optimal gas at which the transaction succeeds by checking the return of the eth_call. if encoded callData is execTransactionFromModule or execTransaction and call is successful safe returns (true) boolean as return. can be used as drop in replacement for eth_estimateGas

### Decoders
decoder package offers utils functions to parse safe specific callData and values.

- `ParseExecTransactionCallData`: extracts encoded `execTransaction` call from the give calldata which can be unpacked into the `execTransaction` type. For example here the safe transaction is wrapped in a relay call, and hence to extract the actual `execTransaction` call this function can be used
- `ParseMultiSendData`: extracts individual safe multisend calls (to, value, operation, data) from the encoded safe transaction

```go
subData, _ := hexutil.Decode(  
    "0x8D80FF0A000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000CE0079D664098E9A97C40695522B9568A67EC752A26600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024BE300447000000000000000000000000000000000000000000000000006A94D74F430000000405D9D1443DFB051D5E8E231E41C911DC8393A40000000000000000000000000000000000000000000000000018838370F340000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",  
)  
calls, err := ParseMultiSendData(subData)
/*
[
  {
    "operation": 0,
    "to": "0x79d664098e9a97c40695522b9568a67ec752a266",
    "value": 0,
    "data": "0xbe300447000000000000000000000000000000000000000000000000006a94d74f430000"
  },
  {
    "operation": 0,
    "to": "0x0405d9d1443dfb051d5e8e231e41c911dc8393a4",
    "value": 6900000000000000,
    "data": "0x"
  }
]
*/
```

### Encoders
encoders offer safe related encoding utils functions.
- `GetEncodedSafeTx`: Gives the `types.SafeTx` from the given list of transactions. handles the multisend encoding internally in case of multiple transactions and the returned type can be used further for signing and gas estimations
- `GetSignedSafeTxn`: adds encoded signatures to the given `core.GnosisSafeTx`
- `GetSafeSignatureHash`: gets the EIP-712 hash to sign which is generated from the given `core.GnosisSafeTx`
- `GetEncodedExecTransaction`: encodes the give `core.GnosisSafeTx` into `execTransaction`
- `TightPack`: TightPack is the go variant of `abi.encodePacked` [solidity function](https://docs.soliditylang.org/en/latest/abi-spec.html#non-standard-packed-mode) which is a non-standard packed mode.
 ```go
import (  
    "github.com/Brahma-fi/go-safe/encoders"  
    "github.com/Brahma-fi/go-safe/types"  
    "github.com/ethereum/go-ethereum/common"  
    "github.com/ethereum/go-ethereum/core/types"  
    "github.com/ethereum/go-ethereum/crypto"  
    "github.com/ethereum/go-ethereum/rlp"  
    "github.com/ethereum/go-ethereum/common/hexutil"  
    "math/big"  
)

func TestNewPackBuilder() {  
    builder := NewPackBuilder()  
    val := new(big.Int).SetInt64(0)  
    data, _ := hexutil.Decode("0xa9059cbb000000000000000000000000c20e3a951e03edd9e17c027324a07d2841cdf983000000000000000000000000000000000000000000000000016345785d8a0000")  
    builder.AddAddress(common.HexToAddress("0x02ABBDbAaa7b1BB64B5c878f7ac17f8DDa169532"))  
    builder.AddUint256(val)  
    builder.AddBytes(data)  
    builder.AddUint8(0)  
    builder.AddAddress(common.HexToAddress("0x4fadb18339be16d57f27a8a65d78ffecf29037d0"))  
    builder.AddUint256(val)  
    builder.AddBytes(common.HexToHash("0x954c8850b002325bd02c8cffb07b6226ff77b3bd16f865827a63ca510ff1559a").Bytes())  
    builder.AddUint32(1694514609)  
    packed, _ := builder.Pack()  
    fmt.Println(hexutil.Encode(packed))  
    // 0x02abbdbaaa7b1bb64b5c878f7ac17f8dda1695320000000000000000000000000000000000000000000000000000000000000000a9059cbb000000000000000000000000c20e3a951e03edd9e17c027324a07d2841cdf983000000000000000000000000000000000000000000000000016345785d8a0000004fadb18339be16d57f27a8a65d78ffecf29037d00000000000000000000000000000000000000000000000000000000000000000954c8850b002325bd02c8cffb07b6226ff77b3bd16f865827a63ca510ff1559a65003db1  
}
 ```

### Safe metadata
metadata package contains the utility functions for fetching safe related metadata like owners, version etc.

- `GetSafeMetadata`: fetch metadata of multiple safes using a multicall, returns version, owners, threshold, nonce and other fields
- Other independent util functions such as `IsValidOwner`, `GetTransactionHash` 


## Installation

To install go-safe, install it as a go module in the project

> go get github.com/Brahma-fi/go-safe

## License

This project is licensed under the [LGPL-3.0](LICENSE).
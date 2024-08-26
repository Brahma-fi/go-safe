gen-abis:
	abigen --abi=./contracts/abis/safe.json --pkg=safe --out=contracts/safe/safe.go
	abigen --abi=./contracts/abis/wallet_registry.json --pkg=walletregistry --out=contracts/walletregistry/wallet_registry.go
	abigen --abi=./contracts/abis/multicall3.json --pkg=multicall --out=contracts/multicall/multicall.go
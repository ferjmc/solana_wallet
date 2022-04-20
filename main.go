package main

import (
	"fmt"
	"solana_wallet/wallet"

	"github.com/portto/solana-go-sdk/rpc"
)

func main() {
	// create a new wallet
	// fernandoWallet := wallet.CreateNewWallet(rpc.DevnetRPCEndpoint)
	// // display the wallet public and private keys
	// fmt.Println("Wallet Address:", fernandoWallet.Account.PublicKey.ToBase58())
	// fmt.Println("Private Key:", fernandoWallet.Account.PrivateKey)
	// // guardo la direccion y borro el cliente
	// direccionFernando := fernandoWallet.Account.PublicKey.ToBase58()
	//fernandoWallet = Wallet{}

	// create a new wallet
	//alejandroWallet := wallet.CreateNewWallet(rpc.DevnetRPCEndpoint)

	// display the wallet public and private keys
	//fmt.Println("Wallet Address:", alejandroWallet.Account.PublicKey.ToBase58())
	//fmt.Println("Private Key:", alejandroWallet.Account.PrivateKey)

	// request for an airdrop
	//fmt.Println(alejandroWallet.RequestAirdrop(1e9))

	// fetch wallet balance
	//fmt.Println(alejandroWallet.GetBalance())

	// fmt.Println(fernandoWallet.GetBalance())

	// // make transfer to fernando
	// fmt.Println(alejandroWallet.Transfer(direccionFernando, 5e8))
	// time.Sleep(time.Second * 20)

	// // fetch wallet balance
	// fmt.Println(alejandroWallet.GetBalance())

	// fmt.Println(fernandoWallet.GetBalance())

	//fmt.Println(nft.NewNft(alejandroWallet.Account))

	// Wallet Address: 2UbvqdqrbAfRaaJpg3P3CREVpRvRKYKRkNNd8D4m5T1D
	// NFT: 8mCsUgzy4qSPBihST8y6iFMR9bQjZmp1KiTgpW4KScnh

	//
	frase := []string{"believe", "deal", "spell", "maid", "emotion", "liberty", "fine", "obtain", "concert", "minor", "wonder", "permit"}
	// // "believe deal spell maid emotion liberty bad obtain concert minor wonder permit"
	wallets := wallet.NewFromMnemonic(frase, rpc.DevnetRPCEndpoint)
	// // wallets2 := wallet.NewFromMnemonic2(frase, rpc.DevnetRPCEndpoint)

	// res, err := wallets.RequestAirdrop(1e9)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)
	//time.Sleep(time.Second * 20)

	balance, err := wallets.GetBalance()
	if err != nil {
		panic(err)
	}
	fmt.Println(balance)

	// res, err = wallets2.RequestAirdrop(1e9)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)
	// //time.Sleep(time.Second * 20)

	// balance, err = wallets2.GetBalance()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(balance)

	//6RWYChT1CnE24gH7jnqxAypmgf3gjtiyF5UwdYgNQsGB

	// w := wallet.NewFromKey("ieoK6wTcHgN92dGSQppYaQQXBHGfazG8gvkXbfdhBZtgxpUgCMUJVQemdEyeGKGoQBmYfkS2wAfyFxo4b3RL7Xd", rpc.TestnetRPCEndpoint)
	// fmt.Println("from key:", w.Account.PublicKey.ToBase58())

	// fmt.Println(w.RequestAirdrop(1e9))

	// // fetch wallet balance
	// fmt.Println(w.GetBalance())

	// // mint NFT
	// fmt.Println(nft.NewNft(w.Account))
}

package wallet

import (
	"context"
	"fmt"
	"strings"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/hdwallet"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/types"
	"github.com/tyler-smith/go-bip39"
)

type Wallet struct {
	Account types.Account
	c       *client.Client
}

func CreateNewWallet(RPCEndpoint string) Wallet {
	return Wallet{
		types.NewAccount(),
		client.NewClient(RPCEndpoint),
	}
}

func (w Wallet) RequestAirdrop(amount uint64) (string, error) {
	// request for SOL using RequestAirdrop()
	txhash, err := w.c.RequestAirdrop(
		context.TODO(),                 // request context
		w.Account.PublicKey.ToBase58(), // wallet address requesting airdrop
		amount,                         // amount of SOL in lamport
	)
	if err != nil {
		return "", err
	}

	return txhash, nil
}

func (w Wallet) GetBalance() (uint64, error) {
	// fetch the balance using GetBalance()
	balance, err := w.c.GetBalance(
		context.TODO(),                 // request context
		w.Account.PublicKey.ToBase58(), // wallet to fetch balance for
	)
	if err != nil {
		return 0, nil
	}

	return balance, nil
}

func (w Wallet) Transfer(receiver string, amount uint64) (string, error) {
	// fetch the most recent blockhash
	response, err := w.c.GetRecentBlockhash(context.TODO())
	if err != nil {
		return "", err
	}

	// make a transfer message with the latest block hash
	message := types.NewMessage(
		types.NewMessageParam{
			w.Account.PublicKey, // public key of the transaction signer
			[]types.Instruction{
				sysprog.Transfer(
					sysprog.TransferParam{
						w.Account.PublicKey,                  // public key of the transaction sender
						common.PublicKeyFromString(receiver), // wallet address of the transaction receiver
						amount,                               // transaction amount in lamport
					},
				),
			},
			response.Blockhash, // recent block hash
		},
	)

	// create a transaction with the message and TX signer
	tx, err := types.NewTransaction(
		types.NewTransactionParam{
			message,
			[]types.Account{w.Account, w.Account},
		},
	)
	if err != nil {
		return "", err
	}

	// send the transaction to the blockchain
	txhash, err := w.c.SendTransaction(context.TODO(), tx)
	if err != nil {
		return "", err
	}

	return txhash, nil
}

func NewFromMnemonic(frase []string, RPCEndpoint string) Wallet {
	mnemonic := strings.Join(frase, " ")

	seed := bip39.NewSeed(mnemonic, "")
	path := `m/44'/501'/0'/0'`
	derivedKey, _ := hdwallet.Derived(path, seed)
	account, _ := types.AccountFromSeed(derivedKey.PrivateKey)
	fmt.Printf("%v => %v\n", path, account.PublicKey.ToBase58())
	//6RWYChT1CnE24gH7jnqxAypmgf3gjtiyF5UwdYgNQsGB

	// for i := 1; i < 100; i++ {
	// 	path := fmt.Sprintf(`m/44'/501'/%d'/0'`, i)
	// 	derivedKey, _ := hdwallet.Derived(path, seed)
	// 	account, _ := types.AccountFromSeed(derivedKey.PrivateKey)
	// 	fmt.Printf("%v => %v\n", path, account.PublicKey.ToBase58())
	// 	if account.PublicKey.ToBase58() == "6RWYChT1CnE24gH7jnqxAypmgf3gjtiyF5UwdYgNQsGB" {
	// 		fmt.Println("EUREKA EN ", i)
	// 		panic("FIN")
	// 	}
	// }

	// 6RWYChT1CnE24gH7jnqxAypmgf3gjtiyF5UwdYgNQsGB
	return Wallet{
		Account: account,
		c:       client.NewClient(RPCEndpoint),
	}
}

func NewFromKey(key string, RCPEndpoint string) Wallet {
	account, err := types.AccountFromBase58(key)
	if err != nil {
		panic(err)
	}
	fmt.Println("account: ", account)
	return Wallet{
		Account: account,
		c:       client.NewClient(RCPEndpoint),
	}
}

// func NewFromMnemonic2(frase []string, RPCEndpoint string) Wallet {
// 	mnemonic := strings.Join(frase, " ")
// 	seed := bip39.NewSeed(mnemonic, "F3rn4nd0") // (mnemonic, password)
// 	account, _ := types.AccountFromSeed(seed[:32])
// 	fmt.Println("account2: ", account.PublicKey.ToBase58())
// 	return Wallet{
// 		Account: account,
// 		c:       client.NewClient(RPCEndpoint),
// 	}
// }

// ieoK6wTcHgN92dGSQppYaQQXBHGfazG8gvkXbfdhBZtgxpUgCMUJVQemdEyeGKGoQBmYfkS2wAfyFxo4b3RL7Xd

package wallet

import (
	"context"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/types"
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

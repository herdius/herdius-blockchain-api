package cmd

import (
	"log"
	"strconv"

	"github.com/herdius/herdius-blockchain-api/cmd/test/testutils"
	"github.com/herdius/herdius-core/crypto/secp256k1"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(transactionCmd)
}

var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Execute to tx to new accounts, Balance will be transfered from master account",
	Run: func(cmd *cobra.Command, args []string) {
		// args[0] represents count
		count, err := strconv.Atoi(args[0])
		if err != nil {
			return
		}
		senderPrivKey, _ := testutils.GetHERKeyPair(masterAccountKey)
		cmdutil := testutils.New(endpoint)
		account, err := cmdutil.GetAccount(senderPrivKey.PubKey().GetAddress())
		if err != nil {
			log.Println("Erront getting account detail")
		}
		log.Println("Last Nonce", account.Nonce)
		nonce := uint64(0)
		if account.Nonce > uint64(0) {
			nonce = account.Nonce + 1
		}
		receiver := secp256k1.GenPrivKey()

		for i := 0; i <= count; i++ {
			tr, err := cmdutil.ToTx(senderPrivKey, accountType, nonce, 1, receiver.PubKey().GetAddress())
			if err != nil {
				log.Println("Error getting account detail")
			}
			log.Println("Tx executed", tr)
			nonce++
		}
	},
}

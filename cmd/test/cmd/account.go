package cmd

import (
	"log"

	"github.com/herdius/herdius-blockchain-api/cmd/test/testutils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(accountCmd)
	accountCmd.AddCommand(addressCmd)

}

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Register account  HER/BTC/ETH and Herdius supported assets",
	Run: func(cmd *cobra.Command, args []string) {
		senderPrivKey, _ := testutils.GetHERKeyPair(accountSecret)
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

		cmdutil.RegisterAccount(senderPrivKey, accountType, nonce, "") //ETH/BTC
	},
}

var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "Get Address using secret key",
	Run: func(cmd *cobra.Command, args []string) {
		senderPrivKey, _ := testutils.GetHERKeyPair(accountSecret)
		log.Println("Address: ", senderPrivKey.PubKey().GetAddress())

	},
}

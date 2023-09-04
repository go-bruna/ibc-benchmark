package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	OSMO_RECEIVER = "osmo1cr00glwjvf7fzv72tfdj4x3ctusq6uu35wp9gx"
	CHAIN_ID_OSMO = "osmo-test-5"
	CHAIN_ID_INCO = "evmos_9000-1"

	DENOM_INCO = "aevmos"
	DENOM_OSMO = "uosmo"

	OSMOS_BINARY = "osmosisd"
	INCO_BINARY  = "evmosd"

	OSMOS_RPC = "https://osmosis-testnet-rpc.polkachu.com:443/"
	INCO_RPC  = "http://16.170.148.40:26657"

	OSMOS_API = "https://osmosis-testnet-rpc.polkachu.com/"
	INCO_API  = "http://16.170.148.40:1317"
)

func IBCTransfer(destChain string, srcChain string, srcChannel string, amount string, receiver string, denom string) (string, error) {
	// hermes tx ft-transfer --timeout-seconds 1000 --dst-chain osmo-test-5 --src-chain evmos_9000-1 --src-port transfer --src-channel channel-0 --amount 10 --denom aevmos --receiver osmo1cr00glwjvf7fzv72tfdj4x3ctusq6uu35wp9gx
	cmdNew := exec.Command("hermes", "tx", "ft-transfer", "--timeout-seconds", "1000", "--dst-chain", destChain, "--src-chain", srcChain, "--src-port", "transfer", "--src-channel", srcChannel, "--amount", amount, "--denom", denom, "--receiver", receiver)
	var errbuf bytes.Buffer
	cmdNew.Stderr = &errbuf

	stdout, err := cmdNew.Output()
	if err != nil {
		return "", fmt.Errorf("error while executing request-transaction: %s: %s", err, errbuf.String())
	}

	stdoutStr := string(stdout)
	fmt.Println(stdoutStr)

	return stdoutStr, nil
}

// QueryBalance returns the balance
func QueryBalance(chainBinary string, address string, rpc string) {
	cmdNew := exec.Command(chainBinary, "q", "bank", "balances", address, "--output", "json", "--node", rpc)
	var errbuf bytes.Buffer
	cmdNew.Stderr = &errbuf

	stdout, err := cmdNew.Output()
	if err != nil {
		fmt.Println(err)

		return
	}

	var response types.QueryBalanceResponse
	if err := json.Unmarshal(stdout, &response); err != nil {
		return
	}

	fmt.Println(response.Balance)
}

func main() {

	// Send 10aevmos from Inco to Osmosis testnet
	// IBCTransfer(CHAIN_ID_OSMO, CHAIN_ID_INCO, "channel-1", "10", OSMO_RECEIVER, DENOM_INCO)

	QueryBalance(OSMOS_BINARY, OSMO_RECEIVER, OSMOS_RPC)
}

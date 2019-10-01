package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"./contracts"

	"github.com/DaveAppleton/etherUtils"
	"github.com/DaveAppleton/memorykeys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

var baseClient *backends.SimulatedBackend

func getClient() (client *backends.SimulatedBackend, err error) {
	if baseClient != nil {
		return baseClient, nil
	}
	funds, _ := etherUtils.StrToEther("10000.0")
	bankerAddress, err := memorykeys.GetAddress("banker")
	if err != nil {
		return nil, err
	}
	baseClient = backends.NewSimulatedBackend(core.GenesisAlloc{
		*bankerAddress: {Balance: funds},
	}, 8000000)
	return baseClient, nil
}

func chkerr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getTime(dateStr string) *big.Int {
	t, err := time.Parse("02/01/06", dateStr)
	chkerr(err)
	return big.NewInt(t.Unix())
}

func currentTime() uint64 {
	client, err := getClient()
	chkerr(err)
	block := client.Blockchain().CurrentBlock()
	return block.Time()
}

func jumpTo(newTime *big.Int) {
	client, err := getClient()
	chkerr(err)
	now := client.Blockchain().CurrentBlock().Time()
	target := newTime.Uint64()
	if now >= target {
		return
	}
	err = client.AdjustTime(time.Duration(target-now) * time.Second)
	chkerr(err)
	client.Commit()
}

func isBiddingOpen(auction *contracts.Auction) {
	biddingOpen, err := auction.InBidding(nil)
	chkerr(err)
	state := "IS NOT"
	if biddingOpen {
		state = "IS"
	}
	fmt.Println("Bidding", state, "open")
}

func main() {
	client, err := getClient()
	if err != nil {
		log.Fatal(err)
	}
	startBids := getTime("13/09/19")
	endBids := getTime("15/09/19")
	startReveal := endBids
	endReveal := getTime("17/09/19")
	minimumBid, _ := etherUtils.StrToEther("4.7")
	wallet, _ := memorykeys.GetAddress("wallet")
	bankTx, _ := memorykeys.GetTransactor("banker")
	txes := make(map[string]*types.Transaction)
	for user := 0; user < 100; user++ {
		userID := fmt.Sprintf("user%02d", user)
		txes[userID], _ = sendEth(client, "banker", userID, new(big.Int).Add(minimumBid, etherUtils.OneEther()), 21000)
	}
	auctionAddress, tx, auctionContract, err := contracts.DeployAuction(bankTx, client, startBids, endBids, startReveal, endReveal, minimumBid, *wallet)
	chkerr(err)
	fmt.Println(auctionAddress.Hex(), tx.Hash().Hex())
	client.Commit()
	min, err := auctionContract.MinimumBid(nil)
	chkerr(err)
	fmt.Println("minumum bid is ", etherUtils.EtherToStr(min))
	fmt.Println("time :", currentTime())
	isBiddingOpen(auctionContract)
	jumpTo(startBids)
	fmt.Println("time :", currentTime())
	isBiddingOpen(auctionContract)
	userTx, _ := memorykeys.GetTransactor("user01")
	hash, err := auctionContract.CalculateHash(nil, minimumBid, []byte("mary had a little lamb"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	userTx.Value = minimumBid
	tx, err = auctionContract.BiddingTime(userTx, hash)
	if err != nil {
		fmt.Println("bid : ", err)
		os.Exit(1)
	}
	client.Commit()
}

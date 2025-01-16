package src

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-eth/models"
	"log"
)

func GetBalanceForAddress(client *ethclient.Client, request *models.BalanceRequest) (*models.BalanceResponse, error) {
	account := common.HexToAddress(request.Address)
	ethBalance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	balances := make(map[string]float64)
	balances["ETH"], _ = BigIntToEthBigFloat(ethBalance, 18).Float64()
	return &models.BalanceResponse{
		Balances: balances,
	}, nil
}

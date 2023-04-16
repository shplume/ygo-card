package service

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"

	"github.com/shplume/ygo-cards/model"
)

func load() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("!!!!!!!!!!!!!!!! Service init err: %s\n", err.Error())
		return
	}

	jsonDir := path.Join(dir, model.JSONDir)
	if _, err := os.Stat(jsonDir); err != nil {
		fmt.Printf("!!!!!!!!!!!!!!!! Service init err: %s\n", err.Error())
		return
	}

	loadCards(jsonDir)

	if len(cards) == 0 {
		fmt.Printf("!!!!!!!!!!!!!!!! Service data load failed\n")
	}

	cardsInfo := make([]model.CardInfo, 0, len(cards))
	storeCardsInfo(cardsInfo)
}

func loadCards(dir string) {
	file, err := os.Open(path.Join(dir, model.CardsFile))
	if err != nil {
		fmt.Printf("Load Cards with error: %s\n", err.Error())
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Load Cards with error: %s\n", err.Error())
		return
	}

	if err := json.Unmarshal(content, &cards); err != nil {
		fmt.Printf("Load Cards with error: %s\n", err.Error())
		return
	}

	fmt.Println("------------Cards Loaded!")
}

func storeCardsInfo(cardsInfo []model.CardInfo) {
	for k, v := range cards {
		id, _ := strconv.Atoi(k)

		b, err := json.Marshal(v)
		if err != nil {
			fmt.Printf("Store Cards(marshal json) with error: %s\n", err.Error())
		}

		cardsInfo = append(cardsInfo, model.CardInfo{
			ID:      uint(id),
			Details: b,
		})
	}

	fmt.Printf("------------Len of data: %d", len(cardsInfo))

	if err := dbConn.Model(&model.CardInfo{}).Create(cardsInfo).Error; err != nil {
		fmt.Printf("Store Cards with error: %s\n", err.Error())
	}
}

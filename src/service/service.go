package service

import (
	"fmt"
	"strconv"
	model2 "test2/src/model"
)

var Soldiers map[string]model2.Soldier

// MatchSoldier 匹配id相同的soldier信息
func MatchSoldier(id string) model2.Soldier {

	return Soldiers[id]
}

// GetRarityUnlockArenaSoldier 输入稀有度，当前解锁阶段，获取该稀有度且已解锁的所有士兵
func GetRarityUnlockArenaSoldier(rarity int, unlockarena int) map[string]model2.Soldier {
	soldier := make(map[string]model2.Soldier)

	for i, s := range Soldiers {
		idn, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Get Error")
		}
		if s.Rarity == rarity && s.UnlockArena >= unlockarena {
			soldier[strconv.Itoa(idn)] = s
		}
	}
	return soldier
}

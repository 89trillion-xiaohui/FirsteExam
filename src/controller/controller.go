package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	service2 "test2/src/service"
)

type Controller struct {
}

// GetRarity 输入士兵id获取稀有度
func (soldierController *Controller) GetRarity(c *gin.Context) {
	id := c.Query("id")
	soldier := service2.MatchSoldier(id)
	c.JSON(http.StatusOK, soldier.Rarity)
}

// GetCombatPoints 输入士兵id获取战力
func (soldierController *Controller) GetCombatPoints(c *gin.Context) {
	id := c.Query("id")
	soldier := service2.MatchSoldier(id)
	c.JSON(http.StatusOK, soldier.CombatPoints)
}

// GetRarityUnlockArena 输入稀有度，当前解锁阶段,获取该稀有度且已解锁的所有士兵
func (soldierController *Controller) GetRarityUnlockArena(c *gin.Context) {
	rarity, _ := strconv.Atoi(c.Query("rarity"))
	unlockArena, _ := strconv.Atoi(c.Query("unlockArena"))
	soldier := service2.GetRarityUnlockArenaSoldier(rarity, unlockArena)
	c.JSON(http.StatusOK, soldier)
}

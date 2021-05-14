package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
	"test2/src/controller"
	"test2/src/model"
	"test2/src/service"
)

//var jsonPath *string	//存储json路径

//var Config map[string] model.Soldier // map存储配置文件

func main() {
	// 解析ini文件
	cfg, err := ini.Load("./src/app.ini")
	if err != nil {
		fmt.Printf("Failed to read file : %v", err)
		os.Exit(1)
	}
	//获取客户端端口
	httpPort := cfg.Section("server").Key("HttpPort").Value()

	JsonInit()

	r := gin.Default()
	soldierController := controller.Controller{}

	r.GET("/rarity", soldierController.GetRarity)

	r.GET("/combatPoints", soldierController.GetCombatPoints)

	r.GET("/rarityAndUnlockArena", soldierController.GetRarityUnlockArena)

	r.Run(":" + httpPort)
}

func readJson(jsonFilePath string) (map[string]model.Soldier, error) {
	var JsonSoldier map[string]model.Soldier

	//读取json文件
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Println("Fail to Open : %", err)
		return nil, err
	}
	defer jsonFile.Close()
	fmt.Println("—————————json文件内容为————————")

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&JsonSoldier)
	if err != nil {
		fmt.Println("Decode failed", err.Error())
	} else {
		fmt.Println("Decode success")
		fmt.Println(JsonSoldier)

	}

	fmt.Println("json文件转到内存")

	// 打印对象结构
	fmt.Println("JsonSoldier= ", JsonSoldier)
	service.Soldiers = JsonSoldier

	// 将对象结构转换为json格式
	newBytes, _ := json.Marshal(&JsonSoldier)
	fmt.Println("————————json文件转换完毕————————")

	// 打印JSON结果
	fmt.Println(string(newBytes))
	return JsonSoldier, nil
}
func JsonInit() {
	jsonSoldier, err := readJson("./src/config.army.model.json")

	//创建新的json文件
	file, err := os.Create("./src/newFile.json")
	if err != nil {
		fmt.Println("创建失败", err.Error())
	}
	defer file.Close()

	//创建json编码器
	encoder := json.NewEncoder(file)

	err = encoder.Encode(jsonSoldier)
	if err != nil {
		fmt.Println("Encoder failed", err.Error())

	} else {
		fmt.Println("Encoder success")
	}
}

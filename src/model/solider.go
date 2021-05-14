package model

type Soldier struct {
	Id           string `json:"id"` //士兵id
	Name         string //`json:"Name"` //名字
	Note         string //`json:"Note"` //注释
	Rarity       int    `json:"Rarity,string,omitempty"`       //士兵稀有度
	UnlockArena  int    `json:"UnlockArena,string,omitempty"`  //解锁阶段
	CombatPoints int    `json:"CombatPoints,string,omitempty"` //战力

}

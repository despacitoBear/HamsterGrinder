package main

import "time"

type Boost struct {
	ID            string `json:"id"`
	Level         int    `json:"level"`
	LastUpgradeAt int64  `json:"lastUpgradeAt"`
}

type Upgrade struct {
	ID                     string `json:"id"`
	Level                  int    `json:"level"`
	LastUpgradeAt          int64  `json:"lastUpgradeAt"`
	SnapshotReferralsCount int    `json:"snapshotReferralsCount"`
}

type Task struct {
	ID          string    `json:"id"`
	CompletedAt time.Time `json:"completedAt"`
	Days        int       `json:"days"`
}

type Skin struct {
	SkinID string    `json:"skinId"`
	BuyAt  time.Time `json:"buyAt"`
}

type SkinDetails struct {
	Available      []Skin `json:"available"`
	SelectedSkinID string `json:"selectedSkinId"`
}

type ClickerUser struct {
	ID                 string                 `json:"id"`
	TotalCoins         float64                `json:"totalCoins"`
	BalanceCoins       float64                `json:"balanceCoins"`
	Level              int                    `json:"level"`
	AvailableTaps      int                    `json:"availableTaps"`
	LastSyncUpdate     int64                  `json:"lastSyncUpdate"`
	ExchangeID         string                 `json:"exchangeId"`
	Boosts             map[string]Boost       `json:"boosts"`
	Upgrades           map[string]Upgrade     `json:"upgrades"`
	Tasks              map[string]Task        `json:"tasks"`
	AirdropTasks       map[string]interface{} `json:"airdropTasks"`
	ReferralsCount     int                    `json:"referralsCount"`
	MaxTaps            int                    `json:"maxTaps"`
	EarnPerTap         int                    `json:"earnPerTap"`
	EarnPassivePerSec  float64                `json:"earnPassivePerSec"`
	EarnPassivePerHour float64                `json:"earnPassivePerHour"`
	LastPassiveEarn    float64                `json:"lastPassiveEarn"`
	TapsRecoverPerSec  int                    `json:"tapsRecoverPerSec"`
	CreatedAt          time.Time              `json:"createdAt"`
	Skin               SkinDetails            `json:"skin"`
}

type Response struct {
	ClickerUser ClickerUser `json:"clickerUser"`
}

type Payload struct {
	Count         int64 `json:"count"`
	AvailableTaps int64 `json:"availableTaps"`
	Timestamp     int64 `json:"timestamp"`
}

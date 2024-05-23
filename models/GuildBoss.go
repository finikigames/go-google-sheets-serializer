package models

type GuildBoss struct {
    Id string `json:"id"`
    SetupId string `json:"setup_id"`
    HealthDenominationCount int `json:"health_denomination_count"`
}

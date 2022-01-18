package Model

import "time"

type PageResult struct{
	Total int
	PageSize int
	PageIndex int
	Result interface{}
}

type BroadCast struct {
	Id int64 `gorm:"column:id" json:"id"`
	GameId int32 `gorm:"column:game_id" json:"gameId"`
	GameVersion string `gorm:"colum:game_version" json:"gameVersion"`
	PartnerId string `gorm:"colum:partner_id" json:"partnerId"`
	WorldId string `gorm:"colum:world_id" json:"worldId"`
	ChannelId string `gorm:"colum:channel_id" json:"channelId"`
	BroadcastPosition string `gorm:"colum:broadcast_position" json:"broadcastPosition"`
	BroadcastText string `gorm:"colum:broadcast_text" json:"broadcastText"`
	startTime time.Time `gorm:"colum:start_time" json:"startTime"`
	EndTime time.Time`gorm:"colum:end_time" json:"endTime"`
}
func (BroadCast) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "broadcastinfo"
}
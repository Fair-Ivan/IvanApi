package Model

import "time"

// 响应数据
type PageResult struct {
	Total     int         `json:"total"`     // 总条数
	PageSize  int         `json:"pageSize"`  // 页码
	PageIndex int         `json:"pageIndex"` // 页签
	Result    interface{} `json:"result"`    // 结果
}

type BroadCast struct {
	Id                int64     `gorm:"column:id" json:"id"`
	GameId            int32     `gorm:"column:game_id" json:"gameId"`
	GameVersion       string    `gorm:"column:game_version" json:"gameVersion"`
	PartnerId         string    `gorm:"column:partner_id" json:"partnerId"`
	WorldId           string    `gorm:"column:world_id" json:"worldId"`
	ChannelId         string    `gorm:"column:channel_id" json:"channelId"`
	BroadcastPosition string    `gorm:"column:broadcast_position" json:"broadcastPosition"`
	BroadcastText     string    `gorm:"column:broadcast_text" json:"broadcastText"`
	startTime         time.Time `gorm:"column:start_time" json:"startTime"`
	EndTime           time.Time `gorm:"column:end_time" json:"endTime"`
}

func (BroadCast) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "broadcastinfo"
}

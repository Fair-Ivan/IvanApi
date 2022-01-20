package Model

import "time"

type BroadCast struct {
	Id                int64     `gorm:"column:id" json:"id"`
	GameId            int       `gorm:"column:game_id" json:"gameId"`
	GameVersion       string    `gorm:"column:game_version" json:"gameVersion"`
	PartnerId         string    `gorm:"column:partner_id" json:"partnerId"`
	WorldId           string    `gorm:"column:world_id" json:"worldId"`
	ChannelId         string    `gorm:"column:channel_id" json:"channelId"`
	BroadcastPosition string    `gorm:"column:broadcast_position" json:"broadcastPosition"`
	BroadcastText     string    `gorm:"column:broadcast_text" json:"broadcastText"`
	StartTime         time.Time `gorm:"column:start_time" json:"startTime"`
	EndTime           time.Time `gorm:"column:end_time" json:"endTime"`
	Frequency         int       `gorm:"column:frequency" json:"frequency"`
	Status            int       `gorm:"column:status" json:"status"`
	JobId             int       `gorm:"column:job_id" json:"jobId"`
	TimingType        int16     `gorm:"column:timing_type" json:"timingType"`
	Sections          string    `gorm:"column:sections" json:"sections"`
	IsDeleted         int       `gorm:"column:is_deleted" json:"isDeleted"`
}

func (BroadCast) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "broadcastinfo"
}

type BroadCastUpdateInput struct {
	Id                int64     `form:"id`
	GameId            int32     `form:"gameId"`
	GameVersion       string    `form:"gameVersion"`
	PartnerId         string    `form:"partnerId"`
	WorldId           string    `form:"worldId"`
	ChannelId         string    `form:"channelId"`
	BroadcastPosition string    `form:"broadcastPosition"`
	BroadcastText     string    `form:"broadcastText"`
	StartTime         time.Time `form:"startTime"`
	EndTime           time.Time `form:"endTime"`
}

type BroadCastRemoveInput struct {
	Id int64 `form:"id"`
}

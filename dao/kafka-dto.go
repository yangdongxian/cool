package dao

type CollectQuery struct {
	Topic string `json:"topic" binding:"required"`
	Value string `json:"value" binding:"required"`
}

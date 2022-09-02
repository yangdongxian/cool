package dao

type QueryParameter struct {
	Env    string `json:"env" binding:"required"`
	DbName string `json:"db_name" binding:"required"`
	Sql    string `json:"sql" binding:"required"`
}

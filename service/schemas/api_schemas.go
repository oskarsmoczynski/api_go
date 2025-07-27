package schemas

type ReadTableSchema struct {
	Table   string   `json:"table" binding:"required"`
	Filters []string `json:"filters"`
	OrderBy []string `json:"order_by"`
	Limit   uint16   `json:"limit"`
}

type CreateEntrySchema struct {
	Table  string                   `json:"table" binding:"required"`
	Values []map[string]interface{} `json:"values" binding:"required"`
}

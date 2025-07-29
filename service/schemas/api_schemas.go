package schemas

type ReadTableSchema struct {
	Table   string   `json:"table" binding:"required"`
	Filters []string `json:"filters"`
	OrderBy []string `json:"order_by"`
	Limit   uint16   `json:"limit"`
}

type CreateEntrySchema struct {
	Table  string           `json:"table" binding:"required"`
	Values []map[string]any `json:"values" binding:"required"`
}

type UpdateEntrySchema struct {
	Table   string   `json:"table" binding:"required"`
	Values  []string `json:"values" binding:"required"`
	Filters []string `json:"filters" binding:"required"`
}

type DeleteEntrySchema struct {
	Table   string   `json:"table" binding:"required"`
	Filters []string `json:"filters" binding:"required"`
}

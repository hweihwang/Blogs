package common

type SQLModel struct {
	ID        uint     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Status    uint       `json:"status" gorm:"column:status;index:idx_status"`
	CreatedAt *DateTime  `json:"created_at" gorm:"column:created_at;index:idx_created_at"`
	UpdatedAt *DateTime `json:"updated_at" gorm:"column:updated_at;index:idx_updated_at"`
}

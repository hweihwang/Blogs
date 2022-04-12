package blogmodel

type LikeFilter string

type Likely interface {
	GetLikeString() string
}

type Filter struct {
	Title       *LikeFilter `json:"title,omitempty" form:"title"`
	Description *LikeFilter `json:"description,omitempty" form:"description"`
	CreateById  *uint       `json:"create_by_id,omitempty" form:"create_by_id"`
}

func (l LikeFilter) GetLikeString() string {
	return "%" + string(l) + "%"
}

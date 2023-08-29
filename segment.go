package avito

type Segment struct {
	Id       int    `json:"-"`
	Seg_name string `json:"seg_name" binding:"required"`
	Chance   int    `json:"chance"`
}

package avito

type Segment struct {
	Seg_name string `json:"seg_name" binding:"required"`
	Chance   int    `json:"chance"`
}

type SegmentRemove struct {
	Seg_name string
}

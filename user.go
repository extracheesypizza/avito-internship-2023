package avito

type User struct {
	Id        int      `json:"id" binding:"required"`
	Seg_names []string `json:"seg_names"`
	TTL       int      `json:"TTL"`
	Year      int      `json:"year"`
	Month     int      `json:"month"`
}

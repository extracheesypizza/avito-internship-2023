package avito

type User struct {
	Id        int      `json:"id" binding:"required"`
	Seg_names []string `json:"seg_names"`
	TTL       int      `json:"TTL"`
	Year      int      `json:"year"`
	Month     int      `json:"month"`
}

type UserAddToSegment struct {
	Id        int
	Seg_names []string
	TTL       int
}

type UserRemoveFromSegment struct {
	Id        int
	Seg_names []string
}

type UserGetActions struct {
	Id    int
	Year  int
	Month int
}

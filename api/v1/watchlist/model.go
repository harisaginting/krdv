package watchlist

type RequestWatchlist struct {
	Name string `json:"name"`
	Note string `json:"note"`
}

type ResponseList struct {
	Items []Watchlist `json:"items"`
	Total int         `json:"total"`
}

type Watchlist struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TotalItem int    `json:"total_item"`
}

type WatchlistDetail struct {
	ID    int              `json:"-" gorm:"column:id"`
	Name  string           `json:"name"`
	Note  string           `json:"note"`
	Movie []WatchlistMovie `json:"movies" gorm:"-"`
}

type WatchlistMovie struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Revenue       string  `json:"revenue"`
	Public        int     `json:"public"`
	BackdropPath  string  `json:"backdrop_path,omitempty"`
	AverageRating float64 `json:"average_rating"`
}

type Movie struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	IDExternal   int    `json:"id_external"`
	Title        string `json:"title" gorm:"type:varchar(255)"`
	BackdropPath string `json:"backdrop_path,omitempty" gorm:"type:varchar(255)"`
	Adult        int    `json:"adult"`
	Overview     string `json:"overview"`
}

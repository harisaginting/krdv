package movie

type ResponseList struct {
	Items []User `json:"items"`
	Total int    `json:"total"`
}

type User struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

type TmdbListMovie struct {
	Page    int         `json:"page"`
	Results []TmdbMovie `json:"results"`
	SortBy  string      `json:"sort_by"`
}

type TmdbMovie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	MediaType        string  `json:"media_type"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      int     `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

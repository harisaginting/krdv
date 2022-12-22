package watchlist

import (
	"context"

	"github.com/harisaginting/krdv/common/log"
	"github.com/harisaginting/krdv/common/utils/helper"
	"github.com/harisaginting/krdv/db/dao"
	"github.com/harisaginting/krdv/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func ProviderRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

func (repo *Repository) Create(ctx context.Context, p RequestWatchlist) (res WatchlistDetail, err error) {
	userId := helper.ForceInt(ctx.Value("userid"))
	username := helper.ForceString(ctx.Value("username"))
	qx := repo.db
	watchlist := dao.Watchlist{
		Name:      p.Name,
		Note:      p.Note,
		UserId:    userId,
		CreatedBy: username,
	}

	qx.Save(&watchlist)
	res = WatchlistDetail{
		ID:   watchlist.ID,
		Name: watchlist.Name,
		Note: watchlist.Note,
	}
	return
}

func (repo *Repository) Update(ctx context.Context, id int, p RequestWatchlist) (res WatchlistDetail, err error) {
	userId := helper.ForceInt(ctx.Value("userid"))
	username := helper.ForceString(ctx.Value("username"))
	qx := repo.db
	now := helper.Now()
	watchlist := dao.Watchlist{
		ID:        id,
		Name:      p.Name,
		Note:      p.Note,
		UserId:    userId,
		UpdatedBy: username,
		UpdatedAt: &now,
	}

	qx.Save(&watchlist)
	res = WatchlistDetail{
		ID:   watchlist.ID,
		Name: watchlist.Name,
		Note: watchlist.Note,
	}
	return
}

func (repo *Repository) Delete(ctx context.Context, id int) (ok bool, err error) {
	qx := repo.db
	watchlist := dao.Watchlist{
		ID: id,
	}
	qx.Delete(&watchlist)
	ok = true
	return
}

func (repo *Repository) Get(ctx context.Context, id int) (res WatchlistDetail, err error) {
	qx := repo.db
	qw := `
	select main.id, main.name, main.note from watchlist as main 
	WHERE main.user_id = ? AND id = ?`
	qx.Debug().Raw(qw, id).First(&res)
	if err != nil {
		log.Error(ctx, err)
		return
	}

	return
}

func (repo *Repository) GetByUser(ctx context.Context, id int) (res WatchlistDetail, err error) {
	userId := helper.ForceInt(ctx.Value("userid"))
	qx := repo.db
	qw := `
	select main.id, main.name, main.note from watchlist as main 
	WHERE main.user_id = ? AND id = ?`
	qx.Debug().Raw(qw, userId, id).First(&res)
	if err != nil {
		log.Error(ctx, err)
		return
	}

	return
}

func (repo *Repository) Page(ctx context.Context, p model.RequestPage) (res []Watchlist, total int, err error) {
	userId := helper.ForceInt(ctx.Value("userid"))
	qx := repo.db
	qw := `
	select a.id, a.name, sum(b.total_item) as total_item from (
		select main.id, main.name from watchlist as main 
		WHERE main.user_id = ? ORDER BY main.name LIMIT ? OFFSET ?
	) as a 
	LEFT JOIN (select watchlist_id , count(id) as total_item from  watchlist_movie group by id ) as b on a.id = b.watchlist_id
	group by a.id, a.name
	order by a.id asc
	`
	rows, err := qx.Debug().Raw(qw, userId, p.Limit, p.Page-1).Rows()
	if err != nil {
		log.Error(ctx, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		qx.ScanRows(rows, &res)
	}

	if len(res) == 0 {
		res = make([]Watchlist, 0)
	}

	qw2 := `
	select count(main.id) as total from watchlist as main WHERE main.user_id = ?`
	qx.Debug().Raw(qw2, userId).First(&total)

	return
}

// MOVIES
func (repo *Repository) GetMovieByWatchlist(ctx context.Context, watchlistId int) (res []WatchlistMovie, err error) {

	qx := repo.db
	qw := `
	select main.id, main.movie_id, a.name, a.description, a.revenue, a.public, a.backdrop_path, a.average_rating from watchlist_movie  as main
	JOIN movie as a on a.id = main.movie_id
	WHERE main.watchlist_id = ? `
	rows, err := qx.Debug().Raw(qw, watchlistId).Rows()
	if err != nil {
		log.Error(ctx, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		qx.ScanRows(rows, &res)
	}
	res = make([]WatchlistMovie, 0)

	return
}

func (repo *Repository) GetMovieByUser(ctx context.Context, id, movieId int) (res Movie, err error) {
	userId := helper.ForceInt(ctx.Value("userid"))
	qx := repo.db
	qw := `
	select main.id, main.id_external, main.title, main.backdrop_path, main.adult, main.overview from movie as main
	JOIN watchlist_movie as a on a.movie_id = main.id_external
	JOIN watchlist as b on b.id = a.watchlist_id
	WHERE b.id = ? AND a.movie_id = ? AND b.user_id = ?`
	qx.Debug().Raw(qw, id, movieId, userId).First(&res)
	if err != nil {
		log.Error(ctx, err)
		return
	}

	return
}

func (repo *Repository) AddMovie(ctx context.Context, id, movieId int) (res Movie, err error) {
	username := helper.ForceString(ctx.Value("username"))
	table := dao.WatchlistMovie{
		MovieId:     movieId,
		WatchlistId: id,
	}
	qx := repo.db
	table.CreatedBy = username
	qx.Save(&table)

	res, err = repo.GetMovieByIdExternal(ctx, table.MovieId)

	return
}

func (repo *Repository) DeleteMovie(ctx context.Context, id, movieId int) (ok bool, err error) {
	qx := repo.db
	qw := `DELETE FROM watchlist_movie WHERE watchlist_id = ? AND movie_id = ?`
	qx.Debug().Exec(qw, id, movieId)
	if err != nil {
		log.Error(ctx, err)
		return
	}
	ok = true
	return
}

func (repo *Repository) GetMovieById(ctx context.Context, movieId int) (res Movie, err error) {
	qx := repo.db
	qw := `
	select main.id, main.id_external, main.title, main.backdrop_path, main.adult, main.overview from movie as main
	WHERE main.id = ?`
	qx.Debug().Raw(qw, movieId).First(&res)
	if err != nil {
		log.Error(ctx, err)
		return
	}

	return
}

func (repo *Repository) GetMovieByIdExternal(ctx context.Context, movieId int) (res Movie, err error) {
	qx := repo.db
	qw := `
	select main.id, main.id_external, main.title, main.backdrop_path, main.adult, main.overview from movie as main
	WHERE main.id_external = ?`
	qx.Debug().Raw(qw, movieId).First(&res)
	if err != nil {
		log.Error(ctx, err)
		return
	}

	return
}

func (repo *Repository) AddMovieToDB(ctx context.Context, p Movie) (res Movie, err error) {
	username := helper.ForceString(ctx.Value("username"))
	var table dao.Movie
	qx := repo.db
	helper.AdjustStructToStruct(p, &table)
	table.CreatedBy = username

	qx.Save(&table)

	res, err = repo.GetMovieById(ctx, table.ID)

	return
}

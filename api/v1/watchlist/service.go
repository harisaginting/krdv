package watchlist

import (
	"context"
	"fmt"

	mMovie "github.com/harisaginting/krdv/api/v1/movie"
	mUser "github.com/harisaginting/krdv/api/v1/user"
	"github.com/harisaginting/krdv/common/log"
	"github.com/harisaginting/krdv/common/utils/helper"
	"github.com/harisaginting/krdv/model"
)

type Service struct {
	repo     Repository
	repoUser mUser.Repository
	movie    mMovie.Service
}

func ProviderService(r Repository) Service {
	return Service{
		repo:     r,
		repoUser: mUser.ProviderRepository(r.db),
		movie:    mMovie.Service{},
	}
}

func (service *Service) Create(ctx context.Context, p RequestWatchlist) (watchlist WatchlistDetail, err error) {
	watchlist, err = service.repo.Create(ctx, p)
	if err != nil {
		log.Error(ctx, err)
		return
	}
	watchlist, err = service.GetByUser(ctx, watchlist.ID)
	return
}

func (service *Service) Update(ctx context.Context, id int, p RequestWatchlist) (watchlist WatchlistDetail, err error) {
	check, _ := service.repo.GetByUser(ctx, id)
	if check.ID == 0 {
		err = fmt.Errorf("watchlist not found")
		return
	}

	watchlist, err = service.repo.Update(ctx, id, p)
	if err != nil {
		log.Error(ctx, err)
		return
	}
	watchlist, err = service.GetByUser(ctx, watchlist.ID)
	return
}

func (service *Service) Delete(ctx context.Context, id int) (ok bool, err error) {
	check, _ := service.repo.GetByUser(ctx, id)
	if check.ID == 0 {
		err = fmt.Errorf("watchlist not found")
		return
	}
	ok, err = service.repo.Delete(ctx, id)
	if err != nil {
		log.Error(ctx, err)
		return
	}
	return
}

func (service *Service) GetByUser(ctx context.Context, id int) (watchlist WatchlistDetail, err error) {
	watchlist, err = service.repo.GetByUser(ctx, id)
	if err != nil {
		log.Error(ctx, err)
		return
	}

	watchlist.Movie, err = service.repo.GetMovieByWatchlist(ctx, watchlist.ID)
	if err != nil {
		log.Error(ctx, err)
		return
	}
	return
}

func (service *Service) List(ctx context.Context, p model.RequestPage) (res ResponseList, err error) {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.Limit == 0 {
		p.Limit = 10
	}

	if p.Sort == "" {
		p.Sort = "title"
	}
	if p.Order == "" {
		p.Order = "asc"
	}

	watchlists, total, err := service.repo.Page(ctx, p)
	if err != nil {
		log.Error(ctx, err)
		return
	}

	res.Items = watchlists
	res.Total = total
	return
}

func (service *Service) SyncMovie(ctx context.Context, movieId int) (movie Movie, err error) {
	movie, err = service.repo.GetMovieByIdExternal(ctx, movieId)
	if movie.ID == 0 {
		tMovie, errIn := service.movie.Get(ctx, movieId)
		if errIn != nil {
			err = errIn
			log.Error(ctx, err)
			return
		}
		if tMovie.IDExternal == 0 {
			err = fmt.Errorf("movie not found")
			log.Error(ctx, err)
			return
		}
		helper.AdjustStructToStruct(tMovie, &movie)
		movie, err = service.repo.AddMovieToDB(ctx, movie)
		return
	}
	return
}

func (service *Service) AddMovie(ctx context.Context, id, movieId int) (movie Movie, err error) {
	check, _ := service.repo.GetMovieByUser(ctx, id, movieId)
	if check.ID != 0 {
		err = fmt.Errorf("movie already available at this watchlist")
		return
	}

	movie, err = service.SyncMovie(ctx, movieId)
	if err != nil {
		log.Error(ctx, err)
		return
	}

	movie, err = service.repo.AddMovie(ctx, id, movie.IDExternal)
	if err != nil {
		log.Error(ctx, err)
		return
	}
	return
}

func (service *Service) DeleteMovie(ctx context.Context, id, movieId int) (ok bool, err error) {
	check, _ := service.repo.GetMovieByUser(ctx, id, movieId)
	if check.ID == 0 {
		err = fmt.Errorf("movie not found at this watchlist")
		return
	}

	ok, err = service.repo.DeleteMovie(ctx, id, movieId)
	if err != nil {
		log.Error(ctx, err)
		return
	}
	return
}

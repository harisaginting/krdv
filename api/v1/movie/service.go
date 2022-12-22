package movie

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/harisaginting/krdv/common/log"
	"github.com/harisaginting/krdv/common/utils/helper"
	"github.com/harisaginting/krdv/db/dao"
	"github.com/harisaginting/krdv/model"
)

type Service struct {
	repo Repository
}

var (
	uriTMDB        string
	imageTMDB      string
	apiAccessToken string
	apiKey         string
)

func init() {
	uriTMDB = helper.MustGetEnv("TMDB_URI")
	imageTMDB = helper.MustGetEnv("TMDB_IMAGE")
	apiKey = helper.MustGetEnv("TMDB_API_KEY")
	apiAccessToken = helper.MustGetEnv("TMDB_ACCESS_TOKEN")

}

func ProviderService(r Repository) Service {
	return Service{
		repo: r,
	}
}

func (service *Service) List(ctx context.Context, p model.RequestPage) (tmbdbList TmdbListMovie, err error) {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.Sort == "" {
		p.Sort = "title"
	}
	if p.Order == "" {
		p.Order = "asc"
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	endpoint := uriTMDB + "/4/list/1"
	client := &http.Client{Transport: tr}
	doScrap, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Error(ctx, err)
		return
	}

	q := doScrap.URL.Query()
	q.Add("sort_by", p.Sort+"."+p.Order)
	q.Add("page", helper.ForceString(p.Page))
	doScrap.URL.RawQuery = q.Encode()

	doScrap.Header.Add("Content-Type", "text/html")
	doScrap.Header.Add("Authorization", "Bearer "+apiAccessToken)
	doScrap.Header.Add("Content-Type", "application/json;charset=utf-8")
	resScrap, err := client.Do(doScrap)
	if err != nil {
		log.Error(ctx, err)
		return
	}
	resBody, err := ioutil.ReadAll(resScrap.Body)
	if err != nil {
		log.Error(ctx, err)
		return
	}
	defer resScrap.Body.Close()
	json.Unmarshal(resBody, &tmbdbList)
	for i, _ := range tmbdbList.Results {
		tmbdbList.Results[i].BackdropPath = imageTMDB + tmbdbList.Results[i].BackdropPath
	}
	return
}

func (service *Service) Get(ctx context.Context, movieId int) (movie dao.Movie, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	endpoint := uriTMDB + "/3/movie/" + helper.ForceString(movieId)
	client := &http.Client{Transport: tr}
	doScrap, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Error(ctx, err)
		return
	}

	q := doScrap.URL.Query()
	q.Add("api_key", apiKey)
	doScrap.URL.RawQuery = q.Encode()

	doScrap.Header.Add("Content-Type", "text/html")
	doScrap.Header.Add("Authorization", "Bearer "+apiAccessToken)
	doScrap.Header.Add("Content-Type", "application/json;charset=utf-8")
	resScrap, err := client.Do(doScrap)
	if err != nil {
		log.Error(ctx, err)
		return
	}
	resBody, err := ioutil.ReadAll(resScrap.Body)
	if err != nil {
		log.Error(ctx, err)
		return
	}
	defer resScrap.Body.Close()
	json.Unmarshal(resBody, &movie)
	movie.IDExternal = movie.ID
	movie.ID = 0
	return
}

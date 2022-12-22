//go:build wireinject
// +build wireinject

package wire

import (
	googleWire "github.com/google/wire"
	mAuth "github.com/harisaginting/krdv/api/v1/auth"
	mMovie "github.com/harisaginting/krdv/api/v1/movie"
	mReport "github.com/harisaginting/krdv/api/v1/report"
	mUser "github.com/harisaginting/krdv/api/v1/user"
	mWatchlist "github.com/harisaginting/krdv/api/v1/watchlist"
	"gorm.io/gorm"
)

func ApiUser(db *gorm.DB) mUser.Controller {
	googleWire.Build(
		mUser.ProviderController,
		mUser.ProviderService,
		mUser.ProviderRepository,
	)
	return mUser.Controller{}
}

func ApiAuth(db *gorm.DB) mAuth.Controller {
	googleWire.Build(
		mAuth.ProviderController,
		mAuth.ProviderService,
		mAuth.ProviderRepository,
	)
	return mAuth.Controller{}
}

func ApiMovie(db *gorm.DB) mMovie.Controller {
	googleWire.Build(
		mMovie.ProviderController,
		mMovie.ProviderService,
		mMovie.ProviderRepository,
	)
	return mMovie.Controller{}
}

func ApiWatchlist(db *gorm.DB) mWatchlist.Controller {
	googleWire.Build(
		mWatchlist.ProviderController,
		mWatchlist.ProviderService,
		mWatchlist.ProviderRepository,
	)
	return mWatchlist.Controller{}
}

func ApiReport(db *gorm.DB) mReport.Controller {
	googleWire.Build(
		mReport.ProviderController,
		mReport.ProviderService,
		mReport.ProviderRepository,
	)
	return mReport.Controller{}
}

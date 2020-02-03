package services

import (
	"api-di/apps/models"
	"time"

	"github.com/sarulabs/di"
	"go.uber.org/zap"
	mgo "gopkg.in/mgo.v2"
)

// Services contains the definitions of the application services.
var Services = []di.Def{
	{
		Name:  "logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return Logger, nil
		},
	},
	{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return Config, nil
		},
	},
	{
		Name:  "mongo-pool",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			//return mgo.DialWithTimeout(fmt.Sprintf("%s:%s", Config.GetString("mongo.host"), Config.GetString("mongo.port")), 5*time.Second)

			return mgo.DialWithTimeout("0.0.0.0:27017", 5*time.Second)
		},
		Close: func(obj interface{}) error {
			obj.(*mgo.Session).Close()
			return nil
		},
	},
	{
		Name:  "mongo",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return ctn.Get("mongo-pool").(*mgo.Session).Copy(), nil
		},
		Close: func(obj interface{}) error {
			obj.(*mgo.Session).Close()
			return nil
		},
	},
	{
		Name:  "repository",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &models.Repository{
				Session: ctn.Get("mongo").(*mgo.Session),
			}, nil
		},
	},
	{
		Name:  "manager",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &models.LinkManager{
				Repo:   ctn.Get("repository").(*models.Repository),
				Logger: ctn.Get("logger").(*zap.Logger),
			}, nil
		},
	},
}

package svc

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/model"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/form/rpc/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	FormClient model.EntryFormModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		FormClient: model.NewEntryFormModel(c.MongoDBConf.URL, c.MongoDBConf.DB, "form"),
	}
}
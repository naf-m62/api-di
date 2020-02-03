package models

import (
	"go.uber.org/zap"
)

type LinkManager struct {
	Repo *Repository
	Logger *zap.Logger
}

func (m *LinkManager) GetAll() (links []*Links, err error) {
	links, err = m.Repo.GetAll()
	if err != nil {
		m.Logger.Error("GetAll error", zap.Error(err))
	}
	if links == nil {
		links = []*Links{}
	}
	return links, nil
}
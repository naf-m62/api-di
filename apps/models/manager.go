package models

import (
	"strconv"
	"time"

	"go.uber.org/zap"

	"api-di/apps/python_generator_links"
)

type LinkManager struct {
	Repo   *Repository
	Logger *zap.Logger
}

// GetAll получить все ссылки
func (m *LinkManager) GetAll() (linksByG []*LinksByGroup, err error) {
	var links []*Link
	links, err = m.Repo.GetAll()
	if err != nil {
		m.Logger.Error("GetAll error", zap.Error(err))
	}
	if links == nil {
		links = []*Link{}
	}
	if linksByG, err = m.sortLinksByGroup(links); err != nil {
		return nil, err
	}
	return linksByG, err
}

// GetById получить ссылку по id
func (m *LinkManager) GetById(id int64) (link *Link, err error) {
	link, err = m.Repo.GetById(id)
	if err != nil {
		m.Logger.Error("GetById error", zap.Error(err))
	}
	return link, err
}

// CreateLink создание ссылки
func (m *LinkManager) CreateLink(l *Link) (err error) {
	// для id используем timestamp
	l.ID = time.Now().Unix()
	if err = m.Repo.CreateLink(l); err != nil {
		m.Logger.Error("CreateLink error", zap.Error(err))
	}
	return err
}

// UpdateLink обновить ссылки
func (m *LinkManager) UpdateLink(l *Link) (err error) {
	if err = m.Repo.UpdateLink(l); err != nil {
		m.Logger.Error("UpdateLink error", zap.Error(err))
	}
	return err
}

// DeleteLink удалить ссылку
func (m *LinkManager) DeleteLink(id int64) (err error) {
	err = m.Repo.DeleteLink(id)
	if err != nil {
		m.Logger.Error("DeleteLink error", zap.Error(err))
	}
	return err
}

// CreateStandLinks создает список ссылок для стенда
func (m *LinkManager) CreateStandLinks(testStandNumber int) (err error) {
	var (
		outLinks python_generator_links.GenerateLinkResponse
		lWrapper = m.Logger.With(zap.String("method", "CreateStandLinks"),
			zap.Int("testStandNumber", testStandNumber))
	)
	outLinks, err = python_generator_links.GenerateLinks(testStandNumber)
	if err != nil {
		lWrapper.Error("GenerateLinks error", zap.Error(err))
		return err
	}

	var dbLinks []Link
	dbLinks = convertToModelLinks(outLinks, testStandNumber)

	for _, l := range dbLinks {
		if err = m.Repo.CreateLink(&l); err != nil {
			lWrapper.Error("CreateLink error", zap.Error(err), zap.String("url", l.Url))
			break
		}
	}
	return err
}

// sortLinksByGroup возвращает ссылки разбитые по группам
func (m *LinkManager) sortLinksByGroup(links []*Link) (linksByG []*LinksByGroup, err error) {
	var groups []string
	//	lg := map[string]LinksByGroup{}
	if groups, err = m.Repo.GetGroups(); err != nil {
		return nil, err
	}

	linksByG = make([]*LinksByGroup, len(groups))

	for _, k := range links {
		for ii, kk := range groups {
			if kk == k.LinkGroup {
				if linksByG[ii] == nil {
					linksByG[ii] = &LinksByGroup{}
				}
				linksByG[ii].Count += 1
				linksByG[ii].Links = append(linksByG[ii].Links, k)
				if linksByG[ii].Name == "" {
					linksByG[ii].Name = k.LinkGroup
				}
				break
			}
		}
	}

	// прибавляем 1 к каждой группе, нужно для rowspan
	for _, k := range linksByG {
		k.Count += 1
	}

	return linksByG, nil
}

// convertToModelLinks конвертирует из списка сгенерированых ссылок, полученного из питон скрипта в слайс ссылок
func convertToModelLinks(outLinkTotal python_generator_links.GenerateLinkResponse, testStandNumber int) (dbLinkTotal []Link) {
	for _, outLink := range outLinkTotal.Links {
		var desc string
		if outLink.Db != "" {
			desc += "db: " + outLink.Db + ", \n"
		}
		if outLink.Login != "" {
			desc += "login: " + outLink.Login + ", \n"
		}
		if outLink.Password != "" {
			desc += "password: " + outLink.Password + ", \n"
		}
		dbLink := Link{
			ID:          time.Now().UnixNano(),
			Url:         outLink.Host,
			LinkGroup:   "test" + strconv.Itoa(testStandNumber),
			Description: desc,
		}
		dbLinkTotal = append(dbLinkTotal, dbLink)
	}
	return dbLinkTotal
}

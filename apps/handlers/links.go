package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"

	"api-di/apps/models"
)

const manager = "manager"

// GetLinkListHandler обработка запроса на получение всех ссылок
func GetLinkListHandler(w http.ResponseWriter, r *http.Request) {
	var (
		res []*models.Link
		lm  *models.LinkManager
		err error
	)

	lm = getLinkManager(w, r)
	res, err = lm.GetAll()
	if err != nil {
		respError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response(w, http.StatusOK, res)
}

// CreateLinkHandler создать ссылку
func CreateLinkHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer func() {
		r.Body.Close()
	}()

	var (
		linkNew *models.Link
		lm      *models.LinkManager
	)
	err = json.Unmarshal(body, &linkNew)

	lm = getLinkManager(w, r)
	err = lm.CreateLink(linkNew)

	if err != nil {
		respError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetLinkHandler получить ссылку по id
func GetLinkHandler(w http.ResponseWriter, r *http.Request) {
	var (
		res *models.Link
		err error
		id  int
		lm  *models.LinkManager
	)
	id, err = strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respError(w, http.StatusBadRequest, "incorrect id: "+err.Error()+"/n")
		return
	}

	lm = getLinkManager(w, r)
	if res, err = lm.GetById(int64(id)); err != nil {
		respError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, res)
	return
}

// UpdateLinkHandler обновить ссылку
func UpdateLinkHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer func() {
		r.Body.Close()
	}()
	var (
		link *models.Link
		lm   *models.LinkManager
		id   int
	)
	err = json.Unmarshal(body, &link)
	id, err = strconv.Atoi(mux.Vars(r)["id"])
	link.ID = int64(id)
	if err != nil {
		respError(w, http.StatusBadRequest, "incorrect id: "+err.Error()+"/n")
		return
	}
	lm = getLinkManager(w, r)
	err = lm.UpdateLink(link)

	if err != nil {
		respError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteLinkHandler удалить ссылку по id
func DeleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		id  int
		lm  *models.LinkManager
	)
	id, err = strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respError(w, http.StatusBadRequest, "incorrect id: "+err.Error()+"/n")
		return
	}

	lm = getLinkManager(w, r)
	if err = lm.DeleteLink(int64(id)); err != nil {
		respError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

// getLinkManager отдает manager
func getLinkManager(w http.ResponseWriter, r *http.Request) (lm *models.LinkManager) {
	c, ok := r.Context().Value(di.ContainerKey("di")).(di.Container)
	if !ok {
		w.Write([]byte("can not get container"))
	}
	if err := c.Fill(manager, &lm); err != nil {
		w.Write([]byte(err.Error()))
	}
	return lm
}

// response записывает ответ
func response(w http.ResponseWriter, code int, data interface{}) {
	var result []byte
	result, _ = json.Marshal(data)
	w.WriteHeader(code)
	w.Write(result)
}

// respError записывает ответ при ошибке
func respError(w http.ResponseWriter, code int, errMessage string) {
	var result []byte
	w.WriteHeader(code)
	result, _ = json.Marshal(map[string]string{"error": errMessage})
	w.Write(result)
}

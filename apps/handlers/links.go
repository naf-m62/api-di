package handlers

import (
	"api-di/apps/models"
	"encoding/json"
	"github.com/sarulabs/di"
	"net/http"
)

func GetLinkListHandler(w http.ResponseWriter, r *http.Request) {
	var res []*models.Links
	res, err := di.Get(r, "manager").(*models.Repository).GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

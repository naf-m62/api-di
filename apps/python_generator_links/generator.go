package python_generator_links

import (
	"encoding/json"
	"os/exec"
	"strconv"
)

type (
	GenerateLinkResponse struct {
		Links []generateLinkResponseItems `json:"links"`
	}

	generateLinkResponseItems struct {
		Host     string `json:"host"`
		Db       string `json:"db"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}
)

// GenerateLinks генерирует ссылки для стенда, на вход получает номер стенда
func GenerateLinks(testStandNumber int) (links GenerateLinkResponse, err error) {
	var out []byte
	// вызывает скрипт питона, который возвращает json со списоком сгенерированных ссылок
	out, err = exec.Command("./apps/python_generator_links/venv/bin/python3.6", "./apps/python_generator_links/generate_link.py", strconv.Itoa(testStandNumber)).Output()
	if err != nil {
		return GenerateLinkResponse{}, err
	}

	outLinks := GenerateLinkResponse{}
	if err = json.Unmarshal(out, &outLinks); err != nil {
		return GenerateLinkResponse{}, err
	}

	return outLinks, nil
}

// convertToModelLinks конвертирует из списка сгенерированых ссылок, полученного из питон скрипта в слайс ссылок
//func convertToModelLinks(outLinkTotal GenerateLinkResponse, testStandNumber int) (dbLinkTotal []models.Link) {
//	for _, outLink := range outLinkTotal.Links {
//		var desc string
//		if outLink.Db != "" {
//			desc += "db: " + outLink.Db + ", \n"
//		}
//		if outLink.Login != "" {
//			desc += "login: " + outLink.Login + ", \n"
//		}
//		if outLink.Password != "" {
//			desc += "password: " + outLink.Password + ", \n"
//		}
//		dbLink := models.Link{
//			ID:          time.Now().Unix(),
//			Url:         outLink.Host,
//			LinkGroup:   "test" + strconv.Itoa(testStandNumber),
//			Description: desc,
//		}
//		dbLinkTotal = append(dbLinkTotal, dbLink)
//	}
//	return dbLinkTotal
//}

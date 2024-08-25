package api

import (
	"encoding/json"
	"errors"
	"footy_res/models"
	"io"
	"net/http"
	"strconv"
)
func FetchData(API_URL string, API_KEY string) (data models.Data,err error){

	var res_data models.Data = models.Data{}

	req, err := http.NewRequest("GET", API_URL, nil)

	req.Header.Add("X-Auth-Token", API_KEY)

	client := &http.Client{}
	res,err := client.Do(req)

	if res.StatusCode != 200 {
		err = errors.New("Status code: " + strconv.Itoa(res.StatusCode))
	}
	defer res.Body.Close()

	body,err := io.ReadAll(res.Body)
	if err != nil {
		return res_data,err
	}
	err = json.Unmarshal(body, &res_data)
	if err != nil {
		return res_data,err
	}

	return res_data,err
}
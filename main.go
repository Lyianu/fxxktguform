package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r := setupRouter()
	r.Run(":80")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/status/:id", idChecker)
	r.Static("/fxxkform", "./static")

	return r
}

func idChecker(c *gin.Context) {

	id := c.Params.ByName("id")
	status, err := validate("5eb21cfc75a03c7a6903df0a", id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": status, "id": id})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status, "id": id})
}

func validate(frmid, value string) (result bool, err error) {

	values := map[string]string{"FRMID": frmid, "NM": "F999", "VAL": value}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://biaodan100.com/web/formview/existValidate",
		"application/json;charset=UTF-8",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var jsonMap map[string]interface{}
	json.Unmarshal(body, &jsonMap)

	return jsonMap["result"].(bool), err
}

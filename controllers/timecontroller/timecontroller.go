package timecontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getTime(name string, API_KEY string) float64 {

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + name + "&appid=" + API_KEY)

	if err != nil {
		log.Fatal(err)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//fmt.Println(string(respData))

	var data map[string]interface{}

	if err = json.Unmarshal(respData, &data); err != nil {
		log.Fatal(err)
	}

	datas := data["timezone"].(float64)

	return datas
}

func convertTime(shifttime float64) time.Time {
	now := time.Now().UTC()
	convertTime := now.Add(time.Second * time.Duration(shifttime))

	//fmt.Println(convertTime)
	return convertTime
}

func Index(c *gin.Context) {
	city := c.Query("city")
	API_KEY := "f7e80336d482f800cfbdba35d08e9539"

	//get shift time
	shifttime := getTime(city, API_KEY)

	//get current time
	currentTime := convertTime(shifttime)

	c.JSON(http.StatusOK, gin.H{"location": city, "timezone": shifttime, "current_time": currentTime})
}

package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type Response []struct {
	Breeds []interface{} `json:"breeds"`
	ID     string        `json:"id"`
	URL    string        `json:"url"`
	Width  int           `json:"width"`
	Height int           `json:"height"`
}

// Golang Method
func (c *MainController) BreedSearch() {

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("get", "https://api.thecatapi.com/v1/images/search", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("x-api-key", "17d94b92-754f-46eb-99a0-65be65b5d18f")

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Response status:", response.Status)

	body, err1 := ioutil.ReadAll(response.Body) // response body is []byte
	if err1 != nil {
		panic(err1)
	}

	//fmt.Println(string(body)) // convert to string before print
	// put the json response in the json to golang struct link
	var result Response
	if err2 := json.Unmarshal(body, &result); err2 != nil { // Parse []byte to go struct pointer
		panic(err2)
	}
	//fmt.Println(PrettyPrint(result))

	// Loop through the struct for the URL

	var image_url string
	var image_id string

	for _, rec := range result {
		image_id = rec.ID
		image_url = rec.URL
		fmt.Println(image_id)
	}

	c.Data["ImgUrl"] = image_url
	c.Data["ImgId"] = image_id
	c.TplName = "breed_search.html"

}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

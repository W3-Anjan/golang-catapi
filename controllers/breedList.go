package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type BreedListController struct {
	beego.Controller
}

type FavBreedIdResponse struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

func (c *BreedListController) FavBreedById() {

	var imgId string = c.GetString("imgid")
	fmt.Println("ID:", imgId)

	baseUrl := "https://api.thecatapi.com/v1/favourites"
	method := "POST"

	postBody, _ := json.Marshal(map[string]interface{}{"image_id": imgId})
	responseBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	req, err := http.NewRequest(method, baseUrl, responseBody)
	if err != nil {
		panic(err)
	}
	req.Header.Add("content-type", "application/json")
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

	fmt.Println(string(body)) // convert to string before print
	var result FavBreedIdResponse
	if err2 := json.Unmarshal(body, &result); err2 != nil { // Parse []byte to go struct pointer
		panic(err2)
	}

	image_id := result.ID

	fmt.Println(image_id)

	c.Data["ImgId"] = image_id
	c.TplName = "breed_fav.html"
}

type AllFabBreedsResponse []struct {
	ID        int         `json:"id"`
	UserID    string      `json:"user_id"`
	ImageID   string      `json:"image_id"`
	SubID     interface{} `json:"sub_id"`
	CreatedAt time.Time   `json:"created_at"`
	Image     struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"image"`
}

func (c *BreedListController) GetFavBreeds() {

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("get", "https://api.thecatapi.com/v1/favourites?limit=12&page=1", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("x-api-key", "17d94b92-754f-46eb-99a0-65be65b5d18f")

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err1 := ioutil.ReadAll(response.Body) // response body is []byte
	if err1 != nil {
		panic(err1)
	}

	// put the json response in the json to golang struct link
	var result AllFabBreedsResponse
	if err2 := json.Unmarshal(body, &result); err2 != nil { // Parse []byte to go struct pointer
		panic(err2)
	}

	imgList := make([]string, len(result))

	for i, rec := range result {
		imgList[i] = rec.Image.URL
	}
	fmt.Println("imgList:", imgList)

	c.Data["ImgList"] = imgList

	c.TplName = "breed_fav.html"
}

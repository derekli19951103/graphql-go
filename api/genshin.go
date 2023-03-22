package api

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const OS_DS_SALT = "6s25p5ox5y14umn1p61aqyyvbvvl3lrt"
const OS_GAME_URL = "https://bbs-api-os.hoyolab.com/"

func generateRandomString(n int) string {
    rand.Seed(time.Now().UnixNano())
    var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}


type GenshinClient struct{
	Uid string
	Server_id string
	Cookies string
}

type GenshinCharaterResponse struct {
	Data  GenshinData `json:"data"`
	Message string `json:"message"`
	Retcode int	`json:"retcode"`
}

type GenshinData struct {
	Avatars []GenshinCharacter `json:"avatars"`	
}

type GenshinCharacter struct {
	Element string `json:"element"`
	Actived_constellation_num int `json:"actived_constellation_num"`
	Name string `json:"name"`
	Level int `json:"level"`
	Rarity int `json:"rarity"`
	Weapon GenshinWeapon
}

type GenshinWeapon struct{
	Name string `json:"name"`
	Rarity int `json:"rarity"`
	Level int `json:"level"`
}

func (gc *GenshinClient) generateDS() string{
	t :=  strconv.Itoa(int(time.Now().Unix()))
    r := generateRandomString(6)
   
	before := fmt.Sprintf("salt=%s&t=%s&r=%s",OS_DS_SALT,t,r,)
	hash := md5.Sum([]byte(before))
	h:= hex.EncodeToString(hash[:])

    return fmt.Sprintf("%s,%s,%s",t,r,h)
}

func (gc *GenshinClient) Fetch(url string, requestBody map[string]any) (GenshinCharaterResponse,error){
	requestDataJSON, err := json.Marshal(requestBody)
    if err != nil {
        return GenshinCharaterResponse{},errors.New("can't parse request body")
    }
	request, err := http.NewRequest("POST", fmt.Sprintf("%s%s",OS_GAME_URL,url), bytes.NewBuffer(requestDataJSON))
	if err != nil {
        return GenshinCharaterResponse{},errors.New("can't format request")
    }

	ds:=gc.generateDS()
	
	// Set the request headers
	 request.Header.Set("x-rpc-app_version", "1.5.1")
	 request.Header.Set("x-rpc-client_type", "5")
	 request.Header.Set("x-rpc-language","en-us")
	 request.Header.Set("ds", ds)
	 request.Header.Set("cookie", gc.Cookies)
	 request.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) miHoYoBBS/2.9.1")

	  // Send the request
	  client := &http.Client{}
	  resp, err := client.Do(request)
	  if err != nil {
		return GenshinCharaterResponse{},errors.New("can't send request")
	  }
  
	  // Handle the response
	  defer resp.Body.Close()


	//  !! No need to depress if not set Accept-Encoding !!

	//   var reader io.ReadCloser
	//   switch resp.Header.Get("Content-Encoding") {
	// 	case "gzip":
	// 		reader, err = gzip.NewReader(resp.Body)
	// 		if err != nil {
	// 			return nil,errors.New("can't parse response body")
	// 		}
	// 		defer reader.Close()
	// 	default:
	// 		fmt.Println("here")
	// 		reader = resp.Body
	//   }

	  bodyBytes, err := io.ReadAll(resp.Body)
	  if err != nil {
		return GenshinCharaterResponse{},errors.New("can't parse response body")
	  }

	  var data GenshinCharaterResponse
	  err = json.Unmarshal(bodyBytes, &data)
	  if err != nil {
		return GenshinCharaterResponse{},errors.New("can't parse response body")
	  }

	  return data, nil
}




   
   
    

   
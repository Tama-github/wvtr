package client

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"wvtrserv/logger"
	"wvtrserv/nanapi/config"
	"wvtrserv/utils"
)

type NanaClient struct {
	Config *config.NanapiConfig
	Header []string
}

func createClient() *NanaClient {
	conf := config.GetNanapiConfig()
	toEncode := conf.ClientUsername + ":" + conf.ClientSecret
	b64 := base64.StdEncoding.EncodeToString([]byte(toEncode))
	client := &NanaClient{
		Config: conf,
		Header: []string{
			"Authorization", "Basic " + b64,
			"Accept", "*/*",
		},
	}

	return client
}

var client *NanaClient = createClient()

func getAscendedWaifusFromDicordID(discordID string) *http.Response {
	methode := "GET"
	params := url.Values{}
	params.Add("discord_id", discordID)
	params.Add("client_id", client.Config.ClientId)
	params.Add("ascended", "1")
	params.Add("blooded", "0")

	reqPath := client.Config.NanapiDomain + "/prod/waicolle/waifus?" + params.Encode()

	response := utils.Fetch(reqPath, methode, url.Values{}, client.Header)

	return response
}

func getAnilistChar(wlist []*Waifu) *http.Response {
	methode := "GET"
	params := url.Values{}
	ids := ""
	for _, w := range wlist {
		ids += w.ID + ","
	}
	params.Add("ids_al", ids)

	reqPath := client.Config.NanapiDomain + "/prod/waicolle/waifus?" + params.Encode()

	response := utils.Fetch(reqPath, methode, url.Values{}, client.Header)

	return response
}

func GetAvailableWaifuToSendToWVTR(discordID string) *http.Response {
	responseWaifu := getAscendedWaifusFromDicordID(discordID)
	var waifus []*Waifu = make([]*Waifu, 0)
	err := json.NewDecoder(responseWaifu.Body).Decode(&waifus) //json.Unmarshal(utils.ReadResponse(responseWaifu), &waifus)
	if err != nil {
		logger.ErrLog.Println("Can't unmarshal waifus : ", err)
		return nil
	}

	return getAnilistChar(waifus)
}

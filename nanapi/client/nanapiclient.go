package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
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
	//params.Add("level", "2")
	params.Add("exclude_custom_image", "1")

	reqPath := client.Config.NanapiDomain + "/prod/waicolle/waifus?" + params.Encode()
	logger.DumpLog.Println(reqPath)
	response := utils.Fetch(reqPath, methode, url.Values{}, client.Header)

	if response == nil {
		return nil
	}

	return response
}

func fetchAnilistCharBulk(wlist []*Waifu, bulksize int) ([]*Waifu, []*JoinWC) {
	rest := []*Waifu{}
	doWlist := wlist
	if len(wlist) > bulksize {
		rest = wlist[bulksize:]
		doWlist = wlist[:bulksize]
	}
	methode := "GET"
	params := url.Values{}
	ids := ""
	logger.DumpLog.Println(doWlist)
	for _, w := range doWlist {
		ids += fmt.Sprintf("%d,", w.Charachter.IdAl)
	}
	if ids != "" {
		params.Add("ids_al", ids[:len(ids)-1])
	}

	reqPath := client.Config.NanapiDomain + "/prod/anilist/charas?" + params.Encode()

	logger.DumpLog.Println("Request anilist charachters to ", reqPath)
	response := utils.Fetch(reqPath, methode, url.Values{}, client.Header)
	logger.DumpLog.Println("Received anilist charachters response")

	var waifus []*CharachterAL = make([]*CharachterAL, 0)
	err := json.NewDecoder(response.Body).Decode(&waifus)

	if err != nil {
		logger.ErrLog.Println("Can't unmarshal waifus : ", err)
		return nil, nil
	}

	//logger.DumpLog.Println("Decoded : ", str)
	var res []*JoinWC = make([]*JoinWC, len(waifus))
	if len(waifus) != len(doWlist) {
		logger.ErrLog.Println("problem happened when getting the anilist characters char/waifu number miss match : ", len(waifus), "/", len(doWlist))
		return nil, nil
	}

	for i := 0; i < len(waifus); i++ {
		res[i] = &JoinWC{
			ID:                doWlist[i].ID,
			IdAl:              waifus[i].IdAl,
			NameUserPreferred: waifus[i].NameUserPreferred,
			ImageLarge:        waifus[i].ImageLarge,
			Rank:              waifus[i].Rank,
		}
	}

	return rest, res
}

func getAnilistChar(wlist []*Waifu) []*JoinWC {
	rest := wlist
	var res []*JoinWC
	bulkSize := 20
	rest, res = fetchAnilistCharBulk(wlist, bulkSize)
	for len(rest) > 0 {
		tmp1, tmp2 := fetchAnilistCharBulk(rest, bulkSize)
		rest = tmp1
		res = append(res, tmp2...)
	}
	return res
}

func GetAvailableWaifuToSendToWVTR(discordID string) []*JoinWC {
	responseWaifu := getAscendedWaifusFromDicordID(discordID)
	if responseWaifu == nil {
		return nil
	}
	var waifus []*Waifu = make([]*Waifu, 0)
	//logger.DumpLog.Println(string(utils.ReadResponse(responseWaifu)))
	err := json.NewDecoder(responseWaifu.Body).Decode(&waifus) //json.Unmarshal(utils.ReadResponse(responseWaifu), &waifus)
	if err != nil {
		logger.ErrLog.Println("Can't Decode waifus : ", err)
		return nil
	}

	logger.DumpLog.Println("Received ascended waifus :", len(waifus))

	return getAnilistChar(waifus)
}

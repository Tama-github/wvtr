package gamelogic

import (
	"encoding/json"
	"strings"
	"wvtrserv/logger"
)

func CreateLogicObjectFromDBObject[T any, V any](dbo *V) *T {
	dboEncoded, err := json.Marshal(dbo)
	if err != nil {
		logger.ErrLog.Println("Error while trying to read : ", dbo, " | ", err)
	}
	var res *T
	err2 := json.NewDecoder(strings.NewReader(string(dboEncoded))).Decode(res)
	if err2 != nil {
		logger.ErrLog.Println("Error while trying to decode : ", string(dboEncoded), " | ", err)
	}
	return res
}

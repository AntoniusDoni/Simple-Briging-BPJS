package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Simple-Briging-BPJS/utils"
)

func main() {
	// now := time.Now().UTC()
	// date := now.Format(utils.YYYYMMDD)
	urlreq := fmt.Sprintf(utils.GETRUJUKAN_BYNO_KA, utils.GET_CLAIM, "0001926061569")
	res, err := utils.GET(&utils.ReqInfo{
		URL: urlreq,
	}, 30*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	var response utils.ResponsePeserta
	json.Unmarshal(res.Body, &response)
	fmt.Println(res)

}

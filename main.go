package main

import (
	"fmt"
	"time"

	"github.com/Simple-Briging-BPJS/utils"
)

func main() {
	now := time.Now().UTC()
	date := now.Format(utils.YYYYMMDD)
	// urlreq := fmt.Sprintf("%sPeserta/nokartu/0001926061569/tglSEP/%s", utils.GET_CLAIM, date)
	urlreq := fmt.Sprintf(utils.GET_BYNO_KARTU, utils.GET_CLAIM, "0001926061569", date)
	res, err := utils.GET(&utils.ReqInfo{
		URL: urlreq,
	}, 30*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Body)
}

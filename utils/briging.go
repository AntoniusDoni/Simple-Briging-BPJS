package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	lzstring "github.com/daku10/go-lz-string"
	"github.com/joho/godotenv"
)

type ReqInfo struct {
	URL         string
	HeaderInfo  HeaderInfoSchema
	Auth        string
	ContentType string
	Body        []byte
}

type HeaderInfoSchema struct {
	ContentType string
	Userkey     string
	XConsID     string
	XSignature  string
	XTimestamp  string
}
type ResDecycptAPI struct {
	Peserta struct {
		NoKartu string `json:"noKartu"`
		Nik     string `json:"nik"`
		Nama    string `json:"nama"`
		Pisa    string `json:"pisa"`
		Sex     string `json:"sex"`
		Mr      struct {
			NoMR      string `json:"noMR"`
			NoTelepon string `json:"noTelepon"`
		} `json:"mr"`
		TglLahir      string `json:"tglLahir"`
		TglCetakKartu string `json:"tglCetakKartu"`
		TglTAT        string `json:"tglTAT"`
		TglTMT        string `json:"tglTMT"`
		StatusPeserta struct {
			Kode       string `json:"kode"`
			Keterangan string `json:"keterangan"`
		} `json:"statusPeserta"`
		ProvUmum struct {
			KdProvider string `json:"kdProvider"`
			NmProvider string `json:"nmProvider"`
		} `json:"provUmum"`
		JenisPeserta struct {
			Kode       string `json:"kode"`
			Keterangan string `json:"keterangan"`
		} `json:"jenisPeserta"`
		HakKelas struct {
			Kode       string `json:"kode"`
			Keterangan string `json:"keterangan"`
		} `json:"hakKelas"`
		Umur struct {
			UmurSekarang      string `json:"umurSekarang"`
			UmurSaatPelayanan string `json:"umurSaatPelayanan"`
		} `json:"umur"`
		Informasi struct {
			Dinsos      interface{} `json:"dinsos"`
			ProlanisPRB string      `json:"prolanisPRB"`
			NoSKTM      interface{} `json:"noSKTM"`
		} `json:"informasi"`
		Cob struct {
			NoAsuransi interface{} `json:"noAsuransi"`
			NmAsuransi interface{} `json:"nmAsuransi"`
			TglTMT     interface{} `json:"tglTMT"`
			TglTAT     interface{} `json:"tglTAT"`
		} `json:"cob"`
	} `json:"peserta"`
}
type ResInfo struct {
	StatusCode int
	Header     http.Header
	Body       []byte
}
type ResposeBodyBriging struct {
	MetaData Body   `json:"metadata"`
	Response string `json:"response"`
}
type ResponseBriging struct {
	MetaData Body   `json:"metadata"`
	Body     []byte `json:"response"`
}

type Body struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GET API request
func GET(reqinf *ReqInfo, timeout time.Duration) (*ResponseBriging, error) {

	req, err := http.NewRequest("GET", reqinf.URL, nil)
	if err != nil {
		return nil, err
	}
	godotenv.Load()
	constid := os.Getenv("CONST_ID")
	Secretkey := os.Getenv("SECRET_KEY")
	userkey := os.Getenv("USER_KEY")

	secondDate := time.Date(1970, 01, 01, 0, 0, 0, 0, time.UTC)
	locInd, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now()
	nowDate := now.In(locInd)
	xTimestamp := int(nowDate.Sub(secondDate).Seconds())
	x := fmt.Sprintf("%s&%d", constid, xTimestamp)
	encodedSignature := GenerateHMAC256(Secretkey, x)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-cons-id", constid)
	req.Header.Add("X-signature", encodedSignature)
	req.Header.Add("X-timestamp", fmt.Sprintf("%d", xTimestamp))
	req.Header.Add("user_key", userkey)
	// execute
	cl := &http.Client{
		Timeout: timeout,
	}
	res, err := cl.Do(req)

	if err != nil {
		fmt.Println("EROOR : ", err)
		return nil, err
	}
	defer res.Body.Close()

	// read body
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERROR READ RESPONSE:", err)
		return nil, err
	}
	key := fmt.Sprintf("%s%s%d", constid, Secretkey, xTimestamp)
	var resBody ResposeBodyBriging
	json.Unmarshal(buf, &resBody)
	if resBody.MetaData.Code != "200" {
		return &ResponseBriging{MetaData: resBody.MetaData, Body: nil}, nil
	}

	shakey := sha256.New()
	shakey.Write([]byte(key))
	keys := shakey.Sum(nil)

	ds, err := AESDecrypt(resBody.Response, keys)

	if err != nil {
		fmt.Println("ERROR AES DECODE :", err)
		return nil, err
	}
	var decyptRes string
	var errs error
	var response ResDecycptAPI
	if len(ds) > 0 {
		decyptRes, errs = lzstring.DecompressFromEncodedURIComponent(string(ds))
		json.Unmarshal([]byte(decyptRes), &response)
		if errs != nil {
			fmt.Println("ERROR LZSTRING DECODE : ", errs)
			return nil, err
		}
	}
	return &ResponseBriging{MetaData: resBody.MetaData, Body: []byte(decyptRes)}, nil
}

func GenerateHMAC256(k, message string) string {
	key := []byte(k)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func AESDecrypt(cryptoText string, key []byte) ([]byte, error) {
	crypt, err := base64.StdEncoding.DecodeString(cryptoText)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
		return nil, err

	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
		return nil, err
	}
	vi := key[:aes.BlockSize]
	ecb := cipher.NewCBCDecrypter(block, vi)
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return PKCS5Trimming(decrypted), nil

}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

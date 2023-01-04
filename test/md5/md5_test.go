package md5

import (
	"crypto/md5"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"testing"
)

func TestMD5(t *testing.T) {

	URL := "https://gchat.qpic.cn/gchatpic_new/3040809965/818848626-2682086032-00FB5731DCAFF37DD940DDAABCD20F10/0?term=3"
	// 请求URL，获取数据
	resp, err := http.Get(URL)
	if err != nil {
		log.Error("请求URL错误 ", err)
		return
	}
	data, _ := io.ReadAll(resp.Body)
	// 对请求到的数据求MD5
	FileMd5 := md5.New()
	FileMd5.Write(data)
	MD5Str := hex.EncodeToString(FileMd5.Sum(nil))

	t.Log(MD5Str)
}

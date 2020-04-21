package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"github.com/robfig/cron"
)

type Data struct {
	Token   string `json:"token"`
	Wxid    string `json:"wxid"`
	Typed   string `json:"type"`
	Message string `json:"message"`
}

func postData(token, wxid, msg string) {
	url := "http://api.nook.vip"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	data := Data{Token: token, Wxid: wxid, Typed: "TXT", Message: encoded}

	b := new(bytes.Buffer)
	//b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(data)

	fmt.Println(string(b.Bytes()))
	req, err := http.NewRequest("POST", url, b)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func getZSXQData(url string) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Host", "api.zsxq.com")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.113 Safari/537.36")
	req.Header.Set("Cookie", "abtest_env=product; zsxq_access_token=AB6747C9-57E5-F136-28BF-8875C41CA7CB_279D582DD5BA5DB7; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22171907422b29ab-02687c6c26b971-6373664-1247616-171907422b3673%22%2C%22%24device_id%22%3A%22171907422b29ab-02687c6c26b971-6373664-1247616-171907422b3673%22%2C%22props%22%3A%7B%7D%7D; UM_distinctid=17191a864e7202-0a007fc5caaf81-6373664-130980-17191a864e89")
	//req.Header.Set("X-Timestamp", "1587316046")
	//req.Header.Set("X-Signature", "14aec71d2d22d58ff713aa80b5460396e91d6694")
	//req.Header.Set("X-Version", "1.10.41")
	//req.Header.Set("X-Request-Id", "abf1a6390-7293-a732-845c-43bfb64ce74")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

func Push2WXGroup() {
	//content := "【测试】\n内容\n　——作者\nhttp://baidu.com"
	//postData("V1gffl5xs6q34mdpbt", "21552176681@chatroom", content)
	//getZSXQData("https://api.zsxq.com/v1.10/topics/548488425851254")
	getZSXQData("https://t.zsxq.com/ufAAUzR")
}

func CronPushJob() {
	// 先push一次
	Push2WXGroup()
	c := cron.New()
	spec := "57 13 * * *"
	if _, err := c.AddFunc(spec, Push2WXGroup); err != nil {
		log.Println("创建定时器失败:", err)
		return
	}
	c.Start()
	log.Println("定时任务启动ok.")
	select {}
	
}

func main() {
	//CronPushJob()
	//getZSXQData("https://api.zsxq.com/v1.10/topics/548488425851254")
	//getZSXQData("https://api.zsxq.com/v1.10/groups/1824528822/topics?scope=all&count=20")
	getZSXQData("https://api.zsxq.com/v1.10/topics/548488425851254")
}

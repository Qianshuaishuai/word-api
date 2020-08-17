package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	loglib "github.com/HYY-yu/LogLib"
	"gitlab.dreamdev.cn/ebag/poetry-api/helper"
)

var (
	curlIdClient  *http.Client
	curlIdClient2 *http.Client
)

type CurlReseponId struct {
	F_id string `json:"F_id"`
}

type CurlReseponIntId struct {
	F_id int `json:"F_id"`
}

func init() {
	curlIdClient = &http.Client{
		Transport: &http.Transport{
			//			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression:    true,
			ResponseHeaderTimeout: time.Second * 0,
		},
	}
	curlIdClient2 = &http.Client{
		Transport: &http.Transport{
			//			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression:    true,
			ResponseHeaderTimeout: time.Second * 0,
		},
	}
}

type MSnowflakCurl struct {
}

//获取发号器发出的ID(string类型,20位)
func (u *MSnowflakCurl) GetId() (id string) {
	id = ""
	uniqueFlag := helper.GetGuid()

	uri := MyConfig.SnowFlakDomain + "/v1/snowflak/id"
	method := "GET"
	req, _ := http.NewRequest(method, uri, nil)
	req.Header.Set("Accept", "application/json")
	auth := MyConfig.SnowFlakAuthUser + ":" + helper.Md5(MyConfig.SnowFlakAuthUserSecurity)
	req.Header.Set("Authentication", auth)

	client := curlIdClient

	//log request
	loglib.GetLogger().LogSnowflakRequest(uniqueFlag, uri, auth)

	resp, err := client.Do(req)
	idObj := CurlReseponId{}
	if err == nil {
		defer resp.Body.Close()
		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			dec := json.NewDecoder(strings.NewReader(string(bodyByte)))
			dec.UseNumber()
			dec.Decode(&idObj)
		}
		if resp.Status == "200 OK" {
			id = idObj.F_id
		}
		//log response
		loglib.GetLogger().LogSnowflakResponse(uniqueFlag, idObj.F_id, resp.Status, string(bodyByte))
	} else {
		//log err
		loglib.GetLogger().LogErr(err, "snowflak module")
	}

	return
}

//获取发号器发出的ID(int类型,16位)
func (u *MSnowflakCurl) GetIntId() (id int) {
	id = 0
	uniqueFlag := helper.GetGuid()

	uri := MyConfig.SnowFlakDomain + "/v1/snowflak/intId"
	method := "GET"
	req, _ := http.NewRequest(method, uri, nil)
	req.Header.Set("Accept", "application/json")
	auth := MyConfig.SnowFlakAuthUser + ":" + helper.Md5(MyConfig.SnowFlakAuthUserSecurity)
	req.Header.Set("Authentication", auth)

	client := curlIdClient2

	//log request
	loglib.GetLogger().LogSnowflakRequest(uniqueFlag, uri, auth)

	resp, err := client.Do(req)
	idObj := CurlReseponIntId{}
	if err == nil {
		defer resp.Body.Close()
		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			dec := json.NewDecoder(strings.NewReader(string(bodyByte)))
			dec.UseNumber()
			dec.Decode(&idObj)
		}
		if resp.Status == "200 OK" {
			id = idObj.F_id
		}
		//log response
		loglib.GetLogger().LogSnowflakResponse(uniqueFlag, strconv.Itoa(idObj.F_id), resp.Status, string(bodyByte))
	} else {
		//log err
		loglib.GetLogger().LogErr(err, "snowflak module")
	}
	return
}

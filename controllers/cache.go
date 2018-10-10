package controllers

import (
	"NeteaseCloudMusic/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/context"
	"net/http"
	"strings"
	"time"
)

var (
	host   = ""
	db     = ""
	key    = ""
	config = ""
)

func init() {
	host = beego.AppConfig.String("redis_host")
	db = beego.AppConfig.String("redis_db")
	key = beego.AppConfig.String("redis_key")

	configMap := make(map[string]string)
	configMap["conn"] = host
	configMap["key"] = key
	configMap["dbNum"] = db

	if configByte, err := json.Marshal(configMap); err != nil {
		fmt.Println(err.Error())
	} else {
		config = string(configByte)
	}
}

func getRedis() (cache.Cache, error) {
	return cache.NewCache("redis", config)
}

func ReadApiCache(ct *context.Context) {
	input := ct.Input

	var reqJson interface{}
	json.Unmarshal(input.RequestBody, &reqJson)

	fmt.Println("RequestBody:", reqJson)

	if redis, err := getRedis(); err != nil {
		beego.Error(err.Error())
		ct.Abort(http.StatusInternalServerError, err.Error())
	} else {
		if cacheJson := redis.Get(input.URI()); cacheJson != nil {

			// TODO: uncompleted
			cacheJsonObj := utils.TransformStructToMap(cacheJson.([]byte))

			requestBody := cacheJsonObj["requestBody"]

			if strings.EqualFold(requestBody.(string), string(input.RequestBody[:])) {
				fmt.Println("IS EQUAL")
			}

			ct.Output.SetStatus(http.StatusOK)
			ct.Output.Body(cacheJson.([]byte))
		}
	}
}

func WriteApiCache(ct *context.Context, response interface{}) {
	input := ct.Input

	if redis, err := getRedis(); err != nil {
		beego.Error(err.Error())
	} else {
		var requestBody interface{}
		json.Unmarshal(input.RequestBody, &requestBody)

		cacheData := make(map[string]interface{})
		cacheData["requestBody"] = requestBody
		cacheData["responseBody"] = response
		cacheByte, _ := json.Marshal(cacheData)

		err := redis.Put(input.URI(), cacheByte, time.Second*60*2)
		if err != nil {
			beego.Error(err.Error())
		}
	}
}

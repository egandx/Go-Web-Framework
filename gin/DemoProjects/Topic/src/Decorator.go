package src

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
)

// CacheDecorator cache decorator
func CacheDecorator(h gin.HandlerFunc, param string, redKeyPattern string, empty interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// redis determine
		getId := context.Param(param)
		redisKey := fmt.Sprintf(redKeyPattern, getId)

		conn := RedisDefaultPool.Get()
		defer conn.Close()

		ret, err := redis.Bytes(conn.Do("get", redisKey))

		if err != nil { //redis hasn't v
			h(context) //执行目标方法
			dbResult, exists := context.Get("dbResult")
			if !exists {
				dbResult = empty
			}

			retData, _ := json.Marshal(dbResult)
			conn.Do("setex",redisKey,20,retData)

			//if topics.TopicID == 0{  //DB 没有匹配到
			//	conn.Do("setex",redisKey,20,retData)
			//}else{ //DB正常数据，缓存50s
			//	conn.Do("setex",redisKey,50,retData)
			//}

			context.JSON(http.StatusOK,dbResult)
			log.Println("从数据库读取")

		} else { //redis cache has v
			json.Unmarshal(ret, &empty)
			context.JSON(http.StatusOK, empty)
			log.Println("从Redis读取")
		}
	}

}

package src

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
	"reflect"
)

// QueryIdCacheDecorator cache decorator
func QueryIdCacheDecorator(h gin.HandlerFunc, param string, redKeyPattern string, empty interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// redis determine
		getId := context.Param(param)
		redisKey := fmt.Sprintf(redKeyPattern, getId)

		conn := RedisDefaultPool.Get()
		defer conn.Close()

		ret, err := redis.Bytes(conn.Do("get", redisKey))

		if err != nil { //redis hasn't v
			h(context) //执行目标方法
			dbResult, exists := context.Get("dbResultTopicDetail")
			if !exists {
				dbResult = empty
			}

			retData, _ := json.Marshal(dbResult)

			obj := reflect.ValueOf(dbResult)
			id := obj.Field(0).Interface() //反射拿值 topicID

			if id == 0 { //DB didn't match the value,Timeout 20s
				conn.Do("setex", redisKey, 20, retData)

			} else { //DB could match the value,Timeout 50s
				conn.Do("setex", redisKey, 50, retData)

			}

			context.JSON(http.StatusOK, dbResult)
			log.Println("从数据库读取")

		} else { //redis cache has v
			json.Unmarshal(ret, &empty)
			context.JSON(http.StatusOK, empty)
			log.Println("从Redis读取")
		}
	}

}

// QueryAllCacheDecorator cache decorator
func QueryAllCacheDecorator(h gin.HandlerFunc, redKeyPattern string, empty interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// redis determine
		redisKey := fmt.Sprintf(redKeyPattern)

		conn := RedisDefaultPool.Get()
		defer conn.Close()

		ret, err := redis.Bytes(conn.Do("get", redisKey))

		if err != nil { //redis hasn't v
			h(context) //执行目标方法
			dbResult, exists := context.Get("dbResultTopicList")
			if !exists {
				dbResult = empty
			}

			retData, _ := json.Marshal(dbResult)

			obj := reflect.ValueOf(dbResult)
			length := obj.Len() //length

			if length == 0 { //DB No data available
				conn.Do("setex", redisKey, 20, retData)

			} else { //DB data available
				conn.Do("setex", redisKey, 50, retData)
			}

			context.JSON(http.StatusOK, dbResult)
			log.Println("从数据库读取")

		} else { //redis cache has v
			json.Unmarshal(ret, &empty)
			context.JSON(http.StatusOK, empty)
			log.Println("从Redis读取")
		}
	}

}

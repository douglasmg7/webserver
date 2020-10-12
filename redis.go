package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
)

func ConnectRedis() {
	// Connect to Redis DB.
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := redisDB.Ping().Result()
	if err != nil || pong != "PONG" {
		log.Panicf("[panic] Couldn't connect to Redis DB. %s", err)
	}
}

// Get
func redisGet(key string) string {
	val, err := redisDB.Get(key).Result()
	if err == redis.Nil {
		return ""
		// log.Printf("Key not exist")
	} else if err != nil {
		// checkError(err)
		log.Printf("RedisGet error: %v", err)
		return ""
	} else {
		return val
		// log.Printf("Val: %+v\n", val)
	}
}

// Set
func redisSet(key string, val string, exp time.Duration) error {
	err := redisDB.Set(key, val, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

// Del
func redisDel(key string) error {
	return redisDB.Del(key).Err()
}

////////////////////////////////////////////////////////////////////////////////
// User id
////////////////////////////////////////////////////////////////////////////////
func makeSessionIDKey(uuid string) string {
	return "sessionid_" + uuid
}

func redisSetUserID(uuid string, userID string) {
	_ = redisDB.Set(makeSessionIDKey(uuid), userID, time.Hour*8760)
}

func redisGetUserID(uuid string) string {
	return redisGet(makeSessionIDKey(uuid))
}

func redisDelUserID(uuid string) error {
	return redisDel(makeSessionIDKey(uuid))
}

////////////////////////////////////////////////////////////////////////////////
// Session
////////////////////////////////////////////////////////////////////////////////
func makeUserIDKey(userID string) string {
	return "userid_" + userID
}

func redisSetSession(userID string, session *Session) {
	sessionJSON, err := json.Marshal(session)
	// if checkError(err) {
	// return
	// }
	if err != nil {
		log.Printf("[error] Marshalling session. Error: %v", err)
	}
	_ = redisDB.Set(makeUserIDKey(userID), sessionJSON, time.Hour*8760)
}

func redisGetSession(userID string) (session *Session) {
	sessionJSON := redisGet(makeUserIDKey(userID))
	if sessionJSON == "" {
		return
	}
	// log.Printf("frSJson: %s\n", frSJson)
	err := json.Unmarshal([]byte(sessionJSON), &session)
	// if checkError(err) {
	// return frS, false
	// }
	if err != nil {
		log.Printf("[error] Unmarshalling session. Error: %v", err)
	}
	return
}

// Get Correios estimate delivery.
func redisDelSession(userID string) error {
	return redisDel(makeUserIDKey(userID))
}

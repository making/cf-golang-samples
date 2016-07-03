// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Sample memcache demonstrates use of a memcached client from App Engine flexible environment.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
	"github.com/garyburd/redigo/redis"
)

func main() {
	var err error
	uri := getConnectionURI()
	log.Printf("Connect %s", uri)
	conn, err := redis.DialURL(uri)
	if err != nil {
		log.Fatal("connection failure: ", err)
	}
	defer conn.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r, conn)
	})
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "4000"
	}
	log.Printf(fmt.Sprintf("Listening at %s", port))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func getConnectionURI() string {
	appEnv, err := cfenv.Current()
	if err != nil {
		log.Fatal("VCAP_APPLICATION and VCAP_SERVICES must be set!")
	}
	redisService, err := appEnv.Services.WithName("redis-db")
	log.Println(redisService)
	if err != nil {
		log.Fatal(err)
	}
	var ok bool
	var hostname string
	hostname, ok = redisService.Credentials["host"].(string)
	if !ok {
		hostname, ok = redisService.Credentials["hostname"].(string)
		if !ok {
			log.Fatal("No valid Redis hostname\n")
		}
	}
	port, ok := redisService.Credentials["port"]
	if !ok {
		log.Fatal("No valid Redis port\n")
	}
	password, ok := redisService.Credentials["password"].(string)
	if !ok {
		log.Fatal("No valid Redis password\n")
	}
	return fmt.Sprintf("redis://:%s@%s:%v", password, hostname, port)
}

const key = "count"

func handle(w http.ResponseWriter, r *http.Request, conn redis.Conn) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	conn.Do("INCR", key)
	value, err := redis.Int(conn.Do("GET", key))
	if err != nil {
		msg := fmt.Sprintf("Could not populate value: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, "Count: %d", value)
	}
}

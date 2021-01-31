// Copyright 2021 Shopee, Inc.
// author pengchengbai
// date 2021/1/23

package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/smtp"
)

func main() {
	fmt.Printf("hello %s", "world")
	fmt.Sprintf("hello %s", "world")


}

var users = make(map[string]int64)

func bytesInUse(username string) int64 { return users[username] /* ... */ }
// Email sender configuration.
// NOTE: never put passwords in source code!
const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"
const template = `Warning: you are using %d bytes of storage, %d%% of your quota.`

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender,
		[]string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendEmail(%s) failed: %s", username, err)
	}
}

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return // OK
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}

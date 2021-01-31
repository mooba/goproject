// Copyright 2021 Shopee, Inc.
// author pengchengbai
// date 2021/1/23

package main

import (
	"strings"
	"testing"
	"unicode"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	// Save and restore original notifyUser
	var saved = notifyUser
	defer func() {
		notifyUser = saved
	}()

	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}
	const user = "joe@example.org"
	// ...simulate a 980MB-used condition...
	users[user] = 980_000_000

	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}

	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}

	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func IsPalindrome(s string) bool {
	var letters = make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

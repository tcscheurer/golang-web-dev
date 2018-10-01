package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u user
	if s, ok := dbSessions[c.Value]; ok { // if , ok idiom
		s.lastActivity = time.Now() // update lastActivity
		dbSessions[c.Value] = s     // Update the value stored in the sessions map
		u = dbUsers[s.un]           // update found user locally
	}
	return u // return the found user from the Users map
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false // If the session cookie not found, User isn't logged in
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now() // If we found the session in the map, update lastActivity
		dbSessions[c.Value] = s     // Put updated value in map
	}
	_, ok = dbUsers[s.un] // store if we found the user locally
	// refresh session
	c.MaxAge = sessionLength // Update the maxAge of cookie, again we're using 30 seconds here
	http.SetCookie(w, c)     // set cookie with response writer, pointer to the cookie we found
	return ok                // return boolean that says we found the user or not
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}

// for demonstration purposes
func showSessions() {
	fmt.Println("********")
	for k, v := range dbSessions { // range over the sessions map and print them
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}

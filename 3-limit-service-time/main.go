//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"sync/atomic"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
/* func HandleRequest(process func(), u *User) bool {
	chProc := make(chan bool)
	go func() {
		process()
		chProc <- true
	}()

	if !u.IsPremium {
		go func() {
			ticker := time.NewTicker(10 * time.Second)
			<-ticker.C
			chProc <- false
		}()
	}

	return <-chProc
}
*/

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed (Advanced Level)
func HandleRequest(process func(), u *User) bool {

	chProc := make(chan bool)
	go func() {
		process()
		chProc <- true
	}()

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			if atomic.AddInt64(&u.TimeUsed, 1) >= 10 && !u.IsPremium {
				return false
			}
		case <-chProc:
			return true
		}
	}
}

func main() {
	RunMockServer()
}

package dms

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	ErrNoDuration = "[DMS]: no duration given"
	ErrNoUsername = "[DMS]: no username given"
	ErrNoPassword = "[DMS]: no password given"
	ErrNoFunction = "[DMS]: no function given"
)

var cyan = color.New(color.FgCyan)

// A DeadManSwitch struct stores values such as Duration and Credentials
type DeadMansSwitch struct {
	// TimerDuration specifies the timespan
	// after which it will run OnTrigger
	// if the Dead Mans Switch was not reset
	TimerDuration time.Duration

	// Username is the username for the Dead
	// Mans Switch which has to be provided
	// via BasicAuth
	Username string

	// Password is the Password for the Dead
	// Mans Switch which has to be provided
	// via BasicAuth
	Password string

	// OnTrigger is the function which gets
	// run once the timer runs out because
	// it was not reset
	OnTrigger func()
}

// SetTimerDuration sets the Timer Duration of the Dead Mans Switch
func (d *DeadMansSwitch) SetTimerDuration(t time.Duration) *DeadMansSwitch {
	d.TimerDuration = t
	return d
}

// SetPassword sets the Password of the Dead Mans Switch
func (d *DeadMansSwitch) SetPassword(p string) *DeadMansSwitch {
	d.Password = p
	return d
}

// SetOnTrigger sets the OnTrigger function of the Dead Mans Switch
func (d *DeadMansSwitch) SetOnTrigger(f func()) *DeadMansSwitch {
	d.OnTrigger = f
	return d
}

// Handler returns a http.HandlerFunc which can be used in every web framework/router
func (d *DeadMansSwitch) Handler() http.HandlerFunc {
	if d.TimerDuration == 0 {
		panic(errors.New(ErrNoDuration))
	}
	if d.Username == "" {
		panic(errors.New(ErrNoUsername))
	}
	if d.Password == "" {
		panic(errors.New(ErrNoPassword))
	}
	if d.OnTrigger == nil {
		panic(errors.New(ErrNoFunction))
	}

	timer := time.AfterFunc(d.TimerDuration, d.OnTrigger)

	return func(rw http.ResponseWriter, r *http.Request) {
		username, password, _ := r.BasicAuth()
		if strings.Compare(d.Password, password) != 0 {
			fmt.Println(strings.Compare(d.Password, password))
			return
		}
		if strings.Compare(d.Username, username) != 0 {
			fmt.Println(strings.Compare(d.Username, username))
			return
		}
		ok := timer.Reset(d.TimerDuration)
		if !ok {
			panic("timer done")
		}
		cyan.Println(time.Now().Format("[DMS: 15:04:05.000]") + " Timer reset")
	}
}

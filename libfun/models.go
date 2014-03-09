package libfun

import "time"

type Account struct {
	Id        int
	Name      string
	LastLogin time.Time
}

/*

You just need a constructor. A common used pattern is

func NewSyncMap() *SyncMap {
    return &SyncMap{hm: make(map[string]string)}
}
In case of more fields inside your struct, starting a goroutine as backend, or registering a finalizer everything could be done in this constructor.

func NewSyncMap() *SyncMap {
    sm := SyncMap{
        hm: make(map[string]string),
        foo: "Bar",
    }

    runtime.SetFinalizer(sm, (*SyncMap).stop)

    go sm.backend()

    return sm
}
*/
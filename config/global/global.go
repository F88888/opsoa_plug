package global

import (
	"opsoa_plug/pkg"
	"sync"
)

var (
	JobList jobList
	Key     string
	Port    int
)

type jobList struct {
	Data map[uint32]pkg.Exec
	Lock sync.Mutex
}

func (j *jobList) Get(k uint32) (pkg.Exec, bool) {
	j.Lock.Lock()
	defer j.Lock.Unlock()
	jobInfo, err := j.Data[k]
	return jobInfo, err
}

func (j *jobList) Set(k uint32, v pkg.Exec) {
	j.Lock.Lock()
	defer j.Lock.Unlock()
	if j.Data == nil {
		j.Data = make(map[uint32]pkg.Exec)
	}
	j.Data[k] = v
}

func (j *jobList) Del(k uint32) {
	j.Lock.Lock()
	defer j.Lock.Unlock()
	if _, err := j.Data[k]; err {
		delete(j.Data, k)
	}
}

func (j *jobList) Scan() map[uint32]pkg.Exec {
	j.Lock.Lock()
	defer j.Lock.Unlock()
	return j.Data
}

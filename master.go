// mqt = mesos query tool
// tool to query info from mesos
// Type of information that can be queried
// - System overall
// - Master
// - Slave
// - frameworks
// - tasks
//
// # Query on partial attribute names
// $ mqt system --attr="build*"

// # Query with attr an value filters
// $ mqt master --attr="cluster" --val="3"

// $ mqt

package mqt

import (
	"encoding/json"
	"fmt"
	"io"
)

type MesosData struct {
	State State
}

type State struct {
	ActivatedSlaves   int64   `json:"activated_slaves"`
	BuildDate         string  `json:"build_date"`
	BuildTime         float64 `json:"build_time"`
	BuildUser         string  `json:"build_user"`
	Cluster           string  `json:"cluster"`
	DeactivatedSlaves int64   `json:"deactivated_slaves"`
	ElectedTime       float64 `json:"elected_time"`
	FailedTasks       int64   `json:"failed_tasks"`
	FinishedTasks     int64   `json:"finished_tasks"`
	GitSha            string  `json:"git_sha"`
	GitTag            string  `json:"git_tag"`
	Hostname          string  `json:"hostname"`
	Id                string  `json:"id"`
	KilledTasks       int64   `json:"killed_tasks"`
	Leader            string  `json:"leader"`
	LogDir            string  `json:"log_dir"`
	LostTasks         int64   `json:"lost_tasks"`
	StagedTasks       int64   `json:"staged_tasks"`
	StartTime         string  `json:"start_time"`
	StartedTasks      string  `json:"started_tasks"`
	Version           string  `json:"version"`
}

// Does a healthcheck on master.
// See http://<master-ip>:<port>/help/master/health
func HealthCheck() bool {
	return true
}

func Query(q string) {

}

func decodeState(data []byte) (mState State, err error) {
	if len(data) == 0 || data == nil {
		return State{}, nil
	}
	err = json.Unmarshal(data, &mState)

	if err != nil {
		return State{}, err
	}
	return
}

func search(attr string, source io.Reader) (string, error) {
	var result interface{}
	err := json.NewDecoder(source).Decode(&result)
	if err != nil {
		return "", err
	}
	return find(attr, "", result), nil
}

func find(node string, val string, tree interface{}) string {
	switch t := tree.(type) {
	case map[string]interface{}:
		for k, v := range t {
			return find(node, val, v)
		}
	case []interface{}:
		//for i, v := range t {}

	case string:
		return t
	case float64:
		return fmt.Sprintf("%.2f", t)
	case bool:
		return fmt.Sprintf("%t", t)
	}

	return ""
}

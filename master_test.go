package mqt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	data := map[string]interface{}{
  		"activated_slaves": 3,
  		"build_date": "2015-01-19 22:32:25",
  		"build_time": 1421706745,
  		"build_user": "root",
   		"deactivated_slaves": 0,
  		"elected_time": 1427074024.07139,
  		"failed_tasks": 0,
  		"finished_tasks": 0,
    }

	f  := find("build_user", "", data)
	assert.Equal(t, f, "root")

}

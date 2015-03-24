package mqt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeState(t *testing.T) {
	data := []byte(`
		{
  			"activated_slaves": 3,
  			"build_date": "2015-01-19 22:32:25",
  			"build_time": 1421706745,
  			"build_user": "root",
   			"deactivated_slaves": 0,
  			"elected_time": 1427074024.07139,
  			"failed_tasks": 0,
  			"finished_tasks": 0
  		}
	`)

	state, err := decodeState(data)
	assert.NoError(t, err)
	assert.NotEmpty(t, state)
	assert.Equal(t, 3, state.ActivatedSlaves)
	assert.Equal(t, 1427074024.07139, state.ElectedTime)
	assert.Equal(t, "root", state.BuildUser)

}

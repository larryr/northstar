package roger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkAssignment(t *testing.T, assignmentObj interface{}) {
	client, _ := NewRClient("localhost", 6311)
	sess, _ := client.GetSession()
	defer sess.Close()
	err := sess.Assign("assignedVar", assignmentObj)
	assert.Equal(t, err, nil)
	obj, err := sess.Eval("assignedVar")
	assert.Equal(t, obj, assignmentObj)
	assert.Equal(t, err, nil)
}

func TestIntArrayAssignment(t *testing.T) {
	checkAssignment(t, []int32{1, 2, 3, 4, 5})
}

func TestFloatArrayAssignment(t *testing.T) {
	checkAssignment(t, []float64{1.1, 2.2, 3.3, 4.4, 5.5})
}

func TestStringArrayAssignment(t *testing.T) {
	checkAssignment(t, []string{"test", "str"})
}

func TestByteArrayAssignment(t *testing.T) {
	checkAssignment(t, []byte{'g', 'o', 'l', 'a', 'n', 'g'})
}

func TestStringAssignment(t *testing.T) {
	checkAssignment(t, "testing")
}

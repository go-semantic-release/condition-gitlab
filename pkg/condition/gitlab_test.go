package condition

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitlabValid(t *testing.T) {
	gl := GitLab{}
	os.Setenv("CI_COMMIT_BRANCH", "")
	err := gl.RunCondition(map[string]string{"defaultBranch": ""})
	assert.EqualError(t, err, "this test run is not running on a branch build")
}

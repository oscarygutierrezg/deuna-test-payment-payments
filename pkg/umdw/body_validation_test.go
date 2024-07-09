package umdw

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequireKeysSuccess(t *testing.T) {
	assert := assert.New(t)
	m := map[string]interface{}{
		"os":    "linux",
		"agent": "go-client",
		"specs": map[string]interface{}{
			"core": "v1.1",
		},
	}
	r := []string{
		"os",
		"agent",
		"specs.core",
	}

	err := requireKeys(m, r)

	assert.Nil(err)
}

func TestRequireKeysError(t *testing.T) {
	assert := assert.New(t)
	m := map[string]interface{}{
		"os":    "linux",
		"agent": "go-client",
		"specs": map[string]interface{}{
			"core": "",
		},
	}
	r := []string{
		"os",
		"agent",
		"specs.core",
	}

	err := requireKeys(m, r)

	assert.NotNil(err)
	assert.Equal("specs.core is required", err.Error())
}

func TestVerificationKeysSuccess(t *testing.T) {
	assert := assert.New(t)
	m := map[string]interface{}{
		"os":    "linux",
		"agent": "go-client",
		"specs": map[string]interface{}{
			"core": "v1.1",
		},
	}
	isLinux := func(os interface{}) bool {
		return os.(string) == "linux"
	}
	isGoAgent := func(agent interface{}) bool {
		return agent.(string) == "go-client"
	}
	iscoreV11 := func(v interface{}) bool {
		return v.(string) == "v1.1"
	}
	vf := VerificationFunctions{
		"os": {
			Func:   isLinux,
			ErrMsg: "OS invalid",
		},
		"agent": {
			Func:   isGoAgent,
			ErrMsg: "Agent invalid",
		},
		"specs.core": {
			Func:   iscoreV11,
			ErrMsg: "Version Core invalid",
		},
	}

	err := verificationKeys(m, vf)

	assert.Nil(err)
}

func TestVerificationKeysError(t *testing.T) {
	assert := assert.New(t)
	m := map[string]interface{}{
		"os":    "linux",
		"agent": "go-client",
		"specs": map[string]interface{}{
			"core": "v1.1",
		},
	}
	isLinux := func(os interface{}) bool {
		return os.(string) == "linux"
	}
	isGoAgent := func(agent interface{}) bool {
		return agent.(string) == "go-client"
	}
	iscoreV12 := func(v interface{}) bool {
		return v.(string) == "v1.2"
	}
	vf := VerificationFunctions{
		"os": {
			Func:   isLinux,
			ErrMsg: "OS invalid",
		},
		"agent": {
			Func:   isGoAgent,
			ErrMsg: "Agent invalid",
		},
		"specs.core": {
			Func:   iscoreV12,
			ErrMsg: "Version Core invalid",
		},
	}

	err := verificationKeys(m, vf)

	assert.NotNil(err)
	assert.Equal("specs.core: Version Core invalid", err.Error())
}

func TestBodyVerificationKeys(t *testing.T) {
	assert := assert.New(t)
	m := map[string]interface{}{
		"os":    "linux",
		"agent": "go-client",
		"specs": map[string]interface{}{
			"core": "v1.1",
		},
	}
	r := []string{
		"os",
		"agent",
		"specs.core",
	}
	vf := VerificationFunctions{
		"os": {
			Func:   func(os interface{}) bool { return os.(string) == "linux" },
			ErrMsg: "OS invalid",
		},
		"agent": {
			Func:   func(agent interface{}) bool { return agent.(string) == "go-client" },
			ErrMsg: "Agent invalid",
		},
		"specs.core": {
			Func:   func(v interface{}) bool { return v.(string) == "v1.1" },
			ErrMsg: "Version Core invalid",
		},
	}

	err := BodyVerificationKeys(m, r, vf)

	assert.Nil(err)
}

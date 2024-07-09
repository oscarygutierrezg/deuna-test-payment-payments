package umdw

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type VerificationFunctions map[string]VerificationKeyFunction

type VerificationKeyFunction struct {
	Func   func(val interface{}) bool
	ErrMsg string
}

func BodyVerificationKeys(body map[string]interface{}, require []string, verify VerificationFunctions) error {
	var err error

	err = requireKeys(body, require)
	if err != nil {
		return err
	}

	err = verificationKeys(body, verify)
	if err != nil {
		return err
	}

	return nil
}

func BodyVerifyFields(c *gin.Context, require []string, verify VerificationFunctions) error {
	if c.IsAborted() {
		return errors.New("c is aborted")
	}

	body := c.Keys[BodyKey].(map[string]interface{})

	return BodyVerificationKeys(body, require, verify)
}

// private functions
func requireKeys(m map[string]interface{}, paths []string) error {
	for _, p := range paths {

		val, err := GetPathFromMap(m, p)
		if err != nil {
			return fmt.Errorf("%s is required", p)
		}

		if s, ok := val.(string); ok && s == "" {
			return fmt.Errorf("%s is required", p)
		}

		if f, ok := val.(float64); ok && f == 0 {
			return fmt.Errorf("%s is required", p)
		}

		if arr, ok := val.([]interface{}); ok && len(arr) == 0 {
			return fmt.Errorf("%s is required", p)
		}
	}

	return nil
}

func verificationKeys(m map[string]interface{}, vf VerificationFunctions) error {
	for k, fk := range vf {

		if val, err := GetPathFromMap(m, k); err == nil {

			if val == "" {
				return nil
			}

			if ok := fk.Func(val); !ok {
				errMsg := fmt.Sprintf("%s: %s", k, fk.ErrMsg)
				return errors.New(errMsg)
			}
		}
	}

	return nil
}

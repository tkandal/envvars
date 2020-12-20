package envvars

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func EnvWrapper(envs ...string) error {
	prefix := strings.ToUpper(filepath.Base(os.Args[0])) + "_"
	for _, env := range envs {
		val, ok := os.LookupEnv(prefix + env)
		if !ok {
			// Not defined
			continue
		}
		err := os.Setenv(env, val)
		if err != nil {
			return fmt.Errorf("set environment variable %s failed; error = %v", env, err)
		}
	}
	return nil
}

package envvars

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// EnvWrapper checks if environment variables are prefixed with the program's name
// plus separator, and translates them to environment variables without the prefix.
func EnvWrapper(sep string, envs ...string) error {
	prefix := strings.ToUpper(filepath.Base(os.Args[0])) + sep
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

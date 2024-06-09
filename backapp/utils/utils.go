package utils

import "os"

func GetEnvWithDefault(key, defaultValue string) string {
	res, ok := os.LookupEnv(key);
	if !ok {
		res = defaultValue;
	}
	return res
}

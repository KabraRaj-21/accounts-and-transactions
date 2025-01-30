package setup

import "os"

func SetupEnvironmentVariables(testEnv *TestEnvironment) {
	os.Setenv("DB_HOST", testEnv.MySQL.Host())
	os.Setenv("DB_PORT", testEnv.MySQL.Port())
	os.Setenv("DB_USER_NAME", testEnv.MySQL.Username())
	os.Setenv("DB_PASSWORD", testEnv.MySQL.Password())
	os.Setenv("DB_NAME", testEnv.MySQL.Database())
}

package settings

import (
	"os"
	"strings"
)

var defaults map[string]string

func init() {
	defaults = map[string]string{
		"db_user":     "root",
		"db_obj":      "sqlite3",
		"db_password": "",
		"db_host":     "127.0.0.1:3306",
		"db_name":     "opt",
		"db_log":      "xorm.log",
		"server_port": "9997",
		"server_addr": "0.0.0.0",
		"ishttps":     "false",
		"jwt_secret":  "", // "" is use random string
	}
}
func Get(key string) string {
	env := strings.TrimSpace(os.Getenv(strings.ToUpper(key)))
	if env != "" {
		return env
	}
	return defaults[key]
}

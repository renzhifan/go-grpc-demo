// Package config 站点配置信息
package config

import "helloworld/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			"port": config.Env("APP_PORT", "3000"),
		}
	})
}

// Package config 站点配置信息
package config

import "helloworld/pkg/config"

func init() {
	config.Add("grpc", func() map[string]interface{} {
		return map[string]interface{}{
			"host": config.Env("GRPC_HOST", "localhost"),
			"port": config.Env("GRPC_PORT", "50051"),
		}
	})
}

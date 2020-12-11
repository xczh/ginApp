package config

import (
	"app/internal/log"
	"path/filepath"
)

var _runtimeConfig = newConfig()

var _runtimeConfigFilePath string

func Initialize(cfgFilePath string) error {
	if cfgFilePath != "" {
		var err error
		if cfgFilePath, err = filepath.Abs(cfgFilePath); err != nil {
			log.Logger().Error("Invalid config file path: ", err)
			return err
		}
	} else {
		log.Logger().Info("No config file found, skipping load config from file")
	}

	newCfg, err := getConfigWithCfgFilePath(cfgFilePath)
	if err != nil {
		log.Logger().Error("Using config error: ", err)
		return err
	}

	log.Logger().Info("Using config file: ", cfgFilePath)
	log.Logger().Debug("runtime config: ", newCfg)

	_runtimeConfig = newCfg
	_runtimeConfigFilePath = cfgFilePath

	return nil
}

func getConfigWithCfgFilePath(cfgFilePath string) (*config, error) {
	newCfg := newConfig()

	// 应用配置文件
	if cfgFilePath != "" {
		if err := newCfg.applyConfigFile(cfgFilePath); err != nil {
			return nil, err
		}
	}

	// 应用环境变量
	newCfg.applyEnvVar("APP_")

	// 验证配置有效性
	if err := newCfg.validate(); err != nil {
		return nil, err
	}

	return newCfg, nil
}

func String() string {
	return _runtimeConfig.String()
}

func GetAddress() string {
	return _runtimeConfig.Address
}

func GetURLPrefix() string {
	return _runtimeConfig.URLPrefix
}

func GetMySQLDSN() string {
	return _runtimeConfig.MySQLDSN
}

func GetRedisDSN() string {
	return _runtimeConfig.RedisDSN
}

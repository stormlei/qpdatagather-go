package cachet

import (
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"time"
)

var cacheT *cache.Cache
var DefaultExpiration time.Duration

func Init(conf *viper.Viper) {
	de := conf.GetInt("cache.defaultExpiration")
	ci := conf.GetInt("cache.cleanupInterval")
	DefaultExpiration = time.Duration(de) * time.Second
	cleanupInterval := time.Duration(ci) * time.Minute
	cacheT = cache.New(DefaultExpiration, cleanupInterval)
}

func GetCache() *cache.Cache {
	return cacheT
}

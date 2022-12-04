package beanstalk

import "github.com/sllt/tao/core/stores/redis"

type (
	Beanstalk struct {
		Endpoint string
		Tube     string
	}

	DqConf struct {
		Beanstalks []Beanstalk
		Redis      redis.RedisConf
	}
)

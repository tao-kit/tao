package beanstalk

import "manlu.org/tao/core/stores/redis"

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

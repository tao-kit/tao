package gormc

import (
	"database/sql"
	"gorm.io/gorm"
	"manlu.org/tao/core/stores/cache"
	"manlu.org/tao/core/stores/redis"
	"manlu.org/tao/core/stores/sqlx"
	"manlu.org/tao/core/syncx"
	"time"
)

// see doc/sql-cache.md
const cacheSafeGapBetweenIndexAndPrimary = time.Second * 5

var (
	// ErrNotFound is an alias of sqlx.ErrNotFound.
	ErrNotFound = gorm.ErrRecordNotFound

	// can't use one SharedCalls per conn, because multiple conns may share the same cache key.
	exclusiveCalls = syncx.NewSharedCalls()
	stats          = cache.NewStat("gormc")
)

type (
	// ExecFn defines the sql exec method.
	ExecFn func(conn sqlx.SqlConn) (sql.Result, error)
	// IndexQueryFn defines the query method that based on unique indexes.
	IndexQueryFn func(conn sqlx.SqlConn, v interface{}) (interface{}, error)
	// PrimaryQueryFn defines the query method that based on primary keys.
	PrimaryQueryFn func(conn sqlx.SqlConn, v, primary interface{}) error
	// QueryFn defines the query method.
	QueryFn func(conn *gorm.DB, v interface{}) error

	// A CachedConn is a DB connection with cache capability.
	CachedConn struct {
		db    *gorm.DB
		cache cache.Cache
	}
)

// NewNodeConn returns a CachedConn with a redis node cache.
func NewNodeConn(db *gorm.DB, rds *redis.Redis, opts ...cache.Option) CachedConn {
	return CachedConn{
		db:    db,
		cache: cache.NewNode(rds, exclusiveCalls, stats, sql.ErrNoRows, opts...),
	}
}

// NewConn returns a CachedConn with a redis cluster cache.
func NewConn(db *gorm.DB, c cache.CacheConf, opts ...cache.Option) CachedConn {
	return CachedConn{
		db:    db,
		cache: cache.New(c, exclusiveCalls, stats, sql.ErrNoRows, opts...),
	}
}

// DelCache deletes cache with keys.
func (cc CachedConn) DelCache(keys ...string) error {
	return cc.cache.Del(keys...)
}

// GetCache unmarshals cache with given key into v.
func (cc CachedConn) GetCache(key string, v interface{}) error {
	return cc.cache.Get(key, v)
}

// QueryRow unmarshals into v with given key and query func.
func (cc CachedConn) QueryRow(v interface{}, key string, query QueryFn) error {
	return cc.cache.Take(v, key, func(v interface{}) error {
		return query(cc.db, v)
	})
}



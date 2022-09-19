package constant

const (
	//value:uid
	REDIS_USER_API_KEY = "user:apikey:%s:userid"

	//value:ratelimit实体
	REDIS_PRI_RATE_LIMIT_KEY = "openapi:prilimit:path:%s:apikey:%s"
	REDIS_PUB_RATE_LIMIT_KEY = "openapi:publimit:path:%s:ip:%s"

	//币对的redis key
	REDIS_SYMBOL_STATUS_KEY = "redis_symbol_status_"

	//market_user做市商用户key
	REDIS_MARKET_USER_KEY = "market_user"
)

const (
	REDIS_USER_API_KEY_ENTITY_KEY = "openapi:userapikey:%s"
	REDIS_USERAPIKEY_UID          = "openapi:userapikey:uid:%s"
	//value:set
	REDIS_BLACKLIST_USER_KEY = "openapi:set:trans:uid:blacklist"
)

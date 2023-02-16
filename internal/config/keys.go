// Package config handles API server configuration binding and loading.
package config

// default configuration elements and keys
const (
	configFileName = "apiserver"

	// configuration options
	keyAppName        = "app_name"
	keyConfigFilePath = "cfg"

	// server related keys
	keyBindAddress      = "server.bind"
	envBindAddress      = "SERVER_BIND"
	keyApiStateOrigin   = "server.origin"
	keyCorsAllowOrigins = "server.cors_origins"

	// server time out related keys
	keyTimeoutRead     = "server.read_timeout"
	keyTimeoutWrite    = "server.write_timeout"
	keyTimeoutIdle     = "server.idle_timeout"
	keyTimeoutHeader   = "server.header_timeout"
	keyTimeoutResolver = "server.resolver_timeout"
	keyMaxParserMemory = "server.mem_max"

	// logging related options
	keyLoggingLevel  = "log.level"
	keyLoggingFormat = "log.format"

	// node connection related options
	keyLachesisUrl = "node.url"
	envLachesisUrl = "NODE_URL"

	// IPFS node connection related options
	keyIpfsUrl           = "ipfs.url"
	envIpfsUrl           = "IPFS_URL"
	keySkipHttpGateways  = "ipfs.skip_http_gateways"
	keyIpfsGateway       = "ipfs.gateway"
	envIpfsGateway       = "IPFS_GATEWAY"
	keyIpfsGatewayBearer = "ipfs.gateway_bearer"
	envIpfsGatewayBearer = "IPFS_GATEWAY_BEARER"
	keyPinataJwt         = "ipfs.pinata_jwt"
	envPinataJwt         = "IPFS_PINATA_JWT"
	keyIpfsFileCacheDir  = "ipfs.file_cache_dir"
	envIpfsFileCacheDir  = "IPFS_FILE_CACHE_DIR"

	// off-chain database related options
	keyMongoUrl      = "db.url"
	envMongoUrl      = "DB_URL"
	keyMongoDatabase = "db.db"

	keySharedMongoUrl      = "shared_db.url"
	keySharedMongoDatabase = "shared_db.db"

	// cache related options
	keyCacheEvictionTime = "cache.eviction"
	keyCacheMaxSize      = "cache.size"

	// authentication related options
	keyAuthBearerSecret = "auth.bearer_secret"
	envAuthBearerSecret = "AUTH_BEARER_SECRET"
	keyAuthNonceSecret  = "auth.nonce_secret"
	envAuthNonceSecret  = "AUTH_NONCE_SECRET"

	// mandatory contracts
	keyWrappedFTM = "contracts.wftm"
)

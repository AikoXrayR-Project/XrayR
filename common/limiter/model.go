package limiter

type RedisConfig struct {
	Limit         int    `mapstructure:"Limit"`
	RedisAddr     string `mapstructure:"RedisAddr"` // host:port
	RedisPassword string `mapstructure:"RedisPassword"`
	RedisDB       int    `mapstructure:"RedisDB"`
	Expiry        int    `mapstructure:"Expiry"` // minute
}

package apiserver

import "time"

type Config struct {
	port                int
	readTimeout         time.Duration
	writeTimeout        time.Duration
	shutdownGracePeriod time.Duration
}

type ServerConfig func(*Config) error

func WithPort(httpPort int) ServerConfig {
	return func(c *Config) error {
		c.port = httpPort
		return nil
	}
}

func WithReadTimeout(readTimeout time.Duration) ServerConfig {
	return func(c *Config) error {
		c.readTimeout = readTimeout
		return nil
	}
}

func WithWriteTimeout(writeTimeout time.Duration) ServerConfig {
	return func(c *Config) error {
		c.writeTimeout = writeTimeout
		return nil
	}
}

func WithShutdownGracePeriod(shutdownGracePeriod time.Duration) ServerConfig {
	return func(c *Config) error {
		c.shutdownGracePeriod = shutdownGracePeriod
		return nil
	}
}

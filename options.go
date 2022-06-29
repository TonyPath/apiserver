package apiserver

import "time"

type Config struct {
	port                int
	readTimeout         time.Duration
	writeTimeout        time.Duration
	shutdownGracePeriod time.Duration
}

type Option func(*Config) error

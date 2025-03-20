package postgres

import "time"

type Configuration func(*Postgres)

func MaxPoolSize(size int) Configuration {
	return func(c *Postgres) {
		c.maxPoolSize = size
	}
}

func ConnAttempts(attempts int) Configuration {
	return func(c *Postgres) {
		c.connAttempts = attempts
	}
}

func ConnTimeout(timeout time.Duration) Configuration {
	return func(c *Postgres) {
		c.connTimeout = timeout
	}
}

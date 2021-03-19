package main

import (
	"fmt"
	"time"
)

type Config struct {
	RetryTimes  int32
	ServiceName string
	ReadTimeout time.Duration
	HTTPS       bool
	// ...
}

type Client struct {
	Config
}

func New(config Config) *Client {
	return &Client{
		config,
	}
}

type Option func(client *Client)

func NewV2(options ...Option) *Client {
	client := &Client{
		Config{
			RetryTimes:  3,
			ReadTimeout: 200 * time.Millisecond,
		},
	}
	for _, f := range options {
		f(client)
	}
	return client
}

func WithSrvName(srvname string) func(client *Client) {
	return func(client *Client) {
		client.ServiceName = srvname
	}
}

func WithReadTimeout(t time.Duration) func(client *Client) {
	return func(client *Client) {
		client.ReadTimeout = t
	}
}

func WithRetryTimes(t int32) func(client *Client) {
	return func(client *Client) {
		client.RetryTimes = t
	}
}

func main() {
	user := New(
		Config{
			ServiceName: "gmu.user.base",
			RetryTimes:  -1,
			ReadTimeout: -1,
		},
	)
	fmt.Println(user)

	_ = NewV2()
	_ = NewV2(WithReadTimeout(2 * time.Millisecond))
	_ = NewV2(WithReadTimeout(2*time.Millisecond), WithRetryTimes(3))
}

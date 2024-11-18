package routes

func ClientRoutes() {
	c, err := client.NewClient(
		client.WithDialTimeout(1*time.Second),
		client.WithMaxConnsPerHost(1024),
		client.WithMaxIdleConnDuration(10*time.Second),
		client.WithMaxConnDuration(10*time.Second),
		client.WithMaxConnWaitTimeout(10*time.Second),
		client.WithKeepAlive(true),
		client.WithClientReadTimeout(10*time.Second),
		client.WithDialer(standard.NewDialer()),
		client.WithResponseBodyStream(true),
		client.WithDisableHeaderNamesNormalizing(true),
		client.WithName("my-client"),
		client.WithNoDefaultUserAgentHeader(true),
		client.WithDisablePathNormalizing(true),
		client.WithRetryConfig(
			retry.WithMaxAttemptTimes(3),
			retry.WithInitDelay(1000),
			retry.WithMaxDelay(10000),
			retry.WithDelayPolicy(retry.DefaultDelayPolicy),
			retry.WithMaxJitter(1000),
		),
		client.WithWriteTimeout(10*time.Second),
		client.WithConnStateObserve(stateFunc, observeInterval),
		client.WithDialFunc(customDialFunc, netpoll.NewDialer()),
		client.WithHostClientConfigHook(func(hc interface{}) error {
			if hct, ok := hc.(*http1.HostClient); ok {
				hct.Addr = "FOO.BAR:443"
			}
			return nil
		})
	)
	if err != nil {
		return
	}
}
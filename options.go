package blastr

type Option func(*Options)

type Options struct {
	relayURLs    []string
	strictErrors bool
}

func WithCustomRelays(relayURLs []string) Option {
	return func(o *Options) {
		o.relayURLs = relayURLs
	}
}

func WithStrictErrors() Option {
	return func(o *Options) {
		o.strictErrors = true
	}
}

func defaultOptions() *Options {
	return &Options{
		relayURLs:    []string{"wss://nostr.mutinywallet.com"},
		strictErrors: false,
	}
}

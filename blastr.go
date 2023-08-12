package blastr

import (
	"context"
	"fmt"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

func New(nsec string, opts ...Option) (*Blastr, error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	var npub, pubkey, privateKey string
	{
		_, sk, err := nip19.Decode(nsec)
		if err != nil {
			return nil, fmt.Errorf("nip19 decode: %w", err)
		}
		privateKey = sk.(string)

		pk, err := nostr.GetPublicKey(privateKey)
		if err != nil {
			return nil, fmt.Errorf("get pubkey: %w", err)
		}
		pubkey = pk

		_npub, err := nip19.EncodePublicKey(pubkey)
		if err != nil {
			return nil, fmt.Errorf("encode pubkey: %w", err)
		}
		npub = _npub
	}

	return &Blastr{
		opts:       options,
		npub:       npub,
		pubkey:     pubkey,
		privateKey: privateKey,
	}, nil
}

type Blastr struct {
	opts                     *Options
	npub, pubkey, privateKey string
}

func (b *Blastr) SendText(ctx context.Context, content string) error {
	event := b.newEvent(content)
	return b.connectAndSend(ctx, event)
}

func (b *Blastr) Send(ctx context.Context, event nostr.Event) error {
	event.PubKey = b.pubkey
	event.CreatedAt = nostr.Now()
	event.Sign(b.privateKey)

	return b.connectAndSend(ctx, event)
}

func (b *Blastr) newEvent(content string) nostr.Event {
	event := nostr.Event{
		PubKey:    b.pubkey,
		CreatedAt: nostr.Now(),
		Kind:      nostr.KindTextNote,
		Tags:      nil,
		Content:   content,
	}
	event.Sign(b.privateKey)

	return event
}

func (b *Blastr) connectAndSend(ctx context.Context, event nostr.Event) error {
	for _, url := range b.opts.relayURLs {
		relay, err := nostr.RelayConnect(ctx, url)
		if err != nil {
			err = fmt.Errorf("connect %v: %w", url, err)
			if b.opts.strictErrors {
				return err
			}

			fmt.Println(err)
			continue
		}
		defer relay.Close()

		_, err = relay.Publish(ctx, event)
		if err != nil {
			err = fmt.Errorf("publish %v: %w", url, err)
			if b.opts.strictErrors {
				return err
			}

			fmt.Println(err)
			continue
		}
	}

	return nil
}

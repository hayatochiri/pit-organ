package pitOrgan

import (
	"golang.org/x/xerrors"
	"net/url"
	"testing"
	"time"
)

func Test_request(t *testing.T) {
	// 到達不能なエンドポイントを指定
	t.Run("ConnectionRefused", func(t *testing.T) {
		connection := newConnection(t, oandaDummy)
		connection.Timeout = time.Nanosecond // 即タイムアウトさせるため最小の待ち時間にする
		_, err := connection.Accounts().Get()

		if _, ok := nakedError(err).(*url.Error); !ok {
			t.Fatalf("Connection was not refused.\n%+v", err)
		}
	})

	// 認証エラー
	t.Run("Unauthorized", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		connection.Token = "hogehoge" // 不正なトークンに書き換え
		_, err := connection.Accounts().Get()

		if _, ok := nakedError(err).(*UnauthorizedError); !ok {
			t.Fatalf("Request was authorized.\n%+v", err)
		}
	})
}

func nakedError(err error) error {
	if u, ok := err.(xerrors.Wrapper); ok {
		return nakedError(u.Unwrap())
	}
	return err
}

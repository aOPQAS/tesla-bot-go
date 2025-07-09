package deps

import (
	"github.com/tesla/tesla-bot-go/internal/pgsql"
	tgnotifier "github.com/tesla/tesla-bot-go/internal/telegram"
)

type Deps struct {
	PG *pgsql.Client
	TG *tgnotifier.TelegramNotifier
}

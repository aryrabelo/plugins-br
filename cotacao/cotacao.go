package cotacao

import (
	"fmt"
	"os"

	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"
)

var (
	url = "http://data.fixer.io/api/latest?base=BRL&access_key=%s"
)

type retorno struct {
	Real struct {
		USD float32 `json:"USD"`
		EUR float32 `json:"EUR"`
		CAD float32 `json:"CAD"`
		GBP float32 `json:"GBP"`
	} `json:"rates"`
}

func cotacao(command *bot.Cmd) (msg string, err error) {
	data := &retorno{}

	url := fmt.Sprintf(url, os.Getenv("FIXER_IO_ACCESS_KEY"))

	err = web.GetJSON(url, data)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Dólar: %.2f, Euro: %.2f, CAD: %.2f, Libra: %.2f",
		1/data.Real.USD,
		1/data.Real.EUR,
		1/data.Real.CAD,
		1/data.Real.GBP), nil
}

func init() {
	bot.RegisterCommand(
		"cotacao",
		"Informa a cotação do Dólar e Euro.",
		"",
		cotacao)
}

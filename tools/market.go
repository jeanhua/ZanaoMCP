package tools

import (
	"os"

	"github.com/jeanhua/ZanaoMCP/zanao"
)

func getClient() *zanao.ZanaoClient {
	token := os.Getenv("ZANAO_TOKEN")
	schoolAlias := os.Getenv("ZANAO_SCHOOL_ALIAS")
	client := zanao.NewZanaoClient(token, schoolAlias)
	return client
}

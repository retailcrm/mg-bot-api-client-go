package v1

import (
	"os"
)

var (
	mgURL   = os.Getenv("MG_URL")
	mgToken = os.Getenv("MG_TOKEN")
)

func client() *MgClient {
	return New(mgURL, mgToken)
}

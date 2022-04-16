package settings

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"os"
)

var S snap.Client

var ServerKey string = os.Getenv("MIDTRANS_SERVER_KEY")

func SetupGlobalMidtransConfig() {
	midtrans.ServerKey = ServerKey
	midtrans.Environment = midtrans.Sandbox
}

func InitializeSnapClient() {
	S.New(ServerKey, midtrans.Sandbox)
}

package kit

import (
	"github.com/go-kit/kit/endpoint"
	httpTransport "github.com/go-kit/kit/transport/http"
)

type EndPointFunc func() (endpoint.Endpoint)

func GetAccountHttpHandler(epFunc EndPointFunc)(serverHanlder *httpTransport.Server){
	endp := epFunc()
	serverHanlder = httpTransport.NewServer(endp, DecodeRequest, EncodeResponse)
	return
}
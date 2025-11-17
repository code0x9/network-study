package gnmi

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/openconfig/gnmic/pkg/api"
	gnmiTarget "github.com/openconfig/gnmic/pkg/api/target"
	"google.golang.org/protobuf/encoding/prototext"
)

var tg *gnmiTarget.Target
var err error

func TestMain(m *testing.M) {
	tg, err = api.NewTarget(
		api.Name("srl1"),
		api.Address("172.20.20.2:57400"),
		api.Username("admin"),
		api.Password("NokiaSrl1!"),
		api.SkipVerify(true),
	)
	if err != nil {
		log.Fatal(err)
	}
	m.Run()
}

func TestCapabilities(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a gNMI client
	err = tg.CreateGNMIClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer tg.Close()

	// send a gNMI capabilities request to the created target
	capResp, err := tg.Capabilities(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prototext.Format(capResp))
}

func TestGetSystemName(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a gNMI client
	err = tg.CreateGNMIClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer tg.Close()

	// create a GetRequest
	getReq, err := api.NewGetRequest(
		api.Path("/system/name"),
		api.Encoding("json_ietf"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prototext.Format(getReq))

	// send the created gNMI GetRequest to the created target
	getResp, err := tg.Get(ctx, getReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prototext.Format(getResp))
}

func TestGetRouterIP(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a gNMI client
	err = tg.CreateGNMIClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer tg.Close()

	// create a GetRequest
	getReq, err := api.NewGetRequest(
		api.Path("/interface/subinterface/ipv4/address/ip-prefix"),
		api.Encoding("json_ietf"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prototext.Format(getReq))

	// send the created gNMI GetRequest to the created target
	getResp, err := tg.Get(ctx, getReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prototext.Format(getResp))
}

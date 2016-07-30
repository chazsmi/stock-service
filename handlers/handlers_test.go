package handlers

import (
	"flag"
	"log"
	"testing"

	"golang.org/x/net/context"

	"github.com/chazsmi/stock-service/config"
	"github.com/chazsmi/stock-service/cs"
)

var (
	FakeRowSku string = "12345"
	FakeRowAmount int32 =  1
}

func init() {
	configFile := flag.String("config_file_path", "config.yml", "path to application config file")
	flag.Parse()
	setUp(configFile)
}

func setUp(c *string) {
	c, err := config.ReadReturn(c)
	if err != nil {
		log.Fatal(err)
	}
	cs.Init(c)
	cs.Insert(FakeRowSku, FakeRowAmount)
}

func tearDown() {

}
		
func TestCheck(t *testing.T) {
	s := Stock{}
	ctx := context.Background()
	req := stock.StockRequest{
		Sku: "123456",
	}
	rsp := stock.StockReadResponse{}
	err := s.Check(ctx, &req, &rsp)
	if err != nil {
		log.Println("Check returned an error ", err)
		t.Fail()
	}
}

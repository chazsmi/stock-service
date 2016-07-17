package handlers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/chazsmi/stock-service/cs"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

type Stock struct{}

func (s *Stock) Check(ctx context.Context, req *api.Request, rsp *api.Response) error {
	sku, ok := req.Get["sku"]
	if !ok {
		return errors.BadRequest("go.micro.api.greeter", "Sku cannot be blank")
	}

	stockResult, err := cs.Get(strings.Join(sku.Values, " "))
	if err != nil {
		return errors.InternalServerError("service.stock.Stock.Check", err.Error())
	}
	b, _ := json.Marshal(map[string]string{
		"sku":    stockResult.Sku,
		"amount": strconv.Itoa(int(stockResult.Amount)),
	})
	rsp.Body = string(b)
	return nil
}

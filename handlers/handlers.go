package handlers

import (
	"github.com/chazsmi/stock-service/cs"
	"github.com/chazsmi/stock-service/proto"
	"github.com/micro/go-micro/errors"
	"golang.org/x/net/context"
)

type Stock struct {
}

func (s *Stock) Check(ctx context.Context, req *stock.StockRequest, rsp *stock.StockResponse) error {
	stockResult, err := cs.Get(req.Sku)
	if err != nil {
		return errors.InternalServerError("service.stock.Stock.Check", err.Error())
	}
	rsp.Sku = stockResult.Sku
	rsp.Amount = stockResult.Amount
	return nil
}

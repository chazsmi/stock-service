package handlers

import (
	"github.com/chazsmi/stock-service/cs"
	"github.com/chazsmi/stock-service/proto"
	"github.com/micro/go-micro/errors"
	"golang.org/x/net/context"
)

type Stock struct {
}

func (s *Stock) Check(ctx context.Context, req *stock.StockRequest, rsp *stock.StockReadResponse) error {
	stockResult, err := cs.Get(req.Sku)
	if err != nil {
		return errors.InternalServerError("charlieplc.Stock.Check", err.Error())
	}
	rsp.Sku = stockResult.Sku
	rsp.Amount = stockResult.Amount
	return nil
}

func (s *Stock) Update(ctx context.Context, req *stock.StockRequest, rsp *stock.StockWriteResponse) error {
	err := cs.Update(req.Sku, req.Amount)
	if err != nil {
		rsp.Success = false
		return errors.InternalServerError("charlieplc.Stock.Update", err.Error())
	}
	rsp.Success = true
	return nil
}

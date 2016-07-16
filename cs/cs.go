package cs

import (
	"log"

	"github.com/chazsmi/stock-service/config"
	proto "github.com/chazsmi/stock-service/proto"
	"github.com/gocql/gocql"
)

var (
	Session *gocql.Session
)

// Opens a Cassandra Session
func Init() {
	c, err := config.ReadReturn(config.File)
	if err != nil {
		log.Fatal(err)
	}
	cluster := gocql.NewCluster(c.CsCluster.Host)
	cluster.Keyspace = "stock"
	cluster.Port = c.CsCluster.Port
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	Session = session
}

func Get(sku string) (*proto.StockResponse, error) {
	var amount int32
	if err := Session.Query(`SELECT amount FROM items WHERE sku = ?`,
		sku).Consistency(gocql.One).Scan(&amount); err != nil {
		return &proto.StockResponse{}, err
	}
	return &proto.StockResponse{
		Sku:    sku,
		Amount: amount,
	}, nil
}

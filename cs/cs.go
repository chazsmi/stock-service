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
func Init(c *config.Config) {
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

func Get(sku string) (*proto.StockReadResponse, error) {
	var amount int32
	if err := Session.Query(`SELECT amount FROM items WHERE sku = ?`,
		sku).Consistency(gocql.One).Scan(&amount); err != nil {
		return &proto.StockReadResponse{}, err
	}
	return &proto.StockReadResponse{
		Sku:    sku,
		Amount: amount,
	}, nil
}

func Update(sku string, amount int32) error {
	if err := Session.Query(`UPDATE items set amount = ? WHERE sku = ?`,
		amount, sku).Consistency(gocql.One).Exec(); err != nil {
		return err
	}
	return nil
}

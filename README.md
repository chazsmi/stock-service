# Example Stock Service        
An example RPC stock service using Go Micro

This service is an example service in the following scenero
<p align="center">
  <img src="stock.png" />
</p>
### Dependencies
- Go Micro
- Consul                       
- Cassandra                    
- RabbitMQ
### To run the service
- Make sure you have Consul installed and running
- Make yourself a copy of the config yaml file specifying the location of your Cassandra and Rabbit instances
``` bash
 cp config.example.yml config.yml
```
- Run the service              
``` Go
 go run main.go --config_file_path=path/to/your/cofig/file
```
### Query the service from the command line
``` Go
 go get github.com/micro/micro
```
``` bash
 micro query charlieplc.stock Stock.Update '{"sku":"123456","amount":1}' 
```
### To do
- DB migrations

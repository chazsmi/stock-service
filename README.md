# Example Stock Service        
An example RPC stock service using Go Micro

### Dependencies
- Go Micro
- Consul                       
- Cassandra                    

### To run the service
- Make sure you have Consul installed and running
- Make yourself a copy of the config yaml file specifying the location of your Cassandra instance
``` bash
 cp config.example.yml config.yml
```
- Run the service              
``` Go
 go run main.go --config_file_path=path/to/your/cofig/file
 ```

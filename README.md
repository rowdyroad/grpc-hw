## REST2gRPC service 

This service has implementation of:
- getting list of measurements with filtration by time period and value range
- getting single measurement by time
- getting daily statistics of measurements (count of measurements, average/min/max of values)

# Start

To start all services `docker-compose up -d`

Open browser http://localhost:8099

Rest swagger documentation at http://localhost:9912/docs/index.html

# Architecture
server - gRPC service. Can be configured to csv/memory storage or postgresql storage.
client - REST service. Bridge from http client to gRPC server.
web - nginx with static frontend and proxy passing to client by api prefix.

# Description
    cmd                     // services directory 
        server              // grpc service
            main.go
        client
            main.go         // rest service
    configs
        client.yaml         // config for rest service
        server.yaml         // config for grpc service
    data
        meterusage.csv      // data source
    deploy                  // docker deploys
        client
            Dockerfile      // docker file for build rest service image
        server
            Dockerfile      // docker file for build grpc service image
        web
            Dockerfile      // docker file for build web frontend and configure nginx
    docs
        client              // directory for swagger documentation generation
    internal
        client
            schema
                schema.go   // REST request/response format
            client.go       // rest service
        grpc                // directory for grpc generation results
        server
            handler
                handler.go 
            server.go
        storage
            csv                 
                csv.go          //data source from CSV file with loading all data into memory
                csv_test.go     // unit tests
            db
                model.go        
                pg.go           //data source from Postgres without loading data into memory
                pg_test.go      // unit tests
    migrations
        0-init.sql              //migration for db storage
    pkg
        client
            client.go           //SDK for grpc service
    proto
        storage.proto           //protobuf file
    web                         // frontend client on react (react-create-app)
        src
            Pages
                Find.js         //page to find measurement value by time
                List.js         //page to load measurements with filtering by time period and value range
                Stats.js        //page to show daily statistics with filtering by time period
            App.css
            App.js
            index.js
            index.css
    .dockerignore           
    .gitignore
    docker-compose.yaml     

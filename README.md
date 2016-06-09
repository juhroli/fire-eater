# fire-eater
A basic web application written in *Go* that consumes HTTP json POSTs and forwards their Body to MongoDB.

## Running with Docker
```
docker build -t fire-eater .
docker run --name mongo -p 27017:27017 -d mongo:3.3
docker run --name fire-eater -p 9081:9081 --link mongo:mongo -d fire-eater
```

## POST example
```
curl -H "Content-Type: application/json" -H "source: collectionName" -d "{\"name\" : \"Gopher\", \"birthday\" : {\"year\" : 2016, \"month\" : 6, \"day\" : 9}}" -X POST http://192.168.99.100:9081/consume
```

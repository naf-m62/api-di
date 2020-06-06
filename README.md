# api-di
api with di container

## run
docker-compose -f docker-compose.yml up --build

## debug
1) docker-compose -f mongo.yml up --build
2) in server.go rewrite  
port = os.Getenv("SERVER_PORT") to port = "8080"




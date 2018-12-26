build:
		go build -o ./gw/gw -i ./gw/*.go
		go build -o ./logic/logic -i ./logic/*.go
		go build -o ./microservice/microservice -i ./microservice/*.go
		docker build -t gw ./gw
		docker build -t logic ./logic
		docker build -t microservice ./microservice		
		rm ./gw/gw	
		rm ./logic/logic
		rm ./microservice/microservice

run:
		docker-compose up

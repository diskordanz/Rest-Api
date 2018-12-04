clean: 
		docker rm db api
build:	
		go build -o api/library -i api/*.go		
		docker build -t db ./db
		docker build -t api ./api
		rm ./api/library
up:
		docker-compose up
down:
		docker-compose down



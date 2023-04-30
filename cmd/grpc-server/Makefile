
# include ../env.mk
_NAME := $(notdir $(CURDIR))
_DATE := $(shell date +%Y-%m-%d-%H-%M-%S)


air:
	air -c .air.toml


fmt:
	./gofmt-w.sh

staticcheck: fmt
	staticcheck ./...
vet: fmt 
	go vet -all ./...  
test: 
	go test -v ./...
	

watch-test: 
	 watch  -x go test -v app/weavedb/*.go
testtest:
	# go test -v  -timeout 30s -run \^TestUpdateDocument\$ weavedb-gprc-cache-proxy/app/weavedb
	# go test -v -watch ./...
	# go test -v  -timeout 30s -run ^TestUpdateDocument$ weavedb-gprc-cache-proxy/app/weavedb

c: # client
	air -c .air.client.toml

build: 
	docker build  --platform linux/amd64  -t $(_NAME) . 

gobuild: 
	go build -o tmp/app src/*.go 
	
gobuild2: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags netgo -a -v -o tmp/app2 src/*.go

push: 
	aws ecr get-login-password --region $(_AWS_MAIN_REGION) |  docker login --username AWS --password-stdin $(_AWS_ACCOUNT_ID).dkr.ecr.$(_AWS_MAIN_REGION).amazonaws.com
	docker tag $(_NAME):latest $(_AWS_ACCOUNT_ID).dkr.ecr.$(_AWS_MAIN_REGION).amazonaws.com/$(_NAME):latest
	docker tag $(_NAME):latest $(_AWS_ACCOUNT_ID).dkr.ecr.$(_AWS_MAIN_REGION).amazonaws.com/$(_NAME):$(_DATE)
	docker push $(_AWS_ACCOUNT_ID).dkr.ecr.$(_AWS_MAIN_REGION).amazonaws.com/$(_NAME):latest
	docker push $(_AWS_ACCOUNT_ID).dkr.ecr.$(_AWS_MAIN_REGION).amazonaws.com/$(_NAME):$(_DATE)
dockerlogin:
	aws ecr get-login-password --region $(_AWS_MAIN_REGION) |  docker login --username AWS --password-stdin $(_AWS_ACCOUNT_ID).dkr.ecr.$(_AWS_MAIN_REGION).amazonaws.com

run: 
	docker run -p 3003:3003 -t $(_NAME)

runsh: 
	docker run -p 3003:3003 -it $(_NAME) /bin/sh


create: 
	aws ecr create-repository --region $(_AWS_MAIN_REGION) \
	--repository-name $(_NAME)  --registry-id $(_AWS_ACCOUNT_ID) --region  $(_AWS_MAIN_REGION) \

url: 
	echo "https://$(_AWS_MAIN_REGION).console.aws.amazon.com/ecr/repositories/private/$(_AWS_ACCOUNT_ID)/$(_NAME)?region=$(_AWS_MAIN_REGION)"


_NAME := indexvm

buildsh:
	./scripts/build.sh
run: 
	./scripts/run.sh

dockerbuild: 
	docker build  --platform linux/amd64  -t $(_NAME) . 



dockerrun: 
	docker run -it $(_NAME)  /bin/bash 





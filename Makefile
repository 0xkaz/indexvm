_NAME := indexvm

buildsh:
	./scripts/build.sh
run: 
	./scripts/run.sh
kill-run: kill
	./scripts/run.sh

watch:
	./scripts/watch.sh

# dockerbuild: 
# 	docker build  --platform linux/amd64  -t $(_NAME) . 



# dockerrun: 
# 	docker run -it $(_NAME)  /bin/bash 



kill:
	ps auxww | grep avalanche | grep -v grep | awk '{print $$2}'|xargs  kill -9 




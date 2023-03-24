FROM golang:1.20-alpine

WORKDIR /indexvm
COPY . .

RUN apk update
RUN apk upgrade
RUN apk add vim
RUN apk add bash


# RUN /indexvm/scripts/build.sh 
# RUN cp /indexvm/build/* /usr/local/bin/


EXPOSE 12352
EXPOSE 12353

CMD ["/bin/bash", "/indexvm/scripts/run.sh"]

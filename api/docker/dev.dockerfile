FROM golang:1.14 

WORKDIR /home/dsdit

COPY go.mod go.sum ./
COPY reflex.conf .

RUN go get github.com/cespare/reflex

ENTRYPOINT ["reflex", "-c", "reflex.conf"]

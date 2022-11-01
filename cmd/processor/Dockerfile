# Build from project root using this command docker build -f cmd/processor/Dockerfile .
FROM golang:1.19

WORKDIR /processor-app

COPY ../go.mod .
COPY ../go.sum .

RUN go mod download

COPY ../lib/processor ./lib/processor
COPY ../cmd/processor ./cmd/processor
COPY ../shared ./shared

RUN ls
RUN ls ./cmd/
RUN ls ./cmd/processor

RUN go build -o processor ./cmd/processor/.
CMD [ "sh", "processor" ]
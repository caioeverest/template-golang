# build stage
FROM golang:1.16.6-alpine as builder
ADD . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags "-X main.gitVersion=$(git rev-parse HEAD ) -w -extldflags "-static""  -o {{.RepositoryName}} main.go

# run stage
FROM scratch
COPY --from=builder /src/application .
CMD ["{{.RepositoryName}}"]

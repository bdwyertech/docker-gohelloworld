FROM golang:1.9.2

WORKDIR /go/src/github.com/bdwyertech/app/
# -- Copy in the Application Code
COPY . .
# -- Build the Application
# RUN CGO_ENABLED=0 \
#     GOOS=linux \
#     go get -d -v ./... \
#     && CGO_ENABLED=0 go build . \
#     && pwd && ls -lah

RUN set -ex \
    && go get -d -v \
    && CGO_ENABLED=0 GOOS=linux go build -v .


# -- Copy over the Compiled Application into an Empty Container
FROM scratch
COPY --from=0 /go/src/github.com/bdwyertech/app/app .
ENTRYPOINT ["/app"]

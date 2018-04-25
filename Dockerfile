# step 1. compile
FROM "sky0621dhub/dockerfile-gowithdep" AS builder

COPY . /go/src/github.com/sky0621/go-stepcounter
WORKDIR /go/src/github.com/sky0621/go-stepcounter
RUN dep ensure
RUN go test ./...
RUN CGO_ENABLED=0 go build -o gostepcounter github.com/sky0621/go-stepcounter

# -----------------------------------------------------------------------------
# step 2. build
FROM scratch
COPY --from=builder /go/src/github.com/sky0621/go-stepcounter/ .
ENTRYPOINT [ "./gostepcounter" ]

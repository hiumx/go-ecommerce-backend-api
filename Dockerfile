FROM --platform=linux/arm64 golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download 

RUN go build -o crm.sd.com ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/crm.sd.com /

ENTRYPOINT [ "/crm.sd.com", "config/local.yaml" ]
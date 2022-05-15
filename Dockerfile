FROM golang:1.17 as builder

WORKDIR /go/src/github.com/cnblvr/testing_forms

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /opt/testing_forms/bin/api ./cmd/api/main.go


FROM alpine:3.15
ENV PATH="/opt/testing_forms/bin:${PATH}"

COPY --from=builder /opt/testing_forms/bin /opt/testing_forms/bin

CMD ["api"]

FROM golang:1.20.0-alpine3.16 AS builder
RUN apk add --no-cache git gcc musl-dev
WORKDIR /go/src/llm_chat_base_service
COPY . .
RUN go mod download
RUN go build -o /go/bin/llm_chat_base_service

FROM alpine:3.16.0
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin/llm_chat_base_service /llm_chat_base_service
RUN mkdir /operation_data
ENTRYPOINT /llm_chat_base_service
LABEL Name=llm_chat_base_service Version=1.0
EXPOSE 7523
CMD ["llm_chat_base_service"]
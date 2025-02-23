# 1. 빌드 단계
FROM golang:alpine AS builder

WORKDIR /build

COPY ./ ./ 

RUN go mod download
RUN go build -o surge .  

WORKDIR /dist
RUN cp /build/surge .

# 2. 실행 단계
FROM alpine AS runtime
WORKDIR /app

COPY --from=builder /dist/surge .
COPY schema ./schema

EXPOSE 7578 

CMD ["./surge"]

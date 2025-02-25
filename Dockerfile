# 1️⃣ 빌드 스테이지
FROM golang:alpine AS builder

# 작업 디렉토리 설정
WORKDIR /build

# 프로젝트 파일 복사
COPY . .

# 의존성 다운로드
RUN go mod download

# Go 애플리케이션 빌드
RUN go build -o surge .

# 배포용 디렉토리 준비
WORKDIR /dist
RUN cp /build/surge . 
RUN mkdir -p /dist/migrations && cp -r /build/migrations /dist/migrations

# 2️⃣ 실행 스테이지
FROM alpine AS runtime

# 작업 디렉토리 설정
WORKDIR /app

# 실행 파일 및 스키마 복사
COPY --from=builder /dist/surge .
COPY --from=builder /dist/migrations ./migrations

# 실행 파일 권한 설정
RUN chmod +x ./surge

# 컨테이너 실행 시 surge 실행
CMD ["./surge"]

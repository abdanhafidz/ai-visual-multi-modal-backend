# Gunakan image dasar Golang versi 1.24.1
FROM golang:1.24.1

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy seluruh kode
COPY . .

# Buat file .env dengan variabel environment menggunakan Hugging Face secrets
RUN --mount=type=secret,id=DB_PASSWORD,mode=0444,required=false \
    --mount=type=secret,id=OPENAI_API_KEY,mode=0444,required=false \
    --mount=type=secret,id=REPLICATE_API_TOKEN,mode=0444,required=false \
    echo "DB_HOST=aws-0-ap-southeast-1.pooler.supabase.com" >> .env && \
    echo "DB_USER=postgres.iuwuiuoisqnfdzlgwurl" >> .env && \
    echo "DB_PASSWORD=$(cat /run/secrets/DB_PASSWORD 2>/dev/null)" >> .env && \
    echo "DB_PORT=5432" >> .env && \
    echo "DB_NAME=kpppl" >> .env && \
    echo "SALT=NZNZtY7dNPz8l0dWINJZLKafWaJrql1s" >> .env && \
    echo "HOST_ADDRESS=0.0.0.0" >> .env && \
    echo "HOST_PORT=7860" >> .env && \
    echo "LOG_PATH=logs" >> .env && \
    echo "EMAIL_VERIFICATION_DURATION=2" >> .env && \
    echo "OPEN_AI_API_KEY=$(cat /run/secrets/OPENAI_API_KEY 2>/dev/null)" >> .env && \
    echo "REPLICATE_API_TOKEN=$(cat /run/secrets/REPLICATE_API_TOKEN 2>/dev/null)" >> .env
# Build aplikasi
RUN go build -o main .

# Expose port untuk Hugging Face Spaces
EXPOSE 7860

# Jalankan aplikasi
CMD ["./main"]

# Temel imaj olarak Go'yu kullanıyoruz.
FROM golang:1.22 as builder

# Çalışma dizinini oluştur ve projenizi buraya kopyalayın
WORKDIR /app

# Go mod ve proje dosyalarını kopyalayın
COPY go.mod go.sum ./
RUN go mod download

# Tüm kaynak dosyaları kopyalayın
COPY . .

# Uygulamayı oluşturun
RUN go build -o kube-log-collector ./main.go

# Küçük bir çalışma imajı oluşturuyoruz
FROM alpine:latest

# İhtiyaç duyulan bağımlılıkları yükleyin
RUN apk --no-cache add ca-certificates

# Çalışma dizinini oluşturun
WORKDIR /root/

# Derlenmiş binary dosyayı kopyalayın
COPY --from=builder /app/kube-log-collector .

# Uygulamayı çalıştır
CMD ["./kube-log-collector"]

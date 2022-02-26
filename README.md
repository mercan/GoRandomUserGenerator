
# Proje Başlığı

Go ile yazılmış rastgele kullanıcı verileri oluşturmak için ücretsiz API

Swagger [Dökümantasyon](https://gorandomusergenerator.herokuapp.com/swagger/)
## Bilgisayarınızda Çalıştırın

Projeyi klonlayın

```bash
  git clone https://github.com/mercan/GoRandomUserGenerator.git
```

Proje dizinine gidin

```bash
  cd GoRandomUserGenerator
```

Sunucuyu çalıştırın

```bash
  go run .
```

  
## API Kullanımı

#### Detaylı kullanıcı getir

```https
  GET /api/v1/user
```

#### Sadece adres getir

```https
  GET api/v1/address
```

#### Sadece kredi kartı getir

```https
  GET api/v1/credit_card
```
  
## Kullanılan Teknolojiler

>**Sunucu:** Go, Fiber, Swagger

  
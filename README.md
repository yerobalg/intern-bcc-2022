# intern-bcc-2022 (Kadoin Aja)
Kelompok 16
Link Dokumentasi API: https://documenter.getpostman.com/view/13667981/UVkvHCGH

List dependensi:
```
   Gin           : go get -u github.com/gin-gonic/gin
   Gin CORS      : go get github.com/gin-contrib/cors
   GORM          : go get -u gorm.io/gorm
   GORM Postgres : go get -u gorm.io/driver/postgres
   CryptoBcrypt  : go get golang.org/x/crypto/bcrypt
   Go JWT        : go get -u github.com/golang-jwt/jwt/v4
   Raja Ongkir   : go get github.com/Bhinneka/go-rajaongkir   
```

# Tutorial Instalasi
1. Clone repo github ini dalam folder yang bernama clean-arch-2
2. Lakukan instalasi pada dependensi diatas   
3. langkapi file env.server dan env.database. Jika sudah, rename dan tambahkan tanda titik di depan, menjadi .env.server dan .env.database
4. Buat database postgres dan jalankan seeder1.sql
5. Setelah sudah, run server untuk melakukan migration.
6. Matikan server, kemudian jalankan seeder2.sql
7. Server siap digunakan

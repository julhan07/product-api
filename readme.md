# Cara menjalankan api
Beberapa langkah yang harus dilakukan untuk menjalankan api ini :

1. clone repository ini terlebih dahulu

2. jalankan docker-compose yang ada dalam folder api ini jika belum menginstall posgresql dan redis

3. ketikkan perintah go mod tidy untuk mengupdate pkg

4. ketikkan perintah go run .

5. Selesai....😁

# APis
Beberapa api yang sudah dibuat pada project ini :

1. api view source-product http://localhost:3000/source-product (GET) : api ini dapat melakukan proses view list source-product dan proses sync data ke tabel destination-product.
2. api view destination-product : http://localhost:3000/destination-product (GET)


# Noted

Saat api dijalankan secara otomatis akan melakukan auto migrate tabel jika database sudah didefinisikan sebelumnya. Proses seeder data juga otomatis dijalanlan jika pengatuan DB_SEED=1 pada file .env


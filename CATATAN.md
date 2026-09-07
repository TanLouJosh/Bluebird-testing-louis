
## Step 1
## Apa yang saya ubah
Sayta buat struct untuk menyimpan data mobil, untuk harga saya menghitung dengan function
## Kenapa begitu
struct untuk simpan semua data mobil dalam satu halaman, harga saya buat function karena ini variable yang banyak modifikasy bukan hanya variable simple
## Apa yang saya pertimbangkan tapi tidak saya pilih, dan kenapa
-
## Yang saya rasa masih kurang
-
## Bagian yang dibantu AI (kalau ada)
Tidak untuk ini

## Step 2
## Apa yang saya ubah
1) Buat struct baru buat trade in transaction
2) Error saya jadikan list sendiri
3) Transaction type saya jadikan enum
## Kenapa begitu
1) Buat struct untuk trade in vair organised
2) Untuk standarkan error sistem, dan isa trigger logic error
3) Transaction type/medotde bayar Biar bisa di refer dan bedakan di struct lain
## Apa yang saya pertimbangkan tapi tidak saya pilih, dan kenapa
-
## Yang saya rasa masih kurang
-
## Bagian yang dibantu AI (kalau ada)
Tidak untuk ini

## Step 3
## Apa yang saya ubah
1) Buat map buaru untuk menyimpan rate
2) Buat function baru untuk menghitung harga secara dynamic
## Kenapa begitu
1) Buat map yang gampang berkembang untuk tiap metode pembayaran baru
2) Buat function yang bisa digunakan ulang dengan map baru
## Apa yang saya pertimbangkan tapi tidak saya pilih, dan kenapa
-
## Yang saya rasa masih kurang
-
## Bagian yang dibantu AI (kalau ada)
Tidak untuk ini

## Apa yang menjadi lebih sulit setelah perubahan ini?
Tiap penambahan metode baru butuh masukkan transaction type baru, masih agak manual

## Step 4
## Apa yang saya ubah
1) Tambahkan fake interface, repo, dan service
## Kenapa begitu
1) Buat interface agar service menjadi lebih luas dan bis adaptasi db pisah"
## Apa yang saya pertimbangkan tapi tidak saya pilih, dan kenapa
-
## Yang saya rasa masih kurang
-
## Bagian yang dibantu AI (kalau ada)
Tidak untuk ini

## Kalau setiap transaksi menanyakan tarif ke database, apa masalahnya? Apa yang akan Anda lakukan?
Bisa ada race condition, harus ada lock/control akses dbnya
banyak transaksi juga bisa sebabkan bottleneck
## Bagaimana Anda memastikan hasil perhitungan tidak berubah setelah perubahan ini?
# latihan web service go

## release 1
- lakukan decode file politician.json dan voting.json ke dalam tipe data struct
- Kemudian, buatlah sebuah database yang akan digunakan untuk memindahkan data file json ke dalam tabel :
1. tabel politicians (politian_id, name, party, location, grade_current)
2. tabel votings (voter_id, politician_id, first_name, last_name, gender, age)
- Relasi antara kedua tabel adalah : voting hanya bisa memiliki satu politician, tapi politician bisa memiliki banyak voting
- Pindahkan data di file decode politicians.json ke tabel politicians
- Pindahkan data di file decode voting.json ke tabel votings

## release 2
-Buatlah sebuah query dari database yang sudah dibuat untuk melakukan :
1. tampilkan semua data politician beserta voting
2. tampilkan semua data voting yang ber gender male dan nama diawali huruf B
4. Tampilkan 1 politician dengan voting teratas di lokasi NY
3. tampilkan 3 politician teratas dengan voting terbanyak


## release 3
- Buatlah sebuah routing dengan output berupa API JSON , penjelasan sebagai berikut :
1. GET "/votings" , untuk menampilkan semua data voting 
2. GET "/votings_male" , untuk menampilkan semua data voting gender male
3. GET "/votings_female" , untuk menampilkan semua data voting gender female
4. GET "/politicians" , untuk menampilkan semua data politicians 
5. GET "/politicians_voting" , untuk menampilkan semua data politician beserta voting. 
6. GET "/politicians_il" , untuk menampilkan semua politician di location IL beserta votingnya
7. GET "/politicians_ny" , untuk menampilkan semua politician di location NY beserta votingnya
8. GET "/politicians_tx" , untuk menampilkan semua politician di location TX beserta votingnya
9. GET "/politicians_la" , untuk menampilkan semua politician di location LA beserta votingnya 
10. GET "/politicians_wa" , untuk menampilkan semua politician di location WA beserta votingnya 
11. GET "/politicians_fl" , untuk menampilkan semua politician di location FL beserta votingnya
12. GET "/politicians_hi" , untuk menampilkan semua politician di location HI beserta votingnya 

## release 4
- buatlah sebuat render html sederhana dengan menampilkan sebuah tabel dari hasil tiap routingan

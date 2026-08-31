# Shipment Tracking API

Proyek ini mensimulasikan sistem pelacakan pengiriman yang biasa 
dipakai platform logistik, seperti Kargo, Waresix, dll. Alurnya:

1. Order masuk
2. Order di-assign ke carrier (ekspedisi pengangkut)
3. Proses pengiriman dicatat sebagai shipment
4. Setiap kali status shipment berubah, dicatat di `shipment_events` 
   (riwayat, append-only — tidak pernah diupdate/dihapus)
5. Ada opsi webhook untuk notifikasi otomatis ke sistem eksternal 
   setiap status berubah
6. `api_keys` untuk autentikasi client B2B yang memanggil API ini

## Fitur
- CRUD dasar untuk data shipment (buat, lihat detail, lihat semua)
- Riwayat perubahan status yang tersimpan lengkap (audit trail)
- Validasi transisi status otomatis (state machine) — mencegah 
  perubahan status yang tidak valid
- Notifikasi webhook otomatis ke sistem eksternal saat status 
  shipment berubah

## Tech Stack
- Go
- Gin (web framework)
- PostgreSQL
- pgx (driver database)

## Cara Menjalankan

### Prasyarat
- Go sudah terinstall
- PostgreSQL sudah terinstall dan berjalan

### Langkah-langkah

1. Clone repository ini

2. Install dependency
```
   go mod tidy
```

3. Setup database
   - Buat database baru bernama `shipment_tracking` di PostgreSQL
   - Aktifkan extension pgcrypto:
```sql
     CREATE EXTENSION IF NOT EXISTS "pgcrypto";
```
   - Buat 6 tabel (`shipments`, `orders`, `carriers`, `shipment_events`, 
     `webhook_subscriptions`, `api_keys`) sesuai skema di bagian
      <img width="1855" height="926" alt="image" src="https://github.com/user-attachments/assets/053a1b74-8275-4bff-99a4-5eaf863c8ac4" />


4. Sesuaikan connection string database di kode (saat ini masih 
   hardcode) sesuai kredensial PostgreSQL di komputer kamu

5. Jalankan aplikasi `go run main.go`

6. Server berjalan di `http://localhost:8080`

## API Endpoints
- POST /shipments
- GET /shipments
- GET /shipments/{id}
- GET /shipments/{id}/history
- PATCH /shipments/{id}/status
- POST /webhooks

## Struktur Projek
<img width="333" height="617" alt="image" src="https://github.com/user-attachments/assets/387443bb-632e-40c0-896e-42df7a7ed3fa" />

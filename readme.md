# Self payroll System

Fitur utama dari web service ini adalah employee dapat melakukan penarikan gaji secara mandiri setiap bulannya.

Lebih detail nya, fitur yang harus dikerjakan adalah:

1. Melakukan manage data jabatan berupa operasi CRUD (create, read, update, delete)
2. Melakukan manage data employee berupa operasi CRUD (Create, Read, Update, Delete)
3. Admin dapat melakukan topup balance perusahaan
4. Melakukan penarikan sallary dengan menyertakan employee ID dan secret ID, besaran salary berdasarkan jabatan yang dimiliki oleh tiap employee
5. Terdapat riwayat topup dan pengurangan balance perusahaan

Untuk menjalankan aplikasi lakukan perintah berikut

## Getting Started

To run the application, follow the instructions below:

1. Clone this repository and navigate to the project directory.

   ```bash
   git clone https://github.com/szczynk/self-payroll.git
   ```

1. Create a `.env` file by running `cp .env.example .env`. Then fill in the `.env` file according to your local environment.

   ```bash
   cp .env.example .env
   ```

1. Run `go mod tidy && go mod vendor`.

   ```bash
   go mod tidy && go mod vendor
   ```

1. Run `go run *.go`.

   ```bash
   go run *.go
   ```

The list of endpoints is available in the [documenter](https://documenter.getpostman.com/view/4080490/2s83Ychhk4).

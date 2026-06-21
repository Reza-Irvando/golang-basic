package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int          // Variabel yang akan diperebutkan
	var mu sync.Mutex        // Membuat Mutex
	var wg sync.WaitGroup    // (Opsional) Hanya untuk menunggu semua goroutine selesai

	// Kita jalankan 1000 pekerja (Goroutine) secara bersamaan!
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		
		go func() {
			defer wg.Done()

			// 🔒 MENGUNCI PINTU
			// Goroutine lain yang sampai di baris ini harus antri
			mu.Lock()
			
			// PROSES MENGUBAH DATA (Aman dari tabrakan)
			counter++ 
			
			// 🔓 MEMBUKA PINTU
			// Mempersilakan antrian Goroutine selanjutnya untuk masuk
			mu.Unlock() 
		}()
	}

	// Menunggu ke-1000 goroutine selesai bekerja
	wg.Wait() 

	// Hasilnya dijamin pasti 1000. 
	// Jika Anda menghapus mu.Lock() dan mu.Unlock(), hasilnya akan acak (misal: 954, 982) karena tabrakan!
	fmt.Println("Hasil akhir counter:", counter)
}
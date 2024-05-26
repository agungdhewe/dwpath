# DwPath
Suatu modul untuk mempermudah pengecekan apakah file/direktori ada atau tidak.

### Mengecek keberadaan File
	exists, _, _ := dwpath.IsFileExists("/home/agungdhewe/filesaya.txt")
	if !exists {
		fmt.Println("file tidak ditemukan")
	}
    
### Mengecek keberadaan direktori
	exists, _ := dwpath.IsDirectoryExists("/Volumes/Development/agungdhewe")
	if !exists {
		fmt.Println("direktori tidak ditemukan")
	}
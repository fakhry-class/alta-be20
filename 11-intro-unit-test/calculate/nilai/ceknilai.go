package nilai

func CheckNilaiSiswa(nilai int) string {
	if nilai >= 85 && nilai <= 100 {
		return "A"
	} else if nilai >= 70 && nilai < 85 {
		return "B+"
	} else if nilai >= 50 && nilai < 70 {
		return "B"
	} else {
		return "E"
	}
}

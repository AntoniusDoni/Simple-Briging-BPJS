package utils

type ResponseSearchRujukan struct {
	MetaData struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"metaData"`
	Response ResponseRujukan `json:"response,omitempty"`
}
type ResponseRujukan struct {
	Rujukan Rujukan `json:"rujukan,omitempty"`
}
type Rujukan struct {
	Diagnosa struct {
		Kode string `json:"kode,omitempty"`
		Nama string `json:"nama,omitempty"`
	} `json:"diagnosa,omitempty"`
	Keluhan     string `json:"keluhan,omitempty"`
	NoKunjungan string `json:"noKunjungan,omitempty"`
	Pelayanan   struct {
		Kode string `json:"kode,omitempty"`
		Nama string `json:"nama,omitempty"`
	} `json:"pelayanan,omitempty"`
	Peserta struct {
		Cob struct {
			NmAsuransi interface{} `json:"nmAsuransi,omitempty"`
			NoAsuransi interface{} `json:"noAsuransi,omitempty"`
			TglTAT     interface{} `json:"tglTAT,omitempty"`
			TglTMT     interface{} `json:"tglTMT,omitempty"`
		} `json:"cob,omitempty"`
		HakKelas struct {
			Keterangan string `json:"keterangan,omitempty"`
			Kode       string `json:"kode,omitempty"`
		} `json:"hakKelas,omitempty"`
		Informasi struct {
			Dinsos      interface{} `json:"dinsos,omitempty"`
			NoSKTM      interface{} `json:"noSKTM,omitempty"`
			ProlanisPRB interface{} `json:"prolanisPRB,omitempty"`
		} `json:"informasi,omitempty"`
		JenisPeserta struct {
			Keterangan string `json:"keterangan,omitempty"`
			Kode       string `json:"kode,omitempty"`
		} `json:"jenisPeserta,omitempty"`
		Mr struct {
			NoMR      string      `json:"noMR,omitempty"`
			NoTelepon interface{} `json:"noTelepon,omitempty"`
		} `json:"mr,omitempty"`
		Nama     string `json:"nama,omitempty"`
		Nik      string `json:"nik,omitempty"`
		NoKartu  string `json:"noKartu,omitempty"`
		Pisa     string `json:"pisa,omitempty"`
		ProvUmum struct {
			KdProvider string `json:"kdProvider,omitempty"`
			NmProvider string `json:"nmProvider,omitempty"`
		} `json:"provUmum,omitempty"`
		Sex           string `json:"sex,omitempty"`
		StatusPeserta struct {
			Keterangan string `json:"keterangan,omitempty"`
			Kode       string `json:"kode,omitempty"`
		} `json:"statusPeserta,omitempty"`
		TglCetakKartu string `json:"tglCetakKartu,omitempty"`
		TglLahir      string `json:"tglLahir,omitempty"`
		TglTAT        string `json:"tglTAT,omitempty"`
		TglTMT        string `json:"tglTMT,omitempty"`
		Umur          struct {
			UmurSaatPelayanan string `json:"umurSaatPelayanan,omitempty"`
			UmurSekarang      string `json:"umurSekarang,omitempty"`
		} `json:"umur,omitempty"`
	} `json:"peserta,omitempty"`
	PoliRujukan struct {
		Kode string `json:"kode,omitempty"`
		Nama string `json:"nama,omitempty"`
	} `json:"poliRujukan,omitempty"`
	ProvPerujuk struct {
		Kode string `json:"kode,omitempty"`
		Nama string `json:"nama,omitempty"`
	} `json:"provPerujuk,omitempty"`
	TglKunjungan string `json:"tglKunjungan,omitempty"`
}

type ResponsePeserta struct {
	Code    string
	Message string
	Peserta struct {
		NoKartu string `json:"noKartu"`
		Nik     string `json:"nik"`
		Nama    string `json:"nama"`
		Pisa    string `json:"pisa"`
		Sex     string `json:"sex"`
		Mr      struct {
			NoMR      string `json:"noMR"`
			NoTelepon string `json:"noTelepon"`
		} `json:"mr"`
		TglLahir      string `json:"tglLahir"`
		TglCetakKartu string `json:"tglCetakKartu"`
		TglTAT        string `json:"tglTAT"`
		TglTMT        string `json:"tglTMT"`
		StatusPeserta struct {
			Kode       string `json:"kode"`
			Keterangan string `json:"keterangan"`
		} `json:"statusPeserta"`
		ProvUmum struct {
			KdProvider string `json:"kdProvider"`
			NmProvider string `json:"nmProvider"`
		} `json:"provUmum"`
		JenisPeserta struct {
			Kode       string `json:"kode"`
			Keterangan string `json:"keterangan"`
		} `json:"jenisPeserta"`
		HakKelas struct {
			Kode       string `json:"kode"`
			Keterangan string `json:"keterangan"`
		} `json:"hakKelas"`
		Umur struct {
			UmurSekarang      string `json:"umurSekarang"`
			UmurSaatPelayanan string `json:"umurSaatPelayanan"`
		} `json:"umur"`
		Informasi struct {
			Dinsos      interface{} `json:"dinsos"`
			ProlanisPRB string      `json:"prolanisPRB"`
			NoSKTM      interface{} `json:"noSKTM"`
		} `json:"informasi"`
		Cob struct {
			NoAsuransi interface{} `json:"noAsuransi"`
			NmAsuransi interface{} `json:"nmAsuransi"`
			TglTMT     interface{} `json:"tglTMT"`
			TglTAT     interface{} `json:"tglTAT"`
		} `json:"cob"`
	} `json:"peserta"`
}

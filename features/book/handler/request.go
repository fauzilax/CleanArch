package handler

import "CleanArch/features/book"

type AddBookRequest struct {
	Judul       string `json:"judul" form:"judul"`
	TahunTerbit int    `json:"tahun_terbit" form:"tahun_terbit"`
	Penulis     string `json:"penulis" form:"penulis"`
}
type UpdateBookRequest struct {
	Judul       string `json:"judul" form:"judul"`
	TahunTerbit int    `json:"tahun_terbit" form:"tahun_terbit"`
	Penulis     string `json:"penulis" form:"penulis"`
}

func ConvToCore(data interface{}) *book.Core {
	res := book.Core{}

	switch data.(type) {
	case AddBookRequest:
		cnv := data.(AddBookRequest)
		res.Judul = cnv.Judul
		res.TahunTerbit = cnv.TahunTerbit
		res.Penulis = cnv.Penulis
	case UpdateBookRequest:
		cnv := data.(UpdateBookRequest)
		res.Judul = cnv.Judul
		res.TahunTerbit = cnv.TahunTerbit
		res.Penulis = cnv.Penulis
	default:
		return nil
	}

	return &res
}

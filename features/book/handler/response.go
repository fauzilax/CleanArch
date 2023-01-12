package handler

import "CleanArch/features/book"

type BookReponse struct {
	ID          uint   `json:"id"`
	Judul       string `json:"title"`
	TahunTerbit int    `json:"published_year"`
	Penulis     string `json:"written by"`
	Pemilik     string `json:"owner"`
}
type BookList struct {
	Judul   string `json:"title"`
	Penulis string `json:"written by"`
	Pemilik string `json:"owner name"`
}
type UpdateBook struct {
	Judul       string `json:"title"`
	TahunTerbit int    `json:"published_year"`
	Penulis     string `json:"written by"`
}

type AddBookReponse struct {
	ID          uint   `json:"id"`
	Judul       string `json:"title"`
	TahunTerbit int    `json:"published_year"`
	Penulis     string `json:"written by"`
}

func AddResponse(data book.Core) AddBookReponse {
	return AddBookReponse{
		ID:          data.ID,
		Judul:       data.Judul,
		TahunTerbit: data.TahunTerbit,
		Penulis:     data.Penulis,
	}
}

func BookListResponse(data book.Core) BookList {
	return BookList{
		Judul:   data.Judul,
		Penulis: data.Penulis,
		Pemilik: data.Pemilik,
	}
}
func MyBookResponse(data book.Core) AddBookReponse {
	return AddBookReponse{
		ID:          data.ID,
		Judul:       data.Judul,
		TahunTerbit: data.TahunTerbit,
		Penulis:     data.Penulis,
	}
}
func UpdateBookResponse(data book.Core) UpdateBook {
	return UpdateBook{
		Judul:       data.Judul,
		TahunTerbit: data.TahunTerbit,
		Penulis:     data.Penulis,
	}
}

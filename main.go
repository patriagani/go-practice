package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

type Categories struct {
	ID          int    `json:"id"`
	Nama        string `json:"nama"`
	Description string `json:"description"`
}

var produk = []Product{
	{ID: 1, Nama: "Indomie Goreng", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Indomie Rebus", Harga: 3500, Stok: 9},
	{ID: 3, Nama: "Nasi Goreng", Harga: 5000, Stok: 8},
}

var kategori = []Categories{
	{ID: 1, Nama: "Makanan", Description: "Semua jenis makanan"},
	{ID: 2, Nama: "Minuman", Description: "Semua jenis minuman"},
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var produkBaru Product
	err := json.NewDecoder(r.Body).Decode(&produkBaru)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	produkBaru.ID = len(produk) + 1
	produk = append(produk, produkBaru)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode((produkBaru))
}

// GET /api/produk/{id}
func getProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Product ID Not Found", http.StatusNotFound)
}

// PUT /api/produk/{id}
func editProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	var produkUpdate Product
	err = json.NewDecoder(r.Body).Decode(&produkUpdate)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	for i := range produk {
		if produk[i].ID == id {
			produkUpdate.ID = id
			produk[i] = produkUpdate
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produkUpdate)
			return
		}
	}
	http.Error(w, "Product ID Not Found", http.StatusNotFound)
}

// DELETE /api/produk/{id}
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	for i, p := range produk {
		if p.ID == id {
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "success delete",
			})
			return
		}
	}
	http.Error(w, "Product ID Not Found", http.StatusNotFound)
}

// POST /api/categories
func createCategory(w http.ResponseWriter, r *http.Request) {
	var kategoriBaru Categories
	err := json.NewDecoder(r.Body).Decode(&kategoriBaru)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	kategoriBaru.ID = len(kategori) + 1
	kategori = append(kategori, kategoriBaru)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode((kategoriBaru))
}

// GET /api/categories/{id}
func getCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	for _, k := range kategori {
		if k.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(k)
			return
		}
	}
	http.Error(w, "Category ID Not Found", http.StatusNotFound)
}

// PUT /api/categories/{id}
func editCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}
	var kategoriUpdate Categories
	err = json.NewDecoder(r.Body).Decode(&kategoriUpdate)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	for i := range kategori {
		if kategori[i].ID == id {
			kategoriUpdate.ID = id
			kategori[i] = kategoriUpdate
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(kategoriUpdate)
			return
		}
	}
	http.Error(w, "Category ID Not Found", http.StatusNotFound)
}

// DELETE /api/categories/{id}
func deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}
	for i, k := range kategori {
		if k.ID == id {
			kategori = append(kategori[:i], kategori[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "success delete",
			})
			return
		}
	}
	http.Error(w, "Category ID Not Found", http.StatusNotFound)
}

func main() {
	//GET /api/produk/{id}
	//PUT /api/produk/{id}
	//DELETE /api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProduct(w, r)
		} else if r.Method == "PUT" {
			editProduct(w, r)
		} else if r.Method == "DELETE" {
			deleteProduct(w, r)
		}

	})

	//GET /api/produk
	//POST /api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			createProduct(w, r)

		}
	})

	//GET /api/categories/{id}
	//PUT /api/categories/{id}
	//DELETE /api/categories/{id}
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategory(w, r)
		} else if r.Method == "PUT" {
			editCategory(w, r)
		} else if r.Method == "DELETE" {
			deleteCategory(w, r)
		}

	})

	//GET /api/categories
	//POST /api/categories
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(kategori)
		} else if r.Method == "POST" {
			createCategory(w, r)

		}
	})
	//GET /health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})
	fmt.Println("Server running")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server running failed")
	}
}

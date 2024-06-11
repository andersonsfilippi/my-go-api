package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Cliente struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// DB variável global para a conexão com o banco de dados
var DB *gorm.DB
var err error

// InitDatabase inicializa a conexão com o banco de dados e faz a migração do schema
func InitDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar no banco de dados: %v", err)
	}

	// Migrar o schema do banco de dados
	DB.AutoMigrate(&Cliente{})
}

// CreateCliente cria um novo cliente no banco de dados
func CreateCliente(w http.ResponseWriter, r *http.Request) {
	var cliente Cliente
	json.NewDecoder(r.Body).Decode(&cliente)
	fmt.Println("Cliente: ", cliente)

	DB.Create(&cliente)
	json.NewEncoder(w).Encode(cliente)
	fmt.Println("Cliente: ", cliente)
}

// GetClientes retorna todos os clientes do banco de dados
func GetClientes(w http.ResponseWriter, r *http.Request) {
	var clientes []Cliente
	DB.Find(&clientes)
	json.NewEncoder(w).Encode(clientes)
}

// setupRoutes configura as rotas da API
func setupRoutes() {
	http.HandleFunc("/clientes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			CreateCliente(w, r)
		case "GET":
			GetClientes(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func main() {
	// Inicializa a conexão com o banco de dados
	InitDatabase()

	// Configura as rotas da API
	setupRoutes()

	// Inicia o servidor HTTP na porta 8080
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

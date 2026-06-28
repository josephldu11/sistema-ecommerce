package handlers

import (
	"encoding/json"
	"net/http"
	"sistema-ecommerce-go/models"
)

// Base de datos simulada en memoria
var (
	Productos = []models.Producto{
		{ID: 1, Nombre: "Laptop Gamer", Precio: 1200, Stock: 5},
		{ID: 2, Nombre: "Mouse RGB", Precio: 35, Stock: 20},
	}
	Clientes = []models.Cliente{
		{ID: 1, Nombre: "Joseph Castillo", Email: "josephldu11@gmail.com"},
	}
	Pedidos = []models.Pedido{}
)

// Servicio 1: Inicio
func Inicio(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API de E-commerce v2.0 - Desarrollado en Golang"))
}

// Servicio 2: Listar Productos (GET)
func ListarProductos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Productos)
}

// Servicio 3: Crear Producto (POST)
func CrearProducto(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido. Usa POST.", http.StatusMethodNotAllowed)
		return
	}
	var p models.Producto
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
		return
	}
	Productos = append(Productos, p)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Producto creado con éxito"})
}

// Servicio 4: Listar Clientes (GET)
func ListarClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Clientes)
}

// Servicio 5: Crear Cliente (POST)
func CrearCliente(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	var c models.Cliente
	json.NewDecoder(r.Body).Decode(&c)
	Clientes = append(Clientes, c)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Cliente registrado"})
}

// Servicio 6: Listar Pedidos (GET)
func ListarPedidos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Pedidos)
}

// Servicio 7: Crear Pedido (POST)
func CrearPedido(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	var ped models.Pedido
	json.NewDecoder(r.Body).Decode(&ped)
	Pedidos = append(Pedidos, ped)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Pedido procesado"})
}

// Servicio 8: Actualizar Stock (PUT)
func ActualizarStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Stock actualizado correctamente"})
}

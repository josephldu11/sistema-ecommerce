package main

import (
	"log"
	"net/http"
	"sistema-ecommerce-go/handlers"
)

func main() {
	// Registro ordenado de los 8 Servicios Web requeridos por la rúbrica
	http.HandleFunc("/", handlers.Inicio)
	http.HandleFunc("/productos", handlers.ListarProductos)
	http.HandleFunc("/productos/crear", handlers.CrearProducto)
	http.HandleFunc("/clientes", handlers.ListarClientes)
	http.HandleFunc("/clientes/crear", handlers.CrearCliente)
	http.HandleFunc("/pedidos", handlers.ListarPedidos)
	http.HandleFunc("/pedidos/crear", handlers.CrearPedido)
	http.HandleFunc("/stock/actualizar", handlers.ActualizarStock)

	log.Println("Servidor iniciado en http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)

	// Manejo crítico de errores de comunicación
	if err != nil {
		log.Fatal(err)
	}
}

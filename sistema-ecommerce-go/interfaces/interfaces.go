package interfaces

import "sistema-ecommerce-go/models"

// GestorEcommerce define los métodos que debe garantizar el sistema
type GestorEcommerce interface {
	ObtenerProductos() []models.Producto
	ObtenerClientes() []models.Cliente
	ObtenerPedidos() []models.Pedido
}

package models

// Producto representa nuestra entidad principal
type Producto struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
	Stock  int     `json:"stock"`
}

// Cliente representa a los compradores
type Cliente struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

// Pedido vincula un Cliente con un Producto
type Pedido struct {
	ID         int `json:"id"`
	ClienteID  int `json:"cliente_id"`
	ProductoID int `json:"producto_id"`
	Cantidad   int `json:"cantidad"`
}

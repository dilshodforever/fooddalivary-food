// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/orderitems": {
            "get": {
                "description": "List all items in an order for a given user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OrderItem"
                ],
                "summary": "List Order Items",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of order items",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetAllItems"
                        }
                    },
                    "400": {
                        "description": "Missing or invalid user ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Add items to an existing order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OrderItem"
                ],
                "summary": "Add Items",
                "parameters": [
                    {
                        "description": "Items data",
                        "name": "items",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.AddItemsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Items added successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.AddItemsResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove items from an existing order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OrderItem"
                ],
                "summary": "Delete Items",
                "parameters": [
                    {
                        "description": "Items deletion data",
                        "name": "items",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.DeleItemsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Items deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.DeleteProductResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "description": "List all orders with optional filtering",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "List Orders",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of orders",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetAllOrdersList"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update the details of an existing order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Update Order",
                "parameters": [
                    {
                        "description": "Updated order data",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.UpdateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order updated successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.UpdateStatusResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new order with the given details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Create Order",
                "parameters": [
                    {
                        "description": "Order data",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order created successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.CreateOrderResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orders/status": {
            "patch": {
                "description": "Update the status of an existing order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Update Order Status",
                "parameters": [
                    {
                        "description": "Status update data",
                        "name": "status",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.UpdateStatusRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order status updated successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.UpdateStatusResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "description": "Retrieve an order's details by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get Order By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order details",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetAllOrders"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an order by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Delete Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.DeleteOrdersByidResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "List all available products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "List Products",
                "responses": {
                    "200": {
                        "description": "List of products",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetAllProductResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update the details of an existing product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Update Product",
                "parameters": [
                    {
                        "description": "Updated product data",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.UpdateProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product updated successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.ProductResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new product with the given details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create Product",
                "parameters": [
                    {
                        "description": "Product data",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.CreateProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product created successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.ProductResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Retrieve a product's details by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get Product By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product details",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetProductResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a product by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete Product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.ProductResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "genprotos.AddItemsRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "genprotos.AddItemsResponse": {
            "type": "object",
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.GetProductResponse"
                    }
                },
                "quantity": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "number"
                }
            }
        },
        "genprotos.CreateOrderRequest": {
            "type": "object",
            "properties": {
                "courier_id": {
                    "type": "string"
                },
                "delivery_address": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "genprotos.CreateOrderResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                }
            }
        },
        "genprotos.CreateProductRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "genprotos.DeleItemsRequest": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "string"
                }
            }
        },
        "genprotos.DeleteOrdersByidResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "succes": {
                    "type": "boolean"
                }
            }
        },
        "genprotos.DeleteProductResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "genprotos.GetAllItems": {
            "type": "object",
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.GetProductResponse"
                    }
                },
                "quantity": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "number"
                }
            }
        },
        "genprotos.GetAllOrders": {
            "type": "object",
            "properties": {
                "courier_id": {
                    "type": "string"
                },
                "delivery_address": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.GetProductResponse"
                    }
                },
                "status": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "total_amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "genprotos.GetAllOrdersList": {
            "type": "object",
            "properties": {
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.GetAllOrders"
                    }
                }
            }
        },
        "genprotos.GetAllProductResponse": {
            "type": "object",
            "properties": {
                "allproducts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.GetProductResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "genprotos.GetProductResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "product_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "genprotos.ProductResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "genprotos.UpdateOrderRequest": {
            "type": "object",
            "properties": {
                "delivery_address": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                }
            }
        },
        "genprotos.UpdateProductRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "genprotos.UpdateStatusRequest": {
            "type": "object",
            "properties": {
                "order_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "genprotos.UpdateStatusResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "succes": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API Gateway",
	Description:      "API documentation for the gateway service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

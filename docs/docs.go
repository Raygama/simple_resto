// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/carts": {
            "get": {
                "description": "Menginisialisasi cart kosong yang nanti akan di isi oleh user id tertentu",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Get empty cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Cart"
                        }
                    }
                }
            }
        },
        "/carts/{cart_id}": {
            "delete": {
                "description": "delete an entire cart along with all its menus",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "delete cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cart_id",
                        "name": "cart_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/carts/{cart_id}/empty": {
            "delete": {
                "description": "remove all menus from an existing cart",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "empty cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cart_id",
                        "name": "cart_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Cart"
                        }
                    }
                }
            }
        },
        "/carts/{cart_id}/menus/{menu_id}": {
            "put": {
                "description": "update menu quantity in an existing cart or remove if quantity is 0",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "update menu in cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cart_id",
                        "name": "cart_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "menu_id",
                        "name": "menu_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "quantity",
                        "name": "quantity",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Cart"
                        }
                    }
                }
            },
            "post": {
                "description": "add menu to an existing cart",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "add menu to cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cart_id",
                        "name": "cart_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "menu_id",
                        "name": "menu_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Cart"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete menu from an existing cart",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "delete menu from cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cart_id",
                        "name": "cart_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "menu_id",
                        "name": "menu_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Cart"
                        }
                    }
                }
            }
        },
        "/carts/{id}/menus": {
            "get": {
                "description": "Get menus from a specific cart using cart id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Get menus by cart id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id cart",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controller.MenuWithQuantity"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "login with username and password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login as a user",
                "parameters": [
                    {
                        "description": "the body to login a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/menus": {
            "get": {
                "description": "Get all list of menus",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "Get all menu",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Menu"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Create new menu",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "Create new menu",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Menu name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Menu price",
                        "name": "price",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Menu image",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization, how to input in swagger: 'Bearer \u003ctoken\u003e' ",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Menu"
                        }
                    }
                }
            }
        },
        "/menus/{id}": {
            "get": {
                "description": "Get a menu by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "Get menu by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Menu ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Menu"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "update menu",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "update menu",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id menu",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "nama menu",
                        "name": "nama",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "harga menu",
                        "name": "price",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "gambar menu",
                        "name": "image",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Authorization, how to input in swagger: 'Bearer \u003ctoken\u003e' ",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Menu"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete a single menu by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "Delete a menu",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id menu",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization, how to input in swagger: 'Bearer \u003ctoken\u003e' ",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            }
        },
        "/menus/{id}/cart-items": {
            "get": {
                "description": "get cart items by menu ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "get cart items by menu ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Menu ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.CartItem"
                            }
                        }
                    }
                }
            }
        },
        "/menus/{id}/carts": {
            "get": {
                "description": "get cart by menu ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "get cart by menu ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Menu ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Cart"
                            }
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "registering a user from public access",
                "tags": [
                    "Auth"
                ],
                "summary": "register a new user",
                "parameters": [
                    {
                        "description": "body to register a new user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controller.MenuWithQuantity": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "imageurl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "controller.RegisterInput": {
            "type": "object",
            "required": [
                "password",
                "role",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Cart": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.CartItem": {
            "type": "object",
            "properties": {
                "cart_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "menu_id": {
                    "type": "integer"
                },
                "qty": {
                    "type": "integer"
                }
            }
        },
        "models.Menu": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "imageurl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

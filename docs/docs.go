// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "login param",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.respResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.User"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/notifier": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notify"
                ],
                "summary": "page query notify event",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page number",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.respResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/response.PageResult"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "list": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/models.Contract"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notify"
                ],
                "summary": "create notify event",
                "parameters": [
                    {
                        "description": "add notifier param",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Contract"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.respResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Contract"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/notifier/history": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notify"
                ],
                "summary": "page query update history",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "contract id",
                        "name": "contractID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page number",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.respResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Contract"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.respResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "models.Contract": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "adminAddr": {
                    "type": "string"
                },
                "contractHistoryArr": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ContractHistory"
                    }
                },
                "createdAt": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastUpdate": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "network": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "integer"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.ContractHistory": {
            "type": "object",
            "properties": {
                "contractId": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "network": {
                    "type": "integer"
                },
                "newOwner": {
                    "type": "string"
                },
                "previousOwner": {
                    "type": "string"
                },
                "updateBlock": {
                    "type": "integer"
                },
                "updateTX": {
                    "type": "string"
                },
                "updateTime": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "contractArr": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Contract"
                    }
                },
                "createdAt": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "integer"
                }
            }
        },
        "request.Contract": {
            "type": "object",
            "properties": {
                "contractAddr": {
                    "type": "string"
                },
                "contractName": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "network": {
                    "type": "integer"
                }
            }
        },
        "request.Login": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "signData": {
                    "type": "string"
                },
                "signature": {
                    "type": "string"
                }
            }
        },
        "response.PageResult": {
            "type": "object",
            "properties": {
                "list": {},
                "pageNum": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "easy-upgrade-backend",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

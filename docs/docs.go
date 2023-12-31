// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Olusola Amoo",
            "url": "https://twitter.com/aosimeon",
            "email": "aosimeon@outlook.com"
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
        "/shorten": {
            "post": {
                "description": "Shorten a url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Shorten a url",
                "parameters": [
                    {
                        "description": "Url to shorten e.g https://www.altschoolafrica.com/",
                        "name": "longUrl",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.UrlPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Url"
                        }
                    }
                }
            }
        },
        "/{short-url}": {
            "get": {
                "description": "Redirect to a url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Redirect to a url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short url",
                        "name": "short-url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "301": {
                        "description": "Moved Permanently"
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "main.UrlPayload": {
            "type": "object",
            "properties": {
                "longUrl": {
                    "type": "string"
                }
            }
        },
        "models.Url": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "longUrl": {
                    "type": "string"
                },
                "shortUrl": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "url-shortner-3b6b.onrender.com",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "AltSchool Capstone Project - URL Shortner",
	Description:      "Shorten long urls.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

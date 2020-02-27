// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/collect/{app_token}": {
            "post": {
                "description": "add by json account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collect"
                ],
                "summary": "Add a batch of records",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Application token",
                        "name": "app_token",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Add records batch",
                        "name": "batchRecords",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app_record.RecordsBatch"
                        }
                    }
                ],
                "responses": {
                    "202": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest_errors.RestErrStruct"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest_errors.RestErrStruct"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest_errors.RestErrStruct"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app_record.Element": {
            "type": "object",
            "properties": {
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app_record.Element"
                    }
                },
                "height": {
                    "type": "integer",
                    "example": 100
                },
                "id": {
                    "type": "integer",
                    "example": 122323
                },
                "name": {
                    "type": "string",
                    "example": "Self Destruct Button"
                },
                "pos_x": {
                    "type": "number",
                    "example": 10
                },
                "pos_y": {
                    "type": "number",
                    "example": 15
                },
                "radius": {
                    "type": "integer",
                    "example": 70
                },
                "touched": {
                    "type": "boolean",
                    "example": true
                },
                "type": {
                    "type": "string",
                    "example": "Button"
                },
                "width": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "app_record.Record": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "screen_layout": {
                    "type": "object",
                    "$ref": "#/definitions/app_record.Element"
                },
                "screen_name": {
                    "type": "string",
                    "example": "Main screen"
                },
                "timestamp": {
                    "type": "integer",
                    "example": 1582759519
                },
                "touch_events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app_record.TouchEvent"
                    }
                },
                "transition_to": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "app_record.RecordsBatch": {
            "type": "object",
            "properties": {
                "app_token": {
                    "type": "string"
                },
                "client_type": {
                    "type": "string",
                    "example": "iPhone10"
                },
                "records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app_record.Record"
                    }
                },
                "session_id": {
                    "type": "string",
                    "example": "dyw234kjbb"
                }
            }
        },
        "app_record.TouchEvent": {
            "type": "object",
            "properties": {
                "pos_x": {
                    "type": "number",
                    "example": 10
                },
                "pos_y": {
                    "type": "number",
                    "example": 20
                },
                "touch_type": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "rest_errors.RestErrStruct": {
            "type": "object",
            "properties": {
                "causes": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                },
                "error": {
                    "type": "string",
                    "example": "internal server error"
                },
                "message": {
                    "type": "string",
                    "example": "error description"
                },
                "status": {
                    "type": "integer",
                    "example": 500
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Collector API",
	Description: "Collector service API",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}

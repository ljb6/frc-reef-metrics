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
        "/all-matches": {
            "get": {
                "description": "Retorna uma lista completa dos dados coletados",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stats"
                ],
                "summary": "Obtém os dados de todas as partidas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.MatchStats"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/match/{match}": {
            "get": {
                "description": "Retorna os dados de uma partida pelo número da partida",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stats"
                ],
                "summary": "Obtém os dados de uma partida",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Número da partida",
                        "name": "match",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MatchStats"
                        }
                    },
                    "400": {
                        "description": "Erro: Número da partida inválido",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Erro interno",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/match/{match}/{team}": {
            "get": {
                "description": "Retorna os dados de um time dentro de uma determinada partida",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stats"
                ],
                "summary": "Obtém os dados de um time em uma partida específica",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Número da partida",
                        "name": "match",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Número do time",
                        "name": "team",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MatchStats"
                        }
                    },
                    "400": {
                        "description": "Erro: Parâmetros inválidos",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Erro interno",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/matches/{team}": {
            "get": {
                "description": "Retorna os dados de um time específico pelo número",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stats"
                ],
                "summary": "Obtém os dados coletados de um time",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Número do time",
                        "name": "team",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MatchStats"
                        }
                    },
                    "400": {
                        "description": "Erro: Número do time inválido",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Erro interno",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.MatchStats": {
            "type": "object",
            "properties": {
                "auto_l1_corals": {
                    "type": "integer"
                },
                "auto_l2_corals": {
                    "type": "integer"
                },
                "auto_l3_corals": {
                    "type": "integer"
                },
                "auto_l4_corals": {
                    "type": "integer"
                },
                "auto_left": {
                    "description": "Autonomous",
                    "type": "boolean"
                },
                "auto_net": {
                    "type": "integer"
                },
                "auto_processor": {
                    "type": "integer"
                },
                "comments": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "end_climb_attempt": {
                    "type": "boolean"
                },
                "end_climb_failed": {
                    "type": "boolean"
                },
                "end_climb_level": {
                    "type": "string"
                },
                "end_fouls": {
                    "type": "integer"
                },
                "end_park": {
                    "description": "End",
                    "type": "boolean"
                },
                "match_level": {
                    "type": "string"
                },
                "match_number": {
                    "type": "integer"
                },
                "name": {
                    "description": "Pre match",
                    "type": "string"
                },
                "played_defense": {
                    "type": "boolean"
                },
                "removed_algae": {
                    "description": "Both",
                    "type": "boolean"
                },
                "robot_failed": {
                    "type": "boolean"
                },
                "start_zone": {
                    "type": "string"
                },
                "team_number": {
                    "type": "integer"
                },
                "tele_l1_corals": {
                    "description": "Teleop",
                    "type": "integer"
                },
                "tele_l2_corals": {
                    "type": "integer"
                },
                "tele_l3_corals": {
                    "type": "integer"
                },
                "tele_l4_corals": {
                    "type": "integer"
                },
                "tele_net": {
                    "type": "integer"
                },
                "tele_processor": {
                    "type": "integer"
                },
                "trapped_in_algae": {
                    "type": "boolean"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "message": {
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

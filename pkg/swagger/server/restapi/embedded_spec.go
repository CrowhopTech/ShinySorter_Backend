// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json",
    "multipart/form-data"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Endpoint definitions for the shiny-sorter image sorting project",
    "title": "shiny-sorter",
    "version": "alpha-v0.0"
  },
  "paths": {
    "/healthz": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "checkHealth",
        "responses": {
          "200": {
            "description": "OK message",
            "schema": {
              "type": "string",
              "enum": [
                "OK"
              ]
            }
          },
          "503": {
            "description": "Server still starting",
            "schema": {
              "type": "string",
              "enum": [
                "Service Unavailable"
              ]
            }
          }
        }
      }
    },
    "/images": {
      "get": {
        "description": "Lists and queries images",
        "operationId": "listImages",
        "parameters": [
          {
            "$ref": "#/parameters/includeTags"
          },
          {
            "$ref": "#/parameters/includeOperator"
          },
          {
            "$ref": "#/parameters/excludeTags"
          },
          {
            "$ref": "#/parameters/excludeOperator"
          },
          {
            "type": "boolean",
            "description": "Whether to filter to tags that have or have not been tagged",
            "name": "hasBeenTagged",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Images were found matching the given query",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/image"
              }
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "No images were found matching the given query. Also returns an empty array for easier parsing",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/image"
              }
            }
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      },
      "post": {
        "description": "Creates a new image entry",
        "operationId": "createImage",
        "parameters": [
          {
            "description": "The new image to create",
            "name": "newImage",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/image"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Image was created successfully",
            "schema": {
              "$ref": "#/definitions/image"
            }
          },
          "400": {
            "description": "Some part of the provided Image was invalid."
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      }
    },
    "/images/contents/{id}": {
      "get": {
        "description": "Gets the image contents with the specified id",
        "produces": [
          "application/octet-stream"
        ],
        "operationId": "getImageContent",
        "parameters": [
          {
            "type": "string",
            "description": "Image ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "type": "boolean",
            "description": "Whether to return the actual contents or a thumbnail",
            "name": "thumb",
            "in": "query",
            "allowEmptyValue": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the image or thumbnail contents",
            "schema": {
              "type": "string",
              "format": "binary"
            },
            "headers": {
              "Content-Type": {
                "type": "string"
              }
            }
          },
          "404": {
            "description": "The given image was not found."
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      },
      "patch": {
        "description": "Sets the image contents for the specified id",
        "consumes": [
          "multipart/form-data"
        ],
        "operationId": "setImageContent",
        "parameters": [
          {
            "type": "file",
            "format": "binary",
            "description": "The file contents to upload.",
            "name": "fileContents",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "Image ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The image contents were modified successfully"
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "The given image was not found."
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      }
    },
    "/images/{id}": {
      "get": {
        "description": "Gets the image metadata with the specified id",
        "operationId": "getImageById",
        "parameters": [
          {
            "type": "string",
            "description": "Image ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the found image.",
            "schema": {
              "$ref": "#/definitions/image"
            }
          },
          "404": {
            "description": "The given image was not found."
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      },
      "patch": {
        "description": "Modifies the image metadata with the specified id",
        "operationId": "patchImageById",
        "parameters": [
          {
            "type": "string",
            "description": "Image ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Patch modifications for the image",
            "name": "patch",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/image"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the modified image.",
            "schema": {
              "$ref": "#/definitions/image"
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      }
    },
    "/questions": {
      "get": {
        "description": "Lists questions",
        "operationId": "listQuestions",
        "responses": {
          "200": {
            "description": "Questions were listed successfully (array may be empty if no questions are registered)",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/question"
              }
            }
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      },
      "post": {
        "description": "Creates a new question",
        "operationId": "createQuestion",
        "parameters": [
          {
            "description": "The new question to create",
            "name": "newQuestion",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/question"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Question was created successfully",
            "schema": {
              "$ref": "#/definitions/question"
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      }
    },
    "/questions/{id}": {
      "delete": {
        "description": "Deletes a question.",
        "operationId": "deleteQuestion",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the question to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Question was deleted successfully"
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      },
      "patch": {
        "description": "Modifies question metadata",
        "operationId": "patchQuestionByID",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the question to modify",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Patch modifications for the question",
            "name": "patch",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/question"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Question was modified successfully",
            "schema": {
              "$ref": "#/definitions/question"
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      }
    },
    "/tags": {
      "get": {
        "description": "Lists tags and their metadata",
        "operationId": "listTags",
        "responses": {
          "200": {
            "description": "Tags were listed successfully (array may be empty if no tags are registered)",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/tag"
              }
            }
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      },
      "post": {
        "description": "Creates a new tag",
        "operationId": "createTag",
        "parameters": [
          {
            "description": "The new tag to create",
            "name": "newTag",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/tag"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Tag was created successfully",
            "schema": {
              "$ref": "#/definitions/tag"
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      }
    },
    "/tags/{id}": {
      "delete": {
        "description": "Deletes a tag. Should also remove it from all images that use it.",
        "operationId": "deleteTag",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the tag to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Tag was deleted successfully",
            "schema": {
              "$ref": "#/definitions/tag"
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      },
      "patch": {
        "description": "Modifies tag metadata such as description, icon, etc.",
        "operationId": "patchTagByID",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the tag to modify",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Patch modifications for the tag",
            "name": "patch",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/tag"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Tag was modified successfully",
            "schema": {
              "$ref": "#/definitions/tag"
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "$ref": "#/responses/genericServerError"
          }
        }
      }
    }
  },
  "definitions": {
    "image": {
      "properties": {
        "hasBeenTagged": {
          "type": "boolean",
          "default": true
        },
        "id": {
          "type": "string",
          "example": "filename.jpg"
        },
        "md5sum": {
          "type": "string",
          "example": "0a8bd0c4863ec1720da0f69d2795d18a"
        },
        "mimeType": {
          "type": "string",
          "example": "image/png"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "integer"
          },
          "example": [
            5,
            7,
            37
          ]
        }
      }
    },
    "question": {
      "properties": {
        "orderingID": {
          "type": "integer"
        },
        "questionID": {
          "type": "integer"
        },
        "questionText": {
          "type": "string"
        },
        "tagOptions": {
          "type": "array",
          "items": {
            "type": "object",
            "required": [
              "tagID",
              "optionText"
            ],
            "properties": {
              "optionText": {
                "type": "string"
              },
              "tagID": {
                "type": "integer"
              }
            }
          }
        }
      },
      "example": {
        "orderingID": 500,
        "questionID": 5,
        "questionText": "What kinds of flowers are present in this picture?",
        "tagOptions": [
          {
            "optionText": "Tulips",
            "tagID": 5
          },
          {
            "optionText": "Roses",
            "tagID": 6
          },
          {
            "optionText": "Violets",
            "tagID": 7
          },
          {
            "optionText": "Daisies",
            "tagID": 8
          }
        ]
      }
    },
    "tag": {
      "properties": {
        "description": {
          "type": "string",
          "example": "This image contains a Tulip"
        },
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string",
          "example": "flower:type:tulip"
        },
        "userFriendlyName": {
          "type": "string",
          "example": "Tulip"
        }
      }
    }
  },
  "parameters": {
    "excludeOperator": {
      "enum": [
        "all",
        "any"
      ],
      "type": "string",
      "default": "all",
      "description": "Whether excludeTags requires all tags to match, or just one",
      "name": "excludeOperator",
      "in": "query"
    },
    "excludeTags": {
      "type": "array",
      "items": {
        "type": "integer"
      },
      "description": "Tags to exclude in this query, referenced by tag ID",
      "name": "excludeTags",
      "in": "query"
    },
    "includeOperator": {
      "enum": [
        "all",
        "any"
      ],
      "type": "string",
      "default": "all",
      "description": "Whether includeTags requires all tags to match, or just one",
      "name": "includeOperator",
      "in": "query"
    },
    "includeTags": {
      "type": "array",
      "items": {
        "type": "integer"
      },
      "description": "Tags to include in this query, referenced by tag ID",
      "name": "includeTags",
      "in": "query"
    }
  },
  "responses": {
    "genericServerError": {
      "description": "Something else went wrong during the request",
      "schema": {
        "type": "string"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json",
    "multipart/form-data"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Endpoint definitions for the shiny-sorter image sorting project",
    "title": "shiny-sorter",
    "version": "alpha-v0.0"
  },
  "paths": {
    "/healthz": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "checkHealth",
        "responses": {
          "200": {
            "description": "OK message",
            "schema": {
              "type": "string",
              "enum": [
                "OK"
              ]
            }
          },
          "503": {
            "description": "Server still starting",
            "schema": {
              "type": "string",
              "enum": [
                "Service Unavailable"
              ]
            }
          }
        }
      }
    },
    "/images": {
      "get": {
        "description": "Lists and queries images",
        "operationId": "listImages",
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "integer"
            },
            "description": "Tags to include in this query, referenced by tag ID",
            "name": "includeTags",
            "in": "query"
          },
          {
            "enum": [
              "all",
              "any"
            ],
            "type": "string",
            "default": "all",
            "description": "Whether includeTags requires all tags to match, or just one",
            "name": "includeOperator",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "integer"
            },
            "description": "Tags to exclude in this query, referenced by tag ID",
            "name": "excludeTags",
            "in": "query"
          },
          {
            "enum": [
              "all",
              "any"
            ],
            "type": "string",
            "default": "all",
            "description": "Whether excludeTags requires all tags to match, or just one",
            "name": "excludeOperator",
            "in": "query"
          },
          {
            "type": "boolean",
            "description": "Whether to filter to tags that have or have not been tagged",
            "name": "hasBeenTagged",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Images were found matching the given query",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/image"
              }
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "No images were found matching the given query. Also returns an empty array for easier parsing",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/image"
              }
            }
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "description": "Creates a new image entry",
        "operationId": "createImage",
        "parameters": [
          {
            "description": "The new image to create",
            "name": "newImage",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/image"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Image was created successfully",
            "schema": {
              "$ref": "#/definitions/image"
            }
          },
          "400": {
            "description": "Some part of the provided Image was invalid."
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/images/contents/{id}": {
      "get": {
        "description": "Gets the image contents with the specified id",
        "produces": [
          "application/octet-stream"
        ],
        "operationId": "getImageContent",
        "parameters": [
          {
            "type": "string",
            "description": "Image ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "type": "boolean",
            "description": "Whether to return the actual contents or a thumbnail",
            "name": "thumb",
            "in": "query",
            "allowEmptyValue": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the image or thumbnail contents",
            "schema": {
              "type": "string",
              "format": "binary"
            },
            "headers": {
              "Content-Type": {
                "type": "string"
              }
            }
          },
          "404": {
            "description": "The given image was not found."
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "patch": {
        "description": "Sets the image contents for the specified id",
        "consumes": [
          "multipart/form-data"
        ],
        "operationId": "setImageContent",
        "parameters": [
          {
            "type": "file",
            "format": "binary",
            "description": "The file contents to upload.",
            "name": "fileContents",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "Image ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The image contents were modified successfully"
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "The given image was not found."
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/images/{id}": {
      "get": {
        "description": "Gets the image metadata with the specified id",
        "operationId": "getImageById",
        "parameters": [
          {
            "type": "string",
            "description": "Image ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the found image.",
            "schema": {
              "$ref": "#/definitions/image"
            }
          },
          "404": {
            "description": "The given image was not found."
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "patch": {
        "description": "Modifies the image metadata with the specified id",
        "operationId": "patchImageById",
        "parameters": [
          {
            "type": "string",
            "description": "Image ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Patch modifications for the image",
            "name": "patch",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/image"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the modified image.",
            "schema": {
              "$ref": "#/definitions/image"
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/questions": {
      "get": {
        "description": "Lists questions",
        "operationId": "listQuestions",
        "responses": {
          "200": {
            "description": "Questions were listed successfully (array may be empty if no questions are registered)",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/question"
              }
            }
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "description": "Creates a new question",
        "operationId": "createQuestion",
        "parameters": [
          {
            "description": "The new question to create",
            "name": "newQuestion",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/question"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Question was created successfully",
            "schema": {
              "$ref": "#/definitions/question"
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/questions/{id}": {
      "delete": {
        "description": "Deletes a question.",
        "operationId": "deleteQuestion",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the question to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Question was deleted successfully"
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "patch": {
        "description": "Modifies question metadata",
        "operationId": "patchQuestionByID",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the question to modify",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Patch modifications for the question",
            "name": "patch",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/question"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Question was modified successfully",
            "schema": {
              "$ref": "#/definitions/question"
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/tags": {
      "get": {
        "description": "Lists tags and their metadata",
        "operationId": "listTags",
        "responses": {
          "200": {
            "description": "Tags were listed successfully (array may be empty if no tags are registered)",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/tag"
              }
            }
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "description": "Creates a new tag",
        "operationId": "createTag",
        "parameters": [
          {
            "description": "The new tag to create",
            "name": "newTag",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/tag"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Tag was created successfully",
            "schema": {
              "$ref": "#/definitions/tag"
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/tags/{id}": {
      "delete": {
        "description": "Deletes a tag. Should also remove it from all images that use it.",
        "operationId": "deleteTag",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the tag to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Tag was deleted successfully",
            "schema": {
              "$ref": "#/definitions/tag"
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "patch": {
        "description": "Modifies tag metadata such as description, icon, etc.",
        "operationId": "patchTagByID",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the tag to modify",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Patch modifications for the tag",
            "name": "patch",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/tag"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Tag was modified successfully",
            "schema": {
              "$ref": "#/definitions/tag"
            }
          },
          "400": {
            "description": "Some part of the request was invalid. More information will be included in the error string",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Something else went wrong during the request",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "QuestionTagOptionsItems0": {
      "type": "object",
      "required": [
        "tagID",
        "optionText"
      ],
      "properties": {
        "optionText": {
          "type": "string"
        },
        "tagID": {
          "type": "integer"
        }
      }
    },
    "image": {
      "properties": {
        "hasBeenTagged": {
          "type": "boolean",
          "default": true
        },
        "id": {
          "type": "string",
          "example": "filename.jpg"
        },
        "md5sum": {
          "type": "string",
          "example": "0a8bd0c4863ec1720da0f69d2795d18a"
        },
        "mimeType": {
          "type": "string",
          "example": "image/png"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "integer"
          },
          "example": [
            5,
            7,
            37
          ]
        }
      }
    },
    "question": {
      "properties": {
        "orderingID": {
          "type": "integer"
        },
        "questionID": {
          "type": "integer"
        },
        "questionText": {
          "type": "string"
        },
        "tagOptions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/QuestionTagOptionsItems0"
          }
        }
      },
      "example": {
        "orderingID": 500,
        "questionID": 5,
        "questionText": "What kinds of flowers are present in this picture?",
        "tagOptions": [
          {
            "optionText": "Tulips",
            "tagID": 5
          },
          {
            "optionText": "Roses",
            "tagID": 6
          },
          {
            "optionText": "Violets",
            "tagID": 7
          },
          {
            "optionText": "Daisies",
            "tagID": 8
          }
        ]
      }
    },
    "tag": {
      "properties": {
        "description": {
          "type": "string",
          "example": "This image contains a Tulip"
        },
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string",
          "example": "flower:type:tulip"
        },
        "userFriendlyName": {
          "type": "string",
          "example": "Tulip"
        }
      }
    }
  },
  "parameters": {
    "excludeOperator": {
      "enum": [
        "all",
        "any"
      ],
      "type": "string",
      "default": "all",
      "description": "Whether excludeTags requires all tags to match, or just one",
      "name": "excludeOperator",
      "in": "query"
    },
    "excludeTags": {
      "type": "array",
      "items": {
        "type": "integer"
      },
      "description": "Tags to exclude in this query, referenced by tag ID",
      "name": "excludeTags",
      "in": "query"
    },
    "includeOperator": {
      "enum": [
        "all",
        "any"
      ],
      "type": "string",
      "default": "all",
      "description": "Whether includeTags requires all tags to match, or just one",
      "name": "includeOperator",
      "in": "query"
    },
    "includeTags": {
      "type": "array",
      "items": {
        "type": "integer"
      },
      "description": "Tags to include in this query, referenced by tag ID",
      "name": "includeTags",
      "in": "query"
    }
  },
  "responses": {
    "genericServerError": {
      "description": "Something else went wrong during the request",
      "schema": {
        "type": "string"
      }
    }
  }
}`))
}

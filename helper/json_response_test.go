package helper

import (
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestJSONSchema(t *testing.T) {
	schema := `
	{
		"id": "create_review_schema",
		"$schema": "http://json-schema.org/draft-04/schema#",
		"description": "create review schema",
		"type": "object",
		"title": "The Review Schema",
		"required": [
			"sku",
			"title",
			"content",
			"rating",
			"merchantId",
			"orderId",
			"soNumber"
		],
		"properties": {
			"sku": {
				"$id": "#/properties/sku",
				"type": "string",
				"title": "The Sku Schema",
				"default": "",
				"examples": [
					"sku3320909365"
				],
				"maxLength": 20,
				"minLength": 8,
				"pattern": "^[A-Za-z0-9]+$"
			},
			"name": {
				"$id": "#/properties/name",
				"type": "string",
				"title": "The Name Schema",
				"default": "",
				"examples": [
					"Wuriyanto"
				],
				"maxLength": 50,
				"pattern": "^[a-zA-Z0-9\\,. \\/\\()-]+$"
			},
			"email": {
				"$id": "#/properties/email",
				"type": "string",
				"title": "The Email Schema",
				"default": "",
				"examples": [
					"wuriyanto@yahoo.com"
				],
				"maxLength": 50,
              	"format": "email"
			},
			"title": {
				"$id": "#/properties/title",
				"type": "string",
				"title": "The Title Schema",
				"default": "",
				"examples": [
					"yoooo pasti keren"
				],
				"pattern": "^[a-zA-Z0-9\\,. \\/\\()-]+$",
				"maxLength": 50,
				"minLength": 4
			},
			"content": {
				"$id": "#/properties/content",
				"type": "string",
				"title": "The Content Schema",
				"default": "",
				"examples": [
					"yooo keren abis"
				],
				"pattern": "^[a-zA-Z0-9\\,. \\/\\()-]+$",
				"maxLength": 1500,
				"minLength": 10
			},
			"rating": {
				"$id": "#/properties/rating",
				"type": "integer",
				"title": "The Rating Schema",
				"examples": [
					3
				],
				"minimum": 1,
				"maximum": 5
			},
			"merchantId": {
				"$id": "#/properties/merchantId",
				"type": "string",
				"title": "The Merchantid Schema",
				"default": "",
				"examples": [
					"M1"
				],
				"pattern": "^[A-Za-z0-9]+$"
			},
			"orderId": {
				"$id": "#/properties/orderId",
				"type": "string",
				"title": "The Orderid Schema",
				"default": "",
				"examples": [
					"1112"
				],
				"pattern": "^[0-9]+$"
			},
			"soNumber": {
				"$id": "#/properties/soNumber",
				"type": "string",
				"title": "The So number Schema",
				"default": "",
				"examples": [
					"SO001"
				],
				"pattern": "^[A-Za-z0-9]+$"
			}
		}
	}
	 `

	t.Run("should return success validate json schema", func(t *testing.T) {

		input := `
		{
		   "sku" : "sku3320909365",
		   "name" : "ryan",
		   "email" : "ryan@yahoo.com",
		   "title" : "yoooo pasti keren",
		   "content" : "yooo keren abis",
		   "rating" : 3,
		   "merchantId" : "M1",
		   "orderId" : "1112",
		   "soNumber": "SO001"
	   }
		`

		err := ValidateJSONSchema(schema, input)

		assert.NoError(t, err)
	})

	t.Run("should return success validate json schema with sku number only", func(t *testing.T) {

		input := `
		{
		   "sku" : "3320909365",
		   "name" : "ryan",
		   "email" : "ryanyahoo.com",
		   "title" : "yoooo pasti keren",
		   "content" : "yooo keren abis",
		   "rating" : 1,
		   "merchantId" : "M1",
		   "orderId" : "1112",
		   "soNumber": "SO001"
	   }
		`

		err := ValidateJSONSchema(schema, input)

		assert.Error(t, err)
	})

	t.Run("should return error validate json schema with minimum rating 0", func(t *testing.T) {

		input := `
		{
		   "sku" : "sku3320909365",
		   "name" : "ryan",
		   "email" : "ryan@yahoo.com",
		   "title" : "yoooo pasti keren",
		   "content" : "yooo keren abis",
		   "rating" : 0,
		   "merchantId" : "M1",
		   "orderId" : "1112",
		   "soNumber": "SO001"
	   }
		`

		err := ValidateJSONSchema(schema, input)

		assert.Error(t, err)
	})

	t.Run("should return error validate json schema with invalid email", func(t *testing.T) {

		input := `
		{
		   "sku" : "sku3320909365",
		   "name" : "ryan",
		   "email" : "ryanyahoo.com",
		   "title" : "yoooo pasti keren",
		   "content" : "yooo keren abis",
		   "rating" : 1,
		   "merchantId" : "M1",
		   "orderId" : "1112",
		   "soNumber": "SO001"
	   }
		`

		err := ValidateJSONSchema(schema, input)

		assert.Error(t, err)
	})

	t.Run("should return error validate json schema with title less than 4", func(t *testing.T) {

		input := `
		{
		   "sku" : "sku3320909365",
		   "name" : "ryan",
		   "email" : "ryanyahoo.com",
		   "title" : "yoo",
		   "content" : "yooo keren abis",
		   "rating" : 1,
		   "merchantId" : "M1",
		   "orderId" : "1112",
		   "soNumber": "SO001"
	   }
		`

		err := ValidateJSONSchema(schema, input)

		assert.Error(t, err)
	})
}

func TestShowHTTPResponse(t *testing.T) {

	c, _ := NewRecorder(t, "http://localhost/test", nil, echo.POST)
	t.Run("#1: Positive", func(t *testing.T) {
		temp := JSONSchemaTemplate{}
		err := temp.ShowHTTPResponse(c)
		assert.NoError(t, err)
	})
}

func TestData(t *testing.T) {

	t.Run("#1: Positive", func(t *testing.T) {
		temp := JSONSchemaTemplate{}
		temp.SetData("data")
		assert.EqualValues(t, "data", temp.Data)
	})
}

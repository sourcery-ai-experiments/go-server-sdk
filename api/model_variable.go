/*
 * DevCycle Bucketing API
 *
 * Documents the DevCycle Bucketing API which provides and API interface to User Bucketing and for generated SDKs.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package api

type BaseVariable struct {
	// Unique key by Project, can be used in the SDK / API to reference by 'key' rather than _id.
	Key string `json:"key"`
	// Variable type
	Type_ string `json:"type"`
	// Variable value can be a string, number, boolean, or JSON
	Value interface{} `json:"value"`
}

type Variable struct {
	BaseVariable
	// Default variable value can be a string, number, boolean, or JSON
	DefaultValue interface{} `json:"defaultValue"`
	// Identifies if variable was returned with the default value
	IsDefaulted bool `json:"isDefaulted"`
}

type ReadOnlyVariable struct {
	BaseVariable
	// unique database id
	Id string `json:"_id"`
}
/*
 * DevCycle Bucketing API
 *
 * Documents the DevCycle Bucketing API which provides and API interface to User Bucketing and for generated SDKs.
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package devcycle

type Feature struct {
	// unique database id
	Id string `json:"_id"`
	// Unique key by Project, can be used in the SDK / API to reference by 'key' rather than _id.
	Key string `json:"key"`
	// Feature type
	Type_ string `json:"type"`
	// Bucketed feature variation
	Variation string `json:"_variation"`
	// Bucketed feature variation key
	VariationKey string `json:"variationKey"`
	// Bucketed feature variation name
	VariationName string `json:"variationName"`
	// Evaluation reasoning
	EvalReason string `json:"evalReason,omitempty"`
}

/*
 * Sonar coverage dashboard
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package model

type RepoCoverage struct {
	Id string `json:"id,omitempty"`

	Branch string `json:"branch,omitempty"`

	LinesToCover int32 `json:"linesToCover,omitempty"`

	UnitTests int32 `json:"unitTests,omitempty"`

	Coverage float64 `json:"coverage,omitempty"`
}

/*
 * PI Web API 2018 Swagger Spec
 *
 * Swagger Spec file that describes PI Web API
 *
 * API version: 1.11.0.640
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gowebapi

type PaginationLinks struct {
	First string `json:"First,omitempty"`

	Previous string `json:"Previous,omitempty"`

	Next string `json:"Next,omitempty"`

	Last string `json:"Last,omitempty"`
}
/*
 * NeXTLinux Enterprise Feeds
 *
 * Enterprise service for normalizing vulnerability data from external sources and making it available via API
 *
 * API version: 0.0.1
 * Contact: dev@next-linux.systems
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ServiceVersion Version information for a service
type ServiceVersion struct {
	Service ServiceVersionService `json:"service,omitempty"`
	Api ServiceVersionApi `json:"api,omitempty"`
	Db ServiceVersionDb `json:"db,omitempty"`
	Engine ServiceVersionEngine `json:"engine,omitempty"`
}

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
// FeedListing A collection of feed items and a 'next_token' for getting the next set of results
type FeedListing struct {
	Feeds []Feed `json:"feeds,omitempty"`
	// The token to include in subsequent requests to get the rest of the data. If empty, this response includes all data.
	NextToken string `json:"next_token,omitempty"`
}

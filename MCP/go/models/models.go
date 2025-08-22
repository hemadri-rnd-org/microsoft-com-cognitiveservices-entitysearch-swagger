package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// SearchResultsAnswer represents the SearchResultsAnswer schema from the OpenAPI specification
type SearchResultsAnswer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// Entities represents the Entities schema from the OpenAPI specification
type Entities struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Querycontext QueryContext `json:"queryContext,omitempty"` // Defines the query context that Bing used for the request.
}

// Response represents the Response schema from the OpenAPI specification
type Response struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
}

// SearchResponse represents the SearchResponse schema from the OpenAPI specification
type SearchResponse struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// Organization represents the Organization schema from the OpenAPI specification
type Organization struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
}

// LodgingBusiness represents the LodgingBusiness schema from the OpenAPI specification
type LodgingBusiness struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Address PostalAddress `json:"address,omitempty"` // Defines a postal address.
	Telephone string `json:"telephone,omitempty"` // The entity's telephone number
	Panoramas []ImageObject `json:"panoramas,omitempty"`
	Pricerange string `json:"priceRange,omitempty"` // $$.
	Tagline string `json:"tagLine,omitempty"`
	Ispermanentlyclosed bool `json:"isPermanentlyClosed,omitempty"`
}

// EntitiesEntityPresentationInfo represents the EntitiesEntityPresentationInfo schema from the OpenAPI specification
type EntitiesEntityPresentationInfo struct {
	Entityscenario string `json:"entityScenario"` // The supported scenario.
	Entitytypedisplayhint string `json:"entityTypeDisplayHint,omitempty"` // A display version of the entity hint. For example, if entityTypeHints is Artist, this field may be set to American Singer.
	Entitytypehints []string `json:"entityTypeHints,omitempty"` // A list of hints that indicate the entity's type. The list could contain a single hint such as Movie or a list of hints such as Place, LocalBusiness, Restaurant. Each successive hint in the array narrows the entity's type.
}

// ResponseBase represents the ResponseBase schema from the OpenAPI specification
type ResponseBase struct {
	TypeField string `json:"_type"`
}

// ContractualRulesAttribution represents the ContractualRulesAttribution schema from the OpenAPI specification
type ContractualRulesAttribution struct {
	TypeField string `json:"_type"`
	Targetpropertyname string `json:"targetPropertyName,omitempty"` // The name of the field that the rule applies to.
}

// CreativeWork represents the CreativeWork schema from the OpenAPI specification
type CreativeWork struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
}

// Thing represents the Thing schema from the OpenAPI specification
type Thing struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// PostalAddress represents the PostalAddress schema from the OpenAPI specification
type PostalAddress struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
}

// Restaurant represents the Restaurant schema from the OpenAPI specification
type Restaurant struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Address PostalAddress `json:"address,omitempty"` // Defines a postal address.
	Telephone string `json:"telephone,omitempty"` // The entity's telephone number
	Panoramas []ImageObject `json:"panoramas,omitempty"`
	Pricerange string `json:"priceRange,omitempty"` // $$.
	Tagline string `json:"tagLine,omitempty"`
	Ispermanentlyclosed bool `json:"isPermanentlyClosed,omitempty"`
}

// EntertainmentBusiness represents the EntertainmentBusiness schema from the OpenAPI specification
type EntertainmentBusiness struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Address PostalAddress `json:"address,omitempty"` // Defines a postal address.
	Telephone string `json:"telephone,omitempty"` // The entity's telephone number
	Tagline string `json:"tagLine,omitempty"`
	Ispermanentlyclosed bool `json:"isPermanentlyClosed,omitempty"`
	Panoramas []ImageObject `json:"panoramas,omitempty"`
	Pricerange string `json:"priceRange,omitempty"` // $$.
}

// Hotel represents the Hotel schema from the OpenAPI specification
type Hotel struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Address PostalAddress `json:"address,omitempty"` // Defines a postal address.
	Telephone string `json:"telephone,omitempty"` // The entity's telephone number
	Pricerange string `json:"priceRange,omitempty"` // $$.
	Tagline string `json:"tagLine,omitempty"`
	Ispermanentlyclosed bool `json:"isPermanentlyClosed,omitempty"`
	Panoramas []ImageObject `json:"panoramas,omitempty"`
}

// Places represents the Places schema from the OpenAPI specification
type Places struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Querycontext QueryContext `json:"queryContext,omitempty"` // Defines the query context that Bing used for the request.
}

// Identifiable represents the Identifiable schema from the OpenAPI specification
type Identifiable struct {
	TypeField string `json:"_type"`
}

// LocalBusiness represents the LocalBusiness schema from the OpenAPI specification
type LocalBusiness struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Address PostalAddress `json:"address,omitempty"` // Defines a postal address.
	Telephone string `json:"telephone,omitempty"` // The entity's telephone number
}

// ContractualRulesContractualRule represents the ContractualRulesContractualRule schema from the OpenAPI specification
type ContractualRulesContractualRule struct {
	TypeField string `json:"_type"`
	Targetpropertyname string `json:"targetPropertyName,omitempty"` // The name of the field that the rule applies to.
}

// Answer represents the Answer schema from the OpenAPI specification
type Answer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// Airport represents the Airport schema from the OpenAPI specification
type Airport struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Address PostalAddress `json:"address,omitempty"` // Defines a postal address.
	Telephone string `json:"telephone,omitempty"` // The entity's telephone number
}

// ContractualRulesLinkAttribution represents the ContractualRulesLinkAttribution schema from the OpenAPI specification
type ContractualRulesLinkAttribution struct {
	TypeField string `json:"_type"`
	Targetpropertyname string `json:"targetPropertyName,omitempty"` // The name of the field that the rule applies to.
	Mustbeclosetocontent bool `json:"mustBeCloseToContent,omitempty"` // A Boolean value that determines whether the contents of the rule must be placed in close proximity to the field that the rule applies to. If true, the contents must be placed in close proximity. If false, or this field does not exist, the contents may be placed at the caller's discretion.
}

// CivicStructure represents the CivicStructure schema from the OpenAPI specification
type CivicStructure struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Address PostalAddress `json:"address,omitempty"` // Defines a postal address.
	Telephone string `json:"telephone,omitempty"` // The entity's telephone number
}

// Place represents the Place schema from the OpenAPI specification
type Place struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
}

// FoodEstablishment represents the FoodEstablishment schema from the OpenAPI specification
type FoodEstablishment struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Address PostalAddress `json:"address,omitempty"` // Defines a postal address.
	Telephone string `json:"telephone,omitempty"` // The entity's telephone number
	Panoramas []ImageObject `json:"panoramas,omitempty"`
	Pricerange string `json:"priceRange,omitempty"` // $$.
	Tagline string `json:"tagLine,omitempty"`
	Ispermanentlyclosed bool `json:"isPermanentlyClosed,omitempty"`
}

// QueryContext represents the QueryContext schema from the OpenAPI specification
type QueryContext struct {
	Askuserforlocation bool `json:"askUserForLocation,omitempty"` // A Boolean value that indicates whether Bing requires the user's location to provide accurate results. If you specified the user's location by using the X-MSEdge-ClientIP and X-Search-Location headers, you can ignore this field. For location aware queries, such as "today's weather" or "restaurants near me" that need the user's location to provide accurate results, this field is set to true. For location aware queries that include the location (for example, "Seattle weather"), this field is set to false. This field is also set to false for queries that are not location aware, such as "best sellers".
	Originalquery string `json:"originalQuery"` // The query string as specified in the request.
	Adultintent bool `json:"adultIntent,omitempty"` // A Boolean value that indicates whether the specified query has adult intent. The value is true if the query has adult intent; otherwise, false.
	Alterationoverridequery string `json:"alterationOverrideQuery,omitempty"` // The query string to use to force Bing to use the original string. For example, if the query string is "saling downwind", the override query string will be "+saling downwind". Remember to encode the query string which results in "%2Bsaling+downwind". This field is included only if the original query string contains a spelling mistake.
	Alteredquery string `json:"alteredQuery,omitempty"` // The query string used by Bing to perform the query. Bing uses the altered query string if the original query string contained spelling mistakes. For example, if the query string is "saling downwind", the altered query string will be "sailing downwind". This field is included only if the original query string contains a spelling mistake.
}

// ContractualRulesTextAttribution represents the ContractualRulesTextAttribution schema from the OpenAPI specification
type ContractualRulesTextAttribution struct {
	Targetpropertyname string `json:"targetPropertyName,omitempty"` // The name of the field that the rule applies to.
	TypeField string `json:"_type"`
	Mustbeclosetocontent bool `json:"mustBeCloseToContent,omitempty"` // A Boolean value that determines whether the contents of the rule must be placed in close proximity to the field that the rule applies to. If true, the contents must be placed in close proximity. If false, or this field does not exist, the contents may be placed at the caller's discretion.
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// StructuredValue represents the StructuredValue schema from the OpenAPI specification
type StructuredValue struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
}

// ContractualRulesMediaAttribution represents the ContractualRulesMediaAttribution schema from the OpenAPI specification
type ContractualRulesMediaAttribution struct {
	TypeField string `json:"_type"`
	Targetpropertyname string `json:"targetPropertyName,omitempty"` // The name of the field that the rule applies to.
	Mustbeclosetocontent bool `json:"mustBeCloseToContent,omitempty"` // A Boolean value that determines whether the contents of the rule must be placed in close proximity to the field that the rule applies to. If true, the contents must be placed in close proximity. If false, or this field does not exist, the contents may be placed at the caller's discretion.
}

// MediaObject represents the MediaObject schema from the OpenAPI specification
type MediaObject struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"`
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
}

// MovieTheater represents the MovieTheater schema from the OpenAPI specification
type MovieTheater struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Address PostalAddress `json:"address,omitempty"` // Defines a postal address.
	Telephone string `json:"telephone,omitempty"` // The entity's telephone number
	Panoramas []ImageObject `json:"panoramas,omitempty"`
	Pricerange string `json:"priceRange,omitempty"` // $$.
	Tagline string `json:"tagLine,omitempty"`
	Ispermanentlyclosed bool `json:"isPermanentlyClosed,omitempty"`
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Code string `json:"code"` // The error code that identifies the category of error.
	Message string `json:"message"` // A description of the error.
	Moredetails string `json:"moreDetails,omitempty"` // A description that provides additional information about the error.
	Parameter string `json:"parameter,omitempty"` // The parameter in the request that caused the error.
	Subcode string `json:"subCode,omitempty"` // The error code that further helps to identify the error.
	Value string `json:"value,omitempty"` // The parameter's value in the request that was not valid.
}

// License represents the License schema from the OpenAPI specification
type License struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"`
}

// ContractualRulesLicenseAttribution represents the ContractualRulesLicenseAttribution schema from the OpenAPI specification
type ContractualRulesLicenseAttribution struct {
	TypeField string `json:"_type"`
	Targetpropertyname string `json:"targetPropertyName,omitempty"` // The name of the field that the rule applies to.
	Mustbeclosetocontent bool `json:"mustBeCloseToContent,omitempty"` // A Boolean value that determines whether the contents of the rule must be placed in close proximity to the field that the rule applies to. If true, the contents must be placed in close proximity. If false, or this field does not exist, the contents may be placed at the caller's discretion.
}

// ImageObject represents the ImageObject schema from the OpenAPI specification
type ImageObject struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Text string `json:"text,omitempty"`
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Height int `json:"height,omitempty"` // The height of the source media object, in pixels.
	Hostpageurl string `json:"hostPageUrl,omitempty"` // URL of the page that hosts the media object.
	Width int `json:"width,omitempty"` // The width of the source media object, in pixels.
	Contenturl string `json:"contentUrl,omitempty"` // Original URL to retrieve the source (file) for the media object (e.g the source URL for the image).
}

// TouristAttraction represents the TouristAttraction schema from the OpenAPI specification
type TouristAttraction struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Address PostalAddress `json:"address,omitempty"` // Defines a postal address.
	Telephone string `json:"telephone,omitempty"` // The entity's telephone number
}

// Intangible represents the Intangible schema from the OpenAPI specification
type Intangible struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Contractualrules []ContractualRulesContractualRule `json:"contractualRules,omitempty"` // A list of rules that you must adhere to if you display the item.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
}

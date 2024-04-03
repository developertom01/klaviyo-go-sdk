# AttributionResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AttributionEnum**](AttributionEnum.md) |  | 
**Id** | **string** | The ID of the attribution | 
**Relationships** | Pointer to [**AttributionResponseObjectResourceRelationships**](AttributionResponseObjectResourceRelationships.md) |  | [optional] 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewAttributionResponseObjectResource

`func NewAttributionResponseObjectResource(type_ AttributionEnum, id string, links ObjectLinks, ) *AttributionResponseObjectResource`

NewAttributionResponseObjectResource instantiates a new AttributionResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributionResponseObjectResourceWithDefaults

`func NewAttributionResponseObjectResourceWithDefaults() *AttributionResponseObjectResource`

NewAttributionResponseObjectResourceWithDefaults instantiates a new AttributionResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AttributionResponseObjectResource) GetType() AttributionEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttributionResponseObjectResource) GetTypeOk() (*AttributionEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttributionResponseObjectResource) SetType(v AttributionEnum)`

SetType sets Type field to given value.


### GetId

`func (o *AttributionResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttributionResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttributionResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *AttributionResponseObjectResource) GetRelationships() AttributionResponseObjectResourceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AttributionResponseObjectResource) GetRelationshipsOk() (*AttributionResponseObjectResourceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AttributionResponseObjectResource) SetRelationships(v AttributionResponseObjectResourceRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AttributionResponseObjectResource) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AttributionResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AttributionResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AttributionResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



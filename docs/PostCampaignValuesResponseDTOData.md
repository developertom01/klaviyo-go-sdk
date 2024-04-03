# PostCampaignValuesResponseDTOData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignValuesReportEnum**](CampaignValuesReportEnum.md) |  | 
**Attributes** | [**PostCampaignValuesResponseDTODataAttributes**](PostCampaignValuesResponseDTODataAttributes.md) |  | 
**Relationships** | Pointer to [**PostCampaignValuesResponseDTODataRelationships**](PostCampaignValuesResponseDTODataRelationships.md) |  | [optional] 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostCampaignValuesResponseDTOData

`func NewPostCampaignValuesResponseDTOData(type_ CampaignValuesReportEnum, attributes PostCampaignValuesResponseDTODataAttributes, links ObjectLinks, ) *PostCampaignValuesResponseDTOData`

NewPostCampaignValuesResponseDTOData instantiates a new PostCampaignValuesResponseDTOData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCampaignValuesResponseDTODataWithDefaults

`func NewPostCampaignValuesResponseDTODataWithDefaults() *PostCampaignValuesResponseDTOData`

NewPostCampaignValuesResponseDTODataWithDefaults instantiates a new PostCampaignValuesResponseDTOData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostCampaignValuesResponseDTOData) GetType() CampaignValuesReportEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCampaignValuesResponseDTOData) GetTypeOk() (*CampaignValuesReportEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCampaignValuesResponseDTOData) SetType(v CampaignValuesReportEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PostCampaignValuesResponseDTOData) GetAttributes() PostCampaignValuesResponseDTODataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostCampaignValuesResponseDTOData) GetAttributesOk() (*PostCampaignValuesResponseDTODataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostCampaignValuesResponseDTOData) SetAttributes(v PostCampaignValuesResponseDTODataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PostCampaignValuesResponseDTOData) GetRelationships() PostCampaignValuesResponseDTODataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PostCampaignValuesResponseDTOData) GetRelationshipsOk() (*PostCampaignValuesResponseDTODataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PostCampaignValuesResponseDTOData) SetRelationships(v PostCampaignValuesResponseDTODataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PostCampaignValuesResponseDTOData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *PostCampaignValuesResponseDTOData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostCampaignValuesResponseDTOData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostCampaignValuesResponseDTOData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



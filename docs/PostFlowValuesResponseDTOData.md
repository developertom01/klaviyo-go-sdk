# PostFlowValuesResponseDTOData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FlowValuesReportEnum**](FlowValuesReportEnum.md) |  | 
**Attributes** | [**PostFlowValuesResponseDTODataAttributes**](PostFlowValuesResponseDTODataAttributes.md) |  | 
**Relationships** | Pointer to [**PostFlowValuesResponseDTODataRelationships**](PostFlowValuesResponseDTODataRelationships.md) |  | [optional] 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostFlowValuesResponseDTOData

`func NewPostFlowValuesResponseDTOData(type_ FlowValuesReportEnum, attributes PostFlowValuesResponseDTODataAttributes, links ObjectLinks, ) *PostFlowValuesResponseDTOData`

NewPostFlowValuesResponseDTOData instantiates a new PostFlowValuesResponseDTOData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostFlowValuesResponseDTODataWithDefaults

`func NewPostFlowValuesResponseDTODataWithDefaults() *PostFlowValuesResponseDTOData`

NewPostFlowValuesResponseDTODataWithDefaults instantiates a new PostFlowValuesResponseDTOData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostFlowValuesResponseDTOData) GetType() FlowValuesReportEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostFlowValuesResponseDTOData) GetTypeOk() (*FlowValuesReportEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostFlowValuesResponseDTOData) SetType(v FlowValuesReportEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PostFlowValuesResponseDTOData) GetAttributes() PostFlowValuesResponseDTODataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostFlowValuesResponseDTOData) GetAttributesOk() (*PostFlowValuesResponseDTODataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostFlowValuesResponseDTOData) SetAttributes(v PostFlowValuesResponseDTODataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PostFlowValuesResponseDTOData) GetRelationships() PostFlowValuesResponseDTODataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PostFlowValuesResponseDTOData) GetRelationshipsOk() (*PostFlowValuesResponseDTODataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PostFlowValuesResponseDTOData) SetRelationships(v PostFlowValuesResponseDTODataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PostFlowValuesResponseDTOData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *PostFlowValuesResponseDTOData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostFlowValuesResponseDTOData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostFlowValuesResponseDTOData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



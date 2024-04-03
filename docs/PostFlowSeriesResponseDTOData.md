# PostFlowSeriesResponseDTOData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FlowSeriesReportEnum**](FlowSeriesReportEnum.md) |  | 
**Attributes** | [**PostFlowSeriesResponseDTODataAttributes**](PostFlowSeriesResponseDTODataAttributes.md) |  | 
**Relationships** | Pointer to [**PostFlowValuesResponseDTODataRelationships**](PostFlowValuesResponseDTODataRelationships.md) |  | [optional] 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostFlowSeriesResponseDTOData

`func NewPostFlowSeriesResponseDTOData(type_ FlowSeriesReportEnum, attributes PostFlowSeriesResponseDTODataAttributes, links ObjectLinks, ) *PostFlowSeriesResponseDTOData`

NewPostFlowSeriesResponseDTOData instantiates a new PostFlowSeriesResponseDTOData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostFlowSeriesResponseDTODataWithDefaults

`func NewPostFlowSeriesResponseDTODataWithDefaults() *PostFlowSeriesResponseDTOData`

NewPostFlowSeriesResponseDTODataWithDefaults instantiates a new PostFlowSeriesResponseDTOData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostFlowSeriesResponseDTOData) GetType() FlowSeriesReportEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostFlowSeriesResponseDTOData) GetTypeOk() (*FlowSeriesReportEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostFlowSeriesResponseDTOData) SetType(v FlowSeriesReportEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PostFlowSeriesResponseDTOData) GetAttributes() PostFlowSeriesResponseDTODataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostFlowSeriesResponseDTOData) GetAttributesOk() (*PostFlowSeriesResponseDTODataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostFlowSeriesResponseDTOData) SetAttributes(v PostFlowSeriesResponseDTODataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PostFlowSeriesResponseDTOData) GetRelationships() PostFlowValuesResponseDTODataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PostFlowSeriesResponseDTOData) GetRelationshipsOk() (*PostFlowValuesResponseDTODataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PostFlowSeriesResponseDTOData) SetRelationships(v PostFlowValuesResponseDTODataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PostFlowSeriesResponseDTOData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *PostFlowSeriesResponseDTOData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostFlowSeriesResponseDTOData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostFlowSeriesResponseDTOData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



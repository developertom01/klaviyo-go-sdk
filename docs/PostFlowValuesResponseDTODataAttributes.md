# PostFlowValuesResponseDTODataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]ValuesData**](ValuesData.md) | An array of all the returned values data. Each object in the array represents one unique grouping and the results for that grouping. | 

## Methods

### NewPostFlowValuesResponseDTODataAttributes

`func NewPostFlowValuesResponseDTODataAttributes(results []ValuesData, ) *PostFlowValuesResponseDTODataAttributes`

NewPostFlowValuesResponseDTODataAttributes instantiates a new PostFlowValuesResponseDTODataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostFlowValuesResponseDTODataAttributesWithDefaults

`func NewPostFlowValuesResponseDTODataAttributesWithDefaults() *PostFlowValuesResponseDTODataAttributes`

NewPostFlowValuesResponseDTODataAttributesWithDefaults instantiates a new PostFlowValuesResponseDTODataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PostFlowValuesResponseDTODataAttributes) GetResults() []ValuesData`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PostFlowValuesResponseDTODataAttributes) GetResultsOk() (*[]ValuesData, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PostFlowValuesResponseDTODataAttributes) SetResults(v []ValuesData)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



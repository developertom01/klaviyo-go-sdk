# ImportErrorResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A code for classifying the error type. | 
**Title** | **string** | A high-level message about the error. | 
**Detail** | **string** | Specific details about the error. | 
**Source** | [**ErrorSource**](ErrorSource.md) |  | 
**OriginalPayload** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewImportErrorResponseObjectResourceAttributes

`func NewImportErrorResponseObjectResourceAttributes(code string, title string, detail string, source ErrorSource, ) *ImportErrorResponseObjectResourceAttributes`

NewImportErrorResponseObjectResourceAttributes instantiates a new ImportErrorResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportErrorResponseObjectResourceAttributesWithDefaults

`func NewImportErrorResponseObjectResourceAttributesWithDefaults() *ImportErrorResponseObjectResourceAttributes`

NewImportErrorResponseObjectResourceAttributesWithDefaults instantiates a new ImportErrorResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ImportErrorResponseObjectResourceAttributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ImportErrorResponseObjectResourceAttributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ImportErrorResponseObjectResourceAttributes) SetCode(v string)`

SetCode sets Code field to given value.


### GetTitle

`func (o *ImportErrorResponseObjectResourceAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ImportErrorResponseObjectResourceAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ImportErrorResponseObjectResourceAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *ImportErrorResponseObjectResourceAttributes) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ImportErrorResponseObjectResourceAttributes) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ImportErrorResponseObjectResourceAttributes) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetSource

`func (o *ImportErrorResponseObjectResourceAttributes) GetSource() ErrorSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ImportErrorResponseObjectResourceAttributes) GetSourceOk() (*ErrorSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ImportErrorResponseObjectResourceAttributes) SetSource(v ErrorSource)`

SetSource sets Source field to given value.


### GetOriginalPayload

`func (o *ImportErrorResponseObjectResourceAttributes) GetOriginalPayload() map[string]interface{}`

GetOriginalPayload returns the OriginalPayload field if non-nil, zero value otherwise.

### GetOriginalPayloadOk

`func (o *ImportErrorResponseObjectResourceAttributes) GetOriginalPayloadOk() (*map[string]interface{}, bool)`

GetOriginalPayloadOk returns a tuple with the OriginalPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPayload

`func (o *ImportErrorResponseObjectResourceAttributes) SetOriginalPayload(v map[string]interface{})`

SetOriginalPayload sets OriginalPayload field to given value.

### HasOriginalPayload

`func (o *ImportErrorResponseObjectResourceAttributes) HasOriginalPayload() bool`

HasOriginalPayload returns a boolean if a field has been set.

### SetOriginalPayloadNil

`func (o *ImportErrorResponseObjectResourceAttributes) SetOriginalPayloadNil(b bool)`

 SetOriginalPayloadNil sets the value for OriginalPayload to be an explicit nil

### UnsetOriginalPayload
`func (o *ImportErrorResponseObjectResourceAttributes) UnsetOriginalPayload()`

UnsetOriginalPayload ensures that no value is present for OriginalPayload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



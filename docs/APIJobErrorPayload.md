# APIJobErrorPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the error. | 
**Code** | **string** | A code for classifying the error type. | 
**Title** | **string** | A high-level message about the error. | 
**Detail** | **string** | Specific details about the error. | 
**Source** | [**ErrorSource**](ErrorSource.md) |  | 

## Methods

### NewAPIJobErrorPayload

`func NewAPIJobErrorPayload(id string, code string, title string, detail string, source ErrorSource, ) *APIJobErrorPayload`

NewAPIJobErrorPayload instantiates a new APIJobErrorPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIJobErrorPayloadWithDefaults

`func NewAPIJobErrorPayloadWithDefaults() *APIJobErrorPayload`

NewAPIJobErrorPayloadWithDefaults instantiates a new APIJobErrorPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIJobErrorPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIJobErrorPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIJobErrorPayload) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *APIJobErrorPayload) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *APIJobErrorPayload) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *APIJobErrorPayload) SetCode(v string)`

SetCode sets Code field to given value.


### GetTitle

`func (o *APIJobErrorPayload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *APIJobErrorPayload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *APIJobErrorPayload) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *APIJobErrorPayload) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *APIJobErrorPayload) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *APIJobErrorPayload) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetSource

`func (o *APIJobErrorPayload) GetSource() ErrorSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *APIJobErrorPayload) GetSourceOk() (*ErrorSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *APIJobErrorPayload) SetSource(v ErrorSource)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



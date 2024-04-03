# ListListResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | A helpful name to label the list | [optional] 
**Created** | Pointer to **NullableTime** | Date and time when the list was created, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm) | [optional] 
**Updated** | Pointer to **NullableTime** | Date and time when the list was last updated, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm) | [optional] 
**OptInProcess** | Pointer to **NullableString** | The opt-in process for this list.  Could be either &#39;single_opt_in&#39; or &#39;double_opt_in&#39;. | [optional] 

## Methods

### NewListListResponseObjectResourceAttributes

`func NewListListResponseObjectResourceAttributes() *ListListResponseObjectResourceAttributes`

NewListListResponseObjectResourceAttributes instantiates a new ListListResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListListResponseObjectResourceAttributesWithDefaults

`func NewListListResponseObjectResourceAttributesWithDefaults() *ListListResponseObjectResourceAttributes`

NewListListResponseObjectResourceAttributesWithDefaults instantiates a new ListListResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListListResponseObjectResourceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListListResponseObjectResourceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListListResponseObjectResourceAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListListResponseObjectResourceAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ListListResponseObjectResourceAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ListListResponseObjectResourceAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreated

`func (o *ListListResponseObjectResourceAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListListResponseObjectResourceAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListListResponseObjectResourceAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListListResponseObjectResourceAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *ListListResponseObjectResourceAttributes) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ListListResponseObjectResourceAttributes) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *ListListResponseObjectResourceAttributes) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ListListResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ListListResponseObjectResourceAttributes) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ListListResponseObjectResourceAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *ListListResponseObjectResourceAttributes) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *ListListResponseObjectResourceAttributes) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetOptInProcess

`func (o *ListListResponseObjectResourceAttributes) GetOptInProcess() string`

GetOptInProcess returns the OptInProcess field if non-nil, zero value otherwise.

### GetOptInProcessOk

`func (o *ListListResponseObjectResourceAttributes) GetOptInProcessOk() (*string, bool)`

GetOptInProcessOk returns a tuple with the OptInProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInProcess

`func (o *ListListResponseObjectResourceAttributes) SetOptInProcess(v string)`

SetOptInProcess sets OptInProcess field to given value.

### HasOptInProcess

`func (o *ListListResponseObjectResourceAttributes) HasOptInProcess() bool`

HasOptInProcess returns a boolean if a field has been set.

### SetOptInProcessNil

`func (o *ListListResponseObjectResourceAttributes) SetOptInProcessNil(b bool)`

 SetOptInProcessNil sets the value for OptInProcess to be an explicit nil

### UnsetOptInProcess
`func (o *ListListResponseObjectResourceAttributes) UnsetOptInProcess()`

UnsetOptInProcess ensures that no value is present for OptInProcess, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



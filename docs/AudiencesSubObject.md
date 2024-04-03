# AudiencesSubObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to **[]string** | A list of included audiences | [optional] 
**Excluded** | Pointer to **[]string** | An optional list of excluded audiences | [optional] 

## Methods

### NewAudiencesSubObject

`func NewAudiencesSubObject() *AudiencesSubObject`

NewAudiencesSubObject instantiates a new AudiencesSubObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudiencesSubObjectWithDefaults

`func NewAudiencesSubObjectWithDefaults() *AudiencesSubObject`

NewAudiencesSubObjectWithDefaults instantiates a new AudiencesSubObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *AudiencesSubObject) GetIncluded() []string`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AudiencesSubObject) GetIncludedOk() (*[]string, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AudiencesSubObject) SetIncluded(v []string)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AudiencesSubObject) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### SetIncludedNil

`func (o *AudiencesSubObject) SetIncludedNil(b bool)`

 SetIncludedNil sets the value for Included to be an explicit nil

### UnsetIncluded
`func (o *AudiencesSubObject) UnsetIncluded()`

UnsetIncluded ensures that no value is present for Included, not even an explicit nil
### GetExcluded

`func (o *AudiencesSubObject) GetExcluded() []string`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *AudiencesSubObject) GetExcludedOk() (*[]string, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *AudiencesSubObject) SetExcluded(v []string)`

SetExcluded sets Excluded field to given value.

### HasExcluded

`func (o *AudiencesSubObject) HasExcluded() bool`

HasExcluded returns a boolean if a field has been set.

### SetExcludedNil

`func (o *AudiencesSubObject) SetExcludedNil(b bool)`

 SetExcludedNil sets the value for Excluded to be an explicit nil

### UnsetExcluded
`func (o *AudiencesSubObject) UnsetExcluded()`

UnsetExcluded ensures that no value is present for Excluded, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TagListOpDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ListEnum**](ListEnum.md) |  | 
**Id** | **string** | The IDs of the lists to link or unlink with the given Tag ID | 

## Methods

### NewTagListOpDataInner

`func NewTagListOpDataInner(type_ ListEnum, id string, ) *TagListOpDataInner`

NewTagListOpDataInner instantiates a new TagListOpDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagListOpDataInnerWithDefaults

`func NewTagListOpDataInnerWithDefaults() *TagListOpDataInner`

NewTagListOpDataInnerWithDefaults instantiates a new TagListOpDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TagListOpDataInner) GetType() ListEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagListOpDataInner) GetTypeOk() (*ListEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagListOpDataInner) SetType(v ListEnum)`

SetType sets Type field to given value.


### GetId

`func (o *TagListOpDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagListOpDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagListOpDataInner) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



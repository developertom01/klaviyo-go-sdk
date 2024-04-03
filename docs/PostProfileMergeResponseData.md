# PostProfileMergeResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProfileEnum**](ProfileEnum.md) |  | 
**Id** | **string** | The ID of the destination profile that was merged into | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostProfileMergeResponseData

`func NewPostProfileMergeResponseData(type_ ProfileEnum, id string, links ObjectLinks, ) *PostProfileMergeResponseData`

NewPostProfileMergeResponseData instantiates a new PostProfileMergeResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostProfileMergeResponseDataWithDefaults

`func NewPostProfileMergeResponseDataWithDefaults() *PostProfileMergeResponseData`

NewPostProfileMergeResponseDataWithDefaults instantiates a new PostProfileMergeResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostProfileMergeResponseData) GetType() ProfileEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostProfileMergeResponseData) GetTypeOk() (*ProfileEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostProfileMergeResponseData) SetType(v ProfileEnum)`

SetType sets Type field to given value.


### GetId

`func (o *PostProfileMergeResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostProfileMergeResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostProfileMergeResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *PostProfileMergeResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostProfileMergeResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostProfileMergeResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



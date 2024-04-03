# PostImageResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ImageEnum**](ImageEnum.md) |  | 
**Id** | **string** | The ID of the image | 
**Attributes** | [**ImageResponseObjectResourceAttributes**](ImageResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostImageResponseData

`func NewPostImageResponseData(type_ ImageEnum, id string, attributes ImageResponseObjectResourceAttributes, links ObjectLinks, ) *PostImageResponseData`

NewPostImageResponseData instantiates a new PostImageResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostImageResponseDataWithDefaults

`func NewPostImageResponseDataWithDefaults() *PostImageResponseData`

NewPostImageResponseDataWithDefaults instantiates a new PostImageResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostImageResponseData) GetType() ImageEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostImageResponseData) GetTypeOk() (*ImageEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostImageResponseData) SetType(v ImageEnum)`

SetType sets Type field to given value.


### GetId

`func (o *PostImageResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostImageResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostImageResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PostImageResponseData) GetAttributes() ImageResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostImageResponseData) GetAttributesOk() (*ImageResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostImageResponseData) SetAttributes(v ImageResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *PostImageResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostImageResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostImageResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



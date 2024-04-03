# ImagePartialUpdateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ImageEnum**](ImageEnum.md) |  | 
**Id** | **string** | The ID of the image | 
**Attributes** | [**ImagePartialUpdateQueryResourceObjectAttributes**](ImagePartialUpdateQueryResourceObjectAttributes.md) |  | 

## Methods

### NewImagePartialUpdateQueryResourceObject

`func NewImagePartialUpdateQueryResourceObject(type_ ImageEnum, id string, attributes ImagePartialUpdateQueryResourceObjectAttributes, ) *ImagePartialUpdateQueryResourceObject`

NewImagePartialUpdateQueryResourceObject instantiates a new ImagePartialUpdateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagePartialUpdateQueryResourceObjectWithDefaults

`func NewImagePartialUpdateQueryResourceObjectWithDefaults() *ImagePartialUpdateQueryResourceObject`

NewImagePartialUpdateQueryResourceObjectWithDefaults instantiates a new ImagePartialUpdateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ImagePartialUpdateQueryResourceObject) GetType() ImageEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImagePartialUpdateQueryResourceObject) GetTypeOk() (*ImageEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImagePartialUpdateQueryResourceObject) SetType(v ImageEnum)`

SetType sets Type field to given value.


### GetId

`func (o *ImagePartialUpdateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImagePartialUpdateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImagePartialUpdateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *ImagePartialUpdateQueryResourceObject) GetAttributes() ImagePartialUpdateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ImagePartialUpdateQueryResourceObject) GetAttributesOk() (*ImagePartialUpdateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ImagePartialUpdateQueryResourceObject) SetAttributes(v ImagePartialUpdateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



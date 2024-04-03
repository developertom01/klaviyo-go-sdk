# DataPrivacyProfileQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProfileEnum**](ProfileEnum.md) |  | 
**Id** | Pointer to **NullableString** | Primary key that uniquely identifies this profile. Generated by Klaviyo. | [optional] 
**Attributes** | [**DataPrivacyProfileQueryResourceObjectAttributes**](DataPrivacyProfileQueryResourceObjectAttributes.md) |  | 

## Methods

### NewDataPrivacyProfileQueryResourceObject

`func NewDataPrivacyProfileQueryResourceObject(type_ ProfileEnum, attributes DataPrivacyProfileQueryResourceObjectAttributes, ) *DataPrivacyProfileQueryResourceObject`

NewDataPrivacyProfileQueryResourceObject instantiates a new DataPrivacyProfileQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataPrivacyProfileQueryResourceObjectWithDefaults

`func NewDataPrivacyProfileQueryResourceObjectWithDefaults() *DataPrivacyProfileQueryResourceObject`

NewDataPrivacyProfileQueryResourceObjectWithDefaults instantiates a new DataPrivacyProfileQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DataPrivacyProfileQueryResourceObject) GetType() ProfileEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataPrivacyProfileQueryResourceObject) GetTypeOk() (*ProfileEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataPrivacyProfileQueryResourceObject) SetType(v ProfileEnum)`

SetType sets Type field to given value.


### GetId

`func (o *DataPrivacyProfileQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataPrivacyProfileQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataPrivacyProfileQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataPrivacyProfileQueryResourceObject) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DataPrivacyProfileQueryResourceObject) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DataPrivacyProfileQueryResourceObject) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *DataPrivacyProfileQueryResourceObject) GetAttributes() DataPrivacyProfileQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DataPrivacyProfileQueryResourceObject) GetAttributesOk() (*DataPrivacyProfileQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DataPrivacyProfileQueryResourceObject) SetAttributes(v DataPrivacyProfileQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


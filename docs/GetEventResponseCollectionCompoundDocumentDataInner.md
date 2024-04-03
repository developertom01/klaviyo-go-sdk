# GetEventResponseCollectionCompoundDocumentDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EventEnum**](EventEnum.md) |  | 
**Id** | **string** | The Event ID | 
**Attributes** | [**EventResponseObjectResourceAttributes**](EventResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships**](GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetEventResponseCollectionCompoundDocumentDataInner

`func NewGetEventResponseCollectionCompoundDocumentDataInner(type_ EventEnum, id string, attributes EventResponseObjectResourceAttributes, links ObjectLinks, ) *GetEventResponseCollectionCompoundDocumentDataInner`

NewGetEventResponseCollectionCompoundDocumentDataInner instantiates a new GetEventResponseCollectionCompoundDocumentDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventResponseCollectionCompoundDocumentDataInnerWithDefaults

`func NewGetEventResponseCollectionCompoundDocumentDataInnerWithDefaults() *GetEventResponseCollectionCompoundDocumentDataInner`

NewGetEventResponseCollectionCompoundDocumentDataInnerWithDefaults instantiates a new GetEventResponseCollectionCompoundDocumentDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) GetType() EventEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) GetTypeOk() (*EventEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) SetType(v EventEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) GetAttributes() EventResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) GetAttributesOk() (*EventResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) SetAttributes(v EventResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) GetRelationships() GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) GetRelationshipsOk() (*GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) SetRelationships(v GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetEventResponseCollectionCompoundDocumentDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



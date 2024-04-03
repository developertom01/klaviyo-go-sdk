# GetProfileResponseCompoundDocumentIncludedInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SegmentEnum**](SegmentEnum.md) |  | 
**Id** | **string** |  | 
**Attributes** | [**SegmentListResponseObjectResourceAttributes**](SegmentListResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewGetProfileResponseCompoundDocumentIncludedInner

`func NewGetProfileResponseCompoundDocumentIncludedInner(type_ SegmentEnum, id string, attributes SegmentListResponseObjectResourceAttributes, links ObjectLinks, ) *GetProfileResponseCompoundDocumentIncludedInner`

NewGetProfileResponseCompoundDocumentIncludedInner instantiates a new GetProfileResponseCompoundDocumentIncludedInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProfileResponseCompoundDocumentIncludedInnerWithDefaults

`func NewGetProfileResponseCompoundDocumentIncludedInnerWithDefaults() *GetProfileResponseCompoundDocumentIncludedInner`

NewGetProfileResponseCompoundDocumentIncludedInnerWithDefaults instantiates a new GetProfileResponseCompoundDocumentIncludedInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetProfileResponseCompoundDocumentIncludedInner) GetType() SegmentEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetProfileResponseCompoundDocumentIncludedInner) GetTypeOk() (*SegmentEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetProfileResponseCompoundDocumentIncludedInner) SetType(v SegmentEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetProfileResponseCompoundDocumentIncludedInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetProfileResponseCompoundDocumentIncludedInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetProfileResponseCompoundDocumentIncludedInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetProfileResponseCompoundDocumentIncludedInner) GetAttributes() SegmentListResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetProfileResponseCompoundDocumentIncludedInner) GetAttributesOk() (*SegmentListResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetProfileResponseCompoundDocumentIncludedInner) SetAttributes(v SegmentListResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetProfileResponseCompoundDocumentIncludedInner) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetProfileResponseCompoundDocumentIncludedInner) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetProfileResponseCompoundDocumentIncludedInner) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



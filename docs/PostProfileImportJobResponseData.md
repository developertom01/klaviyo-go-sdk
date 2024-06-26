# PostProfileImportJobResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProfileBulkImportJobEnum**](ProfileBulkImportJobEnum.md) |  | 
**Id** | **string** | Unique identifier for retrieving the job. Generated by Klaviyo. | 
**Attributes** | [**ProfileImportJobResponseObjectResourceAttributes**](ProfileImportJobResponseObjectResourceAttributes.md) |  | 
**Relationships** | Pointer to [**PostProfileImportJobResponseDataRelationships**](PostProfileImportJobResponseDataRelationships.md) |  | [optional] 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostProfileImportJobResponseData

`func NewPostProfileImportJobResponseData(type_ ProfileBulkImportJobEnum, id string, attributes ProfileImportJobResponseObjectResourceAttributes, links ObjectLinks, ) *PostProfileImportJobResponseData`

NewPostProfileImportJobResponseData instantiates a new PostProfileImportJobResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostProfileImportJobResponseDataWithDefaults

`func NewPostProfileImportJobResponseDataWithDefaults() *PostProfileImportJobResponseData`

NewPostProfileImportJobResponseDataWithDefaults instantiates a new PostProfileImportJobResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostProfileImportJobResponseData) GetType() ProfileBulkImportJobEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostProfileImportJobResponseData) GetTypeOk() (*ProfileBulkImportJobEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostProfileImportJobResponseData) SetType(v ProfileBulkImportJobEnum)`

SetType sets Type field to given value.


### GetId

`func (o *PostProfileImportJobResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostProfileImportJobResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostProfileImportJobResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PostProfileImportJobResponseData) GetAttributes() ProfileImportJobResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostProfileImportJobResponseData) GetAttributesOk() (*ProfileImportJobResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostProfileImportJobResponseData) SetAttributes(v ProfileImportJobResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PostProfileImportJobResponseData) GetRelationships() PostProfileImportJobResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PostProfileImportJobResponseData) GetRelationshipsOk() (*PostProfileImportJobResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PostProfileImportJobResponseData) SetRelationships(v PostProfileImportJobResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PostProfileImportJobResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *PostProfileImportJobResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostProfileImportJobResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostProfileImportJobResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



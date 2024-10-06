# ReservationExpand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Space** | Pointer to **NullableString** |  | [optional] 
**Block** | Pointer to **NullableString** |  | [optional] 
**Cidr** | **string** |  | 
**Desc** | Pointer to **NullableString** |  | [optional] 
**CreatedOn** | **float32** |  | 
**CreatedBy** | **string** |  | 
**SettledOn** | Pointer to **NullableFloat32** |  | [optional] 
**SettledBy** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**Tag** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewReservationExpand

`func NewReservationExpand(id string, cidr string, createdOn float32, createdBy string, status string, ) *ReservationExpand`

NewReservationExpand instantiates a new ReservationExpand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationExpandWithDefaults

`func NewReservationExpandWithDefaults() *ReservationExpand`

NewReservationExpandWithDefaults instantiates a new ReservationExpand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReservationExpand) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReservationExpand) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReservationExpand) SetId(v string)`

SetId sets Id field to given value.


### GetSpace

`func (o *ReservationExpand) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *ReservationExpand) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *ReservationExpand) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *ReservationExpand) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### SetSpaceNil

`func (o *ReservationExpand) SetSpaceNil(b bool)`

 SetSpaceNil sets the value for Space to be an explicit nil

### UnsetSpace
`func (o *ReservationExpand) UnsetSpace()`

UnsetSpace ensures that no value is present for Space, not even an explicit nil
### GetBlock

`func (o *ReservationExpand) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *ReservationExpand) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *ReservationExpand) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *ReservationExpand) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### SetBlockNil

`func (o *ReservationExpand) SetBlockNil(b bool)`

 SetBlockNil sets the value for Block to be an explicit nil

### UnsetBlock
`func (o *ReservationExpand) UnsetBlock()`

UnsetBlock ensures that no value is present for Block, not even an explicit nil
### GetCidr

`func (o *ReservationExpand) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ReservationExpand) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ReservationExpand) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetDesc

`func (o *ReservationExpand) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ReservationExpand) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ReservationExpand) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *ReservationExpand) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### SetDescNil

`func (o *ReservationExpand) SetDescNil(b bool)`

 SetDescNil sets the value for Desc to be an explicit nil

### UnsetDesc
`func (o *ReservationExpand) UnsetDesc()`

UnsetDesc ensures that no value is present for Desc, not even an explicit nil
### GetCreatedOn

`func (o *ReservationExpand) GetCreatedOn() float32`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ReservationExpand) GetCreatedOnOk() (*float32, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ReservationExpand) SetCreatedOn(v float32)`

SetCreatedOn sets CreatedOn field to given value.


### GetCreatedBy

`func (o *ReservationExpand) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ReservationExpand) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ReservationExpand) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetSettledOn

`func (o *ReservationExpand) GetSettledOn() float32`

GetSettledOn returns the SettledOn field if non-nil, zero value otherwise.

### GetSettledOnOk

`func (o *ReservationExpand) GetSettledOnOk() (*float32, bool)`

GetSettledOnOk returns a tuple with the SettledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledOn

`func (o *ReservationExpand) SetSettledOn(v float32)`

SetSettledOn sets SettledOn field to given value.

### HasSettledOn

`func (o *ReservationExpand) HasSettledOn() bool`

HasSettledOn returns a boolean if a field has been set.

### SetSettledOnNil

`func (o *ReservationExpand) SetSettledOnNil(b bool)`

 SetSettledOnNil sets the value for SettledOn to be an explicit nil

### UnsetSettledOn
`func (o *ReservationExpand) UnsetSettledOn()`

UnsetSettledOn ensures that no value is present for SettledOn, not even an explicit nil
### GetSettledBy

`func (o *ReservationExpand) GetSettledBy() string`

GetSettledBy returns the SettledBy field if non-nil, zero value otherwise.

### GetSettledByOk

`func (o *ReservationExpand) GetSettledByOk() (*string, bool)`

GetSettledByOk returns a tuple with the SettledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledBy

`func (o *ReservationExpand) SetSettledBy(v string)`

SetSettledBy sets SettledBy field to given value.

### HasSettledBy

`func (o *ReservationExpand) HasSettledBy() bool`

HasSettledBy returns a boolean if a field has been set.

### SetSettledByNil

`func (o *ReservationExpand) SetSettledByNil(b bool)`

 SetSettledByNil sets the value for SettledBy to be an explicit nil

### UnsetSettledBy
`func (o *ReservationExpand) UnsetSettledBy()`

UnsetSettledBy ensures that no value is present for SettledBy, not even an explicit nil
### GetStatus

`func (o *ReservationExpand) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReservationExpand) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReservationExpand) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTag

`func (o *ReservationExpand) GetTag() map[string]interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ReservationExpand) GetTagOk() (*map[string]interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ReservationExpand) SetTag(v map[string]interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *ReservationExpand) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *ReservationExpand) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *ReservationExpand) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



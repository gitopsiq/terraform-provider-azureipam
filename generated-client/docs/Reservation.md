# Reservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Cidr** | **string** |  | 
**Desc** | Pointer to **NullableString** |  | [optional] 
**CreatedOn** | **float32** |  | 
**CreatedBy** | **string** |  | 
**SettledOn** | Pointer to **NullableFloat32** |  | [optional] 
**SettledBy** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewReservation

`func NewReservation(id string, cidr string, createdOn float32, createdBy string, status string, ) *Reservation`

NewReservation instantiates a new Reservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationWithDefaults

`func NewReservationWithDefaults() *Reservation`

NewReservationWithDefaults instantiates a new Reservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Reservation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Reservation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Reservation) SetId(v string)`

SetId sets Id field to given value.


### GetCidr

`func (o *Reservation) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *Reservation) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *Reservation) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetDesc

`func (o *Reservation) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Reservation) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Reservation) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Reservation) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### SetDescNil

`func (o *Reservation) SetDescNil(b bool)`

 SetDescNil sets the value for Desc to be an explicit nil

### UnsetDesc
`func (o *Reservation) UnsetDesc()`

UnsetDesc ensures that no value is present for Desc, not even an explicit nil
### GetCreatedOn

`func (o *Reservation) GetCreatedOn() float32`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Reservation) GetCreatedOnOk() (*float32, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Reservation) SetCreatedOn(v float32)`

SetCreatedOn sets CreatedOn field to given value.


### GetCreatedBy

`func (o *Reservation) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Reservation) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Reservation) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetSettledOn

`func (o *Reservation) GetSettledOn() float32`

GetSettledOn returns the SettledOn field if non-nil, zero value otherwise.

### GetSettledOnOk

`func (o *Reservation) GetSettledOnOk() (*float32, bool)`

GetSettledOnOk returns a tuple with the SettledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledOn

`func (o *Reservation) SetSettledOn(v float32)`

SetSettledOn sets SettledOn field to given value.

### HasSettledOn

`func (o *Reservation) HasSettledOn() bool`

HasSettledOn returns a boolean if a field has been set.

### SetSettledOnNil

`func (o *Reservation) SetSettledOnNil(b bool)`

 SetSettledOnNil sets the value for SettledOn to be an explicit nil

### UnsetSettledOn
`func (o *Reservation) UnsetSettledOn()`

UnsetSettledOn ensures that no value is present for SettledOn, not even an explicit nil
### GetSettledBy

`func (o *Reservation) GetSettledBy() string`

GetSettledBy returns the SettledBy field if non-nil, zero value otherwise.

### GetSettledByOk

`func (o *Reservation) GetSettledByOk() (*string, bool)`

GetSettledByOk returns a tuple with the SettledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledBy

`func (o *Reservation) SetSettledBy(v string)`

SetSettledBy sets SettledBy field to given value.

### HasSettledBy

`func (o *Reservation) HasSettledBy() bool`

HasSettledBy returns a boolean if a field has been set.

### SetSettledByNil

`func (o *Reservation) SetSettledByNil(b bool)`

 SetSettledByNil sets the value for SettledBy to be an explicit nil

### UnsetSettledBy
`func (o *Reservation) UnsetSettledBy()`

UnsetSettledBy ensures that no value is present for SettledBy, not even an explicit nil
### GetStatus

`func (o *Reservation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Reservation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Reservation) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



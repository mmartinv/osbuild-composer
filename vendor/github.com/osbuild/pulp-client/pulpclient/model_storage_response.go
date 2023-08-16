/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"encoding/json"
)

// checks if the StorageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageResponse{}

// StorageResponse Serializer for information about the storage system
type StorageResponse struct {
	// Total number of bytes
	Total int64 `json:"total"`
	// Number of bytes in use
	Used int64 `json:"used"`
	// Number of free bytes
	Free int64 `json:"free"`
	AdditionalProperties map[string]interface{}
}

type _StorageResponse StorageResponse

// NewStorageResponse instantiates a new StorageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageResponse(total int64, used int64, free int64) *StorageResponse {
	this := StorageResponse{}
	this.Total = total
	this.Used = used
	this.Free = free
	return &this
}

// NewStorageResponseWithDefaults instantiates a new StorageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageResponseWithDefaults() *StorageResponse {
	this := StorageResponse{}
	return &this
}

// GetTotal returns the Total field value
func (o *StorageResponse) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *StorageResponse) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *StorageResponse) SetTotal(v int64) {
	o.Total = v
}

// GetUsed returns the Used field value
func (o *StorageResponse) GetUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *StorageResponse) GetUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value
func (o *StorageResponse) SetUsed(v int64) {
	o.Used = v
}

// GetFree returns the Free field value
func (o *StorageResponse) GetFree() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Free
}

// GetFreeOk returns a tuple with the Free field value
// and a boolean to check if the value has been set.
func (o *StorageResponse) GetFreeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Free, true
}

// SetFree sets field value
func (o *StorageResponse) SetFree(v int64) {
	o.Free = v
}

func (o StorageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["used"] = o.Used
	toSerialize["free"] = o.Free

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageResponse) UnmarshalJSON(bytes []byte) (err error) {
	varStorageResponse := _StorageResponse{}

	if err = json.Unmarshal(bytes, &varStorageResponse); err == nil {
		*o = StorageResponse(varStorageResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "used")
		delete(additionalProperties, "free")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageResponse struct {
	value *StorageResponse
	isSet bool
}

func (v NullableStorageResponse) Get() *StorageResponse {
	return v.value
}

func (v *NullableStorageResponse) Set(val *StorageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageResponse(val *StorageResponse) *NullableStorageResponse {
	return &NullableStorageResponse{value: val, isSet: true}
}

func (v NullableStorageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



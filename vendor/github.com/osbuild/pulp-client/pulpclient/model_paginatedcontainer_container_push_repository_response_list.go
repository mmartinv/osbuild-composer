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

// checks if the PaginatedcontainerContainerPushRepositoryResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedcontainerContainerPushRepositoryResponseList{}

// PaginatedcontainerContainerPushRepositoryResponseList struct for PaginatedcontainerContainerPushRepositoryResponseList
type PaginatedcontainerContainerPushRepositoryResponseList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []ContainerContainerPushRepositoryResponse `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedcontainerContainerPushRepositoryResponseList PaginatedcontainerContainerPushRepositoryResponseList

// NewPaginatedcontainerContainerPushRepositoryResponseList instantiates a new PaginatedcontainerContainerPushRepositoryResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedcontainerContainerPushRepositoryResponseList() *PaginatedcontainerContainerPushRepositoryResponseList {
	this := PaginatedcontainerContainerPushRepositoryResponseList{}
	return &this
}

// NewPaginatedcontainerContainerPushRepositoryResponseListWithDefaults instantiates a new PaginatedcontainerContainerPushRepositoryResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedcontainerContainerPushRepositoryResponseListWithDefaults() *PaginatedcontainerContainerPushRepositoryResponseList {
	this := PaginatedcontainerContainerPushRepositoryResponseList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedcontainerContainerPushRepositoryResponseList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedcontainerContainerPushRepositoryResponseList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedcontainerContainerPushRepositoryResponseList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedcontainerContainerPushRepositoryResponseList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedcontainerContainerPushRepositoryResponseList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedcontainerContainerPushRepositoryResponseList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedcontainerContainerPushRepositoryResponseList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedcontainerContainerPushRepositoryResponseList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedcontainerContainerPushRepositoryResponseList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedcontainerContainerPushRepositoryResponseList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedcontainerContainerPushRepositoryResponseList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedcontainerContainerPushRepositoryResponseList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedcontainerContainerPushRepositoryResponseList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedcontainerContainerPushRepositoryResponseList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedcontainerContainerPushRepositoryResponseList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedcontainerContainerPushRepositoryResponseList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedcontainerContainerPushRepositoryResponseList) GetResults() []ContainerContainerPushRepositoryResponse {
	if o == nil || IsNil(o.Results) {
		var ret []ContainerContainerPushRepositoryResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedcontainerContainerPushRepositoryResponseList) GetResultsOk() ([]ContainerContainerPushRepositoryResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedcontainerContainerPushRepositoryResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ContainerContainerPushRepositoryResponse and assigns it to the Results field.
func (o *PaginatedcontainerContainerPushRepositoryResponseList) SetResults(v []ContainerContainerPushRepositoryResponse) {
	o.Results = v
}

func (o PaginatedcontainerContainerPushRepositoryResponseList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedcontainerContainerPushRepositoryResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedcontainerContainerPushRepositoryResponseList) UnmarshalJSON(bytes []byte) (err error) {
	varPaginatedcontainerContainerPushRepositoryResponseList := _PaginatedcontainerContainerPushRepositoryResponseList{}

	if err = json.Unmarshal(bytes, &varPaginatedcontainerContainerPushRepositoryResponseList); err == nil {
		*o = PaginatedcontainerContainerPushRepositoryResponseList(varPaginatedcontainerContainerPushRepositoryResponseList)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "next")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedcontainerContainerPushRepositoryResponseList struct {
	value *PaginatedcontainerContainerPushRepositoryResponseList
	isSet bool
}

func (v NullablePaginatedcontainerContainerPushRepositoryResponseList) Get() *PaginatedcontainerContainerPushRepositoryResponseList {
	return v.value
}

func (v *NullablePaginatedcontainerContainerPushRepositoryResponseList) Set(val *PaginatedcontainerContainerPushRepositoryResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedcontainerContainerPushRepositoryResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedcontainerContainerPushRepositoryResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedcontainerContainerPushRepositoryResponseList(val *PaginatedcontainerContainerPushRepositoryResponseList) *NullablePaginatedcontainerContainerPushRepositoryResponseList {
	return &NullablePaginatedcontainerContainerPushRepositoryResponseList{value: val, isSet: true}
}

func (v NullablePaginatedcontainerContainerPushRepositoryResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedcontainerContainerPushRepositoryResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



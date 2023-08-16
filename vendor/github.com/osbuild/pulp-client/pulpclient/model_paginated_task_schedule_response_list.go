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

// checks if the PaginatedTaskScheduleResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedTaskScheduleResponseList{}

// PaginatedTaskScheduleResponseList struct for PaginatedTaskScheduleResponseList
type PaginatedTaskScheduleResponseList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []TaskScheduleResponse `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedTaskScheduleResponseList PaginatedTaskScheduleResponseList

// NewPaginatedTaskScheduleResponseList instantiates a new PaginatedTaskScheduleResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedTaskScheduleResponseList() *PaginatedTaskScheduleResponseList {
	this := PaginatedTaskScheduleResponseList{}
	return &this
}

// NewPaginatedTaskScheduleResponseListWithDefaults instantiates a new PaginatedTaskScheduleResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedTaskScheduleResponseListWithDefaults() *PaginatedTaskScheduleResponseList {
	this := PaginatedTaskScheduleResponseList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedTaskScheduleResponseList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTaskScheduleResponseList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedTaskScheduleResponseList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedTaskScheduleResponseList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedTaskScheduleResponseList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedTaskScheduleResponseList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedTaskScheduleResponseList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedTaskScheduleResponseList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedTaskScheduleResponseList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedTaskScheduleResponseList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedTaskScheduleResponseList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedTaskScheduleResponseList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedTaskScheduleResponseList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedTaskScheduleResponseList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedTaskScheduleResponseList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedTaskScheduleResponseList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedTaskScheduleResponseList) GetResults() []TaskScheduleResponse {
	if o == nil || IsNil(o.Results) {
		var ret []TaskScheduleResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTaskScheduleResponseList) GetResultsOk() ([]TaskScheduleResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedTaskScheduleResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []TaskScheduleResponse and assigns it to the Results field.
func (o *PaginatedTaskScheduleResponseList) SetResults(v []TaskScheduleResponse) {
	o.Results = v
}

func (o PaginatedTaskScheduleResponseList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedTaskScheduleResponseList) ToMap() (map[string]interface{}, error) {
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

func (o *PaginatedTaskScheduleResponseList) UnmarshalJSON(bytes []byte) (err error) {
	varPaginatedTaskScheduleResponseList := _PaginatedTaskScheduleResponseList{}

	if err = json.Unmarshal(bytes, &varPaginatedTaskScheduleResponseList); err == nil {
		*o = PaginatedTaskScheduleResponseList(varPaginatedTaskScheduleResponseList)
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

type NullablePaginatedTaskScheduleResponseList struct {
	value *PaginatedTaskScheduleResponseList
	isSet bool
}

func (v NullablePaginatedTaskScheduleResponseList) Get() *PaginatedTaskScheduleResponseList {
	return v.value
}

func (v *NullablePaginatedTaskScheduleResponseList) Set(val *PaginatedTaskScheduleResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedTaskScheduleResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedTaskScheduleResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedTaskScheduleResponseList(val *PaginatedTaskScheduleResponseList) *NullablePaginatedTaskScheduleResponseList {
	return &NullablePaginatedTaskScheduleResponseList{value: val, isSet: true}
}

func (v NullablePaginatedTaskScheduleResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedTaskScheduleResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



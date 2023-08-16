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

// checks if the TagCopy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagCopy{}

// TagCopy Serializer for copying tags from a source repository to a destination repository.
type TagCopy struct {
	// A URI of the repository to copy content from.
	SourceRepository *string `json:"source_repository,omitempty"`
	// A URI of the repository version to copy content from.
	SourceRepositoryVersion *string `json:"source_repository_version,omitempty"`
	// A list of tag names to copy.
	Names []interface{} `json:"names,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TagCopy TagCopy

// NewTagCopy instantiates a new TagCopy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagCopy() *TagCopy {
	this := TagCopy{}
	return &this
}

// NewTagCopyWithDefaults instantiates a new TagCopy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagCopyWithDefaults() *TagCopy {
	this := TagCopy{}
	return &this
}

// GetSourceRepository returns the SourceRepository field value if set, zero value otherwise.
func (o *TagCopy) GetSourceRepository() string {
	if o == nil || IsNil(o.SourceRepository) {
		var ret string
		return ret
	}
	return *o.SourceRepository
}

// GetSourceRepositoryOk returns a tuple with the SourceRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagCopy) GetSourceRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.SourceRepository) {
		return nil, false
	}
	return o.SourceRepository, true
}

// HasSourceRepository returns a boolean if a field has been set.
func (o *TagCopy) HasSourceRepository() bool {
	if o != nil && !IsNil(o.SourceRepository) {
		return true
	}

	return false
}

// SetSourceRepository gets a reference to the given string and assigns it to the SourceRepository field.
func (o *TagCopy) SetSourceRepository(v string) {
	o.SourceRepository = &v
}

// GetSourceRepositoryVersion returns the SourceRepositoryVersion field value if set, zero value otherwise.
func (o *TagCopy) GetSourceRepositoryVersion() string {
	if o == nil || IsNil(o.SourceRepositoryVersion) {
		var ret string
		return ret
	}
	return *o.SourceRepositoryVersion
}

// GetSourceRepositoryVersionOk returns a tuple with the SourceRepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagCopy) GetSourceRepositoryVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SourceRepositoryVersion) {
		return nil, false
	}
	return o.SourceRepositoryVersion, true
}

// HasSourceRepositoryVersion returns a boolean if a field has been set.
func (o *TagCopy) HasSourceRepositoryVersion() bool {
	if o != nil && !IsNil(o.SourceRepositoryVersion) {
		return true
	}

	return false
}

// SetSourceRepositoryVersion gets a reference to the given string and assigns it to the SourceRepositoryVersion field.
func (o *TagCopy) SetSourceRepositoryVersion(v string) {
	o.SourceRepositoryVersion = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *TagCopy) GetNames() []interface{} {
	if o == nil || IsNil(o.Names) {
		var ret []interface{}
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagCopy) GetNamesOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *TagCopy) HasNames() bool {
	if o != nil && !IsNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given []interface{} and assigns it to the Names field.
func (o *TagCopy) SetNames(v []interface{}) {
	o.Names = v
}

func (o TagCopy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagCopy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceRepository) {
		toSerialize["source_repository"] = o.SourceRepository
	}
	if !IsNil(o.SourceRepositoryVersion) {
		toSerialize["source_repository_version"] = o.SourceRepositoryVersion
	}
	if !IsNil(o.Names) {
		toSerialize["names"] = o.Names
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TagCopy) UnmarshalJSON(bytes []byte) (err error) {
	varTagCopy := _TagCopy{}

	if err = json.Unmarshal(bytes, &varTagCopy); err == nil {
		*o = TagCopy(varTagCopy)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "source_repository")
		delete(additionalProperties, "source_repository_version")
		delete(additionalProperties, "names")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTagCopy struct {
	value *TagCopy
	isSet bool
}

func (v NullableTagCopy) Get() *TagCopy {
	return v.value
}

func (v *NullableTagCopy) Set(val *TagCopy) {
	v.value = val
	v.isSet = true
}

func (v NullableTagCopy) IsSet() bool {
	return v.isSet
}

func (v *NullableTagCopy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagCopy(val *TagCopy) *NullableTagCopy {
	return &NullableTagCopy{value: val, isSet: true}
}

func (v NullableTagCopy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagCopy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



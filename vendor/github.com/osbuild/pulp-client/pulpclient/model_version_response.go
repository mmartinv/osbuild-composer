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

// checks if the VersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionResponse{}

// VersionResponse Serializer for the version information of Pulp components
type VersionResponse struct {
	// Name of a versioned component of Pulp
	Component string `json:"component"`
	// Version of the component (e.g. 3.0.0)
	Version string `json:"version"`
	// Python package name providing the component
	Package string `json:"package"`
	// Domain feature compatibility of component
	DomainCompatible bool `json:"domain_compatible"`
	AdditionalProperties map[string]interface{}
}

type _VersionResponse VersionResponse

// NewVersionResponse instantiates a new VersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionResponse(component string, version string, package_ string, domainCompatible bool) *VersionResponse {
	this := VersionResponse{}
	this.Component = component
	this.Version = version
	this.Package = package_
	this.DomainCompatible = domainCompatible
	return &this
}

// NewVersionResponseWithDefaults instantiates a new VersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionResponseWithDefaults() *VersionResponse {
	this := VersionResponse{}
	return &this
}

// GetComponent returns the Component field value
func (o *VersionResponse) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *VersionResponse) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *VersionResponse) SetComponent(v string) {
	o.Component = v
}

// GetVersion returns the Version field value
func (o *VersionResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VersionResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VersionResponse) SetVersion(v string) {
	o.Version = v
}

// GetPackage returns the Package field value
func (o *VersionResponse) GetPackage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Package
}

// GetPackageOk returns a tuple with the Package field value
// and a boolean to check if the value has been set.
func (o *VersionResponse) GetPackageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Package, true
}

// SetPackage sets field value
func (o *VersionResponse) SetPackage(v string) {
	o.Package = v
}

// GetDomainCompatible returns the DomainCompatible field value
func (o *VersionResponse) GetDomainCompatible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DomainCompatible
}

// GetDomainCompatibleOk returns a tuple with the DomainCompatible field value
// and a boolean to check if the value has been set.
func (o *VersionResponse) GetDomainCompatibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainCompatible, true
}

// SetDomainCompatible sets field value
func (o *VersionResponse) SetDomainCompatible(v bool) {
	o.DomainCompatible = v
}

func (o VersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["component"] = o.Component
	toSerialize["version"] = o.Version
	toSerialize["package"] = o.Package
	toSerialize["domain_compatible"] = o.DomainCompatible

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VersionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varVersionResponse := _VersionResponse{}

	if err = json.Unmarshal(bytes, &varVersionResponse); err == nil {
		*o = VersionResponse(varVersionResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "component")
		delete(additionalProperties, "version")
		delete(additionalProperties, "package")
		delete(additionalProperties, "domain_compatible")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVersionResponse struct {
	value *VersionResponse
	isSet bool
}

func (v NullableVersionResponse) Get() *VersionResponse {
	return v.value
}

func (v *NullableVersionResponse) Set(val *VersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionResponse(val *VersionResponse) *NullableVersionResponse {
	return &NullableVersionResponse{value: val, isSet: true}
}

func (v NullableVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



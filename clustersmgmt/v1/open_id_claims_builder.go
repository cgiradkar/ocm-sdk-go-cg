/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/renan-campos/ocm-sdk-go/clustersmgmt/v1

// OpenIDClaimsBuilder contains the data and logic needed to build 'open_ID_claims' objects.
//
// _OpenID_ identity provider claims.
type OpenIDClaimsBuilder struct {
	bitmap_           uint32
	email             []string
	groups            []string
	name              []string
	preferredUsername []string
}

// NewOpenIDClaims creates a new builder of 'open_ID_claims' objects.
func NewOpenIDClaims() *OpenIDClaimsBuilder {
	return &OpenIDClaimsBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *OpenIDClaimsBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// Email sets the value of the 'email' attribute to the given values.
//
//
func (b *OpenIDClaimsBuilder) Email(values ...string) *OpenIDClaimsBuilder {
	b.email = make([]string, len(values))
	copy(b.email, values)
	b.bitmap_ |= 1
	return b
}

// Groups sets the value of the 'groups' attribute to the given values.
//
//
func (b *OpenIDClaimsBuilder) Groups(values ...string) *OpenIDClaimsBuilder {
	b.groups = make([]string, len(values))
	copy(b.groups, values)
	b.bitmap_ |= 2
	return b
}

// Name sets the value of the 'name' attribute to the given values.
//
//
func (b *OpenIDClaimsBuilder) Name(values ...string) *OpenIDClaimsBuilder {
	b.name = make([]string, len(values))
	copy(b.name, values)
	b.bitmap_ |= 4
	return b
}

// PreferredUsername sets the value of the 'preferred_username' attribute to the given values.
//
//
func (b *OpenIDClaimsBuilder) PreferredUsername(values ...string) *OpenIDClaimsBuilder {
	b.preferredUsername = make([]string, len(values))
	copy(b.preferredUsername, values)
	b.bitmap_ |= 8
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *OpenIDClaimsBuilder) Copy(object *OpenIDClaims) *OpenIDClaimsBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	if object.email != nil {
		b.email = make([]string, len(object.email))
		copy(b.email, object.email)
	} else {
		b.email = nil
	}
	if object.groups != nil {
		b.groups = make([]string, len(object.groups))
		copy(b.groups, object.groups)
	} else {
		b.groups = nil
	}
	if object.name != nil {
		b.name = make([]string, len(object.name))
		copy(b.name, object.name)
	} else {
		b.name = nil
	}
	if object.preferredUsername != nil {
		b.preferredUsername = make([]string, len(object.preferredUsername))
		copy(b.preferredUsername, object.preferredUsername)
	} else {
		b.preferredUsername = nil
	}
	return b
}

// Build creates a 'open_ID_claims' object using the configuration stored in the builder.
func (b *OpenIDClaimsBuilder) Build() (object *OpenIDClaims, err error) {
	object = new(OpenIDClaims)
	object.bitmap_ = b.bitmap_
	if b.email != nil {
		object.email = make([]string, len(b.email))
		copy(object.email, b.email)
	}
	if b.groups != nil {
		object.groups = make([]string, len(b.groups))
		copy(object.groups, b.groups)
	}
	if b.name != nil {
		object.name = make([]string, len(b.name))
		copy(object.name, b.name)
	}
	if b.preferredUsername != nil {
		object.preferredUsername = make([]string, len(b.preferredUsername))
		copy(object.preferredUsername, b.preferredUsername)
	}
	return
}

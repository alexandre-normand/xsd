// Copyright 2020 The xgen Authors. All rights reserved.
//
// DO NOT EDIT: generated by xgen XSD generator
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package schema

import (
	"encoding/xml"
)

// CTSchema ...
type CTSchema struct {
	XMLName              xml.Name `xml:"CT_Schema"`
	UriAttr              string   `xml:"uri,attr,omitempty"`
	ManifestLocationAttr string   `xml:"manifestLocation,attr,omitempty"`
	SchemaLocationAttr   string   `xml:"schemaLocation,attr,omitempty"`
	SchemaLanguageAttr   string   `xml:"schemaLanguage,attr,omitempty"`
}

// CTSchemaLibrary ...
type CTSchemaLibrary struct {
	XMLName xml.Name    `xml:"CT_SchemaLibrary"`
	Schema  []*CTSchema `xml:"schema"`
}

// SchemaLibrary ...
type SchemaLibrary *CTSchemaLibrary

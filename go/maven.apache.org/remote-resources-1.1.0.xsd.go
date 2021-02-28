// Code generated by xgen. DO NOT EDIT.

package schema

import (
	"encoding/xml"
)

// RemoteResourcesBundle is Root element of the remote-resources.xml file.
type RemoteResourcesBundle *RemoteResourcesBundle

// RemoteResources ...
type RemoteResources struct {
	XMLName        xml.Name `xml:"remoteResources"`
	RemoteResource []string `xml:"remoteResource"`
}

// RemoteResourcesBundle is Root element of the remote-resources.xml file.
type RemoteResourcesBundle struct {
	RemoteResources *RemoteResources `xml:"remoteResources"`
	SourceEncoding  string           `xml:"sourceEncoding"`
}
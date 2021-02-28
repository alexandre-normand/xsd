// Code generated by xgen. DO NOT EDIT.

package schema

import (
	"encoding/xml"
)

// Archetypecatalog is 0.0.0+
type Archetypecatalog *ArchetypeCatalog

// Archetypes ...
type Archetypes struct {
	XMLName   xml.Name     `xml:"archetypes"`
	Archetype []*Archetype `xml:"archetype"`
}

// ArchetypeCatalog is 0.0.0+
type ArchetypeCatalog struct {
	Archetypes *Archetypes `xml:"archetypes"`
}

// Archetype is The description of the archetype.
type Archetype struct {
	GroupId     string `xml:"groupId"`
	ArtifactId  string `xml:"artifactId"`
	Version     string `xml:"version"`
	Repository  string `xml:"repository"`
	Description string `xml:"description"`
}
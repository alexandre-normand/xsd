// Code generated by xgen. DO NOT EDIT.

package schema

import (
	"encoding/xml"
	"time"
)

// Project ...
type Project struct {
	XMLName               xml.Name      `xml:"project"`
	Extend                string        `xml:"extend"`
	PomVersion            string        `xml:"pomVersion"`
	Id                    string        `xml:"id"`
	Name                  string        `xml:"name"`
	GroupId               string        `xml:"groupId"`
	CurrentVersion        string        `xml:"currentVersion"`
	Organization          *Organization `xml:"organization"`
	InceptionYear         time.Time     `xml:"inceptionYear"`
	Package               string        `xml:"package"`
	Logo                  string        `xml:"logo"`
	GumpRepositoryId      string        `xml:"gumpRepositoryId"`
	Description           string        `xml:"description"`
	ShortDescription      string        `xml:"shortDescription"`
	Url                   string        `xml:"url"`
	IssueTrackingUrl      string        `xml:"issueTrackingUrl"`
	SiteAddress           string        `xml:"siteAddress"`
	SiteDirectory         string        `xml:"siteDirectory"`
	DistributionSite      string        `xml:"distributionSite"`
	DistributionDirectory string        `xml:"distributionDirectory"`
	Repository            *Repository   `xml:"repository"`
	Versions              *Versions     `xml:"versions"`
	Branches              *Branches     `xml:"branches"`
	MailingLists          *MailingLists `xml:"mailingLists"`
	Developers            *Developers   `xml:"developers"`
	Contributors          *Contributors `xml:"contributors"`
	Licenses              *Licenses     `xml:"licenses"`
	Dependencies          *Dependencies `xml:"dependencies"`
	Build                 *Build        `xml:"build"`
	Reports               *Reports      `xml:"reports"`
	Properties            *Properties   `xml:"properties"`
}

// Extend ...
type Extend string

// Connection ...
type Connection string

// DeveloperConnection ...
type DeveloperConnection string

// CurrentVersion ...
type CurrentVersion string

// Description ...
type Description string

// DistributionSite ...
type DistributionSite string

// DistributionDirectory ...
type DistributionDirectory string

// Name ...
type Name string

// GroupId ...
type GroupId string

// ArtifactId ...
type ArtifactId string

// GumpRepositoryId ...
type GumpRepositoryId string

// Id ...
type Id string

// InceptionYear ...
type InceptionYear time.Time

// IssueTrackingUrl ...
type IssueTrackingUrl string

// Logo ...
type Logo string

// Package ...
type Package string

// PomVersion ...
type PomVersion string

// ShortDescription ...
type ShortDescription string

// SiteAddress ...
type SiteAddress string

// SiteDirectory ...
type SiteDirectory string

// Url ...
type Url string

// Repository ...
type Repository struct {
	XMLName             xml.Name `xml:"repository"`
	Connection          string   `xml:"connection"`
	DeveloperConnection string   `xml:"developerConnection"`
	Url                 string   `xml:"url"`
}

// Organization ...
type Organization struct {
	XMLName xml.Name `xml:"organization"`
	Name    string   `xml:"name"`
	Url     string   `xml:"url"`
	Logo    string   `xml:"logo"`
}

// Versions ...
type Versions struct {
	XMLName xml.Name   `xml:"versions"`
	Version []*Version `xml:"version"`
}

// Version ...
type Version struct {
	XMLName xml.Name `xml:"version"`
	Id      string   `xml:"id"`
	Name    string   `xml:"name"`
	Tag     string   `xml:"tag"`
}

// Tag ...
type Tag string

// Branches ...
type Branches struct {
	XMLName xml.Name  `xml:"branches"`
	Branch  []*Branch `xml:"branch"`
}

// Branch ...
type Branch struct {
	XMLName xml.Name `xml:"branch"`
	Tag     string   `xml:"tag"`
}

// MailingLists ...
type MailingLists struct {
	XMLName     xml.Name       `xml:"mailingLists"`
	MailingList []*MailingList `xml:"mailingList"`
}

// MailingList ...
type MailingList struct {
	XMLName     xml.Name `xml:"mailingList"`
	Name        string   `xml:"name"`
	Subscribe   string   `xml:"subscribe"`
	Unsubscribe string   `xml:"unsubscribe"`
	Archive     string   `xml:"archive"`
}

// Subscribe ...
type Subscribe string

// Unsubscribe ...
type Unsubscribe string

// Archive ...
type Archive string

// Developers ...
type Developers struct {
	XMLName   xml.Name     `xml:"developers"`
	Developer []*Developer `xml:"developer"`
}

// Developer ...
type Developer struct {
	XMLName        xml.Name `xml:"developer"`
	ContactDetails *ContactDetails
	Name           string `xml:"name"`
	Id             string `xml:"id"`
}

// ContactDetails ...
type ContactDetails struct {
	XMLName      xml.Name `xml:"contactDetails"`
	Email        string
	Organization string
	Roles        *Roles
	Url          string
	Timezone     float64
}

// Email ...
type Email string

// Roles ...
type Roles struct {
	XMLName xml.Name `xml:"roles"`
	Role    []string `xml:"role"`
}

// Role ...
type Role string

// TimezoneType ...
type TimezoneType float64

// Timezone ...
type Timezone float64

// Contributors ...
type Contributors struct {
	XMLName     xml.Name       `xml:"contributors"`
	Contributor []*Contributor `xml:"contributor"`
}

// Contributor ...
type Contributor struct {
	XMLName        xml.Name `xml:"contributor"`
	ContactDetails *ContactDetails
	Name           string `xml:"name"`
}

// Licenses ...
type Licenses struct {
	XMLName xml.Name   `xml:"licenses"`
	License []*License `xml:"license"`
}

// License ...
type License struct {
	XMLName      xml.Name `xml:"license"`
	Name         string   `xml:"name"`
	Url          string   `xml:"url"`
	Distribution string   `xml:"distribution"`
}

// Distribution ...
type Distribution string

// Dependencies ...
type Dependencies struct {
	XMLName    xml.Name      `xml:"dependencies"`
	Dependency []*Dependency `xml:"dependency"`
}

// Dependency ...
type Dependency struct {
	XMLName    xml.Name    `xml:"dependency"`
	Id         string      `xml:"id"`
	GroupId    string      `xml:"groupId"`
	ArtifactId string      `xml:"artifactId"`
	Version    string      `xml:"version"`
	Jar        string      `xml:"jar"`
	Type       string      `xml:"type"`
	Url        string      `xml:"url"`
	Properties *Properties `xml:"properties"`
}

// Type ...
type Type string

// Jar ...
type Jar string

// Build ...
type Build struct {
	XMLName                            xml.Name             `xml:"build"`
	NagEmailAddress                    string               `xml:"nagEmailAddress"`
	SourceDirectory                    string               `xml:"sourceDirectory"`
	SourceModifications                *SourceModifications `xml:"sourceModifications"`
	UnitTestSourceDirectory            string               `xml:"unitTestSourceDirectory"`
	IntegrationUnitTestSourceDirectory string               `xml:"integrationUnitTestSourceDirectory"`
	AspectSourceDirectory              string               `xml:"aspectSourceDirectory"`
	UnitTest                           *UnitTest            `xml:"unitTest"`
	Resources                          *Resources           `xml:"resources"`
}

// NagEmailAddress ...
type NagEmailAddress string

// SourceDirectory ...
type SourceDirectory string

// SourceModifications ...
type SourceModifications struct {
	XMLName            xml.Name              `xml:"sourceModifications"`
	SourceModification []*SourceModification `xml:"sourceModification"`
}

// SourceModification ...
type SourceModification struct {
	XMLName   xml.Name    `xml:"sourceModification"`
	ClassName string      `xml:"className"`
	Includes  []*Includes `xml:"includes"`
	Excludes  []*Excludes `xml:"excludes"`
}

// UnitTestSourceDirectory ...
type UnitTestSourceDirectory string

// IntegrationUnitTestSourceDirectory ...
type IntegrationUnitTestSourceDirectory string

// AspectSourceDirectory ...
type AspectSourceDirectory string

// UnitTest ...
type UnitTest struct {
	XMLName   xml.Name    `xml:"unitTest"`
	Includes  []*Includes `xml:"includes"`
	Excludes  []*Excludes `xml:"excludes"`
	Resources *Resources  `xml:"resources"`
}

// Includes ...
type Includes struct {
	XMLName xml.Name `xml:"includes"`
	Include []string `xml:"include"`
}

// Excludes ...
type Excludes struct {
	XMLName xml.Name `xml:"excludes"`
	Exclude []string `xml:"exclude"`
}

// Include ...
type Include string

// Exclude ...
type Exclude string

// Resources ...
type Resources struct {
	XMLName  xml.Name    `xml:"resources"`
	Resource []*Resource `xml:"resource"`
}

// Directory ...
type Directory string

// TargetPath ...
type TargetPath string

// Filtering ...
type Filtering bool

// Resource ...
type Resource struct {
	XMLName    xml.Name    `xml:"resource"`
	Directory  string      `xml:"directory"`
	TargetPath string      `xml:"targetPath"`
	Includes   []*Includes `xml:"includes"`
	Excludes   []*Excludes `xml:"excludes"`
	Filtering  bool        `xml:"filtering"`
}

// Reports ...
type Reports struct {
	XMLName xml.Name `xml:"reports"`
	Report  []string `xml:"report"`
}

// Properties ...
type Properties struct {
	XMLName xml.Name `xml:"properties"`
}

// Classloader ...
type Classloader string
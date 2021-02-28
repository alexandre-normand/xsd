// Code generated by xgen. DO NOT EDIT.

package schema

import (
	"encoding/xml"
)

// Component is Describes the component layout and packaging.
type Component *Component

// ModuleSets ...
type ModuleSets struct {
	XMLName   xml.Name     `xml:"moduleSets"`
	ModuleSet []*ModuleSet `xml:"moduleSet"`
}

// FileSets ...
type FileSets struct {
	XMLName xml.Name   `xml:"fileSets"`
	FileSet []*FileSet `xml:"fileSet"`
}

// Files ...
type Files struct {
	XMLName xml.Name    `xml:"files"`
	File    []*FileItem `xml:"file"`
}

// DependencySets ...
type DependencySets struct {
	XMLName       xml.Name         `xml:"dependencySets"`
	DependencySet []*DependencySet `xml:"dependencySet"`
}

// Repositories ...
type Repositories struct {
	XMLName    xml.Name      `xml:"repositories"`
	Repository []*Repository `xml:"repository"`
}

// ContainerDescriptorHandlers ...
type ContainerDescriptorHandlers struct {
	XMLName                    xml.Name                            `xml:"containerDescriptorHandlers"`
	ContainerDescriptorHandler []*ContainerDescriptorHandlerConfig `xml:"containerDescriptorHandler"`
}

// Component is Describes the component layout and packaging.
type Component struct {
	ModuleSets                  *ModuleSets                  `xml:"moduleSets"`
	FileSets                    *FileSets                    `xml:"fileSets"`
	Files                       *Files                       `xml:"files"`
	DependencySets              *DependencySets              `xml:"dependencySets"`
	Repositories                *Repositories                `xml:"repositories"`
	ContainerDescriptorHandlers *ContainerDescriptorHandlers `xml:"containerDescriptorHandlers"`
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

// DependencySet is Sets the output directory relative to the root
//             of the root directory of the assembly. For example,
//             "log" will put the specified files in the log directory,
//             directly beneath the root of the archive.
type DependencySet struct {
	OutputDirectory           string         `xml:"outputDirectory"`
	Includes                  *Includes      `xml:"includes"`
	Excludes                  *Excludes      `xml:"excludes"`
	FileMode                  string         `xml:"fileMode"`
	DirectoryMode             string         `xml:"directoryMode"`
	UseStrictFiltering        bool           `xml:"useStrictFiltering"`
	OutputFileNameMapping     string         `xml:"outputFileNameMapping"`
	Unpack                    bool           `xml:"unpack"`
	UnpackOptions             *UnpackOptions `xml:"unpackOptions"`
	Scope                     string         `xml:"scope"`
	UseProjectArtifact        bool           `xml:"useProjectArtifact"`
	UseProjectAttachments     bool           `xml:"useProjectAttachments"`
	UseTransitiveDependencies bool           `xml:"useTransitiveDependencies"`
	UseTransitiveFiltering    bool           `xml:"useTransitiveFiltering"`
}

// UnpackOptions is Specifies options for including/excluding/filtering items extracted from an archive. (Since 2.2)
type UnpackOptions struct {
	Includes           *Includes `xml:"includes"`
	Excludes           *Excludes `xml:"excludes"`
	Filtered           bool      `xml:"filtered"`
	LineEnding         string    `xml:"lineEnding"`
	UseDefaultExcludes bool      `xml:"useDefaultExcludes"`
	Encoding           string    `xml:"encoding"`
}

// Configuration ...
type Configuration struct {
	XMLName xml.Name `xml:"configuration"`
}

// ContainerDescriptorHandlerConfig is The handler's plexus role-hint, for lookup from the container.
type ContainerDescriptorHandlerConfig struct {
	HandlerName   string         `xml:"handlerName"`
	Configuration *Configuration `xml:"configuration"`
}

// GroupVersionAlignments ...
type GroupVersionAlignments struct {
	XMLName               xml.Name                 `xml:"groupVersionAlignments"`
	GroupVersionAlignment []*GroupVersionAlignment `xml:"groupVersionAlignment"`
}

// Repository is Sets the output directory relative to the root
//             of the root directory of the assembly. For example,
//             "log" will put the specified files in the log directory,
//             directly beneath the root of the archive.
type Repository struct {
	OutputDirectory        string                  `xml:"outputDirectory"`
	Includes               *Includes               `xml:"includes"`
	Excludes               *Excludes               `xml:"excludes"`
	FileMode               string                  `xml:"fileMode"`
	DirectoryMode          string                  `xml:"directoryMode"`
	IncludeMetadata        bool                    `xml:"includeMetadata"`
	GroupVersionAlignments *GroupVersionAlignments `xml:"groupVersionAlignments"`
	Scope                  string                  `xml:"scope"`
}

// GroupVersionAlignment is The version you want to align this group to.
type GroupVersionAlignment struct {
	Id       string    `xml:"id"`
	Version  string    `xml:"version"`
	Excludes *Excludes `xml:"excludes"`
}

// ModuleSet is If set to false, the plugin will exclude sub-modules from processing in this ModuleSet.
//           Otherwise, it will process all sub-modules, each subject to include/exclude rules.
//           Default value is true. (Since 2.2)
type ModuleSet struct {
	UseAllReactorProjects bool            `xml:"useAllReactorProjects"`
	IncludeSubModules     bool            `xml:"includeSubModules"`
	Includes              *Includes       `xml:"includes"`
	Excludes              *Excludes       `xml:"excludes"`
	Sources               *ModuleSources  `xml:"sources"`
	Binaries              *ModuleBinaries `xml:"binaries"`
}

// ModuleBinaries is Sets the output directory relative to the root
//             of the root directory of the assembly. For example,
//             "log" will put the specified files in the log directory,
//             directly beneath the root of the archive.
type ModuleBinaries struct {
	OutputDirectory       string          `xml:"outputDirectory"`
	Includes              *Includes       `xml:"includes"`
	Excludes              *Excludes       `xml:"excludes"`
	FileMode              string          `xml:"fileMode"`
	DirectoryMode         string          `xml:"directoryMode"`
	AttachmentClassifier  string          `xml:"attachmentClassifier"`
	IncludeDependencies   bool            `xml:"includeDependencies"`
	DependencySets        *DependencySets `xml:"dependencySets"`
	Unpack                bool            `xml:"unpack"`
	UnpackOptions         *UnpackOptions  `xml:"unpackOptions"`
	OutputFileNameMapping string          `xml:"outputFileNameMapping"`
}

// ModuleSources is Sets the output directory relative to the root
//             of the root directory of the assembly. For example,
//             "log" will put the specified files in the log directory.
type ModuleSources struct {
	UseDefaultExcludes          bool      `xml:"useDefaultExcludes"`
	OutputDirectory             string    `xml:"outputDirectory"`
	Includes                    *Includes `xml:"includes"`
	Excludes                    *Excludes `xml:"excludes"`
	FileMode                    string    `xml:"fileMode"`
	DirectoryMode               string    `xml:"directoryMode"`
	FileSets                    *FileSets `xml:"fileSets"`
	IncludeModuleDirectory      bool      `xml:"includeModuleDirectory"`
	ExcludeSubModuleDirectories bool      `xml:"excludeSubModuleDirectories"`
	OutputDirectoryMapping      string    `xml:"outputDirectoryMapping"`
}

// FileSet is Sets the output directory relative to the root
//             of the root directory of the assembly. For example,
//             "log" will put the specified files in the log directory.
type FileSet struct {
	UseDefaultExcludes bool      `xml:"useDefaultExcludes"`
	OutputDirectory    string    `xml:"outputDirectory"`
	Includes           *Includes `xml:"includes"`
	Excludes           *Excludes `xml:"excludes"`
	FileMode           string    `xml:"fileMode"`
	DirectoryMode      string    `xml:"directoryMode"`
	Directory          string    `xml:"directory"`
	LineEnding         string    `xml:"lineEnding"`
	Filtered           bool      `xml:"filtered"`
}

// FileItem is Sets whether to determine if the file is filtered.
type FileItem struct {
	Source          string `xml:"source"`
	OutputDirectory string `xml:"outputDirectory"`
	DestName        string `xml:"destName"`
	FileMode        string `xml:"fileMode"`
	LineEnding      string `xml:"lineEnding"`
	Filtered        bool   `xml:"filtered"`
}
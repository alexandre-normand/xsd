// Code generated by xgen. DO NOT EDIT.

typedef Component Component;

// FileSets ...
typedef struct {
	FileSet FileSet[];
} FileSets;

// Files ...
typedef struct {
	FileItem File[];
} Files;

// DependencySets ...
typedef struct {
	DependencySet DependencySet[];
} DependencySets;

// Repositories ...
typedef struct {
	Repository Repository[];
} Repositories;

// ContainerDescriptorHandlers ...
typedef struct {
	ContainerDescriptorHandlerConfig ContainerDescriptorHandler[];
} ContainerDescriptorHandlers;

// Component is Describes the component layout and packaging.
typedef struct {
	FileSets FileSets;
	Files Files;
	DependencySets DependencySets;
	Repositories Repositories;
	ContainerDescriptorHandlers ContainerDescriptorHandlers;
} Component;

// Configuration ...
typedef struct {
} Configuration;

// ContainerDescriptorHandlerConfig is The handler's plexus role-hint, for lookup from the container.
typedef struct {
	char HandlerName;
	Configuration Configuration;
} ContainerDescriptorHandlerConfig;

// GroupVersionAlignments ...
typedef struct {
	GroupVersionAlignment GroupVersionAlignment[];
} GroupVersionAlignments;

// Includes ...
typedef struct {
	char Include[];
} Includes;

// Excludes ...
typedef struct {
	char Exclude[];
} Excludes;

// Repository is If set to true, this property will trigger the creation of repository
//             metadata which will allow the repository to be used as a functional remote
//             repository.
typedef struct {
	bool IncludeMetadata;
	GroupVersionAlignments GroupVersionAlignments;
	char Scope;
	bool UseStrictFiltering;
	bool UseDefaultExcludes;
	char OutputDirectory;
	Includes Includes;
	Excludes Excludes;
	char FileMode;
	char DirectoryMode;
} Repository;

// GroupVersionAlignment is The version you want to align this group to.
typedef struct {
	char Id;
	char Version;
	Excludes Excludes;
} GroupVersionAlignment;

// FileItem is Sets whether to determine if the file is filtered.
typedef struct {
	char Source;
	char OutputDirectory;
	char DestName;
	char FileMode;
	char LineEnding;
	bool Filtered;
} FileItem;

// FileSet is Sets the output directory relative to the root
//             of the root directory of the assembly. For example,
//             "log" will put the specified files in the log directory.
typedef struct {
	char Directory;
	char LineEnding;
	bool Filtered;
	bool UseStrictFiltering;
	bool UseDefaultExcludes;
	char OutputDirectory;
	Includes Includes;
	Excludes Excludes;
	char FileMode;
	char DirectoryMode;
} FileSet;

// DependencySet is Sets the output directory relative to the root
//             of the root directory of the assembly. For example,
//             "log" will put the specified files in the log directory.
typedef struct {
	char OutputFileNameMapping;
	bool Unpack;
	UnpackOptions UnpackOptions;
	char Scope;
	bool UseProjectArtifact;
	bool UseProjectAttachments;
	bool UseTransitiveDependencies;
	bool UseTransitiveFiltering;
	bool UseStrictFiltering;
	bool UseDefaultExcludes;
	char OutputDirectory;
	Includes Includes;
	Excludes Excludes;
	char FileMode;
	char DirectoryMode;
} DependencySet;

// UnpackOptions is Specifies options for including/excluding/filtering items extracted from an archive.
typedef struct {
	Includes Includes;
	Excludes Excludes;
	bool Filtered;
} UnpackOptions;
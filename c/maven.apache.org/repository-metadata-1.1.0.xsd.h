// Code generated by xgen. DO NOT EDIT.

typedef Metadata Metadata;

// Plugins ...
typedef struct {
	Plugin Plugin[];
} Plugins;

// Metadata is Versioning information for the artifact.
typedef struct {
	char ModelVersionAttr; // attr, optional
	char GroupId;
	char ArtifactId;
	char Version;
	Versioning Versioning;
	Plugins Plugins;
} Metadata;

// Plugin is The plugin artifactId
typedef struct {
	char Name;
	char Prefix;
	char ArtifactId;
} Plugin;

// Versions ...
typedef struct {
	char Version[];
} Versions;

// SnapshotVersions ...
typedef struct {
	SnapshotVersion SnapshotVersion[];
} SnapshotVersions;

// Versioning is The current snapshot data in use for this version (artifact snapshots only)
typedef struct {
	char Latest;
	char Release;
	Snapshot Snapshot;
	Versions Versions;
	char LastUpdated;
	SnapshotVersions SnapshotVersions;
} Versioning;

// SnapshotVersion is The timestamp when this version information was last updated. The timestamp is expressed using UTC in the format yyyyMMddHHmmss.
typedef struct {
	char Classifier;
	char Extension;
	char Value;
	char Updated;
} SnapshotVersion;

// Snapshot is Whether to use a local copy instead (with filename that includes the base version)
typedef struct {
	char Timestamp;
	int BuildNumber;
	bool LocalCopy;
} Snapshot;
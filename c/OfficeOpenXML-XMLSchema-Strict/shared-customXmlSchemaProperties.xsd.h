// Code generated by xgen. DO NOT EDIT.

// CT_Schema ...
typedef struct {
	char UriAttr; // attr, optional
	char ManifestLocationAttr; // attr, optional
	char SchemaLocationAttr; // attr, optional
	char SchemaLanguageAttr; // attr, optional
} CT_Schema;

// CT_SchemaLibrary ...
typedef struct {
	CT_Schema Schema[];
} CT_SchemaLibrary;

typedef CT_SchemaLibrary SchemaLibrary;
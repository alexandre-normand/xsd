// Code generated by xgen. DO NOT EDIT.

typedef ChangesDocument Document;

// ChangesDocument is Contains the releases of this project with the actions taken
//             for each of the releases.
typedef struct {
	Properties Properties;
	Body Body;
} ChangesDocument;

// Body is The list of releases for this project.
typedef struct {
	Release Release[];
} Body;

// Release is The list of actions taken for this release.
typedef struct {
	char VersionAttr; // attr, optional
	char DateAttr; // attr, optional
	char DescriptionAttr; // attr, optional
	Action Action[];
} Release;

// Action is A list of contibutors for this issue.
typedef struct {
	char DevAttr; // attr, optional
	char DuetoAttr; // attr, optional
	char DuetoemailAttr; // attr, optional
	char IssueAttr; // attr, optional
	char TypeAttr; // attr, optional
	char SystemAttr; // attr, optional
	char DateAttr; // attr, optional
	FixedIssue Fixes[];
	DueTo Dueto[];
} Action;

// FixedIssue is A fixed issue.
typedef struct {
	char IssueAttr; // attr, optional
} FixedIssue;

// DueTo is Name and Email of the person to be credited for this change. This can be used when a patch is submitted by a non-committer.
typedef struct {
	char NameAttr; // attr, optional
	char EmailAttr; // attr, optional
} DueTo;

// Properties is Page Author
typedef struct {
	char Title;
	Author Author;
} Properties;

// Author is A description of the author page.
typedef struct {
	char EmailAttr; // attr, optional
} Author;
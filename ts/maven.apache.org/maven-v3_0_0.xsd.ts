// Code generated by xgen. DO NOT EDIT.

// Project is The <code>&lt;project&gt;</code> element specifies various attributes 
//          about a project. This is the root element of the project descriptor. 
//          The following table lists all of the possible child elements. Child 
//          elements with children are then documented further in subsequent 
//          sections.
export type Project = Model;

// MailingLists ...
export class MailingLists {
	MailingList: Array<MailingList>;
}

// Developers ...
export class Developers {
	Developer: Array<Developer>;
}

// Contributors ...
export class Contributors {
	Contributor: Array<Contributor>;
}

// Licenses ...
export class Licenses {
	License: Array<License>;
}

// Versions ...
export class Versions {
	Version: Array<Version>;
}

// Branches ...
export class Branches {
	Branch: Array<Branch>;
}

// PackageGroups ...
export class PackageGroups {
	PackageGroup: Array<PackageGroup>;
}

// Reports ...
export class Reports {
	Report: string;
}

// Properties ...
export class Properties {
}

// Dependencies ...
export class Dependencies {
	Dependency: Array<Dependency>;
}

// Model is Optional. The directory on the web server where the final
//             distributions will be published.  This is used when the
//             distributions are
//             <a href="/plugins/dist/index.html">deployed</a>.
export class Model {
	Extend: string;
	PomVersion: string;
	Id: string;
	GroupId: string;
	ArtifactId: string;
	Name: string;
	CurrentVersion: string;
	ShortDescription: string;
	Description: string;
	Url: string;
	Logo: string;
	IssueTrackingUrl: string;
	InceptionYear: string;
	GumpRepositoryId: string;
	SiteAddress: string;
	SiteDirectory: string;
	DistributionSite: string;
	DistributionDirectory: string;
	MailingLists: MailingLists;
	Developers: Developers;
	Contributors: Contributors;
	Licenses: Licenses;
	Versions: Versions;
	Branches: Branches;
	PackageGroups: PackageGroups;
	Reports: Reports;
	Repository: Repository;
	Organization: Organization;
	Properties: Properties;
	Package: string;
	Build: Build;
	Dependencies: Dependencies;
}

// SourceModifications ...
export class SourceModifications {
	SourceModification: Array<SourceModification>;
}

// Resources ...
export class Resources {
	Resource: Array<Resource>;
}

// Build is This element specifies a directory containing integration test    
//              sources of the project.
export class Build {
	NagEmailAddress: string;
	SourceDirectory: string;
	UnitTestSourceDirectory: string;
	AspectSourceDirectory: string;
	IntegrationUnitTestSourceDirectory: string;
	SourceModifications: SourceModifications;
	UnitTest: UnitTest;
	DefaultGoal: string;
	Resources: Resources;
}

// Includes ...
export class Includes {
	Include: string;
}

// Excludes ...
export class Excludes {
	Exclude: string;
}

// UnitTest is 3.0.0
export class UnitTest {
	Resources: Resources;
	Includes: Includes;
	Excludes: Excludes;
}

// Resource is Describe the directory where the resource is stored.
//             The path may be absolute, or relative to the project.xml file.
export class Resource {
	TargetPath: string;
	Filtering: boolean;
	Directory: string;
	Includes: Includes;
	Excludes: Excludes;
}

// SourceModification is Describe the directory where the resource is stored.
//             The path may be absolute, or relative to the project.xml file.
export class SourceModification {
	ClassName: string;
	Property: string;
	Directory: string;
	Includes: Includes;
	Excludes: Excludes;
}

// Organization is The URL to the organization's logo image.  This can be an URL relative
//             to the base directory of the generated web site,
//             (e.g., <code>/images/org-logo.png</code>) or an absolute URL
//             (e.g., <code>http://my.corp/logo.png</code>).  This value is used
//             when generating the project documentation.
export class Organization {
	Name: string;
	Url: string;
	Logo: string;
}

// Roles ...
export class Roles {
	Role: string;
}

// Developer is The URL of the organization.
export class Developer {
	Id: string;
	Name: string;
	Email: string;
	Url: string;
	Organization: string;
	OrganizationUrl: string;
	Roles: Roles;
	Timezone: string;
	Properties: Properties;
}

// Dependency is The type of dependency. This defaults to <code>jar</code>.
//             Known recognised dependency types are:
//             <ul>
//             <li><code>jar</code></li>
//             <li><code>ejb</code></li>
//             <li><code>plugin</code></li>
//             </ul>
export class Dependency {
	Id: string;
	GroupId: string;
	ArtifactId: string;
	Version: string;
	Url: string;
	Jar: string;
	Type: string;
	Properties: Properties;
}

// Repository is The URL to the project's browsable CVS repository.
export class Repository {
	Connection: string;
	DeveloperConnection: string;
	Url: string;
}

// PackageGroup is the description
export class PackageGroup {
	Title: string;
	Packages: string;
}

// Version is A unique identifier for a version.  This ID is
//             used to specify the version that
//             <a href="/plugins/dist/index.html">
//               <code>maven:dist</code>
//             </a> builds.
export class Version {
	Name: string;
	Tag: string;
	Id: string;
}

// License is Addendum information pertaining to this license.
export class License {
	Name: string;
	Url: string;
	Distribution: string;
	Comments: string;
}

// Contributor is The URL of the organization.
export class Contributor {
	Name: string;
	Email: string;
	Url: string;
	Organization: string;
	OrganizationUrl: string;
	Roles: Roles;
	Timezone: string;
	Properties: Properties;
}

// Branch is The branch tag in the version control system (e.g. cvs) used by the 
//             project for the source code associated with this branch of the
//             project.
export class Branch {
	Tag: string;
}

// OtherArchives ...
export class OtherArchives {
	OtherArchive: string;
}

// MailingList is The link to a URL where you can browse the mailing list archive.
export class MailingList {
	Name: string;
	Subscribe: string;
	Unsubscribe: string;
	Post: string;
	Archive: string;
	OtherArchives: OtherArchives;
}
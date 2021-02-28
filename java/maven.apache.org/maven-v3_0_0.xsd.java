// Code generated by xgen. DO NOT EDIT.

package schema;

import java.util.ArrayList;
import java.util.List;
import javax.xml.bind.annotation.XmlAccessType;
import javax.xml.bind.annotation.XmlAccessorType;
import javax.xml.bind.annotation.XmlAttribute;
import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlSchemaType;
import javax.xml.bind.annotation.XmlType;

@XmlAccessorType(XmlAccessType.FIELD)
@XmlElement(required = true, name = "project")
public class Project {
	protected Model Project;
}

// MailingLists ...
public class MailingLists {
	@XmlElement(required = true, name = "mailingList")
	protected List<MailingList> MailingList;
}

// Developers ...
public class Developers {
	@XmlElement(required = true, name = "developer")
	protected List<Developer> Developer;
}

// Contributors ...
public class Contributors {
	@XmlElement(required = true, name = "contributor")
	protected List<Contributor> Contributor;
}

// Licenses ...
public class Licenses {
	@XmlElement(required = true, name = "license")
	protected List<License> License;
}

// Versions ...
public class Versions {
	@XmlElement(required = true, name = "version")
	protected List<Version> Version;
}

// Branches ...
public class Branches {
	@XmlElement(required = true, name = "branch")
	protected List<Branch> Branch;
}

// PackageGroups ...
public class PackageGroups {
	@XmlElement(required = true, name = "packageGroup")
	protected List<PackageGroup> PackageGroup;
}

// Reports ...
public class Reports {
	@XmlElement(required = true, name = "report")
	protected List<String> Report;
}

// Properties ...
public class Properties {
}

// Dependencies ...
public class Dependencies {
	@XmlElement(required = true, name = "dependency")
	protected List<Dependency> Dependency;
}

// Model is Optional. The directory on the web server where the final
//             distributions will be published.  This is used when the
//             distributions are
//             <a href="/plugins/dist/index.html">deployed</a>.
public class Model {
	@XmlElement(required = true, name = "extend")
	protected String Extend;
	@XmlElement(required = true, name = "pomVersion")
	protected String PomVersion;
	@XmlElement(required = true, name = "id")
	protected String Id;
	@XmlElement(required = true, name = "groupId")
	protected String GroupId;
	@XmlElement(required = true, name = "artifactId")
	protected String ArtifactId;
	@XmlElement(required = true, name = "name")
	protected String Name;
	@XmlElement(required = true, name = "currentVersion")
	protected String CurrentVersion;
	@XmlElement(required = true, name = "shortDescription")
	protected String ShortDescription;
	@XmlElement(required = true, name = "description")
	protected String Description;
	@XmlElement(required = true, name = "url")
	protected String Url;
	@XmlElement(required = true, name = "logo")
	protected String Logo;
	@XmlElement(required = true, name = "issueTrackingUrl")
	protected String IssueTrackingUrl;
	@XmlElement(required = true, name = "inceptionYear")
	protected String InceptionYear;
	@XmlElement(required = true, name = "gumpRepositoryId")
	protected String GumpRepositoryId;
	@XmlElement(required = true, name = "siteAddress")
	protected String SiteAddress;
	@XmlElement(required = true, name = "siteDirectory")
	protected String SiteDirectory;
	@XmlElement(required = true, name = "distributionSite")
	protected String DistributionSite;
	@XmlElement(required = true, name = "distributionDirectory")
	protected String DistributionDirectory;
	@XmlElement(required = true, name = "mailingLists")
	protected MailingLists MailingLists;
	@XmlElement(required = true, name = "developers")
	protected Developers Developers;
	@XmlElement(required = true, name = "contributors")
	protected Contributors Contributors;
	@XmlElement(required = true, name = "licenses")
	protected Licenses Licenses;
	@XmlElement(required = true, name = "versions")
	protected Versions Versions;
	@XmlElement(required = true, name = "branches")
	protected Branches Branches;
	@XmlElement(required = true, name = "packageGroups")
	protected PackageGroups PackageGroups;
	@XmlElement(required = true, name = "reports")
	protected Reports Reports;
	@XmlElement(required = true, name = "repository")
	protected Repository Repository;
	@XmlElement(required = true, name = "organization")
	protected Organization Organization;
	@XmlElement(required = true, name = "properties")
	protected Properties Properties;
	@XmlElement(required = true, name = "package")
	protected String Package;
	@XmlElement(required = true, name = "build")
	protected Build Build;
	@XmlElement(required = true, name = "dependencies")
	protected Dependencies Dependencies;
}

// SourceModifications ...
public class SourceModifications {
	@XmlElement(required = true, name = "sourceModification")
	protected List<SourceModification> SourceModification;
}

// Resources ...
public class Resources {
	@XmlElement(required = true, name = "resource")
	protected List<Resource> Resource;
}

// Build is This element specifies a directory containing integration test    
//              sources of the project.
public class Build {
	@XmlElement(required = true, name = "nagEmailAddress")
	protected String NagEmailAddress;
	@XmlElement(required = true, name = "sourceDirectory")
	protected String SourceDirectory;
	@XmlElement(required = true, name = "unitTestSourceDirectory")
	protected String UnitTestSourceDirectory;
	@XmlElement(required = true, name = "aspectSourceDirectory")
	protected String AspectSourceDirectory;
	@XmlElement(required = true, name = "integrationUnitTestSourceDirectory")
	protected String IntegrationUnitTestSourceDirectory;
	@XmlElement(required = true, name = "sourceModifications")
	protected SourceModifications SourceModifications;
	@XmlElement(required = true, name = "unitTest")
	protected UnitTest UnitTest;
	@XmlElement(required = true, name = "defaultGoal")
	protected String DefaultGoal;
	@XmlElement(required = true, name = "resources")
	protected Resources Resources;
}

// Includes ...
public class Includes {
	@XmlElement(required = true, name = "include")
	protected List<String> Include;
}

// Excludes ...
public class Excludes {
	@XmlElement(required = true, name = "exclude")
	protected List<String> Exclude;
}

// UnitTest is 3.0.0
public class UnitTest {
	@XmlElement(required = true, name = "resources")
	protected Resources Resources;
	@XmlElement(required = true, name = "includes")
	protected Includes Includes;
	@XmlElement(required = true, name = "excludes")
	protected Excludes Excludes;
}

// Resource is Describe the directory where the resource is stored.
//             The path may be absolute, or relative to the project.xml file.
public class Resource {
	@XmlElement(required = true, name = "targetPath")
	protected String TargetPath;
	@XmlElement(required = true, name = "filtering")
	protected Boolean Filtering;
	@XmlElement(required = true, name = "directory")
	protected String Directory;
	@XmlElement(required = true, name = "includes")
	protected Includes Includes;
	@XmlElement(required = true, name = "excludes")
	protected Excludes Excludes;
}

// SourceModification is Describe the directory where the resource is stored.
//             The path may be absolute, or relative to the project.xml file.
public class SourceModification {
	@XmlElement(required = true, name = "className")
	protected String ClassName;
	@XmlElement(required = true, name = "property")
	protected String Property;
	@XmlElement(required = true, name = "directory")
	protected String Directory;
	@XmlElement(required = true, name = "includes")
	protected Includes Includes;
	@XmlElement(required = true, name = "excludes")
	protected Excludes Excludes;
}

// Organization is The URL to the organization's logo image.  This can be an URL relative
//             to the base directory of the generated web site,
//             (e.g., <code>/images/org-logo.png</code>) or an absolute URL
//             (e.g., <code>http://my.corp/logo.png</code>).  This value is used
//             when generating the project documentation.
public class Organization {
	@XmlElement(required = true, name = "name")
	protected String Name;
	@XmlElement(required = true, name = "url")
	protected String Url;
	@XmlElement(required = true, name = "logo")
	protected String Logo;
}

// Roles ...
public class Roles {
	@XmlElement(required = true, name = "role")
	protected List<String> Role;
}

// Developer is The URL of the organization.
public class Developer {
	@XmlElement(required = true, name = "id")
	protected String Id;
	@XmlElement(required = true, name = "name")
	protected String Name;
	@XmlElement(required = true, name = "email")
	protected String Email;
	@XmlElement(required = true, name = "url")
	protected String Url;
	@XmlElement(required = true, name = "organization")
	protected String Organization;
	@XmlElement(required = true, name = "organizationUrl")
	protected String OrganizationUrl;
	@XmlElement(required = true, name = "roles")
	protected Roles Roles;
	@XmlElement(required = true, name = "timezone")
	protected String Timezone;
	@XmlElement(required = true, name = "properties")
	protected Properties Properties;
}

// Dependency is The type of dependency. This defaults to <code>jar</code>.
//             Known recognised dependency types are:
//             <ul>
//             <li><code>jar</code></li>
//             <li><code>ejb</code></li>
//             <li><code>plugin</code></li>
//             </ul>
public class Dependency {
	@XmlElement(required = true, name = "id")
	protected String Id;
	@XmlElement(required = true, name = "groupId")
	protected String GroupId;
	@XmlElement(required = true, name = "artifactId")
	protected String ArtifactId;
	@XmlElement(required = true, name = "version")
	protected String Version;
	@XmlElement(required = true, name = "url")
	protected String Url;
	@XmlElement(required = true, name = "jar")
	protected String Jar;
	@XmlElement(required = true, name = "type")
	protected String Type;
	@XmlElement(required = true, name = "properties")
	protected Properties Properties;
}

// Repository is The URL to the project's browsable CVS repository.
public class Repository {
	@XmlElement(required = true, name = "connection")
	protected String Connection;
	@XmlElement(required = true, name = "developerConnection")
	protected String DeveloperConnection;
	@XmlElement(required = true, name = "url")
	protected String Url;
}

// PackageGroup is the description
public class PackageGroup {
	@XmlElement(required = true, name = "title")
	protected String Title;
	@XmlElement(required = true, name = "packages")
	protected String Packages;
}

// Version is A unique identifier for a version.  This ID is
//             used to specify the version that
//             <a href="/plugins/dist/index.html">
//               <code>maven:dist</code>
//             </a> builds.
public class Version {
	@XmlElement(required = true, name = "name")
	protected String Name;
	@XmlElement(required = true, name = "tag")
	protected String Tag;
	@XmlElement(required = true, name = "id")
	protected String Id;
}

// License is Addendum information pertaining to this license.
public class License {
	@XmlElement(required = true, name = "name")
	protected String Name;
	@XmlElement(required = true, name = "url")
	protected String Url;
	@XmlElement(required = true, name = "distribution")
	protected String Distribution;
	@XmlElement(required = true, name = "comments")
	protected String Comments;
}

// Contributor is The URL of the organization.
public class Contributor {
	@XmlElement(required = true, name = "name")
	protected String Name;
	@XmlElement(required = true, name = "email")
	protected String Email;
	@XmlElement(required = true, name = "url")
	protected String Url;
	@XmlElement(required = true, name = "organization")
	protected String Organization;
	@XmlElement(required = true, name = "organizationUrl")
	protected String OrganizationUrl;
	@XmlElement(required = true, name = "roles")
	protected Roles Roles;
	@XmlElement(required = true, name = "timezone")
	protected String Timezone;
	@XmlElement(required = true, name = "properties")
	protected Properties Properties;
}

// Branch is The branch tag in the version control system (e.g. cvs) used by the 
//             project for the source code associated with this branch of the
//             project.
public class Branch {
	@XmlElement(required = true, name = "tag")
	protected String Tag;
}

// OtherArchives ...
public class OtherArchives {
	@XmlElement(required = true, name = "otherArchive")
	protected List<String> OtherArchive;
}

// MailingList is The link to a URL where you can browse the mailing list archive.
public class MailingList {
	@XmlElement(required = true, name = "name")
	protected String Name;
	@XmlElement(required = true, name = "subscribe")
	protected String Subscribe;
	@XmlElement(required = true, name = "unsubscribe")
	protected String Unsubscribe;
	@XmlElement(required = true, name = "post")
	protected String Post;
	@XmlElement(required = true, name = "archive")
	protected String Archive;
	@XmlElement(required = true, name = "otherArchives")
	protected OtherArchives OtherArchives;
}
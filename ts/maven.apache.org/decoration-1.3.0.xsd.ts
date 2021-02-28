// Code generated by xgen. DO NOT EDIT.

// Project is The <code>&lt;project&gt;</code> element is the root of the site decoration descriptor.
//          The following table lists all of the possible child elements.
export type Project = DecorationModel;

// PoweredBy ...
export class PoweredBy {
	Logo: Array<Logo>;
}

// Custom ...
export class Custom {
}

// DecorationModel is Modify the version published display properties.
export class DecorationModel {
	NameAttr: string | null;
	BannerLeft: Banner;
	BannerRight: Banner;
	GoogleAdSenseClient: string;
	GoogleAdSenseSlot: string;
	GoogleAnalyticsAccountId: string;
	PublishDate: PublishDate;
	Version: Version;
	PoweredBy: PoweredBy;
	Skin: Skin;
	Body: Body;
	Custom: Custom;
}

// Version is Modify display properties for version published.
export class Version {
	PositionAttr: string | null;
}

// Banner is The title for the banner image.
export class Banner {
	Name: string;
	Src: string;
	Alt: string;
	Href: string;
	Border: string;
	Width: string;
	Height: string;
	Title: string;
}

// Head ...
export class Head {
}

// Links ...
export class Links {
	Item: Array<LinkItem>;
}

// Breadcrumbs ...
export class Breadcrumbs {
	Item: Array<LinkItem>;
}

// Footer ...
export class Footer {
}

// Body is The main content decoration.
export class Body {
	Head: Head;
	Links: Links;
	Breadcrumbs: Breadcrumbs;
	Menu: Array<Menu>;
	Footer: Footer;
}

// LinkItem is A link in the navigation.
export class LinkItem {
	NameAttr: string | null;
	HrefAttr: string | null;
	ImgAttr: string | null;
	PositionAttr: string | null;
	AltAttr: string | null;
	BorderAttr: string | null;
	WidthAttr: string | null;
	HeightAttr: string | null;
	TargetAttr: string | null;
	TitleAttr: string | null;
}

// Menu is A list of menu item.
export class Menu {
	NameAttr: string | null;
	InheritAttr: string | null;
	InheritAsRefAttr: boolean | null;
	RefAttr: string | null;
	ImgAttr: string | null;
	AltAttr: string | null;
	PositionAttr: string | null;
	BorderAttr: string | null;
	WidthAttr: string | null;
	HeightAttr: string | null;
	TitleAttr: string | null;
	Item: Array<MenuItem>;
}

// MenuItem is A list of menu item.
export class MenuItem {
	CollapseAttr: boolean | null;
	RefAttr: string | null;
	NameAttr: string | null;
	HrefAttr: string | null;
	ImgAttr: string | null;
	PositionAttr: string | null;
	AltAttr: string | null;
	BorderAttr: string | null;
	WidthAttr: string | null;
	HeightAttr: string | null;
	TargetAttr: string | null;
	TitleAttr: string | null;
	Description: string;
	Item: Array<MenuItem>;
}

// Skin is The skin version.
export class Skin {
	GroupId: string;
	ArtifactId: string;
	Version: string;
}

// Logo is Power by logo on the navigation.
export class Logo {
	NameAttr: string | null;
	HrefAttr: string | null;
	ImgAttr: string | null;
	PositionAttr: string | null;
	AltAttr: string | null;
	BorderAttr: string | null;
	WidthAttr: string | null;
	HeightAttr: string | null;
	TargetAttr: string | null;
	TitleAttr: string | null;
}

// PublishDate is Modify display properties for date published.
export class PublishDate {
	PositionAttr: string | null;
	FormatAttr: string | null;
}
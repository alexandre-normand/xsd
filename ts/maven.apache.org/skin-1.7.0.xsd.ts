// Code generated by xgen. DO NOT EDIT.

// Skin is The <code>&lt;skin&gt;</code> element is the root of the skin descriptor.
export type Skin = SkinModel;

// SkinModel is Encoding of text content, like the Velocity template itself.
export class SkinModel {
	Prerequisites: Prerequisites;
	Encoding: string;
}

// Prerequisites is The minimum version of Doxia Sitetools required to use the resulting skin.
export class Prerequisites {
	Doxiasitetools: string;
}
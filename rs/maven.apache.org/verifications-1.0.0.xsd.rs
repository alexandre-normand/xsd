// Code generated by xgen. DO NOT EDIT.

#[macro_use]
extern crate serde_derive;
extern crate serde;
extern crate serde_xml_rs;

use serde_xml_rs::from_reader;


// verifications is Root element of the verifications file.
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct verifications {
	#[serde(rename = "verifications")]
	pub verifications: Verifications,
}


// Files ...
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct Files {
	#[serde(rename = "file")]
	pub file: Vec<File>,
}


// Verifications is Root element of the verifications file.
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct Verifications {
	#[serde(rename = "files")]
	pub files: Files,
}


// File is When this is set to <code>true</code> the plugin checks that the
//             file or directory exists. When set to <code>false</code> it checks
//             that the file or directory does <strong>not</strong> exist.
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct File {
	#[serde(rename = "location")]
	pub location: String,
	#[serde(rename = "contains")]
	pub contains: String,
	#[serde(rename = "exists")]
	pub exists: bool,
}
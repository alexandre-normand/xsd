// Code generated by xgen. DO NOT EDIT.

#[macro_use]
extern crate serde_derive;
extern crate serde;
extern crate serde_xml_rs;

use serde_xml_rs::from_reader;


// DepthValueRecorded ...
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct DepthValueRecorded {
	#[serde(rename = "DepthValue")]
	pub depth_value: f64,
	#[serde(rename = "DepthValueUnits")]
	pub depth_value_units: MandatoryStringType,
	#[serde(rename = "DepthRelativeTo")]
	pub depth_relative_to: MandatoryStringType,
}


// PurgedVolumeRecorded ...
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct PurgedVolumeRecorded {
	#[serde(rename = "PurgedVolume")]
	pub purged_volume: f64,
	#[serde(rename = "PurgedVolumeUnits")]
	pub purged_volume_units: MandatoryStringType,
}


// Measurement ...
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct Measurement {
	#[serde(rename = "DeterminandName")]
	pub determinand_name: MandatoryStringType,
	#[serde(rename = "ResultType")]
	pub result_type: MandatoryStringType,
	#[serde(rename = "ResultValue")]
	pub result_value: f64,
	#[serde(rename = "ResultUnits")]
	pub result_units: MandatoryStringType,
	#[serde(rename = "Qualifier")]
	pub qualifier: MandatoryStringType,
	#[serde(rename = "Comment")]
	pub comment: String,
}


// Sample ...
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct Sample {
	#[serde(rename = "Sampler")]
	pub sampler: String,
	#[serde(rename = "SampleType")]
	pub sample_type: MandatoryStringType,
	#[serde(rename = "CustomerSamplePointName")]
	pub customer_sample_point_name: MandatoryStringType,
	#[serde(rename = "SampleDateTime")]
	pub sample_date_time: u8,
	#[serde(rename = "PurposeTypeName")]
	pub purpose_type_name: MandatoryStringType,
	#[serde(rename = "MaterialName")]
	pub material_name: MandatoryStringType,
	#[serde(rename = "Mechanism")]
	pub mechanism: MandatoryStringType,
	#[serde(rename = "CustomersLabSampleRef")]
	pub customers_lab_sample_ref: MandatoryStringType,
	#[serde(rename = "CustomersLabSampleRefSecondary")]
	pub customers_lab_sample_ref_secondary: Vec<String>,
	#[serde(rename = "Comment")]
	pub comment: String,
	#[serde(rename = "LabName")]
	pub lab_name: MandatoryStringType,
	#[serde(rename = "AnalysisCompleteDateTime")]
	pub analysis_complete_date_time: u8,
	#[serde(rename = "DepthValueRecorded")]
	pub depth_value_recorded: DepthValueRecorded,
	#[serde(rename = "PurgedVolumeRecorded")]
	pub purged_volume_recorded: PurgedVolumeRecorded,
	#[serde(rename = "Measurement")]
	pub measurement: Vec<Measurement>,
}


// FileUpload ...
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct FileUpload {
	#[serde(rename = "Source")]
	pub source: String,
	#[serde(rename = "Sample")]
	pub sample: Vec<Sample>,
	#[serde(rename = "RegulatedCustomerIdentifier")]
	pub regulated_customer_identifier: MandatoryStringType,
	#[serde(rename = "CustomerReference")]
	pub customer_reference: String,
}


// customer_reference is Customer’s own reference.
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct customer_reference {
	#[serde(rename = "CustomerReference")]
	pub customer_reference: String,
}
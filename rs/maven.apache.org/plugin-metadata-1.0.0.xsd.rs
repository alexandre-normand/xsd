// Code generated by xgen. DO NOT EDIT.

#[macro_use]
extern crate serde_derive;
extern crate serde;
extern crate serde_xml_rs;

use serde_xml_rs::from_reader;


// plugin_metadata is Root element of a script-based mojo's plugin metadata bindings.
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct plugin_metadata {
	#[serde(rename = "pluginMetadata")]
	pub plugin_metadata: PluginMetadata,
}


// Mojos ...
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct Mojos {
	#[serde(rename = "mojo")]
	pub mojo: Vec<Mojo>,
}


// PluginMetadata is Root element of a script-based mojo's plugin metadata bindings.
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct PluginMetadata {
	#[serde(rename = "mojos")]
	pub mojos: Mojos,
}


// Components ...
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct Components {
	#[serde(rename = "component")]
	pub component: Vec<Component>,
}


// Parameters ...
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct Parameters {
	#[serde(rename = "parameter")]
	pub parameter: Vec<Parameter>,
}


// Mojo is Information about a sub-execution of the Maven lifecycle which should be processed.
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct Mojo {
	#[serde(rename = "goal")]
	pub goal: String,
	#[serde(rename = "phase")]
	pub phase: String,
	#[serde(rename = "aggregator")]
	pub aggregator: bool,
	#[serde(rename = "requiresDependencyResolution")]
	pub requires_dependency_resolution: String,
	#[serde(rename = "requiresProject")]
	pub requires_project: bool,
	#[serde(rename = "requiresReports")]
	pub requires_reports: bool,
	#[serde(rename = "requiresOnline")]
	pub requires_online: bool,
	#[serde(rename = "inheritByDefault")]
	pub inherit_by_default: bool,
	#[serde(rename = "requiresDirectInvocation")]
	pub requires_direct_invocation: bool,
	#[serde(rename = "execution")]
	pub execution: LifecycleExecution,
	#[serde(rename = "components")]
	pub components: Components,
	#[serde(rename = "parameters")]
	pub parameters: Parameters,
	#[serde(rename = "description")]
	pub description: String,
	#[serde(rename = "deprecated")]
	pub deprecated: String,
	#[serde(rename = "call")]
	pub call: String,
}


// Parameter is A deprecation message for this mojo parameter.
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct Parameter {
	#[serde(rename = "name")]
	pub name: String,
	#[serde(rename = "alias")]
	pub alias: String,
	#[serde(rename = "property")]
	pub property: String,
	#[serde(rename = "required")]
	pub required: bool,
	#[serde(rename = "readonly")]
	pub readonly: bool,
	#[serde(rename = "expression")]
	pub expression: String,
	#[serde(rename = "defaultValue")]
	pub default_value: String,
	#[serde(rename = "type")]
	pub type_attr: String,
	#[serde(rename = "description")]
	pub description: String,
	#[serde(rename = "deprecated")]
	pub deprecated: String,
}


// LifecycleExecution is A goal, not attached to a lifecycle phase, which should be executed ahead of this mojo.
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct LifecycleExecution {
	#[serde(rename = "lifecycle")]
	pub lifecycle: String,
	#[serde(rename = "phase")]
	pub phase: String,
	#[serde(rename = "goal")]
	pub goal: String,
}


// Component is The role-hint to lookup.
#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct Component {
	#[serde(rename = "role")]
	pub role: String,
	#[serde(rename = "hint")]
	pub hint: String,
}
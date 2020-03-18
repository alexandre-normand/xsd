// Copyright 2020 The xgen Authors. All rights reserved.
//
// DO NOT EDIT: generated by xgen XSD generator
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

export class CT_CTName {
	LangAttr: string | null;
	ValAttr: string;
}

export class CT_CTDescription {
	LangAttr: string | null;
	ValAttr: string;
}

export class CT_CTCategory {
	TypeAttr: string;
	PriAttr: number;
}

export class CT_CTCategories {
	Cat: Array<CT_CTCategory>;
}

export type ST_ClrAppMethod = string;

export type ST_HueDir = string;

export class CT_Colors {
	MethAttr: string | null;
	HueDirAttr: string | null;
	AEG_ColorChoice: Array<EG_ColorChoice>;
}

export class CT_CTStyleLabel {
	NameAttr: string;
	FillClrLst: Array<CT_Colors>;
	LinClrLst: Array<CT_Colors>;
	EffectClrLst: Array<CT_Colors>;
	TxLinClrLst: Array<CT_Colors>;
	TxFillClrLst: Array<CT_Colors>;
	TxEffectClrLst: Array<CT_Colors>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export class CT_ColorTransform {
	UniqueIdAttr: string | null;
	MinVerAttr: string | null;
	Title: Array<CT_CTName>;
	Desc: Array<CT_CTDescription>;
	CatLst: Array<CT_CTCategories>;
	StyleLbl: Array<CT_CTStyleLabel>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export type ColorsDef = CT_ColorTransform;

export class CT_ColorTransformHeader {
	UniqueIdAttr: string;
	MinVerAttr: string | null;
	ResIdAttr: number | null;
	Title: Array<CT_CTName>;
	Desc: Array<CT_CTDescription>;
	CatLst: Array<CT_CTCategories>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export type ColorsDefHdr = CT_ColorTransformHeader;

export class CT_ColorTransformHeaderLst {
	ColorsDefHdr: Array<CT_ColorTransformHeader>;
}

export type ColorsDefHdrLst = CT_ColorTransformHeaderLst;

export type ST_PtType = string;

export class CT_Pt {
	ModelIdAttr: ST_ModelId;
	TypeAttr: string | null;
	CxnIdAttr: ST_ModelId | null;
	PrSet: Array<CT_ElemPropSet>;
	SpPr: Array<CT_ShapeProperties>;
	T: Array<CT_TextBody>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export class CT_PtList {
	Pt: Array<CT_Pt>;
}

export type ST_CxnType = string;

export class CT_Cxn {
	ModelIdAttr: ST_ModelId;
	TypeAttr: string | null;
	SrcIdAttr: ST_ModelId;
	DestIdAttr: ST_ModelId;
	SrcOrdAttr: number;
	DestOrdAttr: number;
	ParTransIdAttr: ST_ModelId | null;
	SibTransIdAttr: ST_ModelId | null;
	PresIdAttr: string | null;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export class CT_CxnList {
	Cxn: Array<CT_Cxn>;
}

export class CT_DataModel {
	PtLst: Array<CT_PtList>;
	CxnLst: Array<CT_CxnList>;
	Bg: Array<CT_BackgroundFormatting>;
	Whole: Array<CT_WholeE2oFormatting>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export type DataModel = CT_DataModel;

export class AG_IteratorAttributes {
	AxisAttr: ST_AxisTypes | null;
	PtTypeAttr: ST_ElementTypes | null;
	HideLastTransAttr: ST_Booleans | null;
	StAttr: ST_Ints | null;
	CntAttr: ST_UnsignedInts | null;
	StepAttr: ST_Ints | null;
}

export class AG_ConstraintAttributes {
	TypeAttr: string;
	ForAttr: string | null;
	ForNameAttr: string | null;
	PtTypeAttr: string | null;
}

export class AG_ConstraintRefAttributes {
	RefTypeAttr: string | null;
	RefForAttr: string | null;
	RefForNameAttr: string | null;
	RefPtTypeAttr: string | null;
}

export class CT_Constraint {
	AG_ConstraintAttributes: AG_ConstraintAttributes;
	AG_ConstraintRefAttributes: AG_ConstraintRefAttributes;
	OpAttr: string | null;
	ValAttr: number | null;
	FactAttr: number | null;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export class CT_Constraints {
	Constr: Array<CT_Constraint>;
}

export class CT_NumericRule {
	AG_ConstraintAttributes: AG_ConstraintAttributes;
	ValAttr: number | null;
	FactAttr: number | null;
	MaxAttr: number | null;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export class CT_Rules {
	Rule: Array<CT_NumericRule>;
}

export class CT_PresentationOf {
	AG_IteratorAttributes: AG_IteratorAttributes;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export class ST_LayoutShapeType {
	ST_ShapeType: string;
	ST_OutputShapeType: ST_OutputShapeType;
}

export type ST_Index1 = number;

export class CT_Adj {
	IdxAttr: number;
	ValAttr: number;
}

export class CT_AdjLst {
	Adj: Array<CT_Adj>;
}

export class CT_Shape {
	RotAttr: number | null;
	TypeAttr: ST_LayoutShapeType | null;
	RBlipAttr: string | null;
	ZOrderOffAttr: number | null;
	HideGeomAttr: boolean | null;
	LkTxEntryAttr: boolean | null;
	BlipPhldrAttr: boolean | null;
	AdjLst: Array<CT_AdjLst>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export class CT_Parameter {
	TypeAttr: string;
	ValAttr: ST_ParameterVal;
}

export class CT_Algorithm {
	TypeAttr: string;
	RevAttr: number | null;
	Param: Array<CT_Parameter>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export class CT_LayoutNode {
	NameAttr: string | null;
	StyleLblAttr: string | null;
	ChOrderAttr: string | null;
	MoveWithAttr: string | null;
	Alg: Array<CT_Algorithm>;
	Shape: Array<CT_Shape>;
	PresOf: Array<CT_PresentationOf>;
	ConstrLst: Array<CT_Constraints>;
	RuleLst: Array<CT_Rules>;
	VarLst: Array<CT_LayoutVariablePropertySet>;
	ForEach: Array<CT_ForEach>;
	LayoutNode: Array<CT_LayoutNode>;
	Choose: Array<CT_Choose>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export class CT_ForEach {
	AG_IteratorAttributes: AG_IteratorAttributes;
	NameAttr: string | null;
	RefAttr: string | null;
	Alg: Array<CT_Algorithm>;
	Shape: Array<CT_Shape>;
	PresOf: Array<CT_PresentationOf>;
	ConstrLst: Array<CT_Constraints>;
	RuleLst: Array<CT_Rules>;
	ForEach: Array<CT_ForEach>;
	LayoutNode: Array<CT_LayoutNode>;
	Choose: Array<CT_Choose>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export class CT_When {
	AG_IteratorAttributes: AG_IteratorAttributes;
	NameAttr: string | null;
	FuncAttr: string;
	ArgAttr: ST_FunctionArgument | null;
	OpAttr: string;
	ValAttr: ST_FunctionValue;
	Alg: Array<CT_Algorithm>;
	Shape: Array<CT_Shape>;
	PresOf: Array<CT_PresentationOf>;
	ConstrLst: Array<CT_Constraints>;
	RuleLst: Array<CT_Rules>;
	ForEach: Array<CT_ForEach>;
	LayoutNode: Array<CT_LayoutNode>;
	Choose: Array<CT_Choose>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export class CT_Otherwise {
	NameAttr: string | null;
	Alg: Array<CT_Algorithm>;
	Shape: Array<CT_Shape>;
	PresOf: Array<CT_PresentationOf>;
	ConstrLst: Array<CT_Constraints>;
	RuleLst: Array<CT_Rules>;
	ForEach: Array<CT_ForEach>;
	LayoutNode: Array<CT_LayoutNode>;
	Choose: Array<CT_Choose>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export class CT_Choose {
	NameAttr: string | null;
	If: Array<CT_When>;
	Else: Array<CT_Otherwise>;
}

export class CT_SampleData {
	UseDefAttr: boolean | null;
	DataModel: Array<CT_DataModel>;
}

export class CT_Category {
	TypeAttr: string;
	PriAttr: number;
}

export class CT_Categories {
	Cat: Array<CT_Category>;
}

export class CT_Name {
	LangAttr: string | null;
	ValAttr: string;
}

export class CT_Description {
	LangAttr: string | null;
	ValAttr: string;
}

export class CT_DiagramDefinition {
	UniqueIdAttr: string | null;
	MinVerAttr: string | null;
	DefStyleAttr: string | null;
	Title: Array<CT_Name>;
	Desc: Array<CT_Description>;
	CatLst: Array<CT_Categories>;
	SampData: Array<CT_SampleData>;
	StyleData: Array<CT_SampleData>;
	ClrData: Array<CT_SampleData>;
	LayoutNode: Array<CT_LayoutNode>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export type LayoutDef = CT_DiagramDefinition;

export class CT_DiagramDefinitionHeader {
	UniqueIdAttr: string;
	MinVerAttr: string | null;
	DefStyleAttr: string | null;
	ResIdAttr: number | null;
	Title: Array<CT_Name>;
	Desc: Array<CT_Description>;
	CatLst: Array<CT_Categories>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export type LayoutDefHdr = CT_DiagramDefinitionHeader;

export class CT_DiagramDefinitionHeaderLst {
	LayoutDefHdr: Array<CT_DiagramDefinitionHeader>;
}

export type LayoutDefHdrLst = CT_DiagramDefinitionHeaderLst;

export class CT_RelIds {
	RDmAttr: string;
	RLoAttr: string;
	RQsAttr: string;
	RCsAttr: string;
}

export type RelIds = CT_RelIds;

export class ST_ParameterVal {
	Int: number;
	Double: number;
	ST_ConnectorPoint: ST_ConnectorPoint;
	ST_NodeVerticalAlignment: ST_NodeVerticalAlignment;
	ST_LinearDirection: ST_LinearDirection;
	ST_SecondaryChildAlignment: ST_SecondaryChildAlignment;
	ST_ArrowheadStyle: ST_ArrowheadStyle;
	ST_TextDirection: ST_TextDirection;
	ST_StartingElement: ST_StartingElement;
	String: string;
	ST_TextBlockDirection: ST_TextBlockDirection;
	ST_FallbackDimension: ST_FallbackDimension;
	ST_PyramidAccentPosition: ST_PyramidAccentPosition;
	ST_TextAnchorVertical: ST_TextAnchorVertical;
	ST_Offset: ST_Offset;
	ST_NodeHorizontalAlignment: ST_NodeHorizontalAlignment;
	ST_SecondaryLinearDirection: ST_SecondaryLinearDirection;
	ST_BendPoint: ST_BendPoint;
	ST_RotationPath: ST_RotationPath;
	ST_CenterShapeMapping: ST_CenterShapeMapping;
	ST_GrowDirection: ST_GrowDirection;
	ST_Breakpoint: ST_Breakpoint;
	ST_HierarchyAlignment: ST_HierarchyAlignment;
	ST_VerticalAlignment: ST_VerticalAlignment;
	ST_ChildAlignment: ST_ChildAlignment;
	ST_PyramidAccentTextMargin: ST_PyramidAccentTextMargin;
	ST_TextAnchorHorizontal: ST_TextAnchorHorizontal;
	ST_AutoTextRotation: ST_AutoTextRotation;
	ST_ContinueDirection: ST_ContinueDirection;
	ST_ChildDirection: ST_ChildDirection;
	ST_ConnectorRouting: ST_ConnectorRouting;
	ST_ConnectorDimension: ST_ConnectorDimension;
	ST_DiagramTextAlignment: ST_DiagramTextAlignment;
	ST_FlowDirection: ST_FlowDirection;
	Boolean: boolean;
	ST_DiagramHorizontalAlignment: ST_DiagramHorizontalAlignment;
}

export class ST_ModelId {
	ST_Guid: string;
	Int: number;
}

export class ST_PrSetCustVal {
	ST_Percentage: string;
}

export class CT_ElemPropSet {
	PresAssocIDAttr: ST_ModelId | null;
	PresNameAttr: string | null;
	PresStyleLblAttr: string | null;
	PresStyleIdxAttr: number | null;
	PresStyleCntAttr: number | null;
	LoTypeIdAttr: string | null;
	LoCatIdAttr: string | null;
	QsTypeIdAttr: string | null;
	QsCatIdAttr: string | null;
	CsTypeIdAttr: string | null;
	CsCatIdAttr: string | null;
	Coherent3DOffAttr: boolean | null;
	PhldrTAttr: string | null;
	PhldrAttr: boolean | null;
	CustAngAttr: number | null;
	CustFlipVertAttr: boolean | null;
	CustFlipHorAttr: boolean | null;
	CustSzXAttr: number | null;
	CustSzYAttr: number | null;
	CustScaleXAttr: ST_PrSetCustVal | null;
	CustScaleYAttr: ST_PrSetCustVal | null;
	CustTAttr: boolean | null;
	CustLinFactXAttr: ST_PrSetCustVal | null;
	CustLinFactYAttr: ST_PrSetCustVal | null;
	CustLinFactNeighborXAttr: ST_PrSetCustVal | null;
	CustLinFactNeighborYAttr: ST_PrSetCustVal | null;
	CustRadScaleRadAttr: ST_PrSetCustVal | null;
	CustRadScaleIncAttr: ST_PrSetCustVal | null;
	PresLayoutVars: Array<CT_LayoutVariablePropertySet>;
	Style: Array<CT_ShapeStyle>;
}

export type ST_Direction = string;

export type ST_HierBranchStyle = string;

export type ST_AnimOneStr = string;

export type ST_AnimLvlStr = string;

export class CT_OrgChart {
	ValAttr: boolean | null;
}

export type ST_NodeCount = number;

export class CT_ChildMax {
	ValAttr: number | null;
}

export class CT_ChildPref {
	ValAttr: number | null;
}

export class CT_BulletEnabled {
	ValAttr: boolean | null;
}

export class CT_Direction {
	ValAttr: string | null;
}

export class CT_HierBranchStyle {
	ValAttr: string | null;
}

export class CT_AnimOne {
	ValAttr: string | null;
}

export class CT_AnimLvl {
	ValAttr: string | null;
}

export type ST_ResizeHandlesStr = string;

export class CT_ResizeHandles {
	ValAttr: string | null;
}

export class CT_LayoutVariablePropertySet {
	OrgChart: Array<CT_OrgChart>;
	ChMax: Array<CT_ChildMax>;
	ChPref: Array<CT_ChildPref>;
	BulletEnabled: Array<CT_BulletEnabled>;
	Dir: Array<CT_Direction>;
	HierBranch: Array<CT_HierBranchStyle>;
	AnimOne: Array<CT_AnimOne>;
	AnimLvl: Array<CT_AnimLvl>;
	ResizeHandles: Array<CT_ResizeHandles>;
}

export class CT_SDName {
	LangAttr: string | null;
	ValAttr: string;
}

export class CT_SDDescription {
	LangAttr: string | null;
	ValAttr: string;
}

export class CT_SDCategory {
	TypeAttr: string;
	PriAttr: number;
}

export class CT_SDCategories {
	Cat: Array<CT_SDCategory>;
}

export class CT_TextProps {
	AEG_Text3D: Array<EG_Text3D>;
}

export class CT_StyleLabel {
	NameAttr: string;
	Scene3d: Array<CT_Scene3D>;
	Sp3d: Array<CT_Shape3D>;
	TxPr: Array<CT_TextProps>;
	Style: Array<CT_ShapeStyle>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export class CT_StyleDefinition {
	UniqueIdAttr: string | null;
	MinVerAttr: string | null;
	Title: Array<CT_SDName>;
	Desc: Array<CT_SDDescription>;
	CatLst: Array<CT_SDCategories>;
	Scene3d: Array<CT_Scene3D>;
	StyleLbl: Array<CT_StyleLabel>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export type StyleDef = CT_StyleDefinition;

export class CT_StyleDefinitionHeader {
	UniqueIdAttr: string;
	MinVerAttr: string | null;
	ResIdAttr: number | null;
	Title: Array<CT_SDName>;
	Desc: Array<CT_SDDescription>;
	CatLst: Array<CT_SDCategories>;
	ExtLst: Array<CT_OfficeArtExtensionList>;
}

export type StyleDefHdr = CT_StyleDefinitionHeader;

export class CT_StyleDefinitionHeaderLst {
	StyleDefHdr: Array<CT_StyleDefinitionHeader>;
}

export type StyleDefHdrLst = CT_StyleDefinitionHeaderLst;

export type ST_AlgorithmType = string;

export type ST_AxisType = string;

export type ST_AxisTypes = Array<string>;

export type ST_BoolOperator = string;

export type ST_ChildOrderType = string;

export type ST_ConstraintType = string;

export type ST_ConstraintRelationship = string;

export type ST_ElementType = string;

export type ST_ElementTypes = Array<string>;

export type ST_ParameterId = string;

export type ST_Ints = Array<number>;

export type ST_UnsignedInts = Array<number>;

export type ST_Booleans = Array<boolean>;

export type ST_FunctionType = string;

export type ST_FunctionOperator = string;

export type ST_DiagramHorizontalAlignment = string;

export type ST_VerticalAlignment = string;

export type ST_ChildDirection = string;

export type ST_ChildAlignment = string;

export type ST_SecondaryChildAlignment = string;

export type ST_LinearDirection = string;

export type ST_SecondaryLinearDirection = string;

export type ST_StartingElement = string;

export type ST_RotationPath = string;

export type ST_CenterShapeMapping = string;

export type ST_BendPoint = string;

export type ST_ConnectorRouting = string;

export type ST_ArrowheadStyle = string;

export type ST_ConnectorDimension = string;

export type ST_ConnectorPoint = string;

export type ST_NodeHorizontalAlignment = string;

export type ST_NodeVerticalAlignment = string;

export type ST_FallbackDimension = string;

export type ST_TextDirection = string;

export type ST_PyramidAccentPosition = string;

export type ST_PyramidAccentTextMargin = string;

export type ST_TextBlockDirection = string;

export type ST_TextAnchorHorizontal = string;

export type ST_TextAnchorVertical = string;

export type ST_DiagramTextAlignment = string;

export type ST_AutoTextRotation = string;

export type ST_GrowDirection = string;

export type ST_FlowDirection = string;

export type ST_ContinueDirection = string;

export type ST_Breakpoint = string;

export type ST_Offset = string;

export type ST_HierarchyAlignment = string;

export class ST_FunctionValue {
	Int: number;
	Boolean: boolean;
	ST_Direction: string;
	ST_HierBranchStyle: string;
	ST_AnimOneStr: string;
	ST_AnimLvlStr: string;
	ST_ResizeHandlesStr: string;
}

export type ST_VariableType = string;

export class ST_FunctionArgument {
	ST_VariableType: string;
}

export type ST_OutputShapeType = string;

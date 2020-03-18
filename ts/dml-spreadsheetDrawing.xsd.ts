// Copyright 2020 The xgen Authors. All rights reserved.
//
// DO NOT EDIT: generated by xgen XSD generator
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

export type From = CT_Marker;

export type To = CT_Marker;

export class CT_AnchorClientData {
	FLocksWithSheetAttr: boolean | null;
	FPrintsWithSheetAttr: boolean | null;
}

export class CT_ShapeNonVisual {
	CNvPr: Array<CT_NonVisualDrawingProps>;
	CNvSpPr: Array<CT_NonVisualDrawingShapeProps>;
}

export class CT_Shape {
	MacroAttr: string | null;
	TextlinkAttr: string | null;
	FLocksTextAttr: boolean | null;
	FPublishedAttr: boolean | null;
	NvSpPr: Array<CT_ShapeNonVisual>;
	SpPr: Array<CT_ShapeProperties>;
	Style: Array<CT_ShapeStyle>;
	TxBody: Array<CT_TextBody>;
}

export class CT_ConnectorNonVisual {
	CNvPr: Array<CT_NonVisualDrawingProps>;
	CNvCxnSpPr: Array<CT_NonVisualConnectorProperties>;
}

export class CT_Connector {
	MacroAttr: string | null;
	FPublishedAttr: boolean | null;
	NvCxnSpPr: Array<CT_ConnectorNonVisual>;
	SpPr: Array<CT_ShapeProperties>;
	Style: Array<CT_ShapeStyle>;
}

export class CT_PictureNonVisual {
	CNvPr: Array<CT_NonVisualDrawingProps>;
	CNvPicPr: Array<CT_NonVisualPictureProperties>;
}

export class CT_Picture {
	MacroAttr: string | null;
	FPublishedAttr: boolean | null;
	NvPicPr: Array<CT_PictureNonVisual>;
	BlipFill: Array<CT_BlipFillProperties>;
	SpPr: Array<CT_ShapeProperties>;
	Style: Array<CT_ShapeStyle>;
}

export class CT_GraphicalObjectFrameNonVisual {
	CNvPr: Array<CT_NonVisualDrawingProps>;
	CNvGraphicFramePr: Array<CT_NonVisualGraphicFrameProperties>;
}

export class CT_GraphicalObjectFrame {
	MacroAttr: string | null;
	FPublishedAttr: boolean | null;
	NvGraphicFramePr: Array<CT_GraphicalObjectFrameNonVisual>;
	Xfrm: Array<CT_Transform2D>;
	AGraphic: Array<CT_GraphicalObject>;
}

export class CT_GroupShapeNonVisual {
	CNvPr: Array<CT_NonVisualDrawingProps>;
	CNvGrpSpPr: Array<CT_NonVisualGroupDrawingShapeProps>;
}

export class CT_GroupShape {
	NvGrpSpPr: Array<CT_GroupShapeNonVisual>;
	GrpSpPr: Array<CT_GroupShapeProperties>;
	Sp: Array<CT_Shape>;
	GrpSp: Array<CT_GroupShape>;
	GraphicFrame: Array<CT_GraphicalObjectFrame>;
	CxnSp: Array<CT_Connector>;
	Pic: Array<CT_Picture>;
}

export class EG_ObjectChoices {
	Sp: CT_Shape;
	GrpSp: CT_GroupShape;
	GraphicFrame: CT_GraphicalObjectFrame;
	CxnSp: CT_Connector;
	Pic: CT_Picture;
	ContentPart: CT_Rel;
}

export class CT_Rel {
	RIdAttr: string;
}

export type ST_ColID = number;

export type ST_RowID = number;

export class CT_Marker {
	Col: Array<number>;
	ColOff: Array<ST_Coordinate>;
	Row: Array<number>;
	RowOff: Array<ST_Coordinate>;
}

export type ST_EditAs = string;

export class CT_TwoCellAnchor {
	EditAsAttr: string | null;
	EG_ObjectChoices: EG_ObjectChoices;
	From: Array<CT_Marker>;
	To: Array<CT_Marker>;
	ClientData: Array<CT_AnchorClientData>;
}

export class CT_OneCellAnchor {
	EG_ObjectChoices: EG_ObjectChoices;
	From: Array<CT_Marker>;
	Ext: Array<CT_PositiveSize2D>;
	ClientData: Array<CT_AnchorClientData>;
}

export class CT_AbsoluteAnchor {
	EG_ObjectChoices: EG_ObjectChoices;
	Pos: Array<CT_Point2D>;
	Ext: Array<CT_PositiveSize2D>;
	ClientData: Array<CT_AnchorClientData>;
}

export class EG_Anchor {
	TwoCellAnchor: CT_TwoCellAnchor;
	OneCellAnchor: CT_OneCellAnchor;
	AbsoluteAnchor: CT_AbsoluteAnchor;
}

export class CT_Drawing {
	EG_Anchor: Array<EG_Anchor>;
}

export type WsDr = CT_Drawing;

// Code generated by xgen. DO NOT EDIT.

package schema

import (
	"encoding/xml"
)

// From ...
type From *CTMarker

// To ...
type To *CTMarker

// CTAnchorClientData ...
type CTAnchorClientData struct {
	XMLName              xml.Name `xml:"CT_AnchorClientData"`
	FLocksWithSheetAttr  bool     `xml:"fLocksWithSheet,attr,omitempty"`
	FPrintsWithSheetAttr bool     `xml:"fPrintsWithSheet,attr,omitempty"`
}

// CTShapeNonVisual ...
type CTShapeNonVisual struct {
	XMLName xml.Name                      `xml:"CT_ShapeNonVisual"`
	CNvPr   *CTNonVisualDrawingProps      `xml:"cNvPr"`
	CNvSpPr *CTNonVisualDrawingShapeProps `xml:"cNvSpPr"`
}

// CTShape ...
type CTShape struct {
	XMLName        xml.Name           `xml:"CT_Shape"`
	MacroAttr      string             `xml:"macro,attr,omitempty"`
	TextlinkAttr   string             `xml:"textlink,attr,omitempty"`
	FLocksTextAttr bool               `xml:"fLocksText,attr,omitempty"`
	FPublishedAttr bool               `xml:"fPublished,attr,omitempty"`
	NvSpPr         *CTShapeNonVisual  `xml:"nvSpPr"`
	SpPr           *CTShapeProperties `xml:"spPr"`
	Style          *CTShapeStyle      `xml:"style"`
	TxBody         *CTTextBody        `xml:"txBody"`
}

// CTConnectorNonVisual ...
type CTConnectorNonVisual struct {
	XMLName    xml.Name                        `xml:"CT_ConnectorNonVisual"`
	CNvPr      *CTNonVisualDrawingProps        `xml:"cNvPr"`
	CNvCxnSpPr *CTNonVisualConnectorProperties `xml:"cNvCxnSpPr"`
}

// CTConnector ...
type CTConnector struct {
	XMLName        xml.Name              `xml:"CT_Connector"`
	MacroAttr      string                `xml:"macro,attr,omitempty"`
	FPublishedAttr bool                  `xml:"fPublished,attr,omitempty"`
	NvCxnSpPr      *CTConnectorNonVisual `xml:"nvCxnSpPr"`
	SpPr           *CTShapeProperties    `xml:"spPr"`
	Style          *CTShapeStyle         `xml:"style"`
}

// CTPictureNonVisual ...
type CTPictureNonVisual struct {
	XMLName  xml.Name                      `xml:"CT_PictureNonVisual"`
	CNvPr    *CTNonVisualDrawingProps      `xml:"cNvPr"`
	CNvPicPr *CTNonVisualPictureProperties `xml:"cNvPicPr"`
}

// CTPicture ...
type CTPicture struct {
	XMLName        xml.Name              `xml:"CT_Picture"`
	MacroAttr      string                `xml:"macro,attr,omitempty"`
	FPublishedAttr bool                  `xml:"fPublished,attr,omitempty"`
	NvPicPr        *CTPictureNonVisual   `xml:"nvPicPr"`
	BlipFill       *CTBlipFillProperties `xml:"blipFill"`
	SpPr           *CTShapeProperties    `xml:"spPr"`
	Style          *CTShapeStyle         `xml:"style"`
}

// CTGraphicalObjectFrameNonVisual ...
type CTGraphicalObjectFrameNonVisual struct {
	XMLName           xml.Name                           `xml:"CT_GraphicalObjectFrameNonVisual"`
	CNvPr             *CTNonVisualDrawingProps           `xml:"cNvPr"`
	CNvGraphicFramePr *CTNonVisualGraphicFrameProperties `xml:"cNvGraphicFramePr"`
}

// CTGraphicalObjectFrame ...
type CTGraphicalObjectFrame struct {
	XMLName          xml.Name                         `xml:"CT_GraphicalObjectFrame"`
	MacroAttr        string                           `xml:"macro,attr,omitempty"`
	FPublishedAttr   bool                             `xml:"fPublished,attr,omitempty"`
	NvGraphicFramePr *CTGraphicalObjectFrameNonVisual `xml:"nvGraphicFramePr"`
	Xfrm             *CTTransform2D                   `xml:"xfrm"`
	AGraphic         *CTGraphicalObject               `xml:"a:graphic"`
}

// CTGroupShapeNonVisual ...
type CTGroupShapeNonVisual struct {
	XMLName    xml.Name                           `xml:"CT_GroupShapeNonVisual"`
	CNvPr      *CTNonVisualDrawingProps           `xml:"cNvPr"`
	CNvGrpSpPr *CTNonVisualGroupDrawingShapeProps `xml:"cNvGrpSpPr"`
}

// CTGroupShape ...
type CTGroupShape struct {
	XMLName      xml.Name                  `xml:"CT_GroupShape"`
	NvGrpSpPr    *CTGroupShapeNonVisual    `xml:"nvGrpSpPr"`
	GrpSpPr      *CTGroupShapeProperties   `xml:"grpSpPr"`
	Sp           []*CTShape                `xml:"sp"`
	GrpSp        []*CTGroupShape           `xml:"grpSp"`
	GraphicFrame []*CTGraphicalObjectFrame `xml:"graphicFrame"`
	CxnSp        []*CTConnector            `xml:"cxnSp"`
	Pic          []*CTPicture              `xml:"pic"`
}

// EGObjectChoices ...
type EGObjectChoices struct {
	XMLName      xml.Name `xml:"EG_ObjectChoices"`
	Sp           *CTShape
	GrpSp        *CTGroupShape
	GraphicFrame *CTGraphicalObjectFrame
	CxnSp        *CTConnector
	Pic          *CTPicture
	ContentPart  *CTRel
}

// CTRel ...
type CTRel struct {
	XMLName xml.Name `xml:"CT_Rel"`
	RIdAttr string   `xml:"r:id,attr"`
}

// STColID ...
type STColID int

// STRowID ...
type STRowID int

// CTMarker ...
type CTMarker struct {
	XMLName xml.Name      `xml:"CT_Marker"`
	Col     int           `xml:"col"`
	ColOff  *STCoordinate `xml:"colOff"`
	Row     int           `xml:"row"`
	RowOff  *STCoordinate `xml:"rowOff"`
}

// STEditAs ...
type STEditAs string

// CTTwoCellAnchor ...
type CTTwoCellAnchor struct {
	XMLName         xml.Name `xml:"CT_TwoCellAnchor"`
	EditAsAttr      string   `xml:"editAs,attr,omitempty"`
	EGObjectChoices *EGObjectChoices
	From            *CTMarker           `xml:"from"`
	To              *CTMarker           `xml:"to"`
	ClientData      *CTAnchorClientData `xml:"clientData"`
}

// CTOneCellAnchor ...
type CTOneCellAnchor struct {
	XMLName         xml.Name `xml:"CT_OneCellAnchor"`
	EGObjectChoices *EGObjectChoices
	From            *CTMarker           `xml:"from"`
	Ext             *CTPositiveSize2D   `xml:"ext"`
	ClientData      *CTAnchorClientData `xml:"clientData"`
}

// CTAbsoluteAnchor ...
type CTAbsoluteAnchor struct {
	XMLName         xml.Name `xml:"CT_AbsoluteAnchor"`
	EGObjectChoices *EGObjectChoices
	Pos             *CTPoint2D          `xml:"pos"`
	Ext             *CTPositiveSize2D   `xml:"ext"`
	ClientData      *CTAnchorClientData `xml:"clientData"`
}

// EGAnchor ...
type EGAnchor struct {
	XMLName        xml.Name `xml:"EG_Anchor"`
	TwoCellAnchor  *CTTwoCellAnchor
	OneCellAnchor  *CTOneCellAnchor
	AbsoluteAnchor *CTAbsoluteAnchor
}

// CTDrawing ...
type CTDrawing struct {
	XMLName  xml.Name `xml:"CT_Drawing"`
	EGAnchor []*EGAnchor
}

// WsDr ...
type WsDr *CTDrawing

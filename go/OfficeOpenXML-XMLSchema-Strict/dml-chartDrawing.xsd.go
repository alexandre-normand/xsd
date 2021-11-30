// Code generated by xgen. DO NOT EDIT.

package schema

import (
	"encoding/xml"
)

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

// CTGraphicFrameNonVisual ...
type CTGraphicFrameNonVisual struct {
	XMLName           xml.Name                           `xml:"CT_GraphicFrameNonVisual"`
	CNvPr             *CTNonVisualDrawingProps           `xml:"cNvPr"`
	CNvGraphicFramePr *CTNonVisualGraphicFrameProperties `xml:"cNvGraphicFramePr"`
}

// CTGraphicFrame ...
type CTGraphicFrame struct {
	XMLName          xml.Name                 `xml:"CT_GraphicFrame"`
	MacroAttr        string                   `xml:"macro,attr,omitempty"`
	FPublishedAttr   bool                     `xml:"fPublished,attr,omitempty"`
	NvGraphicFramePr *CTGraphicFrameNonVisual `xml:"nvGraphicFramePr"`
	Xfrm             *CTTransform2D           `xml:"xfrm"`
	AGraphic         *CTGraphicalObject       `xml:"a:graphic"`
}

// CTGroupShapeNonVisual ...
type CTGroupShapeNonVisual struct {
	XMLName    xml.Name                           `xml:"CT_GroupShapeNonVisual"`
	CNvPr      *CTNonVisualDrawingProps           `xml:"cNvPr"`
	CNvGrpSpPr *CTNonVisualGroupDrawingShapeProps `xml:"cNvGrpSpPr"`
}

// CTGroupShape ...
type CTGroupShape struct {
	XMLName      xml.Name                `xml:"CT_GroupShape"`
	NvGrpSpPr    *CTGroupShapeNonVisual  `xml:"nvGrpSpPr"`
	GrpSpPr      *CTGroupShapeProperties `xml:"grpSpPr"`
	Sp           []*CTShape              `xml:"sp"`
	GrpSp        []*CTGroupShape         `xml:"grpSp"`
	GraphicFrame []*CTGraphicFrame       `xml:"graphicFrame"`
	CxnSp        []*CTConnector          `xml:"cxnSp"`
	Pic          []*CTPicture            `xml:"pic"`
}

// EGObjectChoices ...
type EGObjectChoices struct {
	XMLName      xml.Name `xml:"EG_ObjectChoices"`
	Sp           *CTShape
	GrpSp        *CTGroupShape
	GraphicFrame *CTGraphicFrame
	CxnSp        *CTConnector
	Pic          *CTPicture
}

// STMarkerCoordinate ...
type STMarkerCoordinate float64

// CTMarker ...
type CTMarker struct {
	XMLName xml.Name `xml:"CT_Marker"`
	X       float64  `xml:"x"`
	Y       float64  `xml:"y"`
}

// CTRelSizeAnchor ...
type CTRelSizeAnchor struct {
	XMLName         xml.Name `xml:"CT_RelSizeAnchor"`
	EGObjectChoices *EGObjectChoices
	From            *CTMarker `xml:"from"`
	To              *CTMarker `xml:"to"`
}

// CTAbsSizeAnchor ...
type CTAbsSizeAnchor struct {
	XMLName         xml.Name `xml:"CT_AbsSizeAnchor"`
	EGObjectChoices *EGObjectChoices
	From            *CTMarker         `xml:"from"`
	Ext             *CTPositiveSize2D `xml:"ext"`
}

// EGAnchor ...
type EGAnchor struct {
	XMLName       xml.Name `xml:"EG_Anchor"`
	RelSizeAnchor *CTRelSizeAnchor
	AbsSizeAnchor *CTAbsSizeAnchor
}

// CTDrawing ...
type CTDrawing struct {
	XMLName  xml.Name `xml:"CT_Drawing"`
	EGAnchor []*EGAnchor
}

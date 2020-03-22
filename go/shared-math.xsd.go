// Copyright 2020 The xgen Authors. All rights reserved.
//
// DO NOT EDIT: generated by xgen XSD generator
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package schema

import (
	"encoding/xml"
)

// STInteger255 ...
type STInteger255 int

// CTInteger255 ...
type CTInteger255 struct {
	XMLName xml.Name `xml:"CT_Integer255"`
	ValAttr int      `xml:"val,attr"`
}

// STInteger2 ...
type STInteger2 int

// CTInteger2 ...
type CTInteger2 struct {
	XMLName xml.Name `xml:"CT_Integer2"`
	ValAttr int      `xml:"val,attr"`
}

// STSpacingRule ...
type STSpacingRule int

// CTSpacingRule ...
type CTSpacingRule struct {
	XMLName xml.Name `xml:"CT_SpacingRule"`
	ValAttr int      `xml:"val,attr"`
}

// STUnSignedInteger ...
type STUnSignedInteger uint32

// CTUnSignedInteger ...
type CTUnSignedInteger struct {
	XMLName xml.Name `xml:"CT_UnSignedInteger"`
	ValAttr uint32   `xml:"val,attr"`
}

// STChar ...
type STChar string

// CTChar ...
type CTChar struct {
	XMLName xml.Name `xml:"CT_Char"`
	ValAttr string   `xml:"val,attr"`
}

// CTOnOff ...
type CTOnOff struct {
	XMLName xml.Name `xml:"CT_OnOff"`
	ValAttr *STOnOff `xml:"val,attr,omitempty"`
}

// CTString ...
type CTString struct {
	XMLName xml.Name `xml:"CT_String"`
	ValAttr string   `xml:"val,attr,omitempty"`
}

// CTXAlign ...
type CTXAlign struct {
	XMLName xml.Name `xml:"CT_XAlign"`
	ValAttr string   `xml:"val,attr"`
}

// CTYAlign ...
type CTYAlign struct {
	XMLName xml.Name `xml:"CT_YAlign"`
	ValAttr string   `xml:"val,attr"`
}

// STShp ...
type STShp string

// CTShp ...
type CTShp struct {
	XMLName xml.Name `xml:"CT_Shp"`
	ValAttr string   `xml:"val,attr"`
}

// STFType ...
type STFType string

// CTFType ...
type CTFType struct {
	XMLName xml.Name `xml:"CT_FType"`
	ValAttr string   `xml:"val,attr"`
}

// STLimLoc ...
type STLimLoc string

// CTLimLoc ...
type CTLimLoc struct {
	XMLName xml.Name `xml:"CT_LimLoc"`
	ValAttr string   `xml:"val,attr"`
}

// STTopBot ...
type STTopBot string

// CTTopBot ...
type CTTopBot struct {
	XMLName xml.Name `xml:"CT_TopBot"`
	ValAttr string   `xml:"val,attr"`
}

// STScript ...
type STScript string

// CTScript ...
type CTScript struct {
	XMLName xml.Name `xml:"CT_Script"`
	ValAttr string   `xml:"val,attr,omitempty"`
}

// STStyle ...
type STStyle string

// CTStyle ...
type CTStyle struct {
	XMLName xml.Name `xml:"CT_Style"`
	ValAttr string   `xml:"val,attr,omitempty"`
}

// CTManualBreak ...
type CTManualBreak struct {
	XMLName   xml.Name `xml:"CT_ManualBreak"`
	AlnAtAttr int      `xml:"alnAt,attr,omitempty"`
}

// EGScriptStyle ...
type EGScriptStyle struct {
	XMLName xml.Name `xml:"EG_ScriptStyle"`
	Scr     *CTScript
	Sty     *CTStyle
}

// CTRPR ...
type CTRPR struct {
	XMLName       xml.Name `xml:"CT_RPR"`
	EGScriptStyle *EGScriptStyle
	Lit           *CTOnOff       `xml:"lit"`
	Nor           *CTOnOff       `xml:"nor"`
	Brk           *CTManualBreak `xml:"brk"`
	Aln           *CTOnOff       `xml:"aln"`
}

// CTText ...
type CTText struct {
	XMLName      xml.Name `xml:"CT_Text"`
	XmlSpaceAttr *Space   `xml:"xml:space,attr,omitempty"`
}

// CTR ...
type CTR struct {
	XMLName            xml.Name `xml:"CT_R"`
	WEGRPr             *EGRPr
	WEGRunInnerContent *EGRunInnerContent
	RPr                *CTRPR  `xml:"rPr"`
	T                  *CTText `xml:"t"`
}

// CTCtrlPr ...
type CTCtrlPr struct {
	XMLName    xml.Name `xml:"CT_CtrlPr"`
	WEGRPrMath *EGRPrMath
}

// CTAccPr ...
type CTAccPr struct {
	XMLName xml.Name  `xml:"CT_AccPr"`
	Chr     *CTChar   `xml:"chr"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTAcc ...
type CTAcc struct {
	XMLName xml.Name    `xml:"CT_Acc"`
	AccPr   *CTAccPr    `xml:"accPr"`
	E       *CTOMathArg `xml:"e"`
}

// CTBarPr ...
type CTBarPr struct {
	XMLName xml.Name  `xml:"CT_BarPr"`
	Pos     *CTTopBot `xml:"pos"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTBar ...
type CTBar struct {
	XMLName xml.Name    `xml:"CT_Bar"`
	BarPr   *CTBarPr    `xml:"barPr"`
	E       *CTOMathArg `xml:"e"`
}

// CTBoxPr ...
type CTBoxPr struct {
	XMLName xml.Name       `xml:"CT_BoxPr"`
	OpEmu   *CTOnOff       `xml:"opEmu"`
	NoBreak *CTOnOff       `xml:"noBreak"`
	Diff    *CTOnOff       `xml:"diff"`
	Brk     *CTManualBreak `xml:"brk"`
	Aln     *CTOnOff       `xml:"aln"`
	CtrlPr  *CTCtrlPr      `xml:"ctrlPr"`
}

// CTBox ...
type CTBox struct {
	XMLName xml.Name    `xml:"CT_Box"`
	BoxPr   *CTBoxPr    `xml:"boxPr"`
	E       *CTOMathArg `xml:"e"`
}

// CTBorderBoxPr ...
type CTBorderBoxPr struct {
	XMLName    xml.Name  `xml:"CT_BorderBoxPr"`
	HideTop    *CTOnOff  `xml:"hideTop"`
	HideBot    *CTOnOff  `xml:"hideBot"`
	HideLeft   *CTOnOff  `xml:"hideLeft"`
	HideRight  *CTOnOff  `xml:"hideRight"`
	StrikeH    *CTOnOff  `xml:"strikeH"`
	StrikeV    *CTOnOff  `xml:"strikeV"`
	StrikeBLTR *CTOnOff  `xml:"strikeBLTR"`
	StrikeTLBR *CTOnOff  `xml:"strikeTLBR"`
	CtrlPr     *CTCtrlPr `xml:"ctrlPr"`
}

// CTBorderBox ...
type CTBorderBox struct {
	XMLName     xml.Name       `xml:"CT_BorderBox"`
	BorderBoxPr *CTBorderBoxPr `xml:"borderBoxPr"`
	E           *CTOMathArg    `xml:"e"`
}

// CTDPr ...
type CTDPr struct {
	XMLName xml.Name  `xml:"CT_DPr"`
	BegChr  *CTChar   `xml:"begChr"`
	SepChr  *CTChar   `xml:"sepChr"`
	EndChr  *CTChar   `xml:"endChr"`
	Grow    *CTOnOff  `xml:"grow"`
	Shp     *CTShp    `xml:"shp"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTD ...
type CTD struct {
	XMLName xml.Name      `xml:"CT_D"`
	DPr     *CTDPr        `xml:"dPr"`
	E       []*CTOMathArg `xml:"e"`
}

// CTEqArrPr ...
type CTEqArrPr struct {
	XMLName xml.Name           `xml:"CT_EqArrPr"`
	BaseJc  *CTYAlign          `xml:"baseJc"`
	MaxDist *CTOnOff           `xml:"maxDist"`
	ObjDist *CTOnOff           `xml:"objDist"`
	RSpRule *CTSpacingRule     `xml:"rSpRule"`
	RSp     *CTUnSignedInteger `xml:"rSp"`
	CtrlPr  *CTCtrlPr          `xml:"ctrlPr"`
}

// CTEqArr ...
type CTEqArr struct {
	XMLName xml.Name      `xml:"CT_EqArr"`
	EqArrPr *CTEqArrPr    `xml:"eqArrPr"`
	E       []*CTOMathArg `xml:"e"`
}

// CTFPr ...
type CTFPr struct {
	XMLName xml.Name  `xml:"CT_FPr"`
	Type    *CTFType  `xml:"type"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTF ...
type CTF struct {
	XMLName xml.Name    `xml:"CT_F"`
	FPr     *CTFPr      `xml:"fPr"`
	Num     *CTOMathArg `xml:"num"`
	Den     *CTOMathArg `xml:"den"`
}

// CTFuncPr ...
type CTFuncPr struct {
	XMLName xml.Name  `xml:"CT_FuncPr"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTFunc ...
type CTFunc struct {
	XMLName xml.Name    `xml:"CT_Func"`
	FuncPr  *CTFuncPr   `xml:"funcPr"`
	FName   *CTOMathArg `xml:"fName"`
	E       *CTOMathArg `xml:"e"`
}

// CTGroupChrPr ...
type CTGroupChrPr struct {
	XMLName xml.Name  `xml:"CT_GroupChrPr"`
	Chr     *CTChar   `xml:"chr"`
	Pos     *CTTopBot `xml:"pos"`
	VertJc  *CTTopBot `xml:"vertJc"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTGroupChr ...
type CTGroupChr struct {
	XMLName    xml.Name      `xml:"CT_GroupChr"`
	GroupChrPr *CTGroupChrPr `xml:"groupChrPr"`
	E          *CTOMathArg   `xml:"e"`
}

// CTLimLowPr ...
type CTLimLowPr struct {
	XMLName xml.Name  `xml:"CT_LimLowPr"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTLimLow ...
type CTLimLow struct {
	XMLName  xml.Name    `xml:"CT_LimLow"`
	LimLowPr *CTLimLowPr `xml:"limLowPr"`
	E        *CTOMathArg `xml:"e"`
	Lim      *CTOMathArg `xml:"lim"`
}

// CTLimUppPr ...
type CTLimUppPr struct {
	XMLName xml.Name  `xml:"CT_LimUppPr"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTLimUpp ...
type CTLimUpp struct {
	XMLName  xml.Name    `xml:"CT_LimUpp"`
	LimUppPr *CTLimUppPr `xml:"limUppPr"`
	E        *CTOMathArg `xml:"e"`
	Lim      *CTOMathArg `xml:"lim"`
}

// CTMCPr ...
type CTMCPr struct {
	XMLName xml.Name      `xml:"CT_MCPr"`
	Count   *CTInteger255 `xml:"count"`
	McJc    *CTXAlign     `xml:"mcJc"`
}

// CTMC ...
type CTMC struct {
	XMLName xml.Name `xml:"CT_MC"`
	McPr    *CTMCPr  `xml:"mcPr"`
}

// CTMCS ...
type CTMCS struct {
	XMLName xml.Name `xml:"CT_MCS"`
	Mc      []*CTMC  `xml:"mc"`
}

// CTMPr ...
type CTMPr struct {
	XMLName xml.Name           `xml:"CT_MPr"`
	BaseJc  *CTYAlign          `xml:"baseJc"`
	PlcHide *CTOnOff           `xml:"plcHide"`
	RSpRule *CTSpacingRule     `xml:"rSpRule"`
	CGpRule *CTSpacingRule     `xml:"cGpRule"`
	RSp     *CTUnSignedInteger `xml:"rSp"`
	CSp     *CTUnSignedInteger `xml:"cSp"`
	CGp     *CTUnSignedInteger `xml:"cGp"`
	Mcs     *CTMCS             `xml:"mcs"`
	CtrlPr  *CTCtrlPr          `xml:"ctrlPr"`
}

// CTMR ...
type CTMR struct {
	XMLName xml.Name      `xml:"CT_MR"`
	E       []*CTOMathArg `xml:"e"`
}

// CTM ...
type CTM struct {
	XMLName xml.Name `xml:"CT_M"`
	MPr     *CTMPr   `xml:"mPr"`
	Mr      []*CTMR  `xml:"mr"`
}

// CTNaryPr ...
type CTNaryPr struct {
	XMLName xml.Name  `xml:"CT_NaryPr"`
	Chr     *CTChar   `xml:"chr"`
	LimLoc  *CTLimLoc `xml:"limLoc"`
	Grow    *CTOnOff  `xml:"grow"`
	SubHide *CTOnOff  `xml:"subHide"`
	SupHide *CTOnOff  `xml:"supHide"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTNary ...
type CTNary struct {
	XMLName xml.Name    `xml:"CT_Nary"`
	NaryPr  *CTNaryPr   `xml:"naryPr"`
	Sub     *CTOMathArg `xml:"sub"`
	Sup     *CTOMathArg `xml:"sup"`
	E       *CTOMathArg `xml:"e"`
}

// CTPhantPr ...
type CTPhantPr struct {
	XMLName  xml.Name  `xml:"CT_PhantPr"`
	Show     *CTOnOff  `xml:"show"`
	ZeroWid  *CTOnOff  `xml:"zeroWid"`
	ZeroAsc  *CTOnOff  `xml:"zeroAsc"`
	ZeroDesc *CTOnOff  `xml:"zeroDesc"`
	Transp   *CTOnOff  `xml:"transp"`
	CtrlPr   *CTCtrlPr `xml:"ctrlPr"`
}

// CTPhant ...
type CTPhant struct {
	XMLName xml.Name    `xml:"CT_Phant"`
	PhantPr *CTPhantPr  `xml:"phantPr"`
	E       *CTOMathArg `xml:"e"`
}

// CTRadPr ...
type CTRadPr struct {
	XMLName xml.Name  `xml:"CT_RadPr"`
	DegHide *CTOnOff  `xml:"degHide"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTRad ...
type CTRad struct {
	XMLName xml.Name    `xml:"CT_Rad"`
	RadPr   *CTRadPr    `xml:"radPr"`
	Deg     *CTOMathArg `xml:"deg"`
	E       *CTOMathArg `xml:"e"`
}

// CTSPrePr ...
type CTSPrePr struct {
	XMLName xml.Name  `xml:"CT_SPrePr"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTSPre ...
type CTSPre struct {
	XMLName xml.Name    `xml:"CT_SPre"`
	SPrePr  *CTSPrePr   `xml:"sPrePr"`
	Sub     *CTOMathArg `xml:"sub"`
	Sup     *CTOMathArg `xml:"sup"`
	E       *CTOMathArg `xml:"e"`
}

// CTSSubPr ...
type CTSSubPr struct {
	XMLName xml.Name  `xml:"CT_SSubPr"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTSSub ...
type CTSSub struct {
	XMLName xml.Name    `xml:"CT_SSub"`
	SSubPr  *CTSSubPr   `xml:"sSubPr"`
	E       *CTOMathArg `xml:"e"`
	Sub     *CTOMathArg `xml:"sub"`
}

// CTSSubSupPr ...
type CTSSubSupPr struct {
	XMLName xml.Name  `xml:"CT_SSubSupPr"`
	AlnScr  *CTOnOff  `xml:"alnScr"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTSSubSup ...
type CTSSubSup struct {
	XMLName   xml.Name     `xml:"CT_SSubSup"`
	SSubSupPr *CTSSubSupPr `xml:"sSubSupPr"`
	E         *CTOMathArg  `xml:"e"`
	Sub       *CTOMathArg  `xml:"sub"`
	Sup       *CTOMathArg  `xml:"sup"`
}

// CTSSupPr ...
type CTSSupPr struct {
	XMLName xml.Name  `xml:"CT_SSupPr"`
	CtrlPr  *CTCtrlPr `xml:"ctrlPr"`
}

// CTSSup ...
type CTSSup struct {
	XMLName xml.Name    `xml:"CT_SSup"`
	SSupPr  *CTSSupPr   `xml:"sSupPr"`
	E       *CTOMathArg `xml:"e"`
	Sup     *CTOMathArg `xml:"sup"`
}

// EGOMathMathElements ...
type EGOMathMathElements struct {
	XMLName   xml.Name `xml:"EG_OMathMathElements"`
	Acc       *CTAcc
	Bar       *CTBar
	Box       *CTBox
	BorderBox *CTBorderBox
	D         *CTD
	EqArr     *CTEqArr
	F         *CTF
	Func      *CTFunc
	GroupChr  *CTGroupChr
	LimLow    *CTLimLow
	LimUpp    *CTLimUpp
	M         *CTM
	Nary      *CTNary
	Phant     *CTPhant
	Rad       *CTRad
	SPre      *CTSPre
	SSub      *CTSSub
	SSubSup   *CTSSubSup
	SSup      *CTSSup
	R         *CTR
}

// EGOMathElements ...
type EGOMathElements struct {
	XMLName             xml.Name `xml:"EG_OMathElements"`
	EGOMathMathElements *EGOMathMathElements
	WEGPContentMath     *EGPContentMath
}

// CTOMathArgPr ...
type CTOMathArgPr struct {
	XMLName xml.Name    `xml:"CT_OMathArgPr"`
	ArgSz   *CTInteger2 `xml:"argSz"`
}

// CTOMathArg ...
type CTOMathArg struct {
	XMLName         xml.Name `xml:"CT_OMathArg"`
	EGOMathElements []*EGOMathElements
	ArgPr           *CTOMathArgPr `xml:"argPr"`
	CtrlPr          *CTCtrlPr     `xml:"ctrlPr"`
}

// STJc ...
type STJc string

// CTOMathJc ...
type CTOMathJc struct {
	XMLName xml.Name `xml:"CT_OMathJc"`
	ValAttr string   `xml:"val,attr,omitempty"`
}

// CTOMathParaPr ...
type CTOMathParaPr struct {
	XMLName xml.Name   `xml:"CT_OMathParaPr"`
	Jc      *CTOMathJc `xml:"jc"`
}

// CTTwipsMeasure ...
type CTTwipsMeasure struct {
	XMLName xml.Name        `xml:"CT_TwipsMeasure"`
	ValAttr *STTwipsMeasure `xml:"val,attr"`
}

// STBreakBin ...
type STBreakBin string

// CTBreakBin ...
type CTBreakBin struct {
	XMLName xml.Name `xml:"CT_BreakBin"`
	ValAttr string   `xml:"val,attr,omitempty"`
}

// STBreakBinSub ...
type STBreakBinSub string

// CTBreakBinSub ...
type CTBreakBinSub struct {
	XMLName xml.Name `xml:"CT_BreakBinSub"`
	ValAttr string   `xml:"val,attr,omitempty"`
}

// CTMathPr ...
type CTMathPr struct {
	XMLName    xml.Name        `xml:"CT_MathPr"`
	MathFont   *CTString       `xml:"mathFont"`
	BrkBin     *CTBreakBin     `xml:"brkBin"`
	BrkBinSub  *CTBreakBinSub  `xml:"brkBinSub"`
	SmallFrac  *CTOnOff        `xml:"smallFrac"`
	DispDef    *CTOnOff        `xml:"dispDef"`
	LMargin    *CTTwipsMeasure `xml:"lMargin"`
	RMargin    *CTTwipsMeasure `xml:"rMargin"`
	DefJc      *CTOMathJc      `xml:"defJc"`
	PreSp      *CTTwipsMeasure `xml:"preSp"`
	PostSp     *CTTwipsMeasure `xml:"postSp"`
	InterSp    *CTTwipsMeasure `xml:"interSp"`
	IntraSp    *CTTwipsMeasure `xml:"intraSp"`
	WrapIndent *CTTwipsMeasure `xml:"wrapIndent"`
	WrapRight  *CTOnOff        `xml:"wrapRight"`
	IntLim     *CTLimLoc       `xml:"intLim"`
	NaryLim    *CTLimLoc       `xml:"naryLim"`
}

// MathPr ...
type MathPr *CTMathPr

// CTOMathPara ...
type CTOMathPara struct {
	XMLName     xml.Name       `xml:"CT_OMathPara"`
	OMathParaPr *CTOMathParaPr `xml:"oMathParaPr"`
	OMath       []*CTOMath     `xml:"oMath"`
}

// CTOMath ...
type CTOMath struct {
	XMLName         xml.Name `xml:"CT_OMath"`
	EGOMathElements []*EGOMathElements
}

// OMathPara ...
type OMathPara *CTOMathPara

// OMath ...
type OMath *CTOMath

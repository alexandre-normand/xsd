// Copyright 2020 The xgen Authors. All rights reserved.
//
// DO NOT EDIT: generated by xgen XSD generator
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

export type Properties = CT_Properties;

export class CT_Properties {
	Property: Array<CT_Property>;
}

export class CT_Property {
	FmtidAttr: string;
	PidAttr: number;
	NameAttr: string | null;
	LinkTargetAttr: string | null;
	VtVector: Array<CT_Vector>;
	VtArray: Array<CT_Array>;
	VtBlob: Array<Array<any>>;
	VtOblob: Array<Array<any>>;
	VtEmpty: Array<CT_Empty>;
	VtNull: Array<CT_Null>;
	VtI1: Array<Any>;
	VtI2: Array<number>;
	VtI4: Array<number>;
	VtI8: Array<number>;
	VtInt: Array<number>;
	VtUi1: Array<Any>;
	VtUi2: Array<number>;
	VtUi4: Array<number>;
	VtUi8: Array<number>;
	VtUint: Array<number>;
	VtR4: Array<number>;
	VtR8: Array<number>;
	VtDecimal: Array<number>;
	VtLpstr: Array<string>;
	VtLpwstr: Array<string>;
	VtBstr: Array<string>;
	VtDate: Array<string>;
	VtFiletime: Array<string>;
	VtBool: Array<boolean>;
	VtCy: Array<string>;
	VtError: Array<string>;
	VtStream: Array<Array<any>>;
	VtOstream: Array<Array<any>>;
	VtStorage: Array<Array<any>>;
	VtOstorage: Array<Array<any>>;
	VtVstream: Array<CT_Vstream>;
	VtClsid: Array<string>;
}
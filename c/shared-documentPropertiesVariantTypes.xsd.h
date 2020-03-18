// Copyright 2020 The xgen Authors. All rights reserved.
//
// DO NOT EDIT: generated by xgen XSD generator
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

typedef char ST_VectorBaseType;

typedef char ST_ArrayBaseType;

typedef char ST_Cy;

typedef char ST_Error;

typedef struct {
} CT_Empty;

typedef struct {
} CT_Null;

typedef struct {
	char BaseTypeAttr; // attr
	unsigned int SizeAttr; // attr
	CT_Variant Variant;
	char I1[];
	int I2;
	int I4;
	int I8;
	char Ui1;
	unsigned int Ui2;
	unsigned int Ui4;
	unsigned int Ui8;
	float R4;
	float R8;
	char Lpstr;
	char Lpwstr;
	char Bstr;
	char Date;
	char Filetime;
	bool Bool;
	char Cy;
	char Error;
	char Clsid;
} CT_Vector;

typedef struct {
	int LBoundsAttr; // attr
	int UBoundsAttr; // attr
	char BaseTypeAttr; // attr
	CT_Variant Variant;
	char I1[];
	int I2;
	int I4;
	int Int;
	char Ui1;
	unsigned int Ui2;
	unsigned int Ui4;
	unsigned int Uint;
	float R4;
	float R8;
	float Decimal;
	char Bstr;
	char Date;
	bool Bool;
	char Error;
	char Cy;
} CT_Array;

typedef struct {
	CT_Variant Variant;
	CT_Vector Vector;
	CT_Array Array;
	char Blob[];
	char Oblob[];
	CT_Empty Empty;
	CT_Null Null;
	char I1[];
	int I2;
	int I4;
	int I8;
	int Int;
	char Ui1;
	unsigned int Ui2;
	unsigned int Ui4;
	unsigned int Ui8;
	unsigned int Uint;
	float R4;
	float R8;
	float Decimal;
	char Lpstr;
	char Lpwstr;
	char Bstr;
	char Date;
	char Filetime;
	bool Bool;
	char Cy;
	char Error;
	char Stream[];
	char Ostream[];
	char Storage[];
	char Ostorage[];
	CT_Vstream Vstream;
	char Clsid;
} CT_Variant;

typedef struct {
	char VersionAttr; // attr, optional
} CT_Vstream;

typedef CT_Variant Variant;

typedef CT_Vector Vector;

typedef CT_Array Array;

typedef char Blob[];

typedef char Oblob[];

typedef CT_Empty Empty;

typedef CT_Null Null;

typedef char I1[];

typedef int I2;

typedef int I4;

typedef int I8;

typedef int Int;

typedef char Ui1;

typedef unsigned int Ui2;

typedef unsigned int Ui4;

typedef unsigned int Ui8;

typedef unsigned int Uint;

typedef float R4;

typedef float R8;

typedef float Decimal;

typedef char Lpstr;

typedef char Lpwstr;

typedef char Bstr;

typedef char Date;

typedef char Filetime;

typedef bool Bool;

typedef char Cy;

typedef char Error;

typedef char Stream[];

typedef char Ostream[];

typedef char Storage[];

typedef char Ostorage[];

typedef CT_Vstream Vstream;

typedef char Clsid;

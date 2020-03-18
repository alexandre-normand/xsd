// Copyright 2020 The xgen Authors. All rights reserved.
//
// DO NOT EDIT: generated by xgen XSD generator
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

export type ST_SourceType = string;

export class CT_NameListType {
	Person: Array<CT_PersonType>;
}

export class CT_PersonType {
	Last: Array<string>;
	First: Array<string>;
	Middle: Array<string>;
}

export class CT_NameType {
	NameList: Array<CT_NameListType>;
}

export class CT_NameOrCorporateType {
	NameList: Array<CT_NameListType>;
	Corporate: Array<string>;
}

export class CT_AuthorType {
	Artist: Array<CT_NameType>;
	Author: Array<CT_NameOrCorporateType>;
	BookAuthor: Array<CT_NameType>;
	Compiler: Array<CT_NameType>;
	Composer: Array<CT_NameType>;
	Conductor: Array<CT_NameType>;
	Counsel: Array<CT_NameType>;
	Director: Array<CT_NameType>;
	Editor: Array<CT_NameType>;
	Interviewee: Array<CT_NameType>;
	Interviewer: Array<CT_NameType>;
	Inventor: Array<CT_NameType>;
	Performer: Array<CT_NameOrCorporateType>;
	ProducerName: Array<CT_NameType>;
	Translator: Array<CT_NameType>;
	Writer: Array<CT_NameType>;
}

export class CT_SourceType {
	AbbreviatedCaseNumber: Array<string>;
	AlbumTitle: Array<string>;
	Author: Array<CT_AuthorType>;
	BookTitle: Array<string>;
	Broadcaster: Array<string>;
	BroadcastTitle: Array<string>;
	CaseNumber: Array<string>;
	ChapterNumber: Array<string>;
	City: Array<string>;
	Comments: Array<string>;
	ConferenceName: Array<string>;
	CountryRegion: Array<string>;
	Court: Array<string>;
	Day: Array<string>;
	DayAccessed: Array<string>;
	Department: Array<string>;
	Distributor: Array<string>;
	Edition: Array<string>;
	Guid: Array<string>;
	Institution: Array<string>;
	InternetSiteTitle: Array<string>;
	Issue: Array<string>;
	JournalName: Array<string>;
	LCID: Array<string>;
	Medium: Array<string>;
	Month: Array<string>;
	MonthAccessed: Array<string>;
	NumberVolumes: Array<string>;
	Pages: Array<string>;
	PatentNumber: Array<string>;
	PeriodicalTitle: Array<string>;
	ProductionCompany: Array<string>;
	PublicationTitle: Array<string>;
	Publisher: Array<string>;
	RecordingNumber: Array<string>;
	RefOrder: Array<string>;
	Reporter: Array<string>;
	SourceType: Array<string>;
	ShortTitle: Array<string>;
	StandardNumber: Array<string>;
	StateProvince: Array<string>;
	Station: Array<string>;
	Tag: Array<string>;
	Theater: Array<string>;
	ThesisType: Array<string>;
	Title: Array<string>;
	Type: Array<string>;
	URL: Array<string>;
	Version: Array<string>;
	Volume: Array<string>;
	Year: Array<string>;
	YearAccessed: Array<string>;
}

export type Sources = CT_Sources;

export class CT_Sources {
	SelectedStyleAttr: string | null;
	StyleNameAttr: string | null;
	URIAttr: string | null;
	Source: Array<CT_SourceType>;
}
// Code generated from ApiParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package api // ApiParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ApiParserParser struct {
	*antlr.BaseParser
}

var apiparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func apiparserParserInit() {
	staticData := &apiparserParserStaticData
	staticData.literalNames = []string{
		"", "'='", "'('", "')'", "'{'", "'}'", "'*'", "'time.Time'", "'snowflake.ID'",
		"'['", "']'", "'returns'", "'-'", "'/'", "'/:'", "'@doc'", "'@handler'",
		"'interface{}'", "'@server'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ATDOC",
		"ATHANDLER", "INTERFACE", "ATSERVER", "WS", "COMMENT", "LINE_COMMENT",
		"STRING", "RAW_STRING", "LINE_VALUE", "ID",
	}
	staticData.ruleNames = []string{
		"api", "spec", "syntaxLit", "importSpec", "importLit", "importBlock",
		"importBlockValue", "importValue", "infoSpec", "typeSpec", "typeLit",
		"typeBlock", "typeLitBody", "typeBlockBody", "typeStruct", "typeAlias",
		"typeBlockStruct", "typeBlockAlias", "field", "normalField", "anonymousFiled",
		"dataType", "pointerType", "mapType", "arrayType", "serviceSpec", "atServer",
		"serviceApi", "serviceRoute", "atDoc", "atHandler", "route", "body",
		"replybody", "kvLit", "serviceName", "path",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 348, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		1, 0, 5, 0, 76, 8, 0, 10, 0, 12, 0, 79, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 86, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3,
		96, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 106, 8,
		5, 11, 5, 12, 5, 107, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 4, 8, 121, 8, 8, 11, 8, 12, 8, 122, 1, 8, 1, 8, 1, 9, 1,
		9, 3, 9, 129, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 5, 11, 139, 8, 11, 10, 11, 12, 11, 142, 9, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 3, 12, 148, 8, 12, 1, 13, 1, 13, 3, 13, 152, 8, 13, 1, 14, 1, 14,
		1, 14, 3, 14, 157, 8, 14, 1, 14, 1, 14, 5, 14, 161, 8, 14, 10, 14, 12,
		14, 164, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15, 171, 8, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 3, 16, 178, 8, 16, 1, 16, 1, 16, 5, 16,
		182, 8, 16, 10, 16, 12, 16, 185, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		17, 3, 17, 192, 8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 3, 18, 199, 8,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 205, 8, 19, 1, 20, 3, 20, 208, 8,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 3, 21, 221, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25,
		3, 25, 240, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 4, 26, 247, 8, 26,
		11, 26, 12, 26, 248, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5,
		27, 258, 8, 27, 10, 27, 12, 27, 261, 9, 27, 1, 27, 1, 27, 1, 28, 3, 28,
		266, 8, 28, 1, 28, 1, 28, 3, 28, 270, 8, 28, 1, 28, 1, 28, 1, 29, 1, 29,
		3, 29, 276, 8, 29, 1, 29, 4, 29, 279, 8, 29, 11, 29, 12, 29, 280, 1, 29,
		3, 29, 284, 8, 29, 1, 29, 3, 29, 287, 8, 29, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 3, 31, 296, 8, 31, 1, 31, 3, 31, 299, 8, 31, 1, 32,
		1, 32, 3, 32, 303, 8, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 3, 33, 310,
		8, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 3, 35, 320,
		8, 35, 4, 35, 322, 8, 35, 11, 35, 12, 35, 323, 1, 36, 1, 36, 1, 36, 1,
		36, 5, 36, 330, 8, 36, 10, 36, 12, 36, 333, 9, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 3, 36, 339, 8, 36, 4, 36, 341, 8, 36, 11, 36, 12, 36, 342, 1, 36,
		3, 36, 346, 8, 36, 1, 36, 0, 0, 37, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
		56, 58, 60, 62, 64, 66, 68, 70, 72, 0, 0, 358, 0, 77, 1, 0, 0, 0, 2, 85,
		1, 0, 0, 0, 4, 87, 1, 0, 0, 0, 6, 95, 1, 0, 0, 0, 8, 97, 1, 0, 0, 0, 10,
		101, 1, 0, 0, 0, 12, 111, 1, 0, 0, 0, 14, 113, 1, 0, 0, 0, 16, 116, 1,
		0, 0, 0, 18, 128, 1, 0, 0, 0, 20, 130, 1, 0, 0, 0, 22, 134, 1, 0, 0, 0,
		24, 147, 1, 0, 0, 0, 26, 151, 1, 0, 0, 0, 28, 153, 1, 0, 0, 0, 30, 167,
		1, 0, 0, 0, 32, 174, 1, 0, 0, 0, 34, 188, 1, 0, 0, 0, 36, 198, 1, 0, 0,
		0, 38, 200, 1, 0, 0, 0, 40, 207, 1, 0, 0, 0, 42, 220, 1, 0, 0, 0, 44, 222,
		1, 0, 0, 0, 46, 226, 1, 0, 0, 0, 48, 234, 1, 0, 0, 0, 50, 239, 1, 0, 0,
		0, 52, 243, 1, 0, 0, 0, 54, 252, 1, 0, 0, 0, 56, 265, 1, 0, 0, 0, 58, 273,
		1, 0, 0, 0, 60, 288, 1, 0, 0, 0, 62, 291, 1, 0, 0, 0, 64, 300, 1, 0, 0,
		0, 66, 306, 1, 0, 0, 0, 68, 313, 1, 0, 0, 0, 70, 321, 1, 0, 0, 0, 72, 345,
		1, 0, 0, 0, 74, 76, 3, 2, 1, 0, 75, 74, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0,
		77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 1, 1, 0, 0, 0, 79, 77, 1, 0,
		0, 0, 80, 86, 3, 4, 2, 0, 81, 86, 3, 6, 3, 0, 82, 86, 3, 16, 8, 0, 83,
		86, 3, 18, 9, 0, 84, 86, 3, 50, 25, 0, 85, 80, 1, 0, 0, 0, 85, 81, 1, 0,
		0, 0, 85, 82, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 84, 1, 0, 0, 0, 86, 3,
		1, 0, 0, 0, 87, 88, 6, 2, -1, 0, 88, 89, 5, 25, 0, 0, 89, 90, 5, 1, 0,
		0, 90, 91, 6, 2, -1, 0, 91, 92, 5, 22, 0, 0, 92, 5, 1, 0, 0, 0, 93, 96,
		3, 8, 4, 0, 94, 96, 3, 10, 5, 0, 95, 93, 1, 0, 0, 0, 95, 94, 1, 0, 0, 0,
		96, 7, 1, 0, 0, 0, 97, 98, 6, 4, -1, 0, 98, 99, 5, 25, 0, 0, 99, 100, 3,
		14, 7, 0, 100, 9, 1, 0, 0, 0, 101, 102, 6, 5, -1, 0, 102, 103, 5, 25, 0,
		0, 103, 105, 5, 2, 0, 0, 104, 106, 3, 12, 6, 0, 105, 104, 1, 0, 0, 0, 106,
		107, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109,
		1, 0, 0, 0, 109, 110, 5, 3, 0, 0, 110, 11, 1, 0, 0, 0, 111, 112, 3, 14,
		7, 0, 112, 13, 1, 0, 0, 0, 113, 114, 6, 7, -1, 0, 114, 115, 5, 22, 0, 0,
		115, 15, 1, 0, 0, 0, 116, 117, 6, 8, -1, 0, 117, 118, 5, 25, 0, 0, 118,
		120, 5, 2, 0, 0, 119, 121, 3, 68, 34, 0, 120, 119, 1, 0, 0, 0, 121, 122,
		1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 124, 1, 0,
		0, 0, 124, 125, 5, 3, 0, 0, 125, 17, 1, 0, 0, 0, 126, 129, 3, 20, 10, 0,
		127, 129, 3, 22, 11, 0, 128, 126, 1, 0, 0, 0, 128, 127, 1, 0, 0, 0, 129,
		19, 1, 0, 0, 0, 130, 131, 6, 10, -1, 0, 131, 132, 5, 25, 0, 0, 132, 133,
		3, 24, 12, 0, 133, 21, 1, 0, 0, 0, 134, 135, 6, 11, -1, 0, 135, 136, 5,
		25, 0, 0, 136, 140, 5, 2, 0, 0, 137, 139, 3, 26, 13, 0, 138, 137, 1, 0,
		0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0,
		141, 143, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 144, 5, 3, 0, 0, 144,
		23, 1, 0, 0, 0, 145, 148, 3, 28, 14, 0, 146, 148, 3, 30, 15, 0, 147, 145,
		1, 0, 0, 0, 147, 146, 1, 0, 0, 0, 148, 25, 1, 0, 0, 0, 149, 152, 3, 32,
		16, 0, 150, 152, 3, 34, 17, 0, 151, 149, 1, 0, 0, 0, 151, 150, 1, 0, 0,
		0, 152, 27, 1, 0, 0, 0, 153, 154, 6, 14, -1, 0, 154, 156, 5, 25, 0, 0,
		155, 157, 5, 25, 0, 0, 156, 155, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157,
		158, 1, 0, 0, 0, 158, 162, 5, 4, 0, 0, 159, 161, 3, 36, 18, 0, 160, 159,
		1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0,
		0, 0, 163, 165, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 166, 5, 5, 0, 0,
		166, 29, 1, 0, 0, 0, 167, 168, 6, 15, -1, 0, 168, 170, 5, 25, 0, 0, 169,
		171, 5, 1, 0, 0, 170, 169, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 172,
		1, 0, 0, 0, 172, 173, 3, 42, 21, 0, 173, 31, 1, 0, 0, 0, 174, 175, 6, 16,
		-1, 0, 175, 177, 5, 25, 0, 0, 176, 178, 5, 25, 0, 0, 177, 176, 1, 0, 0,
		0, 177, 178, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 183, 5, 4, 0, 0, 180,
		182, 3, 36, 18, 0, 181, 180, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 181,
		1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 186, 1, 0, 0, 0, 185, 183, 1, 0,
		0, 0, 186, 187, 5, 5, 0, 0, 187, 33, 1, 0, 0, 0, 188, 189, 6, 17, -1, 0,
		189, 191, 5, 25, 0, 0, 190, 192, 5, 1, 0, 0, 191, 190, 1, 0, 0, 0, 191,
		192, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 3, 42, 21, 0, 194, 35,
		1, 0, 0, 0, 195, 196, 4, 18, 0, 0, 196, 199, 3, 38, 19, 0, 197, 199, 3,
		40, 20, 0, 198, 195, 1, 0, 0, 0, 198, 197, 1, 0, 0, 0, 199, 37, 1, 0, 0,
		0, 200, 201, 6, 19, -1, 0, 201, 202, 5, 25, 0, 0, 202, 204, 3, 42, 21,
		0, 203, 205, 5, 23, 0, 0, 204, 203, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205,
		39, 1, 0, 0, 0, 206, 208, 5, 6, 0, 0, 207, 206, 1, 0, 0, 0, 207, 208, 1,
		0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 5, 25, 0, 0, 210, 41, 1, 0, 0,
		0, 211, 212, 6, 21, -1, 0, 212, 221, 5, 25, 0, 0, 213, 221, 3, 46, 23,
		0, 214, 221, 3, 48, 24, 0, 215, 221, 5, 17, 0, 0, 216, 221, 5, 7, 0, 0,
		217, 221, 5, 8, 0, 0, 218, 221, 3, 44, 22, 0, 219, 221, 3, 28, 14, 0, 220,
		211, 1, 0, 0, 0, 220, 213, 1, 0, 0, 0, 220, 214, 1, 0, 0, 0, 220, 215,
		1, 0, 0, 0, 220, 216, 1, 0, 0, 0, 220, 217, 1, 0, 0, 0, 220, 218, 1, 0,
		0, 0, 220, 219, 1, 0, 0, 0, 221, 43, 1, 0, 0, 0, 222, 223, 5, 6, 0, 0,
		223, 224, 6, 22, -1, 0, 224, 225, 5, 25, 0, 0, 225, 45, 1, 0, 0, 0, 226,
		227, 6, 23, -1, 0, 227, 228, 5, 25, 0, 0, 228, 229, 5, 9, 0, 0, 229, 230,
		6, 23, -1, 0, 230, 231, 5, 25, 0, 0, 231, 232, 5, 10, 0, 0, 232, 233, 3,
		42, 21, 0, 233, 47, 1, 0, 0, 0, 234, 235, 5, 9, 0, 0, 235, 236, 5, 10,
		0, 0, 236, 237, 3, 42, 21, 0, 237, 49, 1, 0, 0, 0, 238, 240, 3, 52, 26,
		0, 239, 238, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241,
		242, 3, 54, 27, 0, 242, 51, 1, 0, 0, 0, 243, 244, 5, 18, 0, 0, 244, 246,
		5, 2, 0, 0, 245, 247, 3, 68, 34, 0, 246, 245, 1, 0, 0, 0, 247, 248, 1,
		0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 1, 0, 0,
		0, 250, 251, 5, 3, 0, 0, 251, 53, 1, 0, 0, 0, 252, 253, 6, 27, -1, 0, 253,
		254, 5, 25, 0, 0, 254, 255, 3, 70, 35, 0, 255, 259, 5, 4, 0, 0, 256, 258,
		3, 56, 28, 0, 257, 256, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1,
		0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 262, 1, 0, 0, 0, 261, 259, 1, 0, 0,
		0, 262, 263, 5, 5, 0, 0, 263, 55, 1, 0, 0, 0, 264, 266, 3, 58, 29, 0, 265,
		264, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 269, 1, 0, 0, 0, 267, 270,
		3, 52, 26, 0, 268, 270, 3, 60, 30, 0, 269, 267, 1, 0, 0, 0, 269, 268, 1,
		0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272, 3, 62, 31, 0, 272, 57, 1, 0, 0,
		0, 273, 275, 5, 15, 0, 0, 274, 276, 5, 2, 0, 0, 275, 274, 1, 0, 0, 0, 275,
		276, 1, 0, 0, 0, 276, 283, 1, 0, 0, 0, 277, 279, 3, 68, 34, 0, 278, 277,
		1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0,
		0, 0, 281, 284, 1, 0, 0, 0, 282, 284, 5, 22, 0, 0, 283, 278, 1, 0, 0, 0,
		283, 282, 1, 0, 0, 0, 284, 286, 1, 0, 0, 0, 285, 287, 5, 3, 0, 0, 286,
		285, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 59, 1, 0, 0, 0, 288, 289, 5,
		16, 0, 0, 289, 290, 5, 25, 0, 0, 290, 61, 1, 0, 0, 0, 291, 292, 6, 31,
		-1, 0, 292, 293, 5, 25, 0, 0, 293, 295, 3, 72, 36, 0, 294, 296, 3, 64,
		32, 0, 295, 294, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 298, 1, 0, 0, 0,
		297, 299, 3, 66, 33, 0, 298, 297, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299,
		63, 1, 0, 0, 0, 300, 302, 5, 2, 0, 0, 301, 303, 5, 25, 0, 0, 302, 301,
		1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 305, 5, 3,
		0, 0, 305, 65, 1, 0, 0, 0, 306, 307, 5, 11, 0, 0, 307, 309, 5, 2, 0, 0,
		308, 310, 3, 42, 21, 0, 309, 308, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310,
		311, 1, 0, 0, 0, 311, 312, 5, 3, 0, 0, 312, 67, 1, 0, 0, 0, 313, 314, 5,
		25, 0, 0, 314, 315, 6, 34, -1, 0, 315, 316, 5, 24, 0, 0, 316, 69, 1, 0,
		0, 0, 317, 319, 5, 25, 0, 0, 318, 320, 5, 12, 0, 0, 319, 318, 1, 0, 0,
		0, 319, 320, 1, 0, 0, 0, 320, 322, 1, 0, 0, 0, 321, 317, 1, 0, 0, 0, 322,
		323, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 71, 1,
		0, 0, 0, 325, 326, 5, 13, 0, 0, 326, 331, 5, 25, 0, 0, 327, 328, 5, 12,
		0, 0, 328, 330, 5, 25, 0, 0, 329, 327, 1, 0, 0, 0, 330, 333, 1, 0, 0, 0,
		331, 329, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 341, 1, 0, 0, 0, 333,
		331, 1, 0, 0, 0, 334, 335, 5, 14, 0, 0, 335, 338, 5, 25, 0, 0, 336, 337,
		5, 12, 0, 0, 337, 339, 5, 25, 0, 0, 338, 336, 1, 0, 0, 0, 338, 339, 1,
		0, 0, 0, 339, 341, 1, 0, 0, 0, 340, 325, 1, 0, 0, 0, 340, 334, 1, 0, 0,
		0, 341, 342, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343,
		346, 1, 0, 0, 0, 344, 346, 5, 13, 0, 0, 345, 340, 1, 0, 0, 0, 345, 344,
		1, 0, 0, 0, 346, 73, 1, 0, 0, 0, 39, 77, 85, 95, 107, 122, 128, 140, 147,
		151, 156, 162, 170, 177, 183, 191, 198, 204, 207, 220, 239, 248, 259, 265,
		269, 275, 280, 283, 286, 295, 298, 302, 309, 319, 323, 331, 338, 340, 342,
		345,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ApiParserParserInit initializes any static state used to implement ApiParserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewApiParserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ApiParserParserInit() {
	staticData := &apiparserParserStaticData
	staticData.once.Do(apiparserParserInit)
}

// NewApiParserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewApiParserParser(input antlr.TokenStream) *ApiParserParser {
	ApiParserParserInit()
	this := new(ApiParserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &apiparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ApiParser.g4"

	return this
}

// ApiParserParser tokens.
const (
	ApiParserParserEOF          = antlr.TokenEOF
	ApiParserParserT__0         = 1
	ApiParserParserT__1         = 2
	ApiParserParserT__2         = 3
	ApiParserParserT__3         = 4
	ApiParserParserT__4         = 5
	ApiParserParserT__5         = 6
	ApiParserParserT__6         = 7
	ApiParserParserT__7         = 8
	ApiParserParserT__8         = 9
	ApiParserParserT__9         = 10
	ApiParserParserT__10        = 11
	ApiParserParserT__11        = 12
	ApiParserParserT__12        = 13
	ApiParserParserT__13        = 14
	ApiParserParserATDOC        = 15
	ApiParserParserATHANDLER    = 16
	ApiParserParserINTERFACE    = 17
	ApiParserParserATSERVER     = 18
	ApiParserParserWS           = 19
	ApiParserParserCOMMENT      = 20
	ApiParserParserLINE_COMMENT = 21
	ApiParserParserSTRING       = 22
	ApiParserParserRAW_STRING   = 23
	ApiParserParserLINE_VALUE   = 24
	ApiParserParserID           = 25
)

// ApiParserParser rules.
const (
	ApiParserParserRULE_api              = 0
	ApiParserParserRULE_spec             = 1
	ApiParserParserRULE_syntaxLit        = 2
	ApiParserParserRULE_importSpec       = 3
	ApiParserParserRULE_importLit        = 4
	ApiParserParserRULE_importBlock      = 5
	ApiParserParserRULE_importBlockValue = 6
	ApiParserParserRULE_importValue      = 7
	ApiParserParserRULE_infoSpec         = 8
	ApiParserParserRULE_typeSpec         = 9
	ApiParserParserRULE_typeLit          = 10
	ApiParserParserRULE_typeBlock        = 11
	ApiParserParserRULE_typeLitBody      = 12
	ApiParserParserRULE_typeBlockBody    = 13
	ApiParserParserRULE_typeStruct       = 14
	ApiParserParserRULE_typeAlias        = 15
	ApiParserParserRULE_typeBlockStruct  = 16
	ApiParserParserRULE_typeBlockAlias   = 17
	ApiParserParserRULE_field            = 18
	ApiParserParserRULE_normalField      = 19
	ApiParserParserRULE_anonymousFiled   = 20
	ApiParserParserRULE_dataType         = 21
	ApiParserParserRULE_pointerType      = 22
	ApiParserParserRULE_mapType          = 23
	ApiParserParserRULE_arrayType        = 24
	ApiParserParserRULE_serviceSpec      = 25
	ApiParserParserRULE_atServer         = 26
	ApiParserParserRULE_serviceApi       = 27
	ApiParserParserRULE_serviceRoute     = 28
	ApiParserParserRULE_atDoc            = 29
	ApiParserParserRULE_atHandler        = 30
	ApiParserParserRULE_route            = 31
	ApiParserParserRULE_body             = 32
	ApiParserParserRULE_replybody        = 33
	ApiParserParserRULE_kvLit            = 34
	ApiParserParserRULE_serviceName      = 35
	ApiParserParserRULE_path             = 36
)

// IApiContext is an interface to support dynamic dispatch.
type IApiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSpec() []ISpecContext
	Spec(i int) ISpecContext

	// IsApiContext differentiates from other interfaces.
	IsApiContext()
}

type ApiContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApiContext() *ApiContext {
	var p = new(ApiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_api
	return p
}

func (*ApiContext) IsApiContext() {}

func NewApiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApiContext {
	var p = new(ApiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_api

	return p
}

func (s *ApiContext) GetParser() antlr.Parser { return s.parser }

func (s *ApiContext) AllSpec() []ISpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpecContext); ok {
			len++
		}
	}

	tst := make([]ISpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpecContext); ok {
			tst[i] = t.(ISpecContext)
			i++
		}
	}

	return tst
}

func (s *ApiContext) Spec(i int) ISpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecContext)
}

func (s *ApiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterApi(s)
	}
}

func (s *ApiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitApi(s)
	}
}

func (s *ApiContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitApi(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Api() (localctx IApiContext) {
	this := p
	_ = this

	localctx = NewApiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ApiParserParserRULE_api)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserParserATSERVER || _la == ApiParserParserID {
		{
			p.SetState(74)
			p.Spec()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISpecContext is an interface to support dynamic dispatch.
type ISpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SyntaxLit() ISyntaxLitContext
	ImportSpec() IImportSpecContext
	InfoSpec() IInfoSpecContext
	TypeSpec() ITypeSpecContext
	ServiceSpec() IServiceSpecContext

	// IsSpecContext differentiates from other interfaces.
	IsSpecContext()
}

type SpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecContext() *SpecContext {
	var p = new(SpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_spec
	return p
}

func (*SpecContext) IsSpecContext() {}

func NewSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecContext {
	var p = new(SpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_spec

	return p
}

func (s *SpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecContext) SyntaxLit() ISyntaxLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISyntaxLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISyntaxLitContext)
}

func (s *SpecContext) ImportSpec() IImportSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportSpecContext)
}

func (s *SpecContext) InfoSpec() IInfoSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInfoSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInfoSpecContext)
}

func (s *SpecContext) TypeSpec() ITypeSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *SpecContext) ServiceSpec() IServiceSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceSpecContext)
}

func (s *SpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterSpec(s)
	}
}

func (s *SpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitSpec(s)
	}
}

func (s *SpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Spec() (localctx ISpecContext) {
	this := p
	_ = this

	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ApiParserParserRULE_spec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.SyntaxLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.ImportSpec()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.InfoSpec()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.TypeSpec()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(84)
			p.ServiceSpec()
		}

	}

	return localctx
}

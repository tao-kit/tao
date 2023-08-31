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
		"STRING", "RAW_STRING", "LINE_VALUE", "ID", "LetterOrDigit",
	}
	staticData.ruleNames = []string{
		"api", "spec", "syntaxLit", "importSpec", "importLit", "importBlock",
		"importBlockValue", "importValue", "infoSpec", "typeSpec", "typeLit",
		"typeBlock", "typeLitBody", "typeBlockBody", "typeStruct", "typeAlias",
		"typeBlockStruct", "typeBlockAlias", "field", "normalField", "anonymousFiled",
		"dataType", "pointerType", "mapType", "arrayType", "serviceSpec", "atServer",
		"serviceApi", "serviceRoute", "atDoc", "atHandler", "route", "body",
		"replybody", "kvLit", "serviceName", "path", "pathItem",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 355, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 1, 0, 5, 0, 78, 8, 0, 10, 0, 12, 0, 81, 9, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 88, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 3, 3, 98, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		4, 5, 108, 8, 5, 11, 5, 12, 5, 109, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 4, 8, 123, 8, 8, 11, 8, 12, 8, 124, 1, 8,
		1, 8, 1, 9, 1, 9, 3, 9, 131, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 5, 11, 141, 8, 11, 10, 11, 12, 11, 144, 9, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 3, 12, 150, 8, 12, 1, 13, 1, 13, 3, 13, 154, 8, 13,
		1, 14, 1, 14, 1, 14, 3, 14, 159, 8, 14, 1, 14, 1, 14, 5, 14, 163, 8, 14,
		10, 14, 12, 14, 166, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15, 173,
		8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 3, 16, 180, 8, 16, 1, 16, 1,
		16, 5, 16, 184, 8, 16, 10, 16, 12, 16, 187, 9, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 3, 17, 194, 8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 3,
		18, 201, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 207, 8, 19, 1, 20, 3,
		20, 210, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 3, 21, 223, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 25, 3, 25, 242, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 4,
		26, 249, 8, 26, 11, 26, 12, 26, 250, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 5, 27, 260, 8, 27, 10, 27, 12, 27, 263, 9, 27, 1, 27, 1,
		27, 1, 28, 3, 28, 268, 8, 28, 1, 28, 1, 28, 3, 28, 272, 8, 28, 1, 28, 1,
		28, 1, 29, 1, 29, 3, 29, 278, 8, 29, 1, 29, 4, 29, 281, 8, 29, 11, 29,
		12, 29, 282, 1, 29, 3, 29, 286, 8, 29, 1, 29, 3, 29, 289, 8, 29, 1, 30,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 298, 8, 31, 1, 31, 3,
		31, 301, 8, 31, 1, 32, 1, 32, 3, 32, 305, 8, 32, 1, 32, 1, 32, 1, 33, 1,
		33, 1, 33, 3, 33, 312, 8, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 35, 1, 35, 3, 35, 322, 8, 35, 4, 35, 324, 8, 35, 11, 35, 12, 35, 325,
		1, 36, 1, 36, 1, 36, 1, 36, 5, 36, 332, 8, 36, 10, 36, 12, 36, 335, 9,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 341, 8, 36, 4, 36, 343, 8, 36, 11,
		36, 12, 36, 344, 1, 36, 3, 36, 348, 8, 36, 1, 37, 4, 37, 351, 8, 37, 11,
		37, 12, 37, 352, 1, 37, 0, 0, 38, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 62, 64, 66, 68, 70, 72, 74, 0, 1, 1, 0, 25, 26, 365, 0, 79, 1,
		0, 0, 0, 2, 87, 1, 0, 0, 0, 4, 89, 1, 0, 0, 0, 6, 97, 1, 0, 0, 0, 8, 99,
		1, 0, 0, 0, 10, 103, 1, 0, 0, 0, 12, 113, 1, 0, 0, 0, 14, 115, 1, 0, 0,
		0, 16, 118, 1, 0, 0, 0, 18, 130, 1, 0, 0, 0, 20, 132, 1, 0, 0, 0, 22, 136,
		1, 0, 0, 0, 24, 149, 1, 0, 0, 0, 26, 153, 1, 0, 0, 0, 28, 155, 1, 0, 0,
		0, 30, 169, 1, 0, 0, 0, 32, 176, 1, 0, 0, 0, 34, 190, 1, 0, 0, 0, 36, 200,
		1, 0, 0, 0, 38, 202, 1, 0, 0, 0, 40, 209, 1, 0, 0, 0, 42, 222, 1, 0, 0,
		0, 44, 224, 1, 0, 0, 0, 46, 228, 1, 0, 0, 0, 48, 236, 1, 0, 0, 0, 50, 241,
		1, 0, 0, 0, 52, 245, 1, 0, 0, 0, 54, 254, 1, 0, 0, 0, 56, 267, 1, 0, 0,
		0, 58, 275, 1, 0, 0, 0, 60, 290, 1, 0, 0, 0, 62, 293, 1, 0, 0, 0, 64, 302,
		1, 0, 0, 0, 66, 308, 1, 0, 0, 0, 68, 315, 1, 0, 0, 0, 70, 323, 1, 0, 0,
		0, 72, 347, 1, 0, 0, 0, 74, 350, 1, 0, 0, 0, 76, 78, 3, 2, 1, 0, 77, 76,
		1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0,
		80, 1, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 88, 3, 4, 2, 0, 83, 88, 3, 6,
		3, 0, 84, 88, 3, 16, 8, 0, 85, 88, 3, 18, 9, 0, 86, 88, 3, 50, 25, 0, 87,
		82, 1, 0, 0, 0, 87, 83, 1, 0, 0, 0, 87, 84, 1, 0, 0, 0, 87, 85, 1, 0, 0,
		0, 87, 86, 1, 0, 0, 0, 88, 3, 1, 0, 0, 0, 89, 90, 6, 2, -1, 0, 90, 91,
		5, 25, 0, 0, 91, 92, 5, 1, 0, 0, 92, 93, 6, 2, -1, 0, 93, 94, 5, 22, 0,
		0, 94, 5, 1, 0, 0, 0, 95, 98, 3, 8, 4, 0, 96, 98, 3, 10, 5, 0, 97, 95,
		1, 0, 0, 0, 97, 96, 1, 0, 0, 0, 98, 7, 1, 0, 0, 0, 99, 100, 6, 4, -1, 0,
		100, 101, 5, 25, 0, 0, 101, 102, 3, 14, 7, 0, 102, 9, 1, 0, 0, 0, 103,
		104, 6, 5, -1, 0, 104, 105, 5, 25, 0, 0, 105, 107, 5, 2, 0, 0, 106, 108,
		3, 12, 6, 0, 107, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 107, 1, 0,
		0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 5, 3, 0, 0,
		112, 11, 1, 0, 0, 0, 113, 114, 3, 14, 7, 0, 114, 13, 1, 0, 0, 0, 115, 116,
		6, 7, -1, 0, 116, 117, 5, 22, 0, 0, 117, 15, 1, 0, 0, 0, 118, 119, 6, 8,
		-1, 0, 119, 120, 5, 25, 0, 0, 120, 122, 5, 2, 0, 0, 121, 123, 3, 68, 34,
		0, 122, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124,
		125, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 5, 3, 0, 0, 127, 17, 1,
		0, 0, 0, 128, 131, 3, 20, 10, 0, 129, 131, 3, 22, 11, 0, 130, 128, 1, 0,
		0, 0, 130, 129, 1, 0, 0, 0, 131, 19, 1, 0, 0, 0, 132, 133, 6, 10, -1, 0,
		133, 134, 5, 25, 0, 0, 134, 135, 3, 24, 12, 0, 135, 21, 1, 0, 0, 0, 136,
		137, 6, 11, -1, 0, 137, 138, 5, 25, 0, 0, 138, 142, 5, 2, 0, 0, 139, 141,
		3, 26, 13, 0, 140, 139, 1, 0, 0, 0, 141, 144, 1, 0, 0, 0, 142, 140, 1,
		0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 145, 1, 0, 0, 0, 144, 142, 1, 0, 0,
		0, 145, 146, 5, 3, 0, 0, 146, 23, 1, 0, 0, 0, 147, 150, 3, 28, 14, 0, 148,
		150, 3, 30, 15, 0, 149, 147, 1, 0, 0, 0, 149, 148, 1, 0, 0, 0, 150, 25,
		1, 0, 0, 0, 151, 154, 3, 32, 16, 0, 152, 154, 3, 34, 17, 0, 153, 151, 1,
		0, 0, 0, 153, 152, 1, 0, 0, 0, 154, 27, 1, 0, 0, 0, 155, 156, 6, 14, -1,
		0, 156, 158, 5, 25, 0, 0, 157, 159, 5, 25, 0, 0, 158, 157, 1, 0, 0, 0,
		158, 159, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 164, 5, 4, 0, 0, 161,
		163, 3, 36, 18, 0, 162, 161, 1, 0, 0, 0, 163, 166, 1, 0, 0, 0, 164, 162,
		1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 167, 1, 0, 0, 0, 166, 164, 1, 0,
		0, 0, 167, 168, 5, 5, 0, 0, 168, 29, 1, 0, 0, 0, 169, 170, 6, 15, -1, 0,
		170, 172, 5, 25, 0, 0, 171, 173, 5, 1, 0, 0, 172, 171, 1, 0, 0, 0, 172,
		173, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 3, 42, 21, 0, 175, 31,
		1, 0, 0, 0, 176, 177, 6, 16, -1, 0, 177, 179, 5, 25, 0, 0, 178, 180, 5,
		25, 0, 0, 179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 1, 0, 0,
		0, 181, 185, 5, 4, 0, 0, 182, 184, 3, 36, 18, 0, 183, 182, 1, 0, 0, 0,
		184, 187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186,
		188, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 189, 5, 5, 0, 0, 189, 33, 1,
		0, 0, 0, 190, 191, 6, 17, -1, 0, 191, 193, 5, 25, 0, 0, 192, 194, 5, 1,
		0, 0, 193, 192, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0,
		195, 196, 3, 42, 21, 0, 196, 35, 1, 0, 0, 0, 197, 198, 4, 18, 0, 0, 198,
		201, 3, 38, 19, 0, 199, 201, 3, 40, 20, 0, 200, 197, 1, 0, 0, 0, 200, 199,
		1, 0, 0, 0, 201, 37, 1, 0, 0, 0, 202, 203, 6, 19, -1, 0, 203, 204, 5, 25,
		0, 0, 204, 206, 3, 42, 21, 0, 205, 207, 5, 23, 0, 0, 206, 205, 1, 0, 0,
		0, 206, 207, 1, 0, 0, 0, 207, 39, 1, 0, 0, 0, 208, 210, 5, 6, 0, 0, 209,
		208, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 212,
		5, 25, 0, 0, 212, 41, 1, 0, 0, 0, 213, 214, 6, 21, -1, 0, 214, 223, 5,
		25, 0, 0, 215, 223, 3, 46, 23, 0, 216, 223, 3, 48, 24, 0, 217, 223, 5,
		17, 0, 0, 218, 223, 5, 7, 0, 0, 219, 223, 5, 8, 0, 0, 220, 223, 3, 44,
		22, 0, 221, 223, 3, 28, 14, 0, 222, 213, 1, 0, 0, 0, 222, 215, 1, 0, 0,
		0, 222, 216, 1, 0, 0, 0, 222, 217, 1, 0, 0, 0, 222, 218, 1, 0, 0, 0, 222,
		219, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 221, 1, 0, 0, 0, 223, 43, 1,
		0, 0, 0, 224, 225, 5, 6, 0, 0, 225, 226, 6, 22, -1, 0, 226, 227, 5, 25,
		0, 0, 227, 45, 1, 0, 0, 0, 228, 229, 6, 23, -1, 0, 229, 230, 5, 25, 0,
		0, 230, 231, 5, 9, 0, 0, 231, 232, 6, 23, -1, 0, 232, 233, 5, 25, 0, 0,
		233, 234, 5, 10, 0, 0, 234, 235, 3, 42, 21, 0, 235, 47, 1, 0, 0, 0, 236,
		237, 5, 9, 0, 0, 237, 238, 5, 10, 0, 0, 238, 239, 3, 42, 21, 0, 239, 49,
		1, 0, 0, 0, 240, 242, 3, 52, 26, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1,
		0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 244, 3, 54, 27, 0, 244, 51, 1, 0, 0,
		0, 245, 246, 5, 18, 0, 0, 246, 248, 5, 2, 0, 0, 247, 249, 3, 68, 34, 0,
		248, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 250,
		251, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 253, 5, 3, 0, 0, 253, 53, 1,
		0, 0, 0, 254, 255, 6, 27, -1, 0, 255, 256, 5, 25, 0, 0, 256, 257, 3, 70,
		35, 0, 257, 261, 5, 4, 0, 0, 258, 260, 3, 56, 28, 0, 259, 258, 1, 0, 0,
		0, 260, 263, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262,
		264, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 264, 265, 5, 5, 0, 0, 265, 55, 1,
		0, 0, 0, 266, 268, 3, 58, 29, 0, 267, 266, 1, 0, 0, 0, 267, 268, 1, 0,
		0, 0, 268, 271, 1, 0, 0, 0, 269, 272, 3, 52, 26, 0, 270, 272, 3, 60, 30,
		0, 271, 269, 1, 0, 0, 0, 271, 270, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273,
		274, 3, 62, 31, 0, 274, 57, 1, 0, 0, 0, 275, 277, 5, 15, 0, 0, 276, 278,
		5, 2, 0, 0, 277, 276, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 285, 1, 0,
		0, 0, 279, 281, 3, 68, 34, 0, 280, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0,
		0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 286, 1, 0, 0, 0, 284,
		286, 5, 22, 0, 0, 285, 280, 1, 0, 0, 0, 285, 284, 1, 0, 0, 0, 286, 288,
		1, 0, 0, 0, 287, 289, 5, 3, 0, 0, 288, 287, 1, 0, 0, 0, 288, 289, 1, 0,
		0, 0, 289, 59, 1, 0, 0, 0, 290, 291, 5, 16, 0, 0, 291, 292, 5, 25, 0, 0,
		292, 61, 1, 0, 0, 0, 293, 294, 6, 31, -1, 0, 294, 295, 5, 25, 0, 0, 295,
		297, 3, 72, 36, 0, 296, 298, 3, 64, 32, 0, 297, 296, 1, 0, 0, 0, 297, 298,
		1, 0, 0, 0, 298, 300, 1, 0, 0, 0, 299, 301, 3, 66, 33, 0, 300, 299, 1,
		0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 63, 1, 0, 0, 0, 302, 304, 5, 2, 0,
		0, 303, 305, 5, 25, 0, 0, 304, 303, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305,
		306, 1, 0, 0, 0, 306, 307, 5, 3, 0, 0, 307, 65, 1, 0, 0, 0, 308, 309, 5,
		11, 0, 0, 309, 311, 5, 2, 0, 0, 310, 312, 3, 42, 21, 0, 311, 310, 1, 0,
		0, 0, 311, 312, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 314, 5, 3, 0, 0,
		314, 67, 1, 0, 0, 0, 315, 316, 5, 25, 0, 0, 316, 317, 6, 34, -1, 0, 317,
		318, 5, 24, 0, 0, 318, 69, 1, 0, 0, 0, 319, 321, 5, 25, 0, 0, 320, 322,
		5, 12, 0, 0, 321, 320, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 324, 1, 0,
		0, 0, 323, 319, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0,
		325, 326, 1, 0, 0, 0, 326, 71, 1, 0, 0, 0, 327, 328, 5, 13, 0, 0, 328,
		333, 3, 74, 37, 0, 329, 330, 5, 12, 0, 0, 330, 332, 3, 74, 37, 0, 331,
		329, 1, 0, 0, 0, 332, 335, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 333, 334,
		1, 0, 0, 0, 334, 343, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 336, 337, 5, 14,
		0, 0, 337, 340, 3, 74, 37, 0, 338, 339, 5, 12, 0, 0, 339, 341, 3, 74, 37,
		0, 340, 338, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 343, 1, 0, 0, 0, 342,
		327, 1, 0, 0, 0, 342, 336, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 342,
		1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 348, 1, 0, 0, 0, 346, 348, 5, 13,
		0, 0, 347, 342, 1, 0, 0, 0, 347, 346, 1, 0, 0, 0, 348, 73, 1, 0, 0, 0,
		349, 351, 7, 0, 0, 0, 350, 349, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352,
		350, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 75, 1, 0, 0, 0, 40, 79, 87,
		97, 109, 124, 130, 142, 149, 153, 158, 164, 172, 179, 185, 193, 200, 206,
		209, 222, 241, 250, 261, 267, 271, 277, 282, 285, 288, 297, 300, 304, 311,
		321, 325, 333, 340, 342, 344, 347, 352,
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
	ApiParserParserEOF           = antlr.TokenEOF
	ApiParserParserT__0          = 1
	ApiParserParserT__1          = 2
	ApiParserParserT__2          = 3
	ApiParserParserT__3          = 4
	ApiParserParserT__4          = 5
	ApiParserParserT__5          = 6
	ApiParserParserT__6          = 7
	ApiParserParserT__7          = 8
	ApiParserParserT__8          = 9
	ApiParserParserT__9          = 10
	ApiParserParserT__10         = 11
	ApiParserParserT__11         = 12
	ApiParserParserT__12         = 13
	ApiParserParserT__13         = 14
	ApiParserParserATDOC         = 15
	ApiParserParserATHANDLER     = 16
	ApiParserParserINTERFACE     = 17
	ApiParserParserATSERVER      = 18
	ApiParserParserWS            = 19
	ApiParserParserCOMMENT       = 20
	ApiParserParserLINE_COMMENT  = 21
	ApiParserParserSTRING        = 22
	ApiParserParserRAW_STRING    = 23
	ApiParserParserLINE_VALUE    = 24
	ApiParserParserID            = 25
	ApiParserParserLetterOrDigit = 26
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
	ApiParserParserRULE_pathItem         = 37
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
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserParserATSERVER || _la == ApiParserParserID {
		{
			p.SetState(76)
			p.Spec()
		}

		p.SetState(81)
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

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.SyntaxLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.ImportSpec()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.InfoSpec()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(85)
			p.TypeSpec()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(86)
			p.ServiceSpec()
		}

	}

	return localctx
}

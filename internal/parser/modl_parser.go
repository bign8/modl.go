// Code generated from grammar/antlr4/MODLParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MODLParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 31, 327,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 5, 2, 52, 10, 2, 3, 2, 3, 2, 3, 2, 7,
	2, 57, 10, 2, 12, 2, 14, 2, 60, 11, 2, 3, 2, 5, 2, 63, 10, 2, 5, 2, 65,
	10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 73, 10, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 7, 4, 79, 10, 4, 12, 4, 14, 4, 82, 11, 4, 5, 4, 84, 10, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 5, 5, 91, 10, 5, 3, 5, 6, 5, 94, 10, 5, 13,
	5, 14, 5, 95, 3, 5, 3, 5, 5, 5, 100, 10, 5, 3, 5, 7, 5, 103, 10, 5, 12,
	5, 14, 5, 106, 11, 5, 7, 5, 108, 10, 5, 12, 5, 14, 5, 111, 11, 5, 5, 5,
	113, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 6, 6, 119, 10, 6, 13, 6, 14, 6, 120,
	6, 6, 123, 10, 6, 13, 6, 14, 6, 124, 3, 6, 7, 6, 128, 10, 6, 12, 6, 14,
	6, 131, 11, 6, 3, 6, 5, 6, 134, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 5, 7, 143, 10, 7, 3, 8, 3, 8, 5, 8, 147, 10, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 5, 9, 155, 10, 9, 3, 9, 3, 9, 7, 9, 159, 10, 9, 12,
	9, 14, 9, 162, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 7, 10, 169, 10,
	10, 12, 10, 14, 10, 172, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 5, 11, 180, 10, 11, 3, 11, 3, 11, 7, 11, 184, 10, 11, 12, 11, 14, 11,
	187, 11, 11, 3, 11, 3, 11, 3, 12, 6, 12, 192, 10, 12, 13, 12, 14, 12, 193,
	3, 13, 3, 13, 5, 13, 198, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 5, 14, 206, 10, 14, 3, 14, 3, 14, 7, 14, 210, 10, 14, 12, 14, 14, 14,
	213, 11, 14, 3, 14, 3, 14, 3, 15, 6, 15, 218, 10, 15, 13, 15, 14, 15, 219,
	3, 16, 3, 16, 5, 16, 224, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 7, 17, 235, 10, 17, 12, 17, 14, 17, 238, 11, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 244, 10, 17, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 18, 7, 18, 251, 10, 18, 12, 18, 14, 18, 254, 11, 18, 3, 19, 5, 19,
	257, 10, 19, 3, 19, 3, 19, 5, 19, 261, 10, 19, 3, 19, 3, 19, 5, 19, 265,
	10, 19, 3, 19, 3, 19, 5, 19, 269, 10, 19, 7, 19, 271, 10, 19, 12, 19, 14,
	19, 274, 11, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 5, 20, 285, 10, 20, 3, 21, 5, 21, 288, 10, 21, 3, 21, 5, 21, 291,
	10, 21, 3, 21, 3, 21, 3, 21, 7, 21, 296, 10, 21, 12, 21, 14, 21, 299, 11,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 305, 10, 22, 12, 22, 14, 22, 308,
	11, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 317, 10,
	23, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 323, 10, 24, 3, 25, 3, 25, 3, 25,
	2, 2, 26, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 2, 5, 3, 2, 17, 18, 3, 2, 27, 28, 5, 2, 4,
	6, 15, 15, 17, 18, 2, 359, 2, 64, 3, 2, 2, 2, 4, 72, 3, 2, 2, 2, 6, 74,
	3, 2, 2, 2, 8, 87, 3, 2, 2, 2, 10, 122, 3, 2, 2, 2, 12, 142, 3, 2, 2, 2,
	14, 146, 3, 2, 2, 2, 16, 148, 3, 2, 2, 2, 18, 165, 3, 2, 2, 2, 20, 173,
	3, 2, 2, 2, 22, 191, 3, 2, 2, 2, 24, 197, 3, 2, 2, 2, 26, 199, 3, 2, 2,
	2, 28, 217, 3, 2, 2, 2, 30, 223, 3, 2, 2, 2, 32, 225, 3, 2, 2, 2, 34, 247,
	3, 2, 2, 2, 36, 256, 3, 2, 2, 2, 38, 284, 3, 2, 2, 2, 40, 287, 3, 2, 2,
	2, 42, 300, 3, 2, 2, 2, 44, 316, 3, 2, 2, 2, 46, 322, 3, 2, 2, 2, 48, 324,
	3, 2, 2, 2, 50, 52, 5, 4, 3, 2, 51, 50, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2,
	52, 65, 3, 2, 2, 2, 53, 58, 5, 4, 3, 2, 54, 55, 7, 9, 2, 2, 55, 57, 5,
	4, 3, 2, 56, 54, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58,
	59, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 63, 7, 9, 2,
	2, 62, 61, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64, 51,
	3, 2, 2, 2, 64, 53, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 7, 2, 2, 3,
	67, 3, 3, 2, 2, 2, 68, 73, 5, 6, 4, 2, 69, 73, 5, 8, 5, 2, 70, 73, 5, 16,
	9, 2, 71, 73, 5, 12, 7, 2, 72, 68, 3, 2, 2, 2, 72, 69, 3, 2, 2, 2, 72,
	70, 3, 2, 2, 2, 72, 71, 3, 2, 2, 2, 73, 5, 3, 2, 2, 2, 74, 83, 7, 11, 2,
	2, 75, 80, 5, 24, 13, 2, 76, 77, 7, 9, 2, 2, 77, 79, 5, 24, 13, 2, 78,
	76, 3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2,
	2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 83, 75, 3, 2, 2, 2, 83, 84,
	3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 86, 7, 12, 2, 2, 86, 7, 3, 2, 2, 2,
	87, 112, 7, 13, 2, 2, 88, 91, 5, 30, 16, 2, 89, 91, 5, 10, 6, 2, 90, 88,
	3, 2, 2, 2, 90, 89, 3, 2, 2, 2, 91, 109, 3, 2, 2, 2, 92, 94, 7, 9, 2, 2,
	93, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 96, 3,
	2, 2, 2, 96, 99, 3, 2, 2, 2, 97, 100, 5, 30, 16, 2, 98, 100, 5, 10, 6,
	2, 99, 97, 3, 2, 2, 2, 99, 98, 3, 2, 2, 2, 100, 104, 3, 2, 2, 2, 101, 103,
	7, 9, 2, 2, 102, 101, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 102, 3, 2,
	2, 2, 104, 105, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2,
	107, 93, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109, 110,
	3, 2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 112, 90, 3, 2,
	2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 115, 7, 14, 2, 2,
	115, 9, 3, 2, 2, 2, 116, 118, 5, 30, 16, 2, 117, 119, 7, 7, 2, 2, 118,
	117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121,
	3, 2, 2, 2, 121, 123, 3, 2, 2, 2, 122, 116, 3, 2, 2, 2, 123, 124, 3, 2,
	2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 129, 3, 2, 2, 2,
	126, 128, 5, 30, 16, 2, 127, 126, 3, 2, 2, 2, 128, 131, 3, 2, 2, 2, 129,
	127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 133, 3, 2, 2, 2, 131, 129,
	3, 2, 2, 2, 132, 134, 7, 7, 2, 2, 133, 132, 3, 2, 2, 2, 133, 134, 3, 2,
	2, 2, 134, 11, 3, 2, 2, 2, 135, 136, 9, 2, 2, 2, 136, 137, 7, 8, 2, 2,
	137, 143, 5, 14, 8, 2, 138, 139, 7, 18, 2, 2, 139, 143, 5, 6, 4, 2, 140,
	141, 7, 18, 2, 2, 141, 143, 5, 8, 5, 2, 142, 135, 3, 2, 2, 2, 142, 138,
	3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 143, 13, 3, 2, 2, 2, 144, 147, 5, 44,
	23, 2, 145, 147, 5, 32, 17, 2, 146, 144, 3, 2, 2, 2, 146, 145, 3, 2, 2,
	2, 147, 15, 3, 2, 2, 2, 148, 149, 7, 20, 2, 2, 149, 150, 5, 36, 19, 2,
	150, 151, 7, 22, 2, 2, 151, 160, 5, 18, 10, 2, 152, 154, 7, 23, 2, 2, 153,
	155, 5, 36, 19, 2, 154, 153, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 156,
	3, 2, 2, 2, 156, 157, 7, 22, 2, 2, 157, 159, 5, 18, 10, 2, 158, 152, 3,
	2, 2, 2, 159, 162, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2,
	2, 161, 163, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 163, 164, 7, 31, 2, 2, 164,
	17, 3, 2, 2, 2, 165, 170, 5, 4, 3, 2, 166, 167, 7, 9, 2, 2, 167, 169, 5,
	4, 3, 2, 168, 166, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2, 2,
	2, 170, 171, 3, 2, 2, 2, 171, 19, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 173,
	174, 7, 20, 2, 2, 174, 175, 5, 36, 19, 2, 175, 176, 7, 22, 2, 2, 176, 185,
	5, 22, 12, 2, 177, 179, 7, 23, 2, 2, 178, 180, 5, 36, 19, 2, 179, 178,
	3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 7, 22,
	2, 2, 182, 184, 5, 22, 12, 2, 183, 177, 3, 2, 2, 2, 184, 187, 3, 2, 2,
	2, 185, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 188, 3, 2, 2, 2, 187,
	185, 3, 2, 2, 2, 188, 189, 7, 31, 2, 2, 189, 21, 3, 2, 2, 2, 190, 192,
	5, 24, 13, 2, 191, 190, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 191, 3,
	2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 23, 3, 2, 2, 2, 195, 198, 5, 12, 7,
	2, 196, 198, 5, 20, 11, 2, 197, 195, 3, 2, 2, 2, 197, 196, 3, 2, 2, 2,
	198, 25, 3, 2, 2, 2, 199, 200, 7, 20, 2, 2, 200, 201, 5, 36, 19, 2, 201,
	202, 7, 22, 2, 2, 202, 211, 5, 28, 15, 2, 203, 205, 7, 23, 2, 2, 204, 206,
	5, 36, 19, 2, 205, 204, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207, 3,
	2, 2, 2, 207, 208, 7, 22, 2, 2, 208, 210, 5, 28, 15, 2, 209, 203, 3, 2,
	2, 2, 210, 213, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2,
	212, 214, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 214, 215, 7, 31, 2, 2, 215,
	27, 3, 2, 2, 2, 216, 218, 5, 30, 16, 2, 217, 216, 3, 2, 2, 2, 218, 219,
	3, 2, 2, 2, 219, 217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 29, 3, 2,
	2, 2, 221, 224, 5, 46, 24, 2, 222, 224, 5, 26, 14, 2, 223, 221, 3, 2, 2,
	2, 223, 222, 3, 2, 2, 2, 224, 31, 3, 2, 2, 2, 225, 226, 7, 20, 2, 2, 226,
	227, 5, 36, 19, 2, 227, 243, 7, 22, 2, 2, 228, 236, 5, 34, 18, 2, 229,
	230, 7, 23, 2, 2, 230, 231, 5, 36, 19, 2, 231, 232, 7, 22, 2, 2, 232, 233,
	5, 34, 18, 2, 233, 235, 3, 2, 2, 2, 234, 229, 3, 2, 2, 2, 235, 238, 3,
	2, 2, 2, 236, 234, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 239, 3, 2, 2,
	2, 238, 236, 3, 2, 2, 2, 239, 240, 7, 23, 2, 2, 240, 241, 7, 22, 2, 2,
	241, 242, 5, 34, 18, 2, 242, 244, 3, 2, 2, 2, 243, 228, 3, 2, 2, 2, 243,
	244, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 246, 7, 31, 2, 2, 246, 33,
	3, 2, 2, 2, 247, 252, 5, 14, 8, 2, 248, 249, 7, 7, 2, 2, 249, 251, 5, 14,
	8, 2, 250, 248, 3, 2, 2, 2, 251, 254, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2,
	252, 253, 3, 2, 2, 2, 253, 35, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 255, 257,
	7, 29, 2, 2, 256, 255, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 260, 3, 2,
	2, 2, 258, 261, 5, 40, 21, 2, 259, 261, 5, 42, 22, 2, 260, 258, 3, 2, 2,
	2, 260, 259, 3, 2, 2, 2, 261, 272, 3, 2, 2, 2, 262, 264, 9, 3, 2, 2, 263,
	265, 7, 29, 2, 2, 264, 263, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 268,
	3, 2, 2, 2, 266, 269, 5, 40, 21, 2, 267, 269, 5, 42, 22, 2, 268, 266, 3,
	2, 2, 2, 268, 267, 3, 2, 2, 2, 269, 271, 3, 2, 2, 2, 270, 262, 3, 2, 2,
	2, 271, 274, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273,
	37, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 275, 285, 7, 8, 2, 2, 276, 285, 7,
	24, 2, 2, 277, 278, 7, 24, 2, 2, 278, 285, 7, 8, 2, 2, 279, 285, 7, 25,
	2, 2, 280, 281, 7, 25, 2, 2, 281, 285, 7, 8, 2, 2, 282, 283, 7, 29, 2,
	2, 283, 285, 7, 8, 2, 2, 284, 275, 3, 2, 2, 2, 284, 276, 3, 2, 2, 2, 284,
	277, 3, 2, 2, 2, 284, 279, 3, 2, 2, 2, 284, 280, 3, 2, 2, 2, 284, 282,
	3, 2, 2, 2, 285, 39, 3, 2, 2, 2, 286, 288, 7, 18, 2, 2, 287, 286, 3, 2,
	2, 2, 287, 288, 3, 2, 2, 2, 288, 290, 3, 2, 2, 2, 289, 291, 5, 38, 20,
	2, 290, 289, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292,
	297, 5, 44, 23, 2, 293, 294, 7, 23, 2, 2, 294, 296, 5, 44, 23, 2, 295,
	293, 3, 2, 2, 2, 296, 299, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297, 298,
	3, 2, 2, 2, 298, 41, 3, 2, 2, 2, 299, 297, 3, 2, 2, 2, 300, 301, 7, 20,
	2, 2, 301, 306, 5, 36, 19, 2, 302, 303, 9, 3, 2, 2, 303, 305, 5, 36, 19,
	2, 304, 302, 3, 2, 2, 2, 305, 308, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306,
	307, 3, 2, 2, 2, 307, 309, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 309, 310,
	7, 31, 2, 2, 310, 43, 3, 2, 2, 2, 311, 317, 5, 6, 4, 2, 312, 317, 5, 8,
	5, 2, 313, 317, 5, 10, 6, 2, 314, 317, 5, 12, 7, 2, 315, 317, 5, 48, 25,
	2, 316, 311, 3, 2, 2, 2, 316, 312, 3, 2, 2, 2, 316, 313, 3, 2, 2, 2, 316,
	314, 3, 2, 2, 2, 316, 315, 3, 2, 2, 2, 317, 45, 3, 2, 2, 2, 318, 323, 5,
	6, 4, 2, 319, 323, 5, 12, 7, 2, 320, 323, 5, 8, 5, 2, 321, 323, 5, 48,
	25, 2, 322, 318, 3, 2, 2, 2, 322, 319, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2,
	322, 321, 3, 2, 2, 2, 323, 47, 3, 2, 2, 2, 324, 325, 9, 4, 2, 2, 325, 49,
	3, 2, 2, 2, 47, 51, 58, 62, 64, 72, 80, 83, 90, 95, 99, 104, 109, 112,
	120, 124, 129, 133, 142, 146, 154, 160, 170, 179, 185, 193, 197, 205, 211,
	219, 223, 236, 243, 252, 256, 260, 264, 268, 272, 284, 287, 290, 297, 306,
	316, 322,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "','", "", "", "", "", "", "", "", "",
	"", "'{'", "", "'?'", "'/'", "'>'", "'<'", "'*'", "'&'", "'|'", "'!'",
	"", "'}'",
}
var symbolicNames = []string{
	"", "WS", "NULL", "TRUE", "FALSE", "COLON", "EQUALS", "STRUCT_SEP", "ARR_SEP",
	"LBRAC", "RBRAC", "LSBRAC", "RSBRAC", "NUMBER", "COMMENT", "QUOTED", "STRING",
	"HASH_PREFIX", "LCBRAC", "CWS", "QMARK", "FSLASH", "GTHAN", "LTHAN", "ASTERISK",
	"AMP", "PIPE", "EXCLAM", "CCOMMENT", "RCBRAC",
}

var ruleNames = []string{
	"modl", "modl_structure", "modl_map", "modl_array", "modl_nb_array", "modl_pair",
	"modl_value_item", "modl_top_level_conditional", "modl_top_level_conditional_return",
	"modl_map_conditional", "modl_map_conditional_return", "modl_map_item",
	"modl_array_conditional", "modl_array_conditional_return", "modl_array_item",
	"modl_value_conditional", "modl_value_conditional_return", "modl_condition_test",
	"modl_operator", "modl_condition", "modl_condition_group", "modl_value",
	"modl_array_value_item", "modl_primitive",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MODLParser struct {
	*antlr.BaseParser
}

func NewMODLParser(input antlr.TokenStream) *MODLParser {
	this := new(MODLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "MODLParser.g4"

	return this
}

// MODLParser tokens.
const (
	MODLParserEOF         = antlr.TokenEOF
	MODLParserWS          = 1
	MODLParserNULL        = 2
	MODLParserTRUE        = 3
	MODLParserFALSE       = 4
	MODLParserCOLON       = 5
	MODLParserEQUALS      = 6
	MODLParserSTRUCT_SEP  = 7
	MODLParserARR_SEP     = 8
	MODLParserLBRAC       = 9
	MODLParserRBRAC       = 10
	MODLParserLSBRAC      = 11
	MODLParserRSBRAC      = 12
	MODLParserNUMBER      = 13
	MODLParserCOMMENT     = 14
	MODLParserQUOTED      = 15
	MODLParserSTRING      = 16
	MODLParserHASH_PREFIX = 17
	MODLParserLCBRAC      = 18
	MODLParserCWS         = 19
	MODLParserQMARK       = 20
	MODLParserFSLASH      = 21
	MODLParserGTHAN       = 22
	MODLParserLTHAN       = 23
	MODLParserASTERISK    = 24
	MODLParserAMP         = 25
	MODLParserPIPE        = 26
	MODLParserEXCLAM      = 27
	MODLParserCCOMMENT    = 28
	MODLParserRCBRAC      = 29
)

// MODLParser rules.
const (
	MODLParserRULE_modl                              = 0
	MODLParserRULE_modl_structure                    = 1
	MODLParserRULE_modl_map                          = 2
	MODLParserRULE_modl_array                        = 3
	MODLParserRULE_modl_nb_array                     = 4
	MODLParserRULE_modl_pair                         = 5
	MODLParserRULE_modl_value_item                   = 6
	MODLParserRULE_modl_top_level_conditional        = 7
	MODLParserRULE_modl_top_level_conditional_return = 8
	MODLParserRULE_modl_map_conditional              = 9
	MODLParserRULE_modl_map_conditional_return       = 10
	MODLParserRULE_modl_map_item                     = 11
	MODLParserRULE_modl_array_conditional            = 12
	MODLParserRULE_modl_array_conditional_return     = 13
	MODLParserRULE_modl_array_item                   = 14
	MODLParserRULE_modl_value_conditional            = 15
	MODLParserRULE_modl_value_conditional_return     = 16
	MODLParserRULE_modl_condition_test               = 17
	MODLParserRULE_modl_operator                     = 18
	MODLParserRULE_modl_condition                    = 19
	MODLParserRULE_modl_condition_group              = 20
	MODLParserRULE_modl_value                        = 21
	MODLParserRULE_modl_array_value_item             = 22
	MODLParserRULE_modl_primitive                    = 23
)

// IModlContext is an interface to support dynamic dispatch.
type IModlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModlContext differentiates from other interfaces.
	IsModlContext()
}

type ModlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModlContext() *ModlContext {
	var p = new(ModlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl
	return p
}

func (*ModlContext) IsModlContext() {}

func NewModlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModlContext {
	var p = new(ModlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl

	return p
}

func (s *ModlContext) GetParser() antlr.Parser { return s.parser }

func (s *ModlContext) EOF() antlr.TerminalNode {
	return s.GetToken(MODLParserEOF, 0)
}

func (s *ModlContext) AllModl_structure() []IModl_structureContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_structureContext)(nil)).Elem())
	var tst = make([]IModl_structureContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_structureContext)
		}
	}

	return tst
}

func (s *ModlContext) Modl_structure(i int) IModl_structureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_structureContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_structureContext)
}

func (s *ModlContext) AllSTRUCT_SEP() []antlr.TerminalNode {
	return s.GetTokens(MODLParserSTRUCT_SEP)
}

func (s *ModlContext) STRUCT_SEP(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserSTRUCT_SEP, i)
}

func (s *ModlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl(s)
	}
}

func (s *ModlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl(s)
	}
}

func (p *MODLParser) Modl() (localctx IModlContext) {
	localctx = NewModlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MODLParserRULE_modl)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MODLParserLBRAC)|(1<<MODLParserLSBRAC)|(1<<MODLParserQUOTED)|(1<<MODLParserSTRING)|(1<<MODLParserLCBRAC))) != 0 {
			{
				p.SetState(48)
				p.Modl_structure()
			}

		}

	case 2:
		{
			p.SetState(51)
			p.Modl_structure()
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(52)
					p.Match(MODLParserSTRUCT_SEP)
				}
				{
					p.SetState(53)
					p.Modl_structure()
				}

			}
			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MODLParserSTRUCT_SEP {
			{
				p.SetState(59)
				p.Match(MODLParserSTRUCT_SEP)
			}

		}

	}
	{
		p.SetState(64)
		p.Match(MODLParserEOF)
	}

	return localctx
}

// IModl_structureContext is an interface to support dynamic dispatch.
type IModl_structureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_structureContext differentiates from other interfaces.
	IsModl_structureContext()
}

type Modl_structureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_structureContext() *Modl_structureContext {
	var p = new(Modl_structureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_structure
	return p
}

func (*Modl_structureContext) IsModl_structureContext() {}

func NewModl_structureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_structureContext {
	var p = new(Modl_structureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_structure

	return p
}

func (s *Modl_structureContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_structureContext) Modl_map() IModl_mapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_mapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_mapContext)
}

func (s *Modl_structureContext) Modl_array() IModl_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_arrayContext)
}

func (s *Modl_structureContext) Modl_top_level_conditional() IModl_top_level_conditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_top_level_conditionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_top_level_conditionalContext)
}

func (s *Modl_structureContext) Modl_pair() IModl_pairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_pairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_pairContext)
}

func (s *Modl_structureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_structureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_structureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_structure(s)
	}
}

func (s *Modl_structureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_structure(s)
	}
}

func (p *MODLParser) Modl_structure() (localctx IModl_structureContext) {
	localctx = NewModl_structureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MODLParserRULE_modl_structure)

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

	p.SetState(70)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MODLParserLBRAC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Modl_map()
		}

	case MODLParserLSBRAC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Modl_array()
		}

	case MODLParserLCBRAC:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.Modl_top_level_conditional()
		}

	case MODLParserQUOTED, MODLParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(69)
			p.Modl_pair()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IModl_mapContext is an interface to support dynamic dispatch.
type IModl_mapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_mapContext differentiates from other interfaces.
	IsModl_mapContext()
}

type Modl_mapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_mapContext() *Modl_mapContext {
	var p = new(Modl_mapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_map
	return p
}

func (*Modl_mapContext) IsModl_mapContext() {}

func NewModl_mapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_mapContext {
	var p = new(Modl_mapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_map

	return p
}

func (s *Modl_mapContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_mapContext) LBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserLBRAC, 0)
}

func (s *Modl_mapContext) RBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserRBRAC, 0)
}

func (s *Modl_mapContext) AllModl_map_item() []IModl_map_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_map_itemContext)(nil)).Elem())
	var tst = make([]IModl_map_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_map_itemContext)
		}
	}

	return tst
}

func (s *Modl_mapContext) Modl_map_item(i int) IModl_map_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_map_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_map_itemContext)
}

func (s *Modl_mapContext) AllSTRUCT_SEP() []antlr.TerminalNode {
	return s.GetTokens(MODLParserSTRUCT_SEP)
}

func (s *Modl_mapContext) STRUCT_SEP(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserSTRUCT_SEP, i)
}

func (s *Modl_mapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_mapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_mapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_map(s)
	}
}

func (s *Modl_mapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_map(s)
	}
}

func (p *MODLParser) Modl_map() (localctx IModl_mapContext) {
	localctx = NewModl_mapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MODLParserRULE_modl_map)
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
	{
		p.SetState(72)
		p.Match(MODLParserLBRAC)
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MODLParserQUOTED)|(1<<MODLParserSTRING)|(1<<MODLParserLCBRAC))) != 0 {
		{
			p.SetState(73)
			p.Modl_map_item()
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MODLParserSTRUCT_SEP {
			{
				p.SetState(74)
				p.Match(MODLParserSTRUCT_SEP)
			}
			{
				p.SetState(75)
				p.Modl_map_item()
			}

			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(83)
		p.Match(MODLParserRBRAC)
	}

	return localctx
}

// IModl_arrayContext is an interface to support dynamic dispatch.
type IModl_arrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_arrayContext differentiates from other interfaces.
	IsModl_arrayContext()
}

type Modl_arrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_arrayContext() *Modl_arrayContext {
	var p = new(Modl_arrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_array
	return p
}

func (*Modl_arrayContext) IsModl_arrayContext() {}

func NewModl_arrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_arrayContext {
	var p = new(Modl_arrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_array

	return p
}

func (s *Modl_arrayContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_arrayContext) LSBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserLSBRAC, 0)
}

func (s *Modl_arrayContext) RSBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserRSBRAC, 0)
}

func (s *Modl_arrayContext) AllModl_array_item() []IModl_array_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_array_itemContext)(nil)).Elem())
	var tst = make([]IModl_array_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_array_itemContext)
		}
	}

	return tst
}

func (s *Modl_arrayContext) Modl_array_item(i int) IModl_array_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_array_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_array_itemContext)
}

func (s *Modl_arrayContext) AllModl_nb_array() []IModl_nb_arrayContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_nb_arrayContext)(nil)).Elem())
	var tst = make([]IModl_nb_arrayContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_nb_arrayContext)
		}
	}

	return tst
}

func (s *Modl_arrayContext) Modl_nb_array(i int) IModl_nb_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_nb_arrayContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_nb_arrayContext)
}

func (s *Modl_arrayContext) AllSTRUCT_SEP() []antlr.TerminalNode {
	return s.GetTokens(MODLParserSTRUCT_SEP)
}

func (s *Modl_arrayContext) STRUCT_SEP(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserSTRUCT_SEP, i)
}

func (s *Modl_arrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_arrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_arrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_array(s)
	}
}

func (s *Modl_arrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_array(s)
	}
}

func (p *MODLParser) Modl_array() (localctx IModl_arrayContext) {
	localctx = NewModl_arrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MODLParserRULE_modl_array)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(MODLParserLSBRAC)
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MODLParserNULL)|(1<<MODLParserTRUE)|(1<<MODLParserFALSE)|(1<<MODLParserLBRAC)|(1<<MODLParserLSBRAC)|(1<<MODLParserNUMBER)|(1<<MODLParserQUOTED)|(1<<MODLParserSTRING)|(1<<MODLParserLCBRAC))) != 0 {
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(86)
				p.Modl_array_item()
			}

		case 2:
			{
				p.SetState(87)
				p.Modl_nb_array()
			}

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MODLParserSTRUCT_SEP {
			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == MODLParserSTRUCT_SEP {
				{
					p.SetState(90)
					p.Match(MODLParserSTRUCT_SEP)
				}

				p.SetState(93)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(97)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(95)
					p.Modl_array_item()
				}

			case 2:
				{
					p.SetState(96)
					p.Modl_nb_array()
				}

			}
			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(99)
						p.Match(MODLParserSTRUCT_SEP)
					}

				}
				p.SetState(104)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
			}

			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(112)
		p.Match(MODLParserRSBRAC)
	}

	return localctx
}

// IModl_nb_arrayContext is an interface to support dynamic dispatch.
type IModl_nb_arrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_nb_arrayContext differentiates from other interfaces.
	IsModl_nb_arrayContext()
}

type Modl_nb_arrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_nb_arrayContext() *Modl_nb_arrayContext {
	var p = new(Modl_nb_arrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_nb_array
	return p
}

func (*Modl_nb_arrayContext) IsModl_nb_arrayContext() {}

func NewModl_nb_arrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_nb_arrayContext {
	var p = new(Modl_nb_arrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_nb_array

	return p
}

func (s *Modl_nb_arrayContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_nb_arrayContext) AllModl_array_item() []IModl_array_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_array_itemContext)(nil)).Elem())
	var tst = make([]IModl_array_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_array_itemContext)
		}
	}

	return tst
}

func (s *Modl_nb_arrayContext) Modl_array_item(i int) IModl_array_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_array_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_array_itemContext)
}

func (s *Modl_nb_arrayContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(MODLParserCOLON)
}

func (s *Modl_nb_arrayContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserCOLON, i)
}

func (s *Modl_nb_arrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_nb_arrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_nb_arrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_nb_array(s)
	}
}

func (s *Modl_nb_arrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_nb_array(s)
	}
}

func (p *MODLParser) Modl_nb_array() (localctx IModl_nb_arrayContext) {
	localctx = NewModl_nb_arrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MODLParserRULE_modl_nb_array)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(114)
				p.Modl_array_item()
			}
			p.SetState(116)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(115)
						p.Match(MODLParserCOLON)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(118)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(124)
				p.Modl_array_item()
			}

		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(130)
			p.Match(MODLParserCOLON)
		}

	}

	return localctx
}

// IModl_pairContext is an interface to support dynamic dispatch.
type IModl_pairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_pairContext differentiates from other interfaces.
	IsModl_pairContext()
}

type Modl_pairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_pairContext() *Modl_pairContext {
	var p = new(Modl_pairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_pair
	return p
}

func (*Modl_pairContext) IsModl_pairContext() {}

func NewModl_pairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_pairContext {
	var p = new(Modl_pairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_pair

	return p
}

func (s *Modl_pairContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_pairContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(MODLParserEQUALS, 0)
}

func (s *Modl_pairContext) Modl_value_item() IModl_value_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_value_itemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_value_itemContext)
}

func (s *Modl_pairContext) STRING() antlr.TerminalNode {
	return s.GetToken(MODLParserSTRING, 0)
}

func (s *Modl_pairContext) QUOTED() antlr.TerminalNode {
	return s.GetToken(MODLParserQUOTED, 0)
}

func (s *Modl_pairContext) Modl_map() IModl_mapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_mapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_mapContext)
}

func (s *Modl_pairContext) Modl_array() IModl_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_arrayContext)
}

func (s *Modl_pairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_pairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_pairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_pair(s)
	}
}

func (s *Modl_pairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_pair(s)
	}
}

func (p *MODLParser) Modl_pair() (localctx IModl_pairContext) {
	localctx = NewModl_pairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MODLParserRULE_modl_pair)
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

	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MODLParserQUOTED || _la == MODLParserSTRING) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(134)
			p.Match(MODLParserEQUALS)
		}
		{
			p.SetState(135)
			p.Modl_value_item()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.Match(MODLParserSTRING)
		}
		{
			p.SetState(137)
			p.Modl_map()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(138)
			p.Match(MODLParserSTRING)
		}
		{
			p.SetState(139)
			p.Modl_array()
		}

	}

	return localctx
}

// IModl_value_itemContext is an interface to support dynamic dispatch.
type IModl_value_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_value_itemContext differentiates from other interfaces.
	IsModl_value_itemContext()
}

type Modl_value_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_value_itemContext() *Modl_value_itemContext {
	var p = new(Modl_value_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_value_item
	return p
}

func (*Modl_value_itemContext) IsModl_value_itemContext() {}

func NewModl_value_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_value_itemContext {
	var p = new(Modl_value_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_value_item

	return p
}

func (s *Modl_value_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_value_itemContext) Modl_value() IModl_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_valueContext)
}

func (s *Modl_value_itemContext) Modl_value_conditional() IModl_value_conditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_value_conditionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_value_conditionalContext)
}

func (s *Modl_value_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_value_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_value_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_value_item(s)
	}
}

func (s *Modl_value_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_value_item(s)
	}
}

func (p *MODLParser) Modl_value_item() (localctx IModl_value_itemContext) {
	localctx = NewModl_value_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MODLParserRULE_modl_value_item)

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
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(142)
			p.Modl_value()
		}

	case 2:
		{
			p.SetState(143)
			p.Modl_value_conditional()
		}

	}

	return localctx
}

// IModl_top_level_conditionalContext is an interface to support dynamic dispatch.
type IModl_top_level_conditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_top_level_conditionalContext differentiates from other interfaces.
	IsModl_top_level_conditionalContext()
}

type Modl_top_level_conditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_top_level_conditionalContext() *Modl_top_level_conditionalContext {
	var p = new(Modl_top_level_conditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_top_level_conditional
	return p
}

func (*Modl_top_level_conditionalContext) IsModl_top_level_conditionalContext() {}

func NewModl_top_level_conditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_top_level_conditionalContext {
	var p = new(Modl_top_level_conditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_top_level_conditional

	return p
}

func (s *Modl_top_level_conditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_top_level_conditionalContext) LCBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserLCBRAC, 0)
}

func (s *Modl_top_level_conditionalContext) AllModl_condition_test() []IModl_condition_testContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_condition_testContext)(nil)).Elem())
	var tst = make([]IModl_condition_testContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_condition_testContext)
		}
	}

	return tst
}

func (s *Modl_top_level_conditionalContext) Modl_condition_test(i int) IModl_condition_testContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_condition_testContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_condition_testContext)
}

func (s *Modl_top_level_conditionalContext) AllQMARK() []antlr.TerminalNode {
	return s.GetTokens(MODLParserQMARK)
}

func (s *Modl_top_level_conditionalContext) QMARK(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserQMARK, i)
}

func (s *Modl_top_level_conditionalContext) AllModl_top_level_conditional_return() []IModl_top_level_conditional_returnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_top_level_conditional_returnContext)(nil)).Elem())
	var tst = make([]IModl_top_level_conditional_returnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_top_level_conditional_returnContext)
		}
	}

	return tst
}

func (s *Modl_top_level_conditionalContext) Modl_top_level_conditional_return(i int) IModl_top_level_conditional_returnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_top_level_conditional_returnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_top_level_conditional_returnContext)
}

func (s *Modl_top_level_conditionalContext) RCBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserRCBRAC, 0)
}

func (s *Modl_top_level_conditionalContext) AllFSLASH() []antlr.TerminalNode {
	return s.GetTokens(MODLParserFSLASH)
}

func (s *Modl_top_level_conditionalContext) FSLASH(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserFSLASH, i)
}

func (s *Modl_top_level_conditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_top_level_conditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_top_level_conditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_top_level_conditional(s)
	}
}

func (s *Modl_top_level_conditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_top_level_conditional(s)
	}
}

func (p *MODLParser) Modl_top_level_conditional() (localctx IModl_top_level_conditionalContext) {
	localctx = NewModl_top_level_conditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MODLParserRULE_modl_top_level_conditional)
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
	{
		p.SetState(146)
		p.Match(MODLParserLCBRAC)
	}
	{
		p.SetState(147)
		p.Modl_condition_test()
	}
	{
		p.SetState(148)
		p.Match(MODLParserQMARK)
	}
	{
		p.SetState(149)
		p.Modl_top_level_conditional_return()
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MODLParserFSLASH {
		{
			p.SetState(150)
			p.Match(MODLParserFSLASH)
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MODLParserNULL)|(1<<MODLParserTRUE)|(1<<MODLParserFALSE)|(1<<MODLParserEQUALS)|(1<<MODLParserLBRAC)|(1<<MODLParserLSBRAC)|(1<<MODLParserNUMBER)|(1<<MODLParserQUOTED)|(1<<MODLParserSTRING)|(1<<MODLParserLCBRAC)|(1<<MODLParserGTHAN)|(1<<MODLParserLTHAN)|(1<<MODLParserEXCLAM))) != 0 {
			{
				p.SetState(151)
				p.Modl_condition_test()
			}

		}
		{
			p.SetState(154)
			p.Match(MODLParserQMARK)
		}
		{
			p.SetState(155)
			p.Modl_top_level_conditional_return()
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(161)
		p.Match(MODLParserRCBRAC)
	}

	return localctx
}

// IModl_top_level_conditional_returnContext is an interface to support dynamic dispatch.
type IModl_top_level_conditional_returnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_top_level_conditional_returnContext differentiates from other interfaces.
	IsModl_top_level_conditional_returnContext()
}

type Modl_top_level_conditional_returnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_top_level_conditional_returnContext() *Modl_top_level_conditional_returnContext {
	var p = new(Modl_top_level_conditional_returnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_top_level_conditional_return
	return p
}

func (*Modl_top_level_conditional_returnContext) IsModl_top_level_conditional_returnContext() {}

func NewModl_top_level_conditional_returnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_top_level_conditional_returnContext {
	var p = new(Modl_top_level_conditional_returnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_top_level_conditional_return

	return p
}

func (s *Modl_top_level_conditional_returnContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_top_level_conditional_returnContext) AllModl_structure() []IModl_structureContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_structureContext)(nil)).Elem())
	var tst = make([]IModl_structureContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_structureContext)
		}
	}

	return tst
}

func (s *Modl_top_level_conditional_returnContext) Modl_structure(i int) IModl_structureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_structureContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_structureContext)
}

func (s *Modl_top_level_conditional_returnContext) AllSTRUCT_SEP() []antlr.TerminalNode {
	return s.GetTokens(MODLParserSTRUCT_SEP)
}

func (s *Modl_top_level_conditional_returnContext) STRUCT_SEP(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserSTRUCT_SEP, i)
}

func (s *Modl_top_level_conditional_returnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_top_level_conditional_returnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_top_level_conditional_returnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_top_level_conditional_return(s)
	}
}

func (s *Modl_top_level_conditional_returnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_top_level_conditional_return(s)
	}
}

func (p *MODLParser) Modl_top_level_conditional_return() (localctx IModl_top_level_conditional_returnContext) {
	localctx = NewModl_top_level_conditional_returnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MODLParserRULE_modl_top_level_conditional_return)
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
	{
		p.SetState(163)
		p.Modl_structure()
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MODLParserSTRUCT_SEP {
		{
			p.SetState(164)
			p.Match(MODLParserSTRUCT_SEP)
		}
		{
			p.SetState(165)
			p.Modl_structure()
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IModl_map_conditionalContext is an interface to support dynamic dispatch.
type IModl_map_conditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_map_conditionalContext differentiates from other interfaces.
	IsModl_map_conditionalContext()
}

type Modl_map_conditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_map_conditionalContext() *Modl_map_conditionalContext {
	var p = new(Modl_map_conditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_map_conditional
	return p
}

func (*Modl_map_conditionalContext) IsModl_map_conditionalContext() {}

func NewModl_map_conditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_map_conditionalContext {
	var p = new(Modl_map_conditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_map_conditional

	return p
}

func (s *Modl_map_conditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_map_conditionalContext) LCBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserLCBRAC, 0)
}

func (s *Modl_map_conditionalContext) AllModl_condition_test() []IModl_condition_testContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_condition_testContext)(nil)).Elem())
	var tst = make([]IModl_condition_testContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_condition_testContext)
		}
	}

	return tst
}

func (s *Modl_map_conditionalContext) Modl_condition_test(i int) IModl_condition_testContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_condition_testContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_condition_testContext)
}

func (s *Modl_map_conditionalContext) AllQMARK() []antlr.TerminalNode {
	return s.GetTokens(MODLParserQMARK)
}

func (s *Modl_map_conditionalContext) QMARK(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserQMARK, i)
}

func (s *Modl_map_conditionalContext) AllModl_map_conditional_return() []IModl_map_conditional_returnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_map_conditional_returnContext)(nil)).Elem())
	var tst = make([]IModl_map_conditional_returnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_map_conditional_returnContext)
		}
	}

	return tst
}

func (s *Modl_map_conditionalContext) Modl_map_conditional_return(i int) IModl_map_conditional_returnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_map_conditional_returnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_map_conditional_returnContext)
}

func (s *Modl_map_conditionalContext) RCBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserRCBRAC, 0)
}

func (s *Modl_map_conditionalContext) AllFSLASH() []antlr.TerminalNode {
	return s.GetTokens(MODLParserFSLASH)
}

func (s *Modl_map_conditionalContext) FSLASH(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserFSLASH, i)
}

func (s *Modl_map_conditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_map_conditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_map_conditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_map_conditional(s)
	}
}

func (s *Modl_map_conditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_map_conditional(s)
	}
}

func (p *MODLParser) Modl_map_conditional() (localctx IModl_map_conditionalContext) {
	localctx = NewModl_map_conditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MODLParserRULE_modl_map_conditional)
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
	{
		p.SetState(171)
		p.Match(MODLParserLCBRAC)
	}
	{
		p.SetState(172)
		p.Modl_condition_test()
	}
	{
		p.SetState(173)
		p.Match(MODLParserQMARK)
	}
	{
		p.SetState(174)
		p.Modl_map_conditional_return()
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MODLParserFSLASH {
		{
			p.SetState(175)
			p.Match(MODLParserFSLASH)
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MODLParserNULL)|(1<<MODLParserTRUE)|(1<<MODLParserFALSE)|(1<<MODLParserEQUALS)|(1<<MODLParserLBRAC)|(1<<MODLParserLSBRAC)|(1<<MODLParserNUMBER)|(1<<MODLParserQUOTED)|(1<<MODLParserSTRING)|(1<<MODLParserLCBRAC)|(1<<MODLParserGTHAN)|(1<<MODLParserLTHAN)|(1<<MODLParserEXCLAM))) != 0 {
			{
				p.SetState(176)
				p.Modl_condition_test()
			}

		}
		{
			p.SetState(179)
			p.Match(MODLParserQMARK)
		}
		{
			p.SetState(180)
			p.Modl_map_conditional_return()
		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(186)
		p.Match(MODLParserRCBRAC)
	}

	return localctx
}

// IModl_map_conditional_returnContext is an interface to support dynamic dispatch.
type IModl_map_conditional_returnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_map_conditional_returnContext differentiates from other interfaces.
	IsModl_map_conditional_returnContext()
}

type Modl_map_conditional_returnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_map_conditional_returnContext() *Modl_map_conditional_returnContext {
	var p = new(Modl_map_conditional_returnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_map_conditional_return
	return p
}

func (*Modl_map_conditional_returnContext) IsModl_map_conditional_returnContext() {}

func NewModl_map_conditional_returnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_map_conditional_returnContext {
	var p = new(Modl_map_conditional_returnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_map_conditional_return

	return p
}

func (s *Modl_map_conditional_returnContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_map_conditional_returnContext) AllModl_map_item() []IModl_map_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_map_itemContext)(nil)).Elem())
	var tst = make([]IModl_map_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_map_itemContext)
		}
	}

	return tst
}

func (s *Modl_map_conditional_returnContext) Modl_map_item(i int) IModl_map_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_map_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_map_itemContext)
}

func (s *Modl_map_conditional_returnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_map_conditional_returnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_map_conditional_returnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_map_conditional_return(s)
	}
}

func (s *Modl_map_conditional_returnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_map_conditional_return(s)
	}
}

func (p *MODLParser) Modl_map_conditional_return() (localctx IModl_map_conditional_returnContext) {
	localctx = NewModl_map_conditional_returnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MODLParserRULE_modl_map_conditional_return)
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
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MODLParserQUOTED)|(1<<MODLParserSTRING)|(1<<MODLParserLCBRAC))) != 0) {
		{
			p.SetState(188)
			p.Modl_map_item()
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IModl_map_itemContext is an interface to support dynamic dispatch.
type IModl_map_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_map_itemContext differentiates from other interfaces.
	IsModl_map_itemContext()
}

type Modl_map_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_map_itemContext() *Modl_map_itemContext {
	var p = new(Modl_map_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_map_item
	return p
}

func (*Modl_map_itemContext) IsModl_map_itemContext() {}

func NewModl_map_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_map_itemContext {
	var p = new(Modl_map_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_map_item

	return p
}

func (s *Modl_map_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_map_itemContext) Modl_pair() IModl_pairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_pairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_pairContext)
}

func (s *Modl_map_itemContext) Modl_map_conditional() IModl_map_conditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_map_conditionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_map_conditionalContext)
}

func (s *Modl_map_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_map_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_map_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_map_item(s)
	}
}

func (s *Modl_map_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_map_item(s)
	}
}

func (p *MODLParser) Modl_map_item() (localctx IModl_map_itemContext) {
	localctx = NewModl_map_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MODLParserRULE_modl_map_item)

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

	p.SetState(195)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MODLParserQUOTED, MODLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.Modl_pair()
		}

	case MODLParserLCBRAC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			p.Modl_map_conditional()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IModl_array_conditionalContext is an interface to support dynamic dispatch.
type IModl_array_conditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_array_conditionalContext differentiates from other interfaces.
	IsModl_array_conditionalContext()
}

type Modl_array_conditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_array_conditionalContext() *Modl_array_conditionalContext {
	var p = new(Modl_array_conditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_array_conditional
	return p
}

func (*Modl_array_conditionalContext) IsModl_array_conditionalContext() {}

func NewModl_array_conditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_array_conditionalContext {
	var p = new(Modl_array_conditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_array_conditional

	return p
}

func (s *Modl_array_conditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_array_conditionalContext) LCBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserLCBRAC, 0)
}

func (s *Modl_array_conditionalContext) AllModl_condition_test() []IModl_condition_testContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_condition_testContext)(nil)).Elem())
	var tst = make([]IModl_condition_testContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_condition_testContext)
		}
	}

	return tst
}

func (s *Modl_array_conditionalContext) Modl_condition_test(i int) IModl_condition_testContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_condition_testContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_condition_testContext)
}

func (s *Modl_array_conditionalContext) AllQMARK() []antlr.TerminalNode {
	return s.GetTokens(MODLParserQMARK)
}

func (s *Modl_array_conditionalContext) QMARK(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserQMARK, i)
}

func (s *Modl_array_conditionalContext) AllModl_array_conditional_return() []IModl_array_conditional_returnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_array_conditional_returnContext)(nil)).Elem())
	var tst = make([]IModl_array_conditional_returnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_array_conditional_returnContext)
		}
	}

	return tst
}

func (s *Modl_array_conditionalContext) Modl_array_conditional_return(i int) IModl_array_conditional_returnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_array_conditional_returnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_array_conditional_returnContext)
}

func (s *Modl_array_conditionalContext) RCBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserRCBRAC, 0)
}

func (s *Modl_array_conditionalContext) AllFSLASH() []antlr.TerminalNode {
	return s.GetTokens(MODLParserFSLASH)
}

func (s *Modl_array_conditionalContext) FSLASH(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserFSLASH, i)
}

func (s *Modl_array_conditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_array_conditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_array_conditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_array_conditional(s)
	}
}

func (s *Modl_array_conditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_array_conditional(s)
	}
}

func (p *MODLParser) Modl_array_conditional() (localctx IModl_array_conditionalContext) {
	localctx = NewModl_array_conditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MODLParserRULE_modl_array_conditional)
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
	{
		p.SetState(197)
		p.Match(MODLParserLCBRAC)
	}
	{
		p.SetState(198)
		p.Modl_condition_test()
	}
	{
		p.SetState(199)
		p.Match(MODLParserQMARK)
	}
	{
		p.SetState(200)
		p.Modl_array_conditional_return()
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MODLParserFSLASH {
		{
			p.SetState(201)
			p.Match(MODLParserFSLASH)
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MODLParserNULL)|(1<<MODLParserTRUE)|(1<<MODLParserFALSE)|(1<<MODLParserEQUALS)|(1<<MODLParserLBRAC)|(1<<MODLParserLSBRAC)|(1<<MODLParserNUMBER)|(1<<MODLParserQUOTED)|(1<<MODLParserSTRING)|(1<<MODLParserLCBRAC)|(1<<MODLParserGTHAN)|(1<<MODLParserLTHAN)|(1<<MODLParserEXCLAM))) != 0 {
			{
				p.SetState(202)
				p.Modl_condition_test()
			}

		}
		{
			p.SetState(205)
			p.Match(MODLParserQMARK)
		}
		{
			p.SetState(206)
			p.Modl_array_conditional_return()
		}

		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(212)
		p.Match(MODLParserRCBRAC)
	}

	return localctx
}

// IModl_array_conditional_returnContext is an interface to support dynamic dispatch.
type IModl_array_conditional_returnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_array_conditional_returnContext differentiates from other interfaces.
	IsModl_array_conditional_returnContext()
}

type Modl_array_conditional_returnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_array_conditional_returnContext() *Modl_array_conditional_returnContext {
	var p = new(Modl_array_conditional_returnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_array_conditional_return
	return p
}

func (*Modl_array_conditional_returnContext) IsModl_array_conditional_returnContext() {}

func NewModl_array_conditional_returnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_array_conditional_returnContext {
	var p = new(Modl_array_conditional_returnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_array_conditional_return

	return p
}

func (s *Modl_array_conditional_returnContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_array_conditional_returnContext) AllModl_array_item() []IModl_array_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_array_itemContext)(nil)).Elem())
	var tst = make([]IModl_array_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_array_itemContext)
		}
	}

	return tst
}

func (s *Modl_array_conditional_returnContext) Modl_array_item(i int) IModl_array_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_array_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_array_itemContext)
}

func (s *Modl_array_conditional_returnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_array_conditional_returnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_array_conditional_returnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_array_conditional_return(s)
	}
}

func (s *Modl_array_conditional_returnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_array_conditional_return(s)
	}
}

func (p *MODLParser) Modl_array_conditional_return() (localctx IModl_array_conditional_returnContext) {
	localctx = NewModl_array_conditional_returnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MODLParserRULE_modl_array_conditional_return)
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
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MODLParserNULL)|(1<<MODLParserTRUE)|(1<<MODLParserFALSE)|(1<<MODLParserLBRAC)|(1<<MODLParserLSBRAC)|(1<<MODLParserNUMBER)|(1<<MODLParserQUOTED)|(1<<MODLParserSTRING)|(1<<MODLParserLCBRAC))) != 0) {
		{
			p.SetState(214)
			p.Modl_array_item()
		}

		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IModl_array_itemContext is an interface to support dynamic dispatch.
type IModl_array_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_array_itemContext differentiates from other interfaces.
	IsModl_array_itemContext()
}

type Modl_array_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_array_itemContext() *Modl_array_itemContext {
	var p = new(Modl_array_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_array_item
	return p
}

func (*Modl_array_itemContext) IsModl_array_itemContext() {}

func NewModl_array_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_array_itemContext {
	var p = new(Modl_array_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_array_item

	return p
}

func (s *Modl_array_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_array_itemContext) Modl_array_value_item() IModl_array_value_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_array_value_itemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_array_value_itemContext)
}

func (s *Modl_array_itemContext) Modl_array_conditional() IModl_array_conditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_array_conditionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_array_conditionalContext)
}

func (s *Modl_array_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_array_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_array_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_array_item(s)
	}
}

func (s *Modl_array_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_array_item(s)
	}
}

func (p *MODLParser) Modl_array_item() (localctx IModl_array_itemContext) {
	localctx = NewModl_array_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MODLParserRULE_modl_array_item)

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

	p.SetState(221)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MODLParserNULL, MODLParserTRUE, MODLParserFALSE, MODLParserLBRAC, MODLParserLSBRAC, MODLParserNUMBER, MODLParserQUOTED, MODLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.Modl_array_value_item()
		}

	case MODLParserLCBRAC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.Modl_array_conditional()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IModl_value_conditionalContext is an interface to support dynamic dispatch.
type IModl_value_conditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_value_conditionalContext differentiates from other interfaces.
	IsModl_value_conditionalContext()
}

type Modl_value_conditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_value_conditionalContext() *Modl_value_conditionalContext {
	var p = new(Modl_value_conditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_value_conditional
	return p
}

func (*Modl_value_conditionalContext) IsModl_value_conditionalContext() {}

func NewModl_value_conditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_value_conditionalContext {
	var p = new(Modl_value_conditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_value_conditional

	return p
}

func (s *Modl_value_conditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_value_conditionalContext) LCBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserLCBRAC, 0)
}

func (s *Modl_value_conditionalContext) AllModl_condition_test() []IModl_condition_testContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_condition_testContext)(nil)).Elem())
	var tst = make([]IModl_condition_testContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_condition_testContext)
		}
	}

	return tst
}

func (s *Modl_value_conditionalContext) Modl_condition_test(i int) IModl_condition_testContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_condition_testContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_condition_testContext)
}

func (s *Modl_value_conditionalContext) AllQMARK() []antlr.TerminalNode {
	return s.GetTokens(MODLParserQMARK)
}

func (s *Modl_value_conditionalContext) QMARK(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserQMARK, i)
}

func (s *Modl_value_conditionalContext) RCBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserRCBRAC, 0)
}

func (s *Modl_value_conditionalContext) AllModl_value_conditional_return() []IModl_value_conditional_returnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_value_conditional_returnContext)(nil)).Elem())
	var tst = make([]IModl_value_conditional_returnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_value_conditional_returnContext)
		}
	}

	return tst
}

func (s *Modl_value_conditionalContext) Modl_value_conditional_return(i int) IModl_value_conditional_returnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_value_conditional_returnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_value_conditional_returnContext)
}

func (s *Modl_value_conditionalContext) AllFSLASH() []antlr.TerminalNode {
	return s.GetTokens(MODLParserFSLASH)
}

func (s *Modl_value_conditionalContext) FSLASH(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserFSLASH, i)
}

func (s *Modl_value_conditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_value_conditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_value_conditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_value_conditional(s)
	}
}

func (s *Modl_value_conditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_value_conditional(s)
	}
}

func (p *MODLParser) Modl_value_conditional() (localctx IModl_value_conditionalContext) {
	localctx = NewModl_value_conditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MODLParserRULE_modl_value_conditional)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(MODLParserLCBRAC)
	}
	{
		p.SetState(224)
		p.Modl_condition_test()
	}
	{
		p.SetState(225)
		p.Match(MODLParserQMARK)
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MODLParserNULL)|(1<<MODLParserTRUE)|(1<<MODLParserFALSE)|(1<<MODLParserLBRAC)|(1<<MODLParserLSBRAC)|(1<<MODLParserNUMBER)|(1<<MODLParserQUOTED)|(1<<MODLParserSTRING)|(1<<MODLParserLCBRAC))) != 0 {
		{
			p.SetState(226)
			p.Modl_value_conditional_return()
		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(227)
					p.Match(MODLParserFSLASH)
				}
				{
					p.SetState(228)
					p.Modl_condition_test()
				}
				{
					p.SetState(229)
					p.Match(MODLParserQMARK)
				}
				{
					p.SetState(230)
					p.Modl_value_conditional_return()
				}

			}
			p.SetState(236)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
		}

		{
			p.SetState(237)
			p.Match(MODLParserFSLASH)
		}
		{
			p.SetState(238)
			p.Match(MODLParserQMARK)
		}
		{
			p.SetState(239)
			p.Modl_value_conditional_return()
		}

	}
	{
		p.SetState(243)
		p.Match(MODLParserRCBRAC)
	}

	return localctx
}

// IModl_value_conditional_returnContext is an interface to support dynamic dispatch.
type IModl_value_conditional_returnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_value_conditional_returnContext differentiates from other interfaces.
	IsModl_value_conditional_returnContext()
}

type Modl_value_conditional_returnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_value_conditional_returnContext() *Modl_value_conditional_returnContext {
	var p = new(Modl_value_conditional_returnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_value_conditional_return
	return p
}

func (*Modl_value_conditional_returnContext) IsModl_value_conditional_returnContext() {}

func NewModl_value_conditional_returnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_value_conditional_returnContext {
	var p = new(Modl_value_conditional_returnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_value_conditional_return

	return p
}

func (s *Modl_value_conditional_returnContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_value_conditional_returnContext) AllModl_value_item() []IModl_value_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_value_itemContext)(nil)).Elem())
	var tst = make([]IModl_value_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_value_itemContext)
		}
	}

	return tst
}

func (s *Modl_value_conditional_returnContext) Modl_value_item(i int) IModl_value_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_value_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_value_itemContext)
}

func (s *Modl_value_conditional_returnContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(MODLParserCOLON)
}

func (s *Modl_value_conditional_returnContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserCOLON, i)
}

func (s *Modl_value_conditional_returnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_value_conditional_returnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_value_conditional_returnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_value_conditional_return(s)
	}
}

func (s *Modl_value_conditional_returnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_value_conditional_return(s)
	}
}

func (p *MODLParser) Modl_value_conditional_return() (localctx IModl_value_conditional_returnContext) {
	localctx = NewModl_value_conditional_returnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MODLParserRULE_modl_value_conditional_return)
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
	{
		p.SetState(245)
		p.Modl_value_item()
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MODLParserCOLON {
		{
			p.SetState(246)
			p.Match(MODLParserCOLON)
		}
		{
			p.SetState(247)
			p.Modl_value_item()
		}

		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IModl_condition_testContext is an interface to support dynamic dispatch.
type IModl_condition_testContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_condition_testContext differentiates from other interfaces.
	IsModl_condition_testContext()
}

type Modl_condition_testContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_condition_testContext() *Modl_condition_testContext {
	var p = new(Modl_condition_testContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_condition_test
	return p
}

func (*Modl_condition_testContext) IsModl_condition_testContext() {}

func NewModl_condition_testContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_condition_testContext {
	var p = new(Modl_condition_testContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_condition_test

	return p
}

func (s *Modl_condition_testContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_condition_testContext) AllModl_condition() []IModl_conditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_conditionContext)(nil)).Elem())
	var tst = make([]IModl_conditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_conditionContext)
		}
	}

	return tst
}

func (s *Modl_condition_testContext) Modl_condition(i int) IModl_conditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_conditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_conditionContext)
}

func (s *Modl_condition_testContext) AllModl_condition_group() []IModl_condition_groupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_condition_groupContext)(nil)).Elem())
	var tst = make([]IModl_condition_groupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_condition_groupContext)
		}
	}

	return tst
}

func (s *Modl_condition_testContext) Modl_condition_group(i int) IModl_condition_groupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_condition_groupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_condition_groupContext)
}

func (s *Modl_condition_testContext) AllEXCLAM() []antlr.TerminalNode {
	return s.GetTokens(MODLParserEXCLAM)
}

func (s *Modl_condition_testContext) EXCLAM(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserEXCLAM, i)
}

func (s *Modl_condition_testContext) AllAMP() []antlr.TerminalNode {
	return s.GetTokens(MODLParserAMP)
}

func (s *Modl_condition_testContext) AMP(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserAMP, i)
}

func (s *Modl_condition_testContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(MODLParserPIPE)
}

func (s *Modl_condition_testContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserPIPE, i)
}

func (s *Modl_condition_testContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_condition_testContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_condition_testContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_condition_test(s)
	}
}

func (s *Modl_condition_testContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_condition_test(s)
	}
}

func (p *MODLParser) Modl_condition_test() (localctx IModl_condition_testContext) {
	localctx = NewModl_condition_testContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MODLParserRULE_modl_condition_test)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(254)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(253)
			p.Match(MODLParserEXCLAM)
		}

	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(256)
			p.Modl_condition()
		}

	case 2:
		{
			p.SetState(257)
			p.Modl_condition_group()
		}

	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(260)
				_la = p.GetTokenStream().LA(1)

				if !(_la == MODLParserAMP || _la == MODLParserPIPE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			p.SetState(262)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(261)
					p.Match(MODLParserEXCLAM)
				}

			}
			p.SetState(266)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(264)
					p.Modl_condition()
				}

			case 2:
				{
					p.SetState(265)
					p.Modl_condition_group()
				}

			}

		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}

	return localctx
}

// IModl_operatorContext is an interface to support dynamic dispatch.
type IModl_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_operatorContext differentiates from other interfaces.
	IsModl_operatorContext()
}

type Modl_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_operatorContext() *Modl_operatorContext {
	var p = new(Modl_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_operator
	return p
}

func (*Modl_operatorContext) IsModl_operatorContext() {}

func NewModl_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_operatorContext {
	var p = new(Modl_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_operator

	return p
}

func (s *Modl_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_operatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(MODLParserEQUALS, 0)
}

func (s *Modl_operatorContext) GTHAN() antlr.TerminalNode {
	return s.GetToken(MODLParserGTHAN, 0)
}

func (s *Modl_operatorContext) LTHAN() antlr.TerminalNode {
	return s.GetToken(MODLParserLTHAN, 0)
}

func (s *Modl_operatorContext) EXCLAM() antlr.TerminalNode {
	return s.GetToken(MODLParserEXCLAM, 0)
}

func (s *Modl_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_operator(s)
	}
}

func (s *Modl_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_operator(s)
	}
}

func (p *MODLParser) Modl_operator() (localctx IModl_operatorContext) {
	localctx = NewModl_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MODLParserRULE_modl_operator)

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

	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)
			p.Match(MODLParserEQUALS)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Match(MODLParserGTHAN)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(275)
			p.Match(MODLParserGTHAN)
		}
		{
			p.SetState(276)
			p.Match(MODLParserEQUALS)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(277)
			p.Match(MODLParserLTHAN)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(278)
			p.Match(MODLParserLTHAN)
		}
		{
			p.SetState(279)
			p.Match(MODLParserEQUALS)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(280)
			p.Match(MODLParserEXCLAM)
		}
		{
			p.SetState(281)
			p.Match(MODLParserEQUALS)
		}

	}

	return localctx
}

// IModl_conditionContext is an interface to support dynamic dispatch.
type IModl_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_conditionContext differentiates from other interfaces.
	IsModl_conditionContext()
}

type Modl_conditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_conditionContext() *Modl_conditionContext {
	var p = new(Modl_conditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_condition
	return p
}

func (*Modl_conditionContext) IsModl_conditionContext() {}

func NewModl_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_conditionContext {
	var p = new(Modl_conditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_condition

	return p
}

func (s *Modl_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_conditionContext) AllModl_value() []IModl_valueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_valueContext)(nil)).Elem())
	var tst = make([]IModl_valueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_valueContext)
		}
	}

	return tst
}

func (s *Modl_conditionContext) Modl_value(i int) IModl_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_valueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_valueContext)
}

func (s *Modl_conditionContext) STRING() antlr.TerminalNode {
	return s.GetToken(MODLParserSTRING, 0)
}

func (s *Modl_conditionContext) Modl_operator() IModl_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_operatorContext)
}

func (s *Modl_conditionContext) AllFSLASH() []antlr.TerminalNode {
	return s.GetTokens(MODLParserFSLASH)
}

func (s *Modl_conditionContext) FSLASH(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserFSLASH, i)
}

func (s *Modl_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_condition(s)
	}
}

func (s *Modl_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_condition(s)
	}
}

func (p *MODLParser) Modl_condition() (localctx IModl_conditionContext) {
	localctx = NewModl_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MODLParserRULE_modl_condition)
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
	p.SetState(285)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(284)
			p.Match(MODLParserSTRING)
		}

	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MODLParserEQUALS)|(1<<MODLParserGTHAN)|(1<<MODLParserLTHAN)|(1<<MODLParserEXCLAM))) != 0 {
		{
			p.SetState(287)
			p.Modl_operator()
		}

	}
	{
		p.SetState(290)
		p.Modl_value()
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MODLParserFSLASH {
		{
			p.SetState(291)
			p.Match(MODLParserFSLASH)
		}
		{
			p.SetState(292)
			p.Modl_value()
		}

		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IModl_condition_groupContext is an interface to support dynamic dispatch.
type IModl_condition_groupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_condition_groupContext differentiates from other interfaces.
	IsModl_condition_groupContext()
}

type Modl_condition_groupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_condition_groupContext() *Modl_condition_groupContext {
	var p = new(Modl_condition_groupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_condition_group
	return p
}

func (*Modl_condition_groupContext) IsModl_condition_groupContext() {}

func NewModl_condition_groupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_condition_groupContext {
	var p = new(Modl_condition_groupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_condition_group

	return p
}

func (s *Modl_condition_groupContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_condition_groupContext) LCBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserLCBRAC, 0)
}

func (s *Modl_condition_groupContext) AllModl_condition_test() []IModl_condition_testContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_condition_testContext)(nil)).Elem())
	var tst = make([]IModl_condition_testContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_condition_testContext)
		}
	}

	return tst
}

func (s *Modl_condition_groupContext) Modl_condition_test(i int) IModl_condition_testContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_condition_testContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_condition_testContext)
}

func (s *Modl_condition_groupContext) RCBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserRCBRAC, 0)
}

func (s *Modl_condition_groupContext) AllAMP() []antlr.TerminalNode {
	return s.GetTokens(MODLParserAMP)
}

func (s *Modl_condition_groupContext) AMP(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserAMP, i)
}

func (s *Modl_condition_groupContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(MODLParserPIPE)
}

func (s *Modl_condition_groupContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserPIPE, i)
}

func (s *Modl_condition_groupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_condition_groupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_condition_groupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_condition_group(s)
	}
}

func (s *Modl_condition_groupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_condition_group(s)
	}
}

func (p *MODLParser) Modl_condition_group() (localctx IModl_condition_groupContext) {
	localctx = NewModl_condition_groupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MODLParserRULE_modl_condition_group)
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
	{
		p.SetState(298)
		p.Match(MODLParserLCBRAC)
	}
	{
		p.SetState(299)
		p.Modl_condition_test()
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MODLParserAMP || _la == MODLParserPIPE {
		{
			p.SetState(300)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MODLParserAMP || _la == MODLParserPIPE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(301)
			p.Modl_condition_test()
		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(307)
		p.Match(MODLParserRCBRAC)
	}

	return localctx
}

// IModl_valueContext is an interface to support dynamic dispatch.
type IModl_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_valueContext differentiates from other interfaces.
	IsModl_valueContext()
}

type Modl_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_valueContext() *Modl_valueContext {
	var p = new(Modl_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_value
	return p
}

func (*Modl_valueContext) IsModl_valueContext() {}

func NewModl_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_valueContext {
	var p = new(Modl_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_value

	return p
}

func (s *Modl_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_valueContext) Modl_map() IModl_mapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_mapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_mapContext)
}

func (s *Modl_valueContext) Modl_array() IModl_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_arrayContext)
}

func (s *Modl_valueContext) Modl_nb_array() IModl_nb_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_nb_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_nb_arrayContext)
}

func (s *Modl_valueContext) Modl_pair() IModl_pairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_pairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_pairContext)
}

func (s *Modl_valueContext) Modl_primitive() IModl_primitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_primitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_primitiveContext)
}

func (s *Modl_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_value(s)
	}
}

func (s *Modl_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_value(s)
	}
}

func (p *MODLParser) Modl_value() (localctx IModl_valueContext) {
	localctx = NewModl_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MODLParserRULE_modl_value)

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

	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.Modl_map()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)
			p.Modl_array()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(311)
			p.Modl_nb_array()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(312)
			p.Modl_pair()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(313)
			p.Modl_primitive()
		}

	}

	return localctx
}

// IModl_array_value_itemContext is an interface to support dynamic dispatch.
type IModl_array_value_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_array_value_itemContext differentiates from other interfaces.
	IsModl_array_value_itemContext()
}

type Modl_array_value_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_array_value_itemContext() *Modl_array_value_itemContext {
	var p = new(Modl_array_value_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_array_value_item
	return p
}

func (*Modl_array_value_itemContext) IsModl_array_value_itemContext() {}

func NewModl_array_value_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_array_value_itemContext {
	var p = new(Modl_array_value_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_array_value_item

	return p
}

func (s *Modl_array_value_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_array_value_itemContext) Modl_map() IModl_mapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_mapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_mapContext)
}

func (s *Modl_array_value_itemContext) Modl_pair() IModl_pairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_pairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_pairContext)
}

func (s *Modl_array_value_itemContext) Modl_array() IModl_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_arrayContext)
}

func (s *Modl_array_value_itemContext) Modl_primitive() IModl_primitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_primitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_primitiveContext)
}

func (s *Modl_array_value_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_array_value_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_array_value_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_array_value_item(s)
	}
}

func (s *Modl_array_value_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_array_value_item(s)
	}
}

func (p *MODLParser) Modl_array_value_item() (localctx IModl_array_value_itemContext) {
	localctx = NewModl_array_value_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MODLParserRULE_modl_array_value_item)

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

	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(316)
			p.Modl_map()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(317)
			p.Modl_pair()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(318)
			p.Modl_array()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(319)
			p.Modl_primitive()
		}

	}

	return localctx
}

// IModl_primitiveContext is an interface to support dynamic dispatch.
type IModl_primitiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_primitiveContext differentiates from other interfaces.
	IsModl_primitiveContext()
}

type Modl_primitiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_primitiveContext() *Modl_primitiveContext {
	var p = new(Modl_primitiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_primitive
	return p
}

func (*Modl_primitiveContext) IsModl_primitiveContext() {}

func NewModl_primitiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_primitiveContext {
	var p = new(Modl_primitiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_primitive

	return p
}

func (s *Modl_primitiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_primitiveContext) QUOTED() antlr.TerminalNode {
	return s.GetToken(MODLParserQUOTED, 0)
}

func (s *Modl_primitiveContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MODLParserNUMBER, 0)
}

func (s *Modl_primitiveContext) STRING() antlr.TerminalNode {
	return s.GetToken(MODLParserSTRING, 0)
}

func (s *Modl_primitiveContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MODLParserTRUE, 0)
}

func (s *Modl_primitiveContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MODLParserFALSE, 0)
}

func (s *Modl_primitiveContext) NULL() antlr.TerminalNode {
	return s.GetToken(MODLParserNULL, 0)
}

func (s *Modl_primitiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_primitiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_primitiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_primitive(s)
	}
}

func (s *Modl_primitiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_primitive(s)
	}
}

func (p *MODLParser) Modl_primitive() (localctx IModl_primitiveContext) {
	localctx = NewModl_primitiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MODLParserRULE_modl_primitive)
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
	{
		p.SetState(322)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MODLParserNULL)|(1<<MODLParserTRUE)|(1<<MODLParserFALSE)|(1<<MODLParserNUMBER)|(1<<MODLParserQUOTED)|(1<<MODLParserSTRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

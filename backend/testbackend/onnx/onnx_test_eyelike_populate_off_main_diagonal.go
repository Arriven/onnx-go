package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("EyeLike", "TestEyelikePopulateOffMainDiagonal", NewTestEyelikePopulateOffMainDiagonal)
}

// NewTestEyelikePopulateOffMainDiagonal version: 3.
func NewTestEyelikePopulateOffMainDiagonal() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "EyeLike",
		Title:  "TestEyelikePopulateOffMainDiagonal",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x7c, 0xa, 0x27, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x7, 0x45, 0x79, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x2a, 0xc, 0xa, 0x5, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x2a, 0x8, 0xa, 0x1, 0x6b, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x27, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x79, 0x65, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x64, 0x69, 0x61, 0x67, 0x6f, 0x6e, 0x61, 0x6c, 0x5a, 0x13, 0xa, 0x1, 0x78, 0x12, 0xe, 0xa, 0xc, 0x8, 0x6, 0x12, 0x8, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x13, 0xa, 0x1, 0x79, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "EyeLike",
		     Attributes: ([]*pb.AttributeProto) (len=2 cap=2) {
		    (*pb.AttributeProto)(0xc000119600)(name:"dtype" type:INT i:1 ),
		    (*pb.AttributeProto)(0xc000119700)(name:"k" type:INT i:1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(4, 5),
				tensor.WithBacking([]int32{44, 47, 64, 67, 67, 9, 83, 21, 36, 87, 70, 88, 88, 12, 58, 65, 39, 87, 46, 88}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(4, 5),
				tensor.WithBacking([]float32{0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1}),
			),
		},
	}
}

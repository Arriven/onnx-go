package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Split", "TestSplitVariableParts2d", NewTestSplitVariableParts2d)
}

// NewTestSplitVariableParts2d version: 3.
func NewTestSplitVariableParts2d() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Split",
		Title:  "TestSplitVariableParts2d",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xb0, 0x1, 0xa, 0x3f, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x31, 0x12, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x32, 0x22, 0x5, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x2a, 0xb, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x2a, 0xe, 0xa, 0x5, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x40, 0x2, 0x40, 0x4, 0xa0, 0x1, 0x7, 0x12, 0x1c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x32, 0x64, 0x5a, 0x17, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x6, 0x62, 0x1a, 0xa, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x31, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0x62, 0x1a, 0xa, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x32, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"input"},
		     Output:    []string{"output_1", "output_2"},
		     Name:      "",
		     OpType:    "Split",
		     Attributes: ([]*pb.AttributeProto) (len=2 cap=2) {
		    (*pb.AttributeProto)(0xc000126600)(name:"axis" type:INT i:1 ),
		    (*pb.AttributeProto)(0xc000126700)(name:"split" type:INTS ints:2 ints:4 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 6),
				tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 2),
				tensor.WithBacking([]float32{1, 2, 7, 8}),
			),

			tensor.New(
				tensor.WithShape(2, 4),
				tensor.WithBacking([]float32{3, 4, 5, 6, 9, 10, 11, 12}),
			),
		},
	}
}

package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("ReverseSequence", "TestReversesequenceBatch", NewTestReversesequenceBatch)
}

// NewTestReversesequenceBatch version: 5.
func NewTestReversesequenceBatch() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "ReverseSequence",
		Title:  "TestReversesequenceBatch",
		ModelB: []byte{0x8, 0x5, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xb0, 0x1, 0xa, 0x4b, 0xa, 0x1, 0x78, 0xa, 0xd, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x73, 0x12, 0x1, 0x79, 0x22, 0xf, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x2a, 0x11, 0xa, 0xa, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x18, 0x0, 0xa0, 0x1, 0x2, 0x2a, 0x10, 0xa, 0x9, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5a, 0x13, 0xa, 0x1, 0x78, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x4, 0x5a, 0x1b, 0xa, 0xd, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x4, 0x62, 0x13, 0xa, 0x1, 0x79, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0xa},

		/*

		   &pb.NodeProto{
		     Input:     []string{"x", "sequence_lens"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "ReverseSequence",
		     Attributes: ([]*pb.AttributeProto) (len=2 cap=2) {
		    (*pb.AttributeProto)(0xc000358700)(name:"batch_axis" type:INT ),
		    (*pb.AttributeProto)(0xc000358800)(name:"time_axis" type:INT i:1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(4, 4),
				tensor.WithBacking([]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
			),

			tensor.New(
				tensor.WithShape(4),
				tensor.WithBacking([]int64{1, 2, 3, 4}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(4, 4),
				tensor.WithBacking([]float32{0, 1, 2, 3, 5, 4, 6, 7, 10, 9, 8, 11, 15, 14, 13, 12}),
			),
		},
	}
}

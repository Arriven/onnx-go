package onnx

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gorgonia.org/tensor"
)

func TestNewTensor(t *testing.T) {
	tp := &pb.TensorProto{
		Dims:      []int64{1, 1, 1, 1},
		DataType:  pb.TensorProto_DataType_value["FLOAT"],
		FloatData: []float32{1},
	}
	b, err := proto.Marshal(tp)
	if err != nil {
		t.Fatal(err)
	}
	ts, err := NewTensor(b)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tensor.ElEq(ts, tensor.New(tensor.WithShape(1, 1, 1, 1), tensor.Of(tensor.Float32), tensor.WithBacking([]float32{1})))
	if err != nil {
		t.Fatal(err)
	}
}

package stdops

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

import (
	"context"
	"runtime/trace"

	gcontext "gorgonia.org/gorgonia/internal/context"
	"gorgonia.org/gorgonia/values"
	"gorgonia.org/tensor"
)

// powOp is the base op for elementwise exponentiation.
type powOp struct{ binop }

// String implements fmt.Stringer.
func (op powOp) String() string { return "^" }

// Do performs elementwise exponentiation.
func (op powOp) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := gcontext.Handle(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.Pow(a, b, tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise exponentiation but with a preallocated return value.
// PreallocDo allows pow to implement ops.PreallocOp.
func (op powOp) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := gcontext.Handle(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.Pow(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// powVV is a tensor-tensor elementwise exponentiation.
type powVV struct {
	powOp
	binopVV
}

// powVS is a tensor-scalar elementwise exponentiation.
type powVS struct {
	powOp
	binopVS
}

// String implements fmt.Stringer.
func (op powVS) String() string { return "^·" }

// powSV is a scalar-tensor elementwise exponentiation.
type powSV struct {
	powOp
	binopSV
}

// String implements fmt.Stringer.
func (op powSV) String() string { return "·^" }

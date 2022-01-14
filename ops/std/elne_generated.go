package stdops

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

import (
	"context"
	"runtime/trace"

	"github.com/chewxy/hm"
	gctx "gorgonia.org/gorgonia/internal/context"
	"gorgonia.org/gorgonia/types"
	"gorgonia.org/gorgonia/values"
	"gorgonia.org/tensor"
)

// elNeOp is the base op for elementwise not-equal-to.
type elNeOp struct {
	binop
	retSame bool
}

// String implements fmt.Stringer.
func (op elNeOp) String() string { return "≠" }

// Do performs elementwise not-equal-to.
func (op elNeOp) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := gctx.Handle(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	// Do the actual operation
	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.ElNe(a, b, tensor.WithContext(ctx2), tensor.AsSameType())
	} else {
		retVal, err = tensor.ElNe(a, b, tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise not-equal-to but with a preallocated return value.
// PreallocDo allows elNe to implement ops.PreallocOp.
func (op elNeOp) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := gctx.Handle(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.ElNe(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2), tensor.AsSameType())
	} else {
		retVal, err = tensor.ElNe(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}                                           // DiffWRT returns {false, false} for elNe
func (op elNeOp) DiffWRT(inputs int) []bool { return twofalses }

// elNeVV is a tensor-tensor elementwise not-equal-to.
type elNeVV struct {
	elNeOp
	binopVV
}

// Type returns the type: (·) : a → a → a or (·) :  a → a → b
func (op elNeVV) Type() hm.Type {
	a := hm.TypeVariable('a') // (T U) or U
	if op.retSame {
		return types.NewFunc(a, a, a)
	}
	b := types.MakeDependent(a, tensor.Bool) // (T Bool) or Bool
	return types.NewFunc(a, a, b)
}

// elNeVS is a tensor-scalar elementwise not-equal-to.
type elNeVS struct {
	elNeOp
	binopVS
}

// String implements fmt.Stringer.
func (op elNeVS) String() string { return "≠·" }

// Type returns the type: (·) : a → b → a or (·) :  a → b → c
func (op elNeVS) Type() hm.Type {
	a := hm.TypeVariable('a') // (T U) or U
	b := hm.TypeVariable('b') // U
	if op.retSame {
		return types.NewFunc(a, b, a)
	}
	c := types.MakeDependent(a, tensor.Bool) // (T Bool) or Bool
	return types.NewFunc(a, b, c)
}

// elNeSV is a scalar-tensor elementwise not-equal-to.
type elNeSV struct {
	elNeOp
	binopSV
}

// String implements fmt.Stringer.
func (op elNeSV) String() string { return "·≠" }

// Type returns the type: (·) : a → b → b or (·) :  a → b → c
func (op elNeSV) Type() hm.Type {
	a := hm.TypeVariable('a') // U
	b := hm.TypeVariable('b') // (T U) or U
	if op.retSame {
		return types.NewFunc(a, b, b)
	}
	c := types.MakeDependent(b, tensor.Bool) // (T Bool) or Bool
	return types.NewFunc(a, b, c)
}

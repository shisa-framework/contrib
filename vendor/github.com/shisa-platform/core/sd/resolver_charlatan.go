// generated by "charlatan -output=./resolver_charlatan.go Resolver".  DO NOT EDIT.

package sd

import "github.com/ansel1/merry"

import "reflect"

// ResolverResolveInvocation represents a single call of FakeResolver.Resolve
type ResolverResolveInvocation struct {
	Parameters struct {
		Name string
	}
	Results struct {
		Ident1 []string
		Ident2 merry.Error
	}
}

// NewResolverResolveInvocation creates a new instance of ResolverResolveInvocation
func NewResolverResolveInvocation(name string, ident1 []string, ident2 merry.Error) *ResolverResolveInvocation {
	invocation := new(ResolverResolveInvocation)

	invocation.Parameters.Name = name

	invocation.Results.Ident1 = ident1
	invocation.Results.Ident2 = ident2

	return invocation
}

// ResolverTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type ResolverTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeResolver is a mock implementation of Resolver for testing.
Use it in your tests as in this example:

	package example

	func TestWithResolver(t *testing.T) {
		f := &sd.FakeResolver{
			ResolveHook: func(name string) (ident1 []string, ident2 merry.Error) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeResolve ...
		f.AssertResolveCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeResolve.
*/
type FakeResolver struct {
	ResolveHook func(string) ([]string, merry.Error)

	ResolveCalls []*ResolverResolveInvocation
}

// NewFakeResolverDefaultPanic returns an instance of FakeResolver with all hooks configured to panic
func NewFakeResolverDefaultPanic() *FakeResolver {
	return &FakeResolver{
		ResolveHook: func(string) (ident1 []string, ident2 merry.Error) {
			panic("Unexpected call to Resolver.Resolve")
		},
	}
}

// NewFakeResolverDefaultFatal returns an instance of FakeResolver with all hooks configured to call t.Fatal
func NewFakeResolverDefaultFatal(t ResolverTestingT) *FakeResolver {
	return &FakeResolver{
		ResolveHook: func(string) (ident1 []string, ident2 merry.Error) {
			t.Fatal("Unexpected call to Resolver.Resolve")
			return
		},
	}
}

// NewFakeResolverDefaultError returns an instance of FakeResolver with all hooks configured to call t.Error
func NewFakeResolverDefaultError(t ResolverTestingT) *FakeResolver {
	return &FakeResolver{
		ResolveHook: func(string) (ident1 []string, ident2 merry.Error) {
			t.Error("Unexpected call to Resolver.Resolve")
			return
		},
	}
}

func (f *FakeResolver) Reset() {
	f.ResolveCalls = []*ResolverResolveInvocation{}
}

func (_f1 *FakeResolver) Resolve(name string) (ident1 []string, ident2 merry.Error) {
	if _f1.ResolveHook == nil {
		panic("Resolver.Resolve() called but FakeResolver.ResolveHook is nil")
	}

	invocation := new(ResolverResolveInvocation)
	_f1.ResolveCalls = append(_f1.ResolveCalls, invocation)

	invocation.Parameters.Name = name

	ident1, ident2 = _f1.ResolveHook(name)

	invocation.Results.Ident1 = ident1
	invocation.Results.Ident2 = ident2

	return
}

// SetResolveStub configures Resolver.Resolve to always return the given values
func (_f2 *FakeResolver) SetResolveStub(ident1 []string, ident2 merry.Error) {
	_f2.ResolveHook = func(string) ([]string, merry.Error) {
		return ident1, ident2
	}
}

// SetResolveInvocation configures Resolver.Resolve to return the given results when called with the given parameters
// If no match is found for an invocation the result(s) of the fallback function are returned
func (_f3 *FakeResolver) SetResolveInvocation(calls_f4 []*ResolverResolveInvocation, fallback_f5 func() ([]string, merry.Error)) {
	_f3.ResolveHook = func(name string) (ident1 []string, ident2 merry.Error) {
		for _, call := range calls_f4 {
			if reflect.DeepEqual(call.Parameters.Name, name) {
				ident1 = call.Results.Ident1
				ident2 = call.Results.Ident2

				return
			}
		}

		return fallback_f5()
	}
}

// ResolveCalled returns true if FakeResolver.Resolve was called
func (f *FakeResolver) ResolveCalled() bool {
	return len(f.ResolveCalls) != 0
}

// AssertResolveCalled calls t.Error if FakeResolver.Resolve was not called
func (f *FakeResolver) AssertResolveCalled(t ResolverTestingT) {
	t.Helper()
	if len(f.ResolveCalls) == 0 {
		t.Error("FakeResolver.Resolve not called, expected at least one")
	}
}

// ResolveNotCalled returns true if FakeResolver.Resolve was not called
func (f *FakeResolver) ResolveNotCalled() bool {
	return len(f.ResolveCalls) == 0
}

// AssertResolveNotCalled calls t.Error if FakeResolver.Resolve was called
func (f *FakeResolver) AssertResolveNotCalled(t ResolverTestingT) {
	t.Helper()
	if len(f.ResolveCalls) != 0 {
		t.Error("FakeResolver.Resolve called, expected none")
	}
}

// ResolveCalledOnce returns true if FakeResolver.Resolve was called exactly once
func (f *FakeResolver) ResolveCalledOnce() bool {
	return len(f.ResolveCalls) == 1
}

// AssertResolveCalledOnce calls t.Error if FakeResolver.Resolve was not called exactly once
func (f *FakeResolver) AssertResolveCalledOnce(t ResolverTestingT) {
	t.Helper()
	if len(f.ResolveCalls) != 1 {
		t.Errorf("FakeResolver.Resolve called %d times, expected 1", len(f.ResolveCalls))
	}
}

// ResolveCalledN returns true if FakeResolver.Resolve was called at least n times
func (f *FakeResolver) ResolveCalledN(n int) bool {
	return len(f.ResolveCalls) >= n
}

// AssertResolveCalledN calls t.Error if FakeResolver.Resolve was called less than n times
func (f *FakeResolver) AssertResolveCalledN(t ResolverTestingT, n int) {
	t.Helper()
	if len(f.ResolveCalls) < n {
		t.Errorf("FakeResolver.Resolve called %d times, expected >= %d", len(f.ResolveCalls), n)
	}
}

// ResolveCalledWith returns true if FakeResolver.Resolve was called with the given values
func (_f6 *FakeResolver) ResolveCalledWith(name string) (found bool) {
	for _, call := range _f6.ResolveCalls {
		if reflect.DeepEqual(call.Parameters.Name, name) {
			found = true
			break
		}
	}

	return
}

// AssertResolveCalledWith calls t.Error if FakeResolver.Resolve was not called with the given values
func (_f7 *FakeResolver) AssertResolveCalledWith(t ResolverTestingT, name string) {
	t.Helper()
	var found bool
	for _, call := range _f7.ResolveCalls {
		if reflect.DeepEqual(call.Parameters.Name, name) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeResolver.Resolve not called with expected parameters")
	}
}

// ResolveCalledOnceWith returns true if FakeResolver.Resolve was called exactly once with the given values
func (_f8 *FakeResolver) ResolveCalledOnceWith(name string) bool {
	var count int
	for _, call := range _f8.ResolveCalls {
		if reflect.DeepEqual(call.Parameters.Name, name) {
			count++
		}
	}

	return count == 1
}

// AssertResolveCalledOnceWith calls t.Error if FakeResolver.Resolve was not called exactly once with the given values
func (_f9 *FakeResolver) AssertResolveCalledOnceWith(t ResolverTestingT, name string) {
	t.Helper()
	var count int
	for _, call := range _f9.ResolveCalls {
		if reflect.DeepEqual(call.Parameters.Name, name) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeResolver.Resolve called %d times with expected parameters, expected one", count)
	}
}

// ResolveResultsForCall returns the result values for the first call to FakeResolver.Resolve with the given values
func (_f10 *FakeResolver) ResolveResultsForCall(name string) (ident1 []string, ident2 merry.Error, found bool) {
	for _, call := range _f10.ResolveCalls {
		if reflect.DeepEqual(call.Parameters.Name, name) {
			ident1 = call.Results.Ident1
			ident2 = call.Results.Ident2
			found = true
			break
		}
	}

	return
}

// generated by "charlatan -output=./sentrycapturer_charlatan.go capturer".  DO NOT EDIT.

package crash

import "reflect"

import "github.com/getsentry/raven-go"

// capturerCaptureInvocation represents a single call of Fakecapturer.Capture
type capturerCaptureInvocation struct {
	Parameters struct {
		Ident1 *raven.Packet
		Ident2 map[string]string
	}
	Results struct {
		Ident3 string
		Ident4 chan error
	}
}

// NewcapturerCaptureInvocation creates a new instance of capturerCaptureInvocation
func NewcapturerCaptureInvocation(ident1 *raven.Packet, ident2 map[string]string, ident3 string, ident4 chan error) *capturerCaptureInvocation {
	invocation := new(capturerCaptureInvocation)

	invocation.Parameters.Ident1 = ident1
	invocation.Parameters.Ident2 = ident2

	invocation.Results.Ident3 = ident3
	invocation.Results.Ident4 = ident4

	return invocation
}

// capturerCloseInvocation represents a single call of Fakecapturer.Close
type capturerCloseInvocation struct {
}

// capturerTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type capturerTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
Fakecapturer is a mock implementation of capturer for testing.
Use it in your tests as in this example:

	package example

	func TestWithcapturer(t *testing.T) {
		f := &crash.Fakecapturer{
			CaptureHook: func(ident1 *raven.Packet, ident2 map[string]string) (ident3 string, ident4 chan error) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeCapture ...
		f.AssertCaptureCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeCapture.
*/
type Fakecapturer struct {
	CaptureHook func(*raven.Packet, map[string]string) (string, chan error)
	CloseHook   func()

	CaptureCalls []*capturerCaptureInvocation
	CloseCalls   []*capturerCloseInvocation
}

// NewFakecapturerDefaultPanic returns an instance of Fakecapturer with all hooks configured to panic
func NewFakecapturerDefaultPanic() *Fakecapturer {
	return &Fakecapturer{
		CaptureHook: func(*raven.Packet, map[string]string) (ident3 string, ident4 chan error) {
			panic("Unexpected call to capturer.Capture")
		},
		CloseHook: func() {
			panic("Unexpected call to capturer.Close")
		},
	}
}

// NewFakecapturerDefaultFatal returns an instance of Fakecapturer with all hooks configured to call t.Fatal
func NewFakecapturerDefaultFatal(t capturerTestingT) *Fakecapturer {
	return &Fakecapturer{
		CaptureHook: func(*raven.Packet, map[string]string) (ident3 string, ident4 chan error) {
			t.Fatal("Unexpected call to capturer.Capture")
			return
		},
		CloseHook: func() {
			t.Fatal("Unexpected call to capturer.Close")
			return
		},
	}
}

// NewFakecapturerDefaultError returns an instance of Fakecapturer with all hooks configured to call t.Error
func NewFakecapturerDefaultError(t capturerTestingT) *Fakecapturer {
	return &Fakecapturer{
		CaptureHook: func(*raven.Packet, map[string]string) (ident3 string, ident4 chan error) {
			t.Error("Unexpected call to capturer.Capture")
			return
		},
		CloseHook: func() {
			t.Error("Unexpected call to capturer.Close")
			return
		},
	}
}

func (f *Fakecapturer) Reset() {
	f.CaptureCalls = []*capturerCaptureInvocation{}
	f.CloseCalls = []*capturerCloseInvocation{}
}

func (_f1 *Fakecapturer) Capture(ident1 *raven.Packet, ident2 map[string]string) (ident3 string, ident4 chan error) {
	if _f1.CaptureHook == nil {
		panic("capturer.Capture() called but Fakecapturer.CaptureHook is nil")
	}

	invocation := new(capturerCaptureInvocation)
	_f1.CaptureCalls = append(_f1.CaptureCalls, invocation)

	invocation.Parameters.Ident1 = ident1
	invocation.Parameters.Ident2 = ident2

	ident3, ident4 = _f1.CaptureHook(ident1, ident2)

	invocation.Results.Ident3 = ident3
	invocation.Results.Ident4 = ident4

	return
}

// SetCaptureStub configures capturer.Capture to always return the given values
func (_f2 *Fakecapturer) SetCaptureStub(ident3 string, ident4 chan error) {
	_f2.CaptureHook = func(*raven.Packet, map[string]string) (string, chan error) {
		return ident3, ident4
	}
}

// SetCaptureInvocation configures capturer.Capture to return the given results when called with the given parameters
// If no match is found for an invocation the result(s) of the fallback function are returned
func (_f3 *Fakecapturer) SetCaptureInvocation(calls_f4 []*capturerCaptureInvocation, fallback_f5 func() (string, chan error)) {
	_f3.CaptureHook = func(ident1 *raven.Packet, ident2 map[string]string) (ident3 string, ident4 chan error) {
		for _, call := range calls_f4 {
			if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
				ident3 = call.Results.Ident3
				ident4 = call.Results.Ident4

				return
			}
		}

		return fallback_f5()
	}
}

// CaptureCalled returns true if Fakecapturer.Capture was called
func (f *Fakecapturer) CaptureCalled() bool {
	return len(f.CaptureCalls) != 0
}

// AssertCaptureCalled calls t.Error if Fakecapturer.Capture was not called
func (f *Fakecapturer) AssertCaptureCalled(t capturerTestingT) {
	t.Helper()
	if len(f.CaptureCalls) == 0 {
		t.Error("Fakecapturer.Capture not called, expected at least one")
	}
}

// CaptureNotCalled returns true if Fakecapturer.Capture was not called
func (f *Fakecapturer) CaptureNotCalled() bool {
	return len(f.CaptureCalls) == 0
}

// AssertCaptureNotCalled calls t.Error if Fakecapturer.Capture was called
func (f *Fakecapturer) AssertCaptureNotCalled(t capturerTestingT) {
	t.Helper()
	if len(f.CaptureCalls) != 0 {
		t.Error("Fakecapturer.Capture called, expected none")
	}
}

// CaptureCalledOnce returns true if Fakecapturer.Capture was called exactly once
func (f *Fakecapturer) CaptureCalledOnce() bool {
	return len(f.CaptureCalls) == 1
}

// AssertCaptureCalledOnce calls t.Error if Fakecapturer.Capture was not called exactly once
func (f *Fakecapturer) AssertCaptureCalledOnce(t capturerTestingT) {
	t.Helper()
	if len(f.CaptureCalls) != 1 {
		t.Errorf("Fakecapturer.Capture called %d times, expected 1", len(f.CaptureCalls))
	}
}

// CaptureCalledN returns true if Fakecapturer.Capture was called at least n times
func (f *Fakecapturer) CaptureCalledN(n int) bool {
	return len(f.CaptureCalls) >= n
}

// AssertCaptureCalledN calls t.Error if Fakecapturer.Capture was called less than n times
func (f *Fakecapturer) AssertCaptureCalledN(t capturerTestingT, n int) {
	t.Helper()
	if len(f.CaptureCalls) < n {
		t.Errorf("Fakecapturer.Capture called %d times, expected >= %d", len(f.CaptureCalls), n)
	}
}

// CaptureCalledWith returns true if Fakecapturer.Capture was called with the given values
func (_f6 *Fakecapturer) CaptureCalledWith(ident1 *raven.Packet, ident2 map[string]string) (found bool) {
	for _, call := range _f6.CaptureCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			found = true
			break
		}
	}

	return
}

// AssertCaptureCalledWith calls t.Error if Fakecapturer.Capture was not called with the given values
func (_f7 *Fakecapturer) AssertCaptureCalledWith(t capturerTestingT, ident1 *raven.Packet, ident2 map[string]string) {
	t.Helper()
	var found bool
	for _, call := range _f7.CaptureCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			found = true
			break
		}
	}

	if !found {
		t.Error("Fakecapturer.Capture not called with expected parameters")
	}
}

// CaptureCalledOnceWith returns true if Fakecapturer.Capture was called exactly once with the given values
func (_f8 *Fakecapturer) CaptureCalledOnceWith(ident1 *raven.Packet, ident2 map[string]string) bool {
	var count int
	for _, call := range _f8.CaptureCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			count++
		}
	}

	return count == 1
}

// AssertCaptureCalledOnceWith calls t.Error if Fakecapturer.Capture was not called exactly once with the given values
func (_f9 *Fakecapturer) AssertCaptureCalledOnceWith(t capturerTestingT, ident1 *raven.Packet, ident2 map[string]string) {
	t.Helper()
	var count int
	for _, call := range _f9.CaptureCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("Fakecapturer.Capture called %d times with expected parameters, expected one", count)
	}
}

// CaptureResultsForCall returns the result values for the first call to Fakecapturer.Capture with the given values
func (_f10 *Fakecapturer) CaptureResultsForCall(ident1 *raven.Packet, ident2 map[string]string) (ident3 string, ident4 chan error, found bool) {
	for _, call := range _f10.CaptureCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			ident3 = call.Results.Ident3
			ident4 = call.Results.Ident4
			found = true
			break
		}
	}

	return
}

func (_f11 *Fakecapturer) Close() {
	if _f11.CloseHook == nil {
		panic("capturer.Close() called but Fakecapturer.CloseHook is nil")
	}

	invocation := new(capturerCloseInvocation)
	_f11.CloseCalls = append(_f11.CloseCalls, invocation)

	_f11.CloseHook()

	return
}

// CloseCalled returns true if Fakecapturer.Close was called
func (f *Fakecapturer) CloseCalled() bool {
	return len(f.CloseCalls) != 0
}

// AssertCloseCalled calls t.Error if Fakecapturer.Close was not called
func (f *Fakecapturer) AssertCloseCalled(t capturerTestingT) {
	t.Helper()
	if len(f.CloseCalls) == 0 {
		t.Error("Fakecapturer.Close not called, expected at least one")
	}
}

// CloseNotCalled returns true if Fakecapturer.Close was not called
func (f *Fakecapturer) CloseNotCalled() bool {
	return len(f.CloseCalls) == 0
}

// AssertCloseNotCalled calls t.Error if Fakecapturer.Close was called
func (f *Fakecapturer) AssertCloseNotCalled(t capturerTestingT) {
	t.Helper()
	if len(f.CloseCalls) != 0 {
		t.Error("Fakecapturer.Close called, expected none")
	}
}

// CloseCalledOnce returns true if Fakecapturer.Close was called exactly once
func (f *Fakecapturer) CloseCalledOnce() bool {
	return len(f.CloseCalls) == 1
}

// AssertCloseCalledOnce calls t.Error if Fakecapturer.Close was not called exactly once
func (f *Fakecapturer) AssertCloseCalledOnce(t capturerTestingT) {
	t.Helper()
	if len(f.CloseCalls) != 1 {
		t.Errorf("Fakecapturer.Close called %d times, expected 1", len(f.CloseCalls))
	}
}

// CloseCalledN returns true if Fakecapturer.Close was called at least n times
func (f *Fakecapturer) CloseCalledN(n int) bool {
	return len(f.CloseCalls) >= n
}

// AssertCloseCalledN calls t.Error if Fakecapturer.Close was called less than n times
func (f *Fakecapturer) AssertCloseCalledN(t capturerTestingT, n int) {
	t.Helper()
	if len(f.CloseCalls) < n {
		t.Errorf("Fakecapturer.Close called %d times, expected >= %d", len(f.CloseCalls), n)
	}
}
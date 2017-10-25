// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type TableWriter struct {
	SetHeaderStub        func([]string)
	setHeaderMutex       sync.RWMutex
	setHeaderArgsForCall []struct {
		arg1 []string
	}
	AppendStub        func([]string)
	appendMutex       sync.RWMutex
	appendArgsForCall []struct {
		arg1 []string
	}
	SetAlignmentStub        func(int)
	setAlignmentMutex       sync.RWMutex
	setAlignmentArgsForCall []struct {
		arg1 int
	}
	RenderStub                      func()
	renderMutex                     sync.RWMutex
	renderArgsForCall               []struct{}
	SetAutoFormatHeadersStub        func(bool)
	setAutoFormatHeadersMutex       sync.RWMutex
	setAutoFormatHeadersArgsForCall []struct {
		arg1 bool
	}
	SetAutoWrapTextStub        func(bool)
	setAutoWrapTextMutex       sync.RWMutex
	setAutoWrapTextArgsForCall []struct {
		arg1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TableWriter) SetHeader(arg1 []string) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setHeaderMutex.Lock()
	fake.setHeaderArgsForCall = append(fake.setHeaderArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("SetHeader", []interface{}{arg1Copy})
	fake.setHeaderMutex.Unlock()
	if fake.SetHeaderStub != nil {
		fake.SetHeaderStub(arg1)
	}
}

func (fake *TableWriter) SetHeaderCallCount() int {
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	return len(fake.setHeaderArgsForCall)
}

func (fake *TableWriter) SetHeaderArgsForCall(i int) []string {
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	return fake.setHeaderArgsForCall[i].arg1
}

func (fake *TableWriter) Append(arg1 []string) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.appendMutex.Lock()
	fake.appendArgsForCall = append(fake.appendArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("Append", []interface{}{arg1Copy})
	fake.appendMutex.Unlock()
	if fake.AppendStub != nil {
		fake.AppendStub(arg1)
	}
}

func (fake *TableWriter) AppendCallCount() int {
	fake.appendMutex.RLock()
	defer fake.appendMutex.RUnlock()
	return len(fake.appendArgsForCall)
}

func (fake *TableWriter) AppendArgsForCall(i int) []string {
	fake.appendMutex.RLock()
	defer fake.appendMutex.RUnlock()
	return fake.appendArgsForCall[i].arg1
}

func (fake *TableWriter) SetAlignment(arg1 int) {
	fake.setAlignmentMutex.Lock()
	fake.setAlignmentArgsForCall = append(fake.setAlignmentArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("SetAlignment", []interface{}{arg1})
	fake.setAlignmentMutex.Unlock()
	if fake.SetAlignmentStub != nil {
		fake.SetAlignmentStub(arg1)
	}
}

func (fake *TableWriter) SetAlignmentCallCount() int {
	fake.setAlignmentMutex.RLock()
	defer fake.setAlignmentMutex.RUnlock()
	return len(fake.setAlignmentArgsForCall)
}

func (fake *TableWriter) SetAlignmentArgsForCall(i int) int {
	fake.setAlignmentMutex.RLock()
	defer fake.setAlignmentMutex.RUnlock()
	return fake.setAlignmentArgsForCall[i].arg1
}

func (fake *TableWriter) Render() {
	fake.renderMutex.Lock()
	fake.renderArgsForCall = append(fake.renderArgsForCall, struct{}{})
	fake.recordInvocation("Render", []interface{}{})
	fake.renderMutex.Unlock()
	if fake.RenderStub != nil {
		fake.RenderStub()
	}
}

func (fake *TableWriter) RenderCallCount() int {
	fake.renderMutex.RLock()
	defer fake.renderMutex.RUnlock()
	return len(fake.renderArgsForCall)
}

func (fake *TableWriter) SetAutoFormatHeaders(arg1 bool) {
	fake.setAutoFormatHeadersMutex.Lock()
	fake.setAutoFormatHeadersArgsForCall = append(fake.setAutoFormatHeadersArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("SetAutoFormatHeaders", []interface{}{arg1})
	fake.setAutoFormatHeadersMutex.Unlock()
	if fake.SetAutoFormatHeadersStub != nil {
		fake.SetAutoFormatHeadersStub(arg1)
	}
}

func (fake *TableWriter) SetAutoFormatHeadersCallCount() int {
	fake.setAutoFormatHeadersMutex.RLock()
	defer fake.setAutoFormatHeadersMutex.RUnlock()
	return len(fake.setAutoFormatHeadersArgsForCall)
}

func (fake *TableWriter) SetAutoFormatHeadersArgsForCall(i int) bool {
	fake.setAutoFormatHeadersMutex.RLock()
	defer fake.setAutoFormatHeadersMutex.RUnlock()
	return fake.setAutoFormatHeadersArgsForCall[i].arg1
}

func (fake *TableWriter) SetAutoWrapText(arg1 bool) {
	fake.setAutoWrapTextMutex.Lock()
	fake.setAutoWrapTextArgsForCall = append(fake.setAutoWrapTextArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("SetAutoWrapText", []interface{}{arg1})
	fake.setAutoWrapTextMutex.Unlock()
	if fake.SetAutoWrapTextStub != nil {
		fake.SetAutoWrapTextStub(arg1)
	}
}

func (fake *TableWriter) SetAutoWrapTextCallCount() int {
	fake.setAutoWrapTextMutex.RLock()
	defer fake.setAutoWrapTextMutex.RUnlock()
	return len(fake.setAutoWrapTextArgsForCall)
}

func (fake *TableWriter) SetAutoWrapTextArgsForCall(i int) bool {
	fake.setAutoWrapTextMutex.RLock()
	defer fake.setAutoWrapTextMutex.RUnlock()
	return fake.setAutoWrapTextArgsForCall[i].arg1
}

func (fake *TableWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	fake.appendMutex.RLock()
	defer fake.appendMutex.RUnlock()
	fake.setAlignmentMutex.RLock()
	defer fake.setAlignmentMutex.RUnlock()
	fake.renderMutex.RLock()
	defer fake.renderMutex.RUnlock()
	fake.setAutoFormatHeadersMutex.RLock()
	defer fake.setAutoFormatHeadersMutex.RUnlock()
	fake.setAutoWrapTextMutex.RLock()
	defer fake.setAutoWrapTextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *TableWriter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"os"
	"sync"
)

type IOReadWriter struct {
	ReadDirStub        func(string) ([]os.FileInfo, error)
	readDirMutex       sync.RWMutex
	readDirArgsForCall []struct {
		arg1 string
	}
	readDirReturns struct {
		result1 []os.FileInfo
		result2 error
	}
	readDirReturnsOnCall map[int]struct {
		result1 []os.FileInfo
		result2 error
	}
	ReadFileStub        func(string) ([]byte, error)
	readFileMutex       sync.RWMutex
	readFileArgsForCall []struct {
		arg1 string
	}
	readFileReturns struct {
		result1 []byte
		result2 error
	}
	readFileReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	RemoveStub        func(name string) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		name string
	}
	removeReturns struct {
		result1 error
	}
	removeReturnsOnCall map[int]struct {
		result1 error
	}
	StatStub        func(string) (os.FileInfo, error)
	statMutex       sync.RWMutex
	statArgsForCall []struct {
		arg1 string
	}
	statReturns struct {
		result1 os.FileInfo
		result2 error
	}
	statReturnsOnCall map[int]struct {
		result1 os.FileInfo
		result2 error
	}
	WriteFileStub        func(string, []byte, os.FileMode) error
	writeFileMutex       sync.RWMutex
	writeFileArgsForCall []struct {
		arg1 string
		arg2 []byte
		arg3 os.FileMode
	}
	writeFileReturns struct {
		result1 error
	}
	writeFileReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *IOReadWriter) ReadDir(arg1 string) ([]os.FileInfo, error) {
	fake.readDirMutex.Lock()
	ret, specificReturn := fake.readDirReturnsOnCall[len(fake.readDirArgsForCall)]
	fake.readDirArgsForCall = append(fake.readDirArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReadDir", []interface{}{arg1})
	fake.readDirMutex.Unlock()
	if fake.ReadDirStub != nil {
		return fake.ReadDirStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.readDirReturns.result1, fake.readDirReturns.result2
}

func (fake *IOReadWriter) ReadDirCallCount() int {
	fake.readDirMutex.RLock()
	defer fake.readDirMutex.RUnlock()
	return len(fake.readDirArgsForCall)
}

func (fake *IOReadWriter) ReadDirArgsForCall(i int) string {
	fake.readDirMutex.RLock()
	defer fake.readDirMutex.RUnlock()
	return fake.readDirArgsForCall[i].arg1
}

func (fake *IOReadWriter) ReadDirReturns(result1 []os.FileInfo, result2 error) {
	fake.ReadDirStub = nil
	fake.readDirReturns = struct {
		result1 []os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *IOReadWriter) ReadDirReturnsOnCall(i int, result1 []os.FileInfo, result2 error) {
	fake.ReadDirStub = nil
	if fake.readDirReturnsOnCall == nil {
		fake.readDirReturnsOnCall = make(map[int]struct {
			result1 []os.FileInfo
			result2 error
		})
	}
	fake.readDirReturnsOnCall[i] = struct {
		result1 []os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *IOReadWriter) ReadFile(arg1 string) ([]byte, error) {
	fake.readFileMutex.Lock()
	ret, specificReturn := fake.readFileReturnsOnCall[len(fake.readFileArgsForCall)]
	fake.readFileArgsForCall = append(fake.readFileArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReadFile", []interface{}{arg1})
	fake.readFileMutex.Unlock()
	if fake.ReadFileStub != nil {
		return fake.ReadFileStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.readFileReturns.result1, fake.readFileReturns.result2
}

func (fake *IOReadWriter) ReadFileCallCount() int {
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	return len(fake.readFileArgsForCall)
}

func (fake *IOReadWriter) ReadFileArgsForCall(i int) string {
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	return fake.readFileArgsForCall[i].arg1
}

func (fake *IOReadWriter) ReadFileReturns(result1 []byte, result2 error) {
	fake.ReadFileStub = nil
	fake.readFileReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *IOReadWriter) ReadFileReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.ReadFileStub = nil
	if fake.readFileReturnsOnCall == nil {
		fake.readFileReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.readFileReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *IOReadWriter) Remove(name string) error {
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Remove", []interface{}{name})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeReturns.result1
}

func (fake *IOReadWriter) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *IOReadWriter) RemoveArgsForCall(i int) string {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.removeArgsForCall[i].name
}

func (fake *IOReadWriter) RemoveReturns(result1 error) {
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *IOReadWriter) RemoveReturnsOnCall(i int, result1 error) {
	fake.RemoveStub = nil
	if fake.removeReturnsOnCall == nil {
		fake.removeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IOReadWriter) Stat(arg1 string) (os.FileInfo, error) {
	fake.statMutex.Lock()
	ret, specificReturn := fake.statReturnsOnCall[len(fake.statArgsForCall)]
	fake.statArgsForCall = append(fake.statArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Stat", []interface{}{arg1})
	fake.statMutex.Unlock()
	if fake.StatStub != nil {
		return fake.StatStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.statReturns.result1, fake.statReturns.result2
}

func (fake *IOReadWriter) StatCallCount() int {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return len(fake.statArgsForCall)
}

func (fake *IOReadWriter) StatArgsForCall(i int) string {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return fake.statArgsForCall[i].arg1
}

func (fake *IOReadWriter) StatReturns(result1 os.FileInfo, result2 error) {
	fake.StatStub = nil
	fake.statReturns = struct {
		result1 os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *IOReadWriter) StatReturnsOnCall(i int, result1 os.FileInfo, result2 error) {
	fake.StatStub = nil
	if fake.statReturnsOnCall == nil {
		fake.statReturnsOnCall = make(map[int]struct {
			result1 os.FileInfo
			result2 error
		})
	}
	fake.statReturnsOnCall[i] = struct {
		result1 os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *IOReadWriter) WriteFile(arg1 string, arg2 []byte, arg3 os.FileMode) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.writeFileMutex.Lock()
	ret, specificReturn := fake.writeFileReturnsOnCall[len(fake.writeFileArgsForCall)]
	fake.writeFileArgsForCall = append(fake.writeFileArgsForCall, struct {
		arg1 string
		arg2 []byte
		arg3 os.FileMode
	}{arg1, arg2Copy, arg3})
	fake.recordInvocation("WriteFile", []interface{}{arg1, arg2Copy, arg3})
	fake.writeFileMutex.Unlock()
	if fake.WriteFileStub != nil {
		return fake.WriteFileStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.writeFileReturns.result1
}

func (fake *IOReadWriter) WriteFileCallCount() int {
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	return len(fake.writeFileArgsForCall)
}

func (fake *IOReadWriter) WriteFileArgsForCall(i int) (string, []byte, os.FileMode) {
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	return fake.writeFileArgsForCall[i].arg1, fake.writeFileArgsForCall[i].arg2, fake.writeFileArgsForCall[i].arg3
}

func (fake *IOReadWriter) WriteFileReturns(result1 error) {
	fake.WriteFileStub = nil
	fake.writeFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *IOReadWriter) WriteFileReturnsOnCall(i int, result1 error) {
	fake.WriteFileStub = nil
	if fake.writeFileReturnsOnCall == nil {
		fake.writeFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IOReadWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readDirMutex.RLock()
	defer fake.readDirMutex.RUnlock()
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *IOReadWriter) recordInvocation(key string, args []interface{}) {
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

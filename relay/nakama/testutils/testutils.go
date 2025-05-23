package testutils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"testing"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/mock"

	"pkg.world.dev/world-engine/relay/nakama/mocks"
)

// This file contains helpers that are common across all tests.
//
// The mocks/ directory was generated with mockery. The Nakama interfaces are unlikely to change, so regenerating the
// mocks will likely not be required. That being said, here are instructions for regenerating the mocks.
// Install mockery. On Mac:
//
//	$ brew install mockery
//
// Run mockery:
//
//	$ mockery
//
// The configuration file at .mockery.yaml will be used to generate the Nakama mocks.

// AnyContext is used to make mock expectations more readable.
// someMock.On("SomeFunction", anyContext...) makes it clear that the first parameter is supposed to be a context.
// context.valueCtx is the type returned by context.Background.
var AnyContext = mock.AnythingOfType("*context.valueCtx")

// Ensure that FakeLogger implements runtime.Logger (this will produce a compile-time error if it doesn't).
var _ runtime.Logger = (*FakeLogger)(nil)

type keyTuple struct {
	collection string
	key        string
	userID     string
}

type storeValue struct {
	Value   string
	Version string
}

type FakeLogger struct {
	runtime.Logger
	mu     sync.Mutex
	Errors []string
}

func (l *FakeLogger) Debug(string, ...interface{}) {}
func (l *FakeLogger) Info(string, ...interface{})  {}
func (l *FakeLogger) Warn(string, ...interface{})  {}

// Capture error messages
//
//nolint:goprintffuncname // [not important]
func (l *FakeLogger) Error(format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Errors = append(l.Errors, fmt.Sprintf(format, args...))
}

// GetErrors A method to retrieve captured errors for assertions.
func (l *FakeLogger) GetErrors() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Errors
}

// CtxWithUserID saves the given user ID to the background context in a location that Nakama expects to find user IDs.
func CtxWithUserID(userID string) context.Context {
	ctx := context.Background()
	//nolint:staticcheck // this is how Nakama reads userIDs from the context.
	return context.WithValue(ctx, runtime.RUNTIME_CTX_USER_ID, userID)
}

func MockMatchStoreWrite(collection, key, userID string) any {
	return mock.MatchedBy(func(writes []*runtime.StorageWrite) bool {
		if len(writes) != 1 {
			return false
		}
		if collection != "" && writes[0].Collection != collection {
			return false
		}
		if key != "" && writes[0].Key != key {
			return false
		}
		if userID != "" && writes[0].UserID != userID {
			return false
		}
		return true
	})
}

// MockMatchStoreRead creates a mock.Matcher (suitable for use as a variadic argument into an "On" method).
// It should be used when your test is expecting a call to "StorageRead", and it verifies that the single
// read request matches the given collection/key/userID. Use the empty string ("") to skip any of the comparisons.
func MockMatchStoreRead(collection, key, userID string) any {
	return mock.MatchedBy(func(reads []*runtime.StorageRead) bool {
		if len(reads) != 1 {
			return false
		}
		if collection != "" && reads[0].Collection != collection {
			return false
		}
		if key != "" && reads[0].Key != key {
			return false
		}
		if userID != "" && reads[0].UserID != userID {
			return false
		}
		return true
	})
}

// MockNoopLogger returns a mock logger that ignores all log messages.
func MockNoopLogger(t *testing.T) runtime.Logger {
	mockLog := mocks.NewMockLogger(t)
	mockLog.On("Error", mock.Anything).Return().Maybe()
	mockLog.On("Debug", mock.Anything).Return().Maybe()
	mockLog.On("Info", mock.Anything).Return().Maybe()
	return mockLog
}

func MockMatchReadKey(key string) interface{} {
	return mock.MatchedBy(func(storeRead []*runtime.StorageRead) bool {
		if len(storeRead) != 1 {
			return false
		}
		return storeRead[0].Key == key
	})
}

func MockMatchWriteKey(key string) interface{} {
	return mock.MatchedBy(func(storeWrite []*runtime.StorageWrite) bool {
		if len(storeWrite) != 1 {
			return false
		}
		return storeWrite[0].Key == key
	})
}

// FakeNakamaModule is a fake implementation of runtime.NakamaModule that ONLY implements the StorageRead and
// StorageWrite methods. Under the hood, a map is used to map collection/key/userID tuples onto the stored values.
// Calling other methods on the NakamaModule interface will panic. In addition, searching for values (e.g. specifying
// a collection, but no user ID) will not return the correct results.
//
// This Fake implements the atomic guarantees of StorageRead and StorageWrite.
// In addition, Version field is populated during StorageRead which allows for compare-and-swap writes.
type FakeNakamaModule struct { //nolint:decorder
	runtime.NakamaModule
	sync.Mutex
	versionItr   int
	store        map[keyTuple]storeValue
	errsToReturn []error
}

func NewFakeNakamaModule() *FakeNakamaModule {
	return &FakeNakamaModule{
		store: map[keyTuple]storeValue{},
	}
}

// WithError modifies the FakeNakamaModule to return the given error the next time StorageRead or StorageWrite is
// called.
func (f *FakeNakamaModule) WithError(err error) *FakeNakamaModule {
	f.errsToReturn = append(f.errsToReturn, err)
	return f
}

func (f *FakeNakamaModule) StorageRead(_ context.Context, reads []*runtime.StorageRead) ([]*api.StorageObject, error) {
	f.Lock()
	defer f.Unlock()
	if len(f.errsToReturn) > 0 {
		var err error
		err, f.errsToReturn = f.errsToReturn[0], f.errsToReturn[1:]
		return nil, err
	}
	var results []*api.StorageObject
	for _, read := range reads {
		key := keyTuple{read.Collection, read.Key, read.UserID}
		value, ok := f.store[key]
		if ok {
			results = append(results, &api.StorageObject{
				Collection: read.Collection,
				Key:        read.Key,
				UserId:     read.UserID,
				Version:    value.Version,
				Value:      value.Value,
			})
		}
	}
	return results, nil
}

func (f *FakeNakamaModule) nextVersion() string {
	f.versionItr++
	return strconv.Itoa(f.versionItr)
}

func (f *FakeNakamaModule) StorageWrite(
	_ context.Context, writes []*runtime.StorageWrite,
) (
	[]*api.StorageObjectAck, error,
) {
	f.Lock()
	defer f.Unlock()
	if len(f.errsToReturn) > 0 {
		var err error
		err, f.errsToReturn = f.errsToReturn[0], f.errsToReturn[1:]
		return nil, err
	}
	for _, write := range writes {
		key := keyTuple{write.Collection, write.Key, write.UserID}
		currValue, ok := f.store[key]
		if write.Version == "*" {
			// The "*" version instructs Nakama to only save the item if it doesn't already exist.
			// See: https://heroiclabs.com/docs/nakama/concepts/storage/collections/#conditional-writes
			if ok {
				return nil, errors.New("item already exists in storage")
			}
		} else if write.Version != "" {
			// If this write request has a version, and it doesn't match the saved version, return an error.
			// A missing currValue will also result in an error here.
			if currValue.Version != write.Version {
				return nil, errors.New("write failed due to out of date version")
			}
		}

		// It is a Nakama requirement that this value field is JSON encoded
		if !json.Valid([]byte(write.Value)) {
			return nil, errors.New("value field must be JSON encoded")
		}
	}
	// All write items can safely be stored in the storage map.
	for _, write := range writes {
		key := keyTuple{write.Collection, write.Key, write.UserID}
		f.store[key] = storeValue{
			Value:   write.Value,
			Version: f.nextVersion(),
		}
	}
	return nil, nil
}

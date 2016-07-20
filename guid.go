package gox

import (
	"strings"

	"github.com/satori/go.uuid"
)

func Uuid0() uuid.UUID                          { return uuid.Nil }
func Uuid1() uuid.UUID                          { return uuid.NewV1() }
func Uuid3(ns uuid.UUID, name string) uuid.UUID { return uuid.NewV3(ns, name) }
func Uuid4() uuid.UUID                          { return uuid.NewV4() }
func Uuid5(ns uuid.UUID, name string) uuid.UUID { return uuid.NewV5(ns, name) }

type GUID string

func Guid() GUID                 { return GUID(Uuid1().String()) }
func (p GUID) String() string    { return string(p) }
func (p GUID) StrNoDash() string { return strings.Replace(string(p), "-", "", -1) }
func (p GUID) IsEmpty() bool     { return len(p) == 0 }
func (p GUID) IsZero() bool      { return string(p) == uuid.Nil.String() }

var (
	GuidEmpty = GUID("")
	GuidZero  = GUID(uuid.Nil.String())
)

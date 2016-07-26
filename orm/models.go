package orm

import "encoding/json"

type Jsonb struct {
	Id   string          `Id`
	Json json.RawMessage `Json`
}

package convert

import (
	"database/sql"
)


func StrToNullString(str string) (data sql.NullString){
	if str == "" {
		data.Valid = false
	} else {
		data.Valid = true
	}
	data.String = str
	return
}
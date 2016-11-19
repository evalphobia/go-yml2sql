package yml2sql

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"bitbucket.org/pkg/inflect"
	"gopkg.in/yaml.v2"
)

const (
	NamingTypeDir  = "dir"
	NamingTypeFile = "file"
)

var (
	pluralFlag   = false
	nullableFlag = false
	namingType   = NamingTypeDir
)

// CreateStatementByFile returns SQL INSERT statement
func CreateStatementByFile(file string) string {
	tableName := getTableName(file)
	m := createMapData(file)
	schema, values := schemaStrings(m)

	return fmt.Sprintf("INSERT INTO\n  `%s`(%s)\nVALUES\n  %s; \n",
		tableName,
		schema,
		values)
}

// SetPlural set plural flag
func SetPlural(b bool) {
	pluralFlag = b
}

func isPlural() bool {
	return pluralFlag
}

// SetNullable set nullable flag
func SetNullable(b bool) {
	nullableFlag = b
}

func isNullable() bool {
	return nullableFlag
}

// SetNamingTypeDir set naming type to dir or file
func SetNamingTypeDir(b bool) {
	if b {
		namingType = NamingTypeDir
	} else {
		namingType = NamingTypeFile
	}
}

func isNamingTypeDir() bool {
	return namingType == NamingTypeDir
}

func createMapData(path string) []map[string]interface{} {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	var m []map[string]interface{}
	yaml.Unmarshal(buf, &m)
	return m
}

func getTableName(path string) string {
	paths := strings.Split(path, "/")

	var name string
	if isNamingTypeDir() {
		name = paths[len(paths)-2]
	} else {
		fileName := paths[len(paths)-1]
		name = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	}
	// add 's' when plural flag is set
	if isPlural() {
		return inflect.Pluralize(name)
	}
	return name
}

func schemaStrings(m []map[string]interface{}) (string, string) {
	keys := getKeys(m[0])
	schema := encodeKeys(keys)
	values := encodeValues(keys, m)
	return schema, values
}

func getKeys(m map[string]interface{}) []string {
	var keys []string
	for key := range m {
		keys = append(keys, fmt.Sprint(key))
	}
	sort.Strings(keys)
	return keys
}

func encodeKeys(keys []string) string {
	var result []string
	for _, key := range keys {
		result = append(result, fmt.Sprintf("`%s`", key))
	}
	return strings.Join(result, ", ")
}

func encodeValues(keys []string, m []map[string]interface{}) string {
	var result []string
	for _, row := range m {
		var v []string
		for _, key := range keys {
			v = append(v, toString(row[key]))
		}
		result = append(result,
			fmt.Sprintf("(%s)", strings.Join(v, ",")))
	}
	return strings.Join(result, ",\n  ")
}

func toString(value interface{}) string {
	switch t := value.(type) {
	case string:
		if isNullString(t) {
			return "NULL"
		}
		return "'" + t + "'"
	case int:
		return fmt.Sprint(t)
	case nil:
		if isNullable() {
			return "NULL"
		} else {
			return "''"
		}
	default:
		return fmt.Sprint(t)
	}
}

func isNullString(v string) bool {
	str := strings.ToLower(v)
	switch {
	case str == "null":
		return true
	case str == "nil":
		return true
	case str == "" && isNullable():
		return true
	}
	return false
}

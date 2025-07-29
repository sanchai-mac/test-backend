package container_mock

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"test-backend/internal/config"
	"test-backend/internal/container"
	"test-backend/internal/infrastructure/database"

	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.uber.org/dig"
)

// Setup uber.go Container for testing
type TestContainer struct {
	*dig.Container

	t testing.TB
}

func (c TestContainer) RequireProvide(f interface{}, opts ...dig.ProvideOption) {
	c.t.Helper()

	require.NoError(c.t, c.Provide(f, opts...), "failed to provide")
}

func (c TestContainer) RequireInvoke(f interface{}, opts ...dig.InvokeOption) {
	c.t.Helper()

	require.NoError(c.t, c.Invoke(f, opts...), "failed to invoke")
}

func (c TestContainer) RequireDecorate(f interface{}, opts ...dig.DecorateOption) {
	c.t.Helper()

	require.NoError(c.t, c.Decorate(f, opts...), "failed to decorate")
}

// Json to Struct
func (c TestContainer) JsonToStruct(filePath string, db *database.DB) {
	fileContent, err := ioutil.ReadFile(filePath)
	require.NoError(c.t, err, fmt.Sprintf("%v doesn't exist.", filePath))

	var mockData []struct {
		TableName string `json:"table_name"`
		Records   []map[string]any
	}

	err = json.Unmarshal(fileContent, &mockData)
	require.NoError(c.t, err)

	for _, data := range mockData {
		for _, rec := range data.Records {
			formatKey, formatVals, err := formatData(rec)
			require.NoError(c.t, err, "Error formatting data")

			query := fmt.Sprintf(`INSERT INTO %s %s VALUES %s`, data.TableName, formatKey, formatVals)
			log.Printf("Executing query: %s\n", query)
		}
	}
}

func formatData(data map[string]any) (string, string, error) {
	var keys, vals []string
	for key, value := range data {
		switch v := value.(type) {
		case int, float64, bool:
			vals = append(vals, fmt.Sprintf("%v", v))
		case string:
			vals = append(vals, fmt.Sprintf("'%s'", v))
		case nil:
			vals = append(vals, fmt.Sprintf("%s", "NULL"))
		default:
			return "", "", errors.New("invalid value type")
		}
		keys = append(keys, key)
	}
	formatKey := fmt.Sprintf("(%s)", strings.Join(keys, ", "))
	formatVals := fmt.Sprintf("(%s)", strings.Join(vals, ", "))

	return formatKey, formatVals, nil
}

// Csv to Struct
func (c TestContainer) CsvToStruct(filePath string, to interface{}) ([]interface{}, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Unable to parse file as CSV for "+filePath, err)
	}
	if len(records) <= 1 {
		return []interface{}{}, nil
	}

	headers := records[0]
	body := records[1:]
	var mapped []interface{}
	for _, row := range body {
		structValue := reflect.ValueOf(to).Elem()

		for j, value := range row {
			colVal := value
			if strings.ToUpper(colVal) == "NULL" {
				continue
			}
			header := headers[j]
			field := structValue.FieldByName(header)
			if !field.IsValid() {
				return nil, fmt.Errorf("no such field: %s in obj", header)
			}
			switch field.Interface().(type) {
			case string:
				field.SetString(colVal)
			case *string:
				field.Set(reflect.ValueOf(&colVal))
			case uint64, int32, *int64:
				parsed, err := strconv.Atoi(colVal)
				parsedValue := reflect.ValueOf(int64(parsed))
				if err != nil {
					return nil, err
				}
				if parsedValue.CanConvert(field.Type()) {
					converted := parsedValue.Convert(field.Type())
					field.Set(converted)
				} else {
					t := reflect.New(parsedValue.Type())
					t.Elem().Set(parsedValue)
					field.Set(t)
				}
			case float64:
				parsed, err := strconv.ParseFloat(colVal, 64)
				parsedValue := reflect.ValueOf(parsed)
				if err != nil {
					return nil, err
				}
				if parsedValue.CanConvert(field.Type()) {
					converted := parsedValue.Convert(field.Type())
					field.Set(converted)
				} else {
					return nil, fmt.Errorf("can't convert to type: %s", field.Type().String())
				}
			case *float64, *float32:
				parsed, err := strconv.ParseFloat(colVal, 64)
				parsedValue := reflect.ValueOf(&parsed)
				switch field.Interface().(type) {
				case *float32:
					parsd32 := float32(parsed)
					parsedValue = reflect.ValueOf(&parsd32)
				}

				if err != nil {
					return nil, err
				}
				field.Set(parsedValue)
			case time.Time, *time.Time:
				var parsedT time.Time
				if parsedT, err = time.Parse("2006-01-02 15:04:05 Z07:00", colVal); err != nil {
					parsedT, err = time.Parse("2006-01-02", colVal)
				}
				if err != nil {
					return nil, err
				}
				if field.Type() == reflect.TypeOf(new(time.Time)) {
					field.Set(reflect.ValueOf(&parsedT))
				} else {
					field.Set(reflect.ValueOf(parsedT))
				}
			case bool, *bool:
				boolVal := (colVal == "1" || strings.ToLower(colVal) == "true")
				if field.Type() == reflect.TypeOf(new(bool)) {
					field.Set(reflect.ValueOf(&boolVal))
				} else {
					field.SetBool(boolVal)
				}
			default:
				return nil, fmt.Errorf("CsvToStruct can't mapped type: %s", field.Type().String())
			}
		}
		mapped = append(mapped, structValue.Interface())
	}

	return mapped, nil
}

func NewTestContainer(t testing.TB, opts ...dig.Option) *TestContainer {
	c := container.NewContainer().Container
	c.Decorate(func(cfg *config.Configuration, db *database.DB) *database.DB {
		return db
	})
	return &TestContainer{
		t:         t,
		Container: c,
	}
}

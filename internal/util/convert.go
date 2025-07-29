package util

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"
)

// ConvertStructToMap ...
func ConvertStructToMap(data interface{}) map[string]interface{} {
	dataBytes, _ := json.Marshal(data)
	mapped := map[string]interface{}{}
	json.Unmarshal(dataBytes, &mapped)
	return mapped
}

// ConvertStringToMap ...
func ConvertStringToMap(raw string, dest *map[string]interface{}) {
	json.Unmarshal([]byte(raw), dest)
}

// ConvertMapToString ...
func ConvertMapToString(data map[string]interface{}) string {
	bConverted, _ := json.Marshal(data)
	return string(bConverted)
}

// ConvertToReader ...
func ConvertToReader(data interface{}) (*bytes.Reader, error) {
	dataBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(dataBytes), nil
}

// ConvertStructToJSONString ...
func ConvertStructToJSONString(data interface{}) string {
	strJson, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}

	return string(strJson)
}

// ConvertStructToPrettyJSONString ...
func ConvertStructToPrettyJSONString(data interface{}) string {
	strJson, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return err.Error()
	}

	return string(strJson)
}

// ConvertJSONByteToStruct ...
func ConvertJSONByteToStruct(data []byte, target interface{}) error {
	return json.Unmarshal(data, target)
}

// ConvertStructToJSONByte ...
func ConvertStructToJSONByte(data interface{}) ([]byte, error) {
	dataBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}

// ConvertRequestToString ...
func ConvertRequestToString(req *http.Request) (string, error) {
	data, err := httputil.DumpRequest(req, true)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// ConvertResponseToString ...
func ConvertResponseToString(resp *http.Response) (string, error) {
	data, err := httputil.DumpResponse(resp, true)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// ConvertPointerToFloat64 ...
func ConvertPointerToFloat64(float *float64) float64 {
	if float == nil {
		return 0
	}

	return *float
}

// ConvertFloat64ToPointer ...
func ConvertFloat64ToPointer(float float64) *float64 {
	if float == 0 {
		return nil
	}

	return &float
}

func ConvertYNToBool(a string) bool {
	if a == "Y" {
		return true
	} else {
		return false
	}
}

func ConvertToYNBoolPtr(a *string) *bool {
	if a != nil {
		var r bool
		if *a == "Y" {
			r = true
		} else {
			r = false
		}
		return &r
	}
	return nil
}

func SubString(s string, max int) string {
	if len(s) < max {
		return s
	}
	return s[:max]
}

// StringToInt ...
func StringToInt(text string) int {
	number, err := strconv.Atoi(text)
	if err != nil {
		return 0
	}
	return number
}

// StringToInt32 ...
func StringToInt32(text string) int32 {
	number, err := strconv.Atoi(text)
	if err != nil {
		return 0
	}
	return int32(number)
}

// findMinAndMax...
func FindMinAndMax(a []int32) (min int32, max int32) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

// ConvertPointerToString ...
func ConvertPointerToString(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

// ConvertStringToPointer ...
func ConvertStringToPointer(str string) *string {
	if str == "" {
		return nil
	}

	return &str
}

// ConvertPointerToBool ...
func ConvertPointerToBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// ConvertStringToDate ...
func ConvertStringToDate(date string) time.Time {
	t, _ := time.Parse("20060102", date)
	return t
}

// ConvertDateToString ...
func ConvertDateToString(date time.Time) string {
	strDate := date.Format("20060102")
	return strDate
}

// ConvertStringToDateTime ...
func ConvertStringToDateTime(date string) time.Time {
	t, _ := time.Parse("200601021504", date)
	return t
}

// GetSAReferenceNo .
func GetCloseSAReferenceNo() string {
	t := time.Now().Format("20060102150405")
	s := "PTMC" + t
	return s
}

func MapSlice(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// ConvertPointerToInt64 ...
func ConvertPointerToInt64(input *int64) int64 {
	if input == nil {
		return 0
	}

	return *input
}

// ConvertStringToInt64...
func ConvertStringToInt64(str string) int64 {
	i, _ := strconv.Atoi(str)
	return int64(i)
}

// ConvertStringToInt...
func ConvertStringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return int(i)
}

// ConvertStringToUInt64...
func ConvertStringToUInt64(str string) uint64 {
	i, _ := strconv.Atoi(str)
	return uint64(i)
}

func ConvertUint64(p *int) *uint64 {
	var in *uint64
	if p != nil {
		ui := uint64(*p)
		in = &ui
	}
	return in
}

func HashMD5OnBytes(data []byte) (string, error) {
	byteHashed := md5.Sum(data)
	fileHashed := base64.StdEncoding.EncodeToString(byteHashed[:])
	return fileHashed, nil
}

func ToTimeString(t time.Time) string {
	return t.Format("1504")
}

func ToPtrStr(str string) *string {
	r := &str
	return r
}

func ToDateString(t time.Time) string {
	return t.Format("20060102")
}

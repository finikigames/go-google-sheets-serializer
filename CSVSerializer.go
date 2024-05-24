package csvserializer

import (
    "fmt"
    "strconv"
    "reflect"
    "encoding/json"
)

func convertToStruct[T any](records [][]string) ([]T, error) {
    header := records[0]
    var result []T

    for _, record := range records[1:] {
        if len(record) < len(header) {
            return nil, fmt.Errorf("not enough data in row: %v", record)
        }

        var objMap = make(map[string]interface{})
        for i, value := range record {
            headerName := header[i]
            objMap[headerName] = value
        }

        var obj T
        v := reflect.ValueOf(&obj).Elem()
        for i := 0; i < v.NumField(); i++ {
            field := v.Field(i)
            if field.Kind() == reflect.Int {
                if strValue, ok := objMap[header[i]].(string); ok {
                    intValue, err := strconv.Atoi(strValue)
                    if err != nil {
                        return nil, err
                    }
                    objMap[header[i]] = intValue
                }
            }
        }

        objBytes, err := json.Marshal(objMap)
        if err != nil {
            return nil, fmt.Errorf("error marshaling to JSON: %v", err)
        }

        err = json.Unmarshal(objBytes, &obj)
        if err != nil {
            return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
        }

        result = append(result, obj)
    }

    return result, nil
}

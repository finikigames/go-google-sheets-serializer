package main

import (
    "fmt"
    "log"
    "strconv"
    "reflect"
    "encoding/json"
    "encoding/csv"
    "net/http"
    "main/models"
)

func main() {    
    //var csvUrl = "https://docs.google.com/spreadsheets/d/e/2PACX-1vRdOZ-Hxj351d42GDMWxaif4LFt4Z4P2M9qPm6hPPaggY8fX8Ebz-unOAVD9ehT1eY04GHngiZGGIBK/pub?gid=986548007&single=true&output=csv";
    var csvUrl = "https://docs.google.com/spreadsheets/d/e/2PACX-1vRdOZ-Hxj351d42GDMWxaif4LFt4Z4P2M9qPm6hPPaggY8fX8Ebz-unOAVD9ehT1eY04GHngiZGGIBK/pub?gid=2017613902&single=true&output=csv"
    
    resp, err := http.Get(csvUrl)
    if err != nil {
        fmt.Println("Cannot load CSV: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("Cannot load CSV: %v", err)
    }

    reader := csv.NewReader(resp.Body)

    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Cannot load CSV: %v", err)
    }

    //var items []models.GuildBoss
    var items []models.SeasonConfig
    if len(records) == 0 {
        fmt.Println("No data found.")
    } else {
        items, err = convertToStruct[models.SeasonConfig](records)
        if err != nil {
            log.Fatalf("Error converting data: %v", err)
        }

        fmt.Printf("Items: %+v\n", items)
    }
}

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

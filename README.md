Функция convertToStruct() в CSVSerializer.go десериализует данные из CSV в структуру Go.

Действия для проверки работоспособности:
1. Определите структуру, которая будет использоваться для десериализации данных (var items []models.Boss)
2. Generic функция convertToStruct:
    - Принимает список records (полученные по ссылки строки)
    - Использует заголовок из первой строки для создания карты значений
    - Преобразует карту значений в JSON и затем в структуру
    - Обрабатывает специальные поля, такие как целые числа

Example
```csharp
import (
    "log"
    "encoding/csv"
    "net/http"
    "main/models"
)

func main() {    
    var csvUrl = "path/to/csv";
    
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

    var items []models.Boss
    if len(records) == 0 {
        fmt.Println("No data found.")
    } else {
        items, err = convertToStruct[models.Boss](records)
        if err != nil {
            log.Fatalf("Error converting data: %v", err)
        }

        fmt.Printf("Items: %+v\n", items)
    }
}
```

Models example:
```csharp
type Boss struct {
    Id string `json:"id"`
    SetupId string `json:"setup_id"`
    Health int `json:"health"`
}
```

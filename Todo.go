package main
import "time"

type Todo struct {
    Name    string  'json:"name"'
    Completed   bool    'json:"completed"'
    Due     time.Time   'Jsons:"due"'
}

type Todos []Todo

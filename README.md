# Json library

## Candidate functions :

| Func Name |                                    Prototype                                     |                Description                 | Comments |
|:---------:|:--------------------------------------------------------------------------------:|:------------------------------------------:|:--------:|
|  Marshal  |  Marshal(content map\[string]string \| []map\[string]string) (string, error) {}  |     Returns the JSON encoding of a map     |   N/A    |
| Unmarshal | Unmarshal(content string) (map\[string]string \| []map\[string]string, error) {} | Returns the map representation of the JSON |   N/A    |
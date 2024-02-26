# Json library

[![Go Report Card](https://goreportcard.com/badge/github.com/Eclalang/json)](https://goreportcard.com/report/github.com/Eclalang/json)
[![codecov](https://codecov.io/gh/Eclalang/json/graph/badge.svg?token=YNCIYERVBO)](https://codecov.io/gh/Eclalang/json)

## Candidate functions :

| Func Name |                                    Prototype                                     |                Description                 | Comments |
|:---------:|:--------------------------------------------------------------------------------:|:------------------------------------------:|:--------:|
|  Marshal  |  Marshal(content map\[string]string \| []map\[string]string) (string, error) {}  |     Returns the JSON encoding of a map     |   N/A    |
| Unmarshal | Unmarshal(content string) (map\[string]string \| []map\[string]string, error) {} | Returns the map representation of the JSON |   N/A    |
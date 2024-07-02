package main

import (
	"fmt"
	"time"
)

type SymbolValue interface{} // this should be defined to a more specific interface IRL

type Symbol struct {
	key   string
	value SymbolValue
}

type SymbolTable map[string]Symbol

func (st SymbolTable) IsEmpty() bool {
	return len(st) == 0
}

func (st SymbolTable) Insert(symbol Symbol) {
	st[symbol.key] = Symbol{key: symbol.key, value: symbol.value}
}

func (st SymbolTable) Has(key string) bool {
	_, exists := st[key]
	return exists
}

func (st SymbolTable) Peek(key string) SymbolValue {
	symbol := st[key]
	return symbol.value
}

func (st SymbolTable) Delete(key string) {
	delete(st, key)
}

func main() {
	start := time.Now()

	symbolTable := &SymbolTable{}

	fmt.Println(symbolTable.IsEmpty()) // true
	fmt.Println(symbolTable)           // map[]
	symbolTable.Insert(Symbol{key: "testKey", value: 2})
	symbolTable.Insert(Symbol{key: "testKey2", value: "Test string"})
	fmt.Println(symbolTable.IsEmpty())        // false
	fmt.Println(symbolTable)                  // map[testKey:{testKey 2} testKey2:{testKey2 Test string}]
	fmt.Println(symbolTable.Peek("testKey"))  // 2
	fmt.Println(symbolTable.Has("testKey3"))  // false
	fmt.Println(symbolTable.Peek("testKey3")) // nil
	symbolTable.Delete("testKey2")
	fmt.Println(symbolTable) // map[testKey:{testKey 2}]

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}

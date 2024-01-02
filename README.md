# in-memory cache (Go)

Установка пакета (в терминале):
```shell
go get github.com/denisakhmetovdev/inMemoryCache  
```

## Добавления пакета в проект
`import github.com/denisakhmetovdev/inMemoryCache/pkg/inMemoryCache`
```go
package main

import (
"github.com/denisakhmetovdev/inMemoryCache/pkg/inMemoryCache"
)

func main(){
	// Ваш код
}
```

## Использование пакета
- Создание нового кеш-хранилища

```go
func main(){
	newCache := inMemoryCache.New()
}
```

- Создание новой пары ключ-значение используйте метод `Set()`. Ключом должна быть строка,
 а значением могут любые типы данных. Например:
```go
newCache.Set("userId": 1002)
newCache.Set("orders": []int{123, 234, 345})
```

- Получение значения по ключу - метод `Get()`. `Get()` принимает строку и возвращает значение по ключу и `bool`
```go
orders, ok := newCache.Get("orders")
if ok{
	fmt.Println(orders)
} else {
	fmt.Println("Такого ключа в кеше нет")
}
```

- Удалить пару ключ-значение - метод `Delete()`, который принимает ключ и ничего не возвращает. Если такого ключа нет,
 ничего не произойдёт и кеш останется прежним.
```go
newCache.Delete("orders")
```

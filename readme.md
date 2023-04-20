<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# gobidirectionalmap

```go
import "github.com/graytonio/go-bidirectional-map"
```

## Index

- [type Map](<#type-map>)
  - [func NewMap[T comparable, K comparable](input map[T]K) *Map[T, K]](<#func-newmap>)
  - [func (t *Map[T, K]) Associate(key T, value K)](<#func-mapt-k-associate>)
  - [func (t *Map[T, K]) GetP(key T) (K, bool)](<#func-mapt-k-getp>)
  - [func (t *Map[T, K]) GetS(key K) (T, bool)](<#func-mapt-k-gets>)


## type Map

A two way associative array the initial type is considered the primary key the second type is considered the value or secondary key

```go
type Map[T comparable, K comparable] struct {
    // contains filtered or unexported fields
}
```

### func NewMap

```go
func NewMap[T comparable, K comparable](input map[T]K) *Map[T, K]
```

Create a new map based on an existing map The key on this map will be considered the primary key and the value type the secondary key

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"

	bidmap "github.com/graytonio/go-bidirectional-map"
)

func main() {
	myMap := bidmap.NewMap(map[string]int{
		"foo": 3,
		"bar": 4,
	})

	fmt.Println(myMap.GetP("foo"))
}
```

#### Output

```
3 true
```

</p>
</details>

### func \(\*Map\[T, K\]\) Associate

```go
func (t *Map[T, K]) Associate(key T, value K)
```

Create a new association between the primary key and secondary key

### func \(\*Map\[T, K\]\) GetP

```go
func (t *Map[T, K]) GetP(key T) (K, bool)
```

Lookup a value based on the Primary key type

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"

	bidmap "github.com/graytonio/go-bidirectional-map"
)

func main() {
	myMap := bidmap.NewMap(map[string]int{
		"foo": 3,
		"bar": 4,
	})

	fmt.Println(myMap.GetP("foo"))
}
```

#### Output

```
3 true
```

</p>
</details>

### func \(\*Map\[T, K\]\) GetS

```go
func (t *Map[T, K]) GetS(key K) (T, bool)
```

Return a value based on the Secondary key type

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"

	bidmap "github.com/graytonio/go-bidirectional-map"
)

func main() {
	myMap := bidmap.NewMap(map[string]int{
		"foo": 3,
		"bar": 4,
	})

	fmt.Println(myMap.GetS(3))
}
```

#### Output

```
foo true
```

</p>
</details>



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)

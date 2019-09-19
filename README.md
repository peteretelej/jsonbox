# jsonbox Go Client SDK
[![GoDoc](https://godoc.org/github.com/peteretelej/jsonbox?status.svg)](https://godoc.org/github.com/peteretelej/jsonbox)

Go wrapper / Client SDK for [jsonbox](https://github.com/vasanthv/jsonbox)

- See [examples](./examples) for sample usage


## Usage

Import the package
```
import "github.com/peteretelej/jsonbox"
```

Use `NewClient` to get a new jsonbox Client to use
```
cl,err := jsonbox.NewClient("https://jsonbox.io/")
```

**Create**
```
//  Create a record
val := []byte(`{"name": "Jon Snow", "age": 25}`)
out, err := cl.Create("demobox_6d9e326c183fde7b",val)

// Create multiple records
val := []byte(`[{"name": "Daenerys Targaryen", "age": 25}, {"name": "Arya Stark", "age": 16}]`)
out, err := cl.Create("demobox_6d9e326c183fde7b",val)
```

**Read**
```
out, err := cl.Read("demobox_6d9e326c183fde7b")
fmt.Printf("%s",out)

// Query records for a boxid
out, err := cl.Read("demobox_6d9e326c183fde7b?query_key=name&query_value=arya%20stark")
fmt.Printf("%s",out)

```

**Update**
```
val := []byte(`{"name": "Arya Stark", "age": 18}`)
out, err := cl.Update("demobox_6d9e326c183fde7b","5d776b75fd6d3d6cb1d45c53",val)
fmt.Printf("%s",out)
```

**Delete** record from _BOXID_
```
err := cl.Delete("demobox_6d9e326c183fde7b","5d776b75fd6d3d6cb1d45c53")
```

**DeleteAll** records for a _BOX_ID_
```
err := cl.DeleteAll(BOXID)
```

**List IDs** for all records for a _BOX_ID_
```
ids,err := cl.IDs("demobox_6d9e326c183fde7b")
fmt.Printf("%s",ids)
```

See full example usage at [examples/full](./examples/full/main.go)

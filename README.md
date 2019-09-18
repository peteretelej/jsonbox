# jsonbox Go SDK
Go wrapper for [jsonbox](https://github.com/vasanthv/jsonbox)

See [examples](./examples)


Importing the package
```
import "github.com/peteretelej/jsonbox"
```

## Usage

Use `NewClient` to get a new jsonbox Client to use
```
cl,err := jsonbox.NewClient("https://jsonbox.io/")
```

Create
```
//  Create a record
val := []byte(`{"name": "Jon Snow", "age": 25}`)
out, err := cl.Create("demobox_6d9e326c183fde7b",val)

// Create multiple records
val := []byte(`[{"name": "Daenerys Targaryen", "age": 25}, {"name": "Arya Stark", "age": 16}]`)
out, err := cl.Create("demobox_6d9e326c183fde7b",val)
```

Read
```
out, err := cl.Read("demobox_6d9e326c183fde7b")
fmt.Printf("%s",out)

// Query records for a boxid
out, err := cl.Read("demobox_6d9e326c183fde7b?query_key=name&query_value=arya%20stark")
fmt.Printf("%s",out)

```

Update
```
val := []byte(`{"name": "Arya Stark", "age": 18}`)
out, err := cl.Update("demobox_6d9e326c183fde7b","5d776b75fd6d3d6cb1d45c53",val)
fmt.Printf("%s",out)
```

Delete
```
out, err := cl.Delete("demobox_6d9e326c183fde7b","5d776b75fd6d3d6cb1d45c53")
fmt.Printf("%s",out)
```

See full example usage at [examples/full](./examples/full/main.go)
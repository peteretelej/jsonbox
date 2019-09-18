/*
Package jsondoc wraps jsonbox API calls.

Create a new Client for use

	cl,err := jsonbox.NewClient("https://jsonbox.io/")

Create or add a record

	val := []byte(`{"name": "Jon Snow", "age": 25}`)
	out, err := cl.Create("demobox_6d9e326c183fde7b",val)

Read record(s)

	out, err := cl.Read("demobox_6d9e326c183fde7b")
	fmt.Printf("%s",out)

Query records

	out, err := cl.Read("demobox_6d9e326c183fde7b?query_key=name&query_value=arya%20stark")
	fmt.Printf("%s",out)

Update record

	val := []byte(`{"name": "Arya Stark", "age": 18}`)
	out, err := cl.Update("demobox_6d9e326c183fde7b","5d776b75fd6d3d6cb1d45c53",val)
	fmt.Printf("%s",out)

Delete record

	err := cl.Delete("demobox_6d9e326c183fde7b","5d776b75fd6d3d6cb1d45c53")

Delete All Records

	err := cl.Delete(BOX_ID)

List IDs for all records

	ids,err := cl.IDs("demobox_6d9e326c183fde7b")
	fmt.Printf("%s",ids)

*/
package jsonbox
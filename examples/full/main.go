package main

import (
	"fmt"
	"log"

	"github.com/peteretelej/jsonbox"
)

// TODO: change this value for your own tests
const demoBox = "demoboxgo_6d9e326c183fde7b"

func main() {

	cl, err := jsonbox.NewClient("https://jsonbox.io")
	if err != nil {
		log.Fatalf("Failed to create jsonbox client: %v", err)
	}

	// CREATE record in a BOX_ID
	record := []byte(`{"name": "John Doe", "age": 25}`)
	out, err := cl.Create(demoBox, record)
	if err != nil {
		log.Fatalf("failed to create new record for demobox %s: %v", demoBox, err)
	}
	fmt.Printf("CREATE - Created record for demobox %s: \n\t%s \n", demoBox, out)

	// get record ID from response
	id, err := jsonbox.GetRecordID(out)
	if err != nil {
		log.Fatalf("Failed to fetch recordId for newly created record: %v", err)
	}
	fmt.Printf("GetRecordID - record id for new record is %s\n", id)

	// READ records in BOX_ID
	out, err = cl.Read(demoBox)
	if err != nil {
		log.Fatalf("failed to READ record from demobox %s: %v", demoBox, err)
	}
	fmt.Printf("READ - record(s) in demobox %s: \n\t%s\n", demoBox, out)

	// UPDATE record in BOX_ID
	records := []byte(`{"name": "Jane Doe", "age": 30}`)
	out, err = cl.Update(demoBox, id, records)
	if err != nil {
		log.Fatalf("failed to UPDATE record %s/%s: %v", demoBox, id, err)
	}
	fmt.Printf("UPDATE - updated record %s/%s: \n\t%s\n", demoBox, id, out)

	// DELETE record in BOX_ID
	err = cl.Delete(demoBox, id)
	if err != nil {
		log.Fatalf("failed to DELETE record %s/%s: %v", demoBox, id, err)
	}
	fmt.Printf("DELETE - deleted single record %s/%s: \n\t%s\n", demoBox, id, out)

	// CREATE multiple records
	records = []byte(`[{"name": "John Doe", "age": 25},{"name": "Jane Do", "age": 30},{"name": "Jon Doe", "age": 30}]`)
	out, err = cl.Create(demoBox, records)
	if err != nil {
		log.Fatalf("failed to create multiple record for demobox %s: %v", demoBox, err)
	}
	fmt.Printf("CREATE - Created multiple records for demobox %s: \n\t%s \n", demoBox, out)

	// READ records in BOX_ID
	if out, err = cl.Read(demoBox + `?query_key=name&query_type=endswith&query_value=Doe`); err == nil {
		fmt.Printf("READ QUERY - Query records with name endswith 'Doe' in demobox %s: \n\t%s\n", demoBox, out)
	}

	// DELETING ALL RECORDS in a box
	if err = cl.DeleteAll(demoBox); err != nil {
		log.Fatalf("failed to DELETE ALL records for %s: %v", demoBox, err)
	}
	fmt.Printf("DELETE ALL - Deleted all records for  %s: \n\t%s \n", demoBox, out)

}

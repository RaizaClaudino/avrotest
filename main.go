package main

import (
	"bytes"
	"fmt"

	"github.com/hamba/avro"
)

type SomeEvent struct {
	Field float32 `avro:"field"`
}

type SomeNestedEvent struct {
	Test SomeEvent `avro:"teste"`
}

func main() {
	schemaSimples := `{"type":"record", "name":"test", "namespace": "org.hamba.avro", "fields":[{"name": "field", "type": "float"}]}`
	nestedSchema := `{"type":"record", "name":"testNested", "namespace": "org.hamba.avro", "fields":[{"name": "teste", "type": "org.hamba.avro.test"}]}`

	var newNestedEvent SomeNestedEvent
	event := SomeEvent{
		Field: 0.1,
	}
	nestedEvent := SomeNestedEvent{
		Test: event,
	}
	var avroSchema avro.Schema

	schemas := [2]string{schemaSimples, nestedSchema}

	for _, s := range schemas {
		avroSchema, _ = avro.Parse(s)
	}
	buf := bytes.NewBuffer([]byte{})
	encoder := avro.NewEncoderForSchema(avroSchema, buf)

	err := encoder.Encode(nestedEvent)
	fmt.Println(err, "ERR1")

	decoder := avro.NewDecoderForSchema(avroSchema, bytes.NewReader(buf.Bytes()))

	err = decoder.Decode(&newNestedEvent)
	fmt.Println(err, "ERR")
	// dataBytes, err := avro.Marshal(avroSchema, nestedEvent)
	// if err != nil {
	// 	fmt.Print(err)
	// }

	//avro.Unmarshal(avroSchema, dataBytes, newNestedEvent)

	// fmt.Println(nestedEvent)
	// fmt.Println(avroSchema.String())
	// fmt.Println(buf)
	fmt.Println(newNestedEvent)

}

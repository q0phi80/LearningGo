Types
You can create custom types using Go's simple types as a starting point
Syntax: type <name> <type>
An ID type based on a string will be type id string

Structs
Collecting different bit of data into a single data structure that you can design and control.
A custom type that you can name and specify the field properties and their types.
Syntax:
type <name> struct { 
<fieldName1> <type> 
<fieldName2> <type> 
 … 
<fieldNameN> <type> 
} 
To access a field on a struct: <structValue>.<fieldName>
To set a value: <structValue>.<fieldName> = <fieldName>
To read a value: value = <structValue>.<fieldName>
Once you've defined your custom struct type, you can use it to create a value.
Using embedding, you can add fields to a struct from other structs.
When you embed, the fields from the embedded struct get promoted. 
Once promoted, a field acts as if it's defined on the target struct.
To embed a struct, you add a struct type name to another struct without giving it a field name.
type <name> struct {
  <Type>
}

Interface
Describes a type with no functions.
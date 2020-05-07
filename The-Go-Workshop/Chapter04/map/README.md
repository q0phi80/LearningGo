Map
Arrays and slices are similar and can sometimes be interchangeable.
Map, howerver, is quite different and not interchangeable with array and slice.
With a map, you can set, get and delete a data quickly
You can access indivdidual elements of a map using [ and ]
Shouldn't use a map as an ordered list.
Syntax to define a map: map[<key_type>]<value_type>
Can also use make to create maps
You can't control or check the capacity of a map.
You must specify keys when initializing a map with data
- map[<key_type>]<value_type>{<key1>: <value>, ... <keyN>: <valueN>}
Once defined, you can set values like so: <map>[<key>] = <value>
It is good practice to avoid defining a map using var.
When looping over a map, you should use the range keyword.

#### Variable Declaration Notes

- Variables must first be initialized with the ':=' operator
- We can initialize a variable outside of a function, we just can't assign a value to it.
- Variable can be initialize and assigned with the ':=' operator
- To reassign the value of an operation you use the '=' operator

#### Function Declaration

- Every function that returns a value must indicate what type of value it is returning.(ex. string, bool, int, float64)
- Files in the same package can freely call functions defined in other files. Without having to import.

##### Slices and Arrays

- Array: fixed length list of things
- Slice: Array that can grow or shrink
- Every element in a slice must be of the same type

### Custom type declaration

- `type _name of your custome type_ underlying type`
- Underlying type are string, int, float, etc.

## Receiver Functions

- `func (d duck) quack() { // do something }`
- The name of this function is quack()
- To call func you would do `d.quack()`
- Don't ever reference a receiver value as 'this' or 'self'

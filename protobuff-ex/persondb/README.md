## PROTO BUFFER

How do we create new type from primitive type ?

Suppose we want to represents an Person Information type.

How do we create an new type?

Classical OOP - Provides concept of `CLASS`. Class helps you to create a new
 type.

```javascript
class Person {
    constructor(name, age, email) {
        this.name = name;
        this.age = age;
        this.email = email;
    }
}
```
GO - Named type with `type definition syntax`

```go
package p
type Person struct {
    name string
    age uint8
    email string
}
```

So What construct does the **proto buffer** provides to construct a new type.
`message`.

# db
Database framework that uses the [bee ORM](github.com/astaxie/beego/orm) at it's base.
ORM's like Entity Framework were used as the main source of inspiration. 

Built to enable easier creation and management of data.

## Features
* Nested Inserts & Updates
* Simple syntax to load related objects. To force eager loading. 
* Reduced usage of 'interface{}' as params
* Databases operate on a single Context (Ctx)
* Simplified reading when working with multiple databases. [dbName].Ctx.[tableName].[operation]
* Soft delete only. The `Deleted` column's value will be set to true.

### Note 
This project was built for use with PostgreSQL, and is tested as such.
Currently 'postgre' is hard coded as the database provider, but this can be changed as any provider supported by beego/orm can be used.

## Installation
* Ensure PostrgreSQL is installed. (version 9.6 & above)
* Create database (no need to create the tables in SQL as they will be generated by the ORM.)
* $ go get github.com/astaxie/beego/orm
* $ go get github.com/lib/pq (database driver)
* $ go get github.com/louisevanderlith/db

## Usage
* Create `dbcontext` for "sample" db.
*db/sample/dbcontext.go*
```go
package sample

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/db"
)

type Context struct {
    TableX *db.Set
}

var Ctx *Context

func NewDatabase(){
    dbSource := "host=localhost port=5432 user=postgres password=pass123 dbname=sample sslmode=disable"

    // register tables with beego/orm
    orm.RegisterModel(new(TableX))

    db.SyncDatabase(dbSource)

    // populate context object
    Ctx = &Context{
        TableX: db.NewSet(TableX{})
    }
}
```
* Create `model` for table
*db/tablex.go*
```go
package sample

import "github.com/louisevanderlith/db"

type TableX struct {
    db.Record       // Adds Id, CreateDate & Deleted columns. This also aids record management
    Name    string  `orm:"size(50)"`
    Age     int     `orm:"null"`
    Child   *TableX `orm:"rel(fk)"`
}

// Validate - Every table must have a validate function.
func (t TableX) Validate() (bool, error){
    return true, nil
}
```
* Use database and perform operations.
*main.go*
```go
package main

import (
    "./db/sample"
    _ "github.com/lib/pq"
)

func main(){
    // init database
    sample.NewDatabase()

    // create
    obj := &sample.TableX{
        Name: "TEST",
        Age: 66,
    }

    sample.Ctx.TableX.Create(obj)

    // read one and load related child (filter by Id)
    filterId := sample.Table{}
    filterId.Id = 1
    
    record, err := sample.Ctx.TableX.ReadOne(&filterId, "Child")

    // read all (filter by age = 66)
    var resultSet []*sample.TableX
    filter := &sample.TableX{
        Age: 66,
    }

    sample.Ctx.TableX.Read(filter, &resultSet)

    // see "set.go" for more operations and how to call them.
}
```
# go-yml2sql

Create SQL INSERT statement from YAML file.  
Suits for fixture data.

# Quick Usage

## command line
```sh
$ go get github.com/evalphobia/go-yml2sql
$ go-yml2sql -y <yaml file>

INSERT INTO `foo`(`col1`, `col2`) VALUES('1', 'a'), ('2', 'b'), ('3', 'c');
```

### Options

- `-y`: yaml file path.
- `-t`: table naming type. `file` is from file name, `dir` is from dir name of the file. (default=`dir`)
- `-plural`: flag for making table name pluralize or not. (default=`true`)


## go

```go
import(
    "github.com/evalphobia/go-yml2sql/yml2sql"
)

func main(
    path := "/path/to/bank/fixture.yml"

    yml2sql.SetPlural(true)
    yml2sql.SetNamingTypeDir(true)

    stmt := yml2sql.CreateStatementByFile(path)

    // execute insert
    _, err := db.Exec(stmt)

    if err != nil {
        panic(err.Error())
    }
)
```

# License

Apache License, Version 2.0

# go-yml2sql

Create SQL INSERT statement from YAML file.  
Suits for fixture data.

# Quick Usage

```sh
$ go get github.com/evalphobia/go-yml2sql
$ go-yml2sql -y <yaml file>

INSERT INTO `foo`(`col1`, `col2`) VALUES('1', 'a'), ('2', 'b'), ('3', 'c');
```

## Options

- `-y`: yaml file path.
- `-t`: table naming type. `file` is from file name, `dir` is from dir name of the file. (default=`dir`)
- `-plural`: flag for making table name pluralize or not. (default=`true`)

# License

Apache License, Version 2.0

# Gorm interface for Gorf

* Gorm support of Gorf.
* Inbuilt support for the sqlitedb

# Quickstart

Installation

```bash
go get github.com/go-gorf/gorf/backends/gormi
```

add gormi to the Gorf settings.go
```` go
gorf.Settings.DbBackends = gormi.NewSqliteBackend("db.sqlite")
````

# Releases

# backends/gormi/v0.0.1 (2023-06-18)

* **Initial release**: initial release


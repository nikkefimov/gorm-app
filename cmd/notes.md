Program flow.
Execute application -> Database Connection -> Define database model -> Migrate database -> CRUD operations -> Close Connection

1. The GO application started.
2. The GO applicastion establishes a connection to the MySQL database server by using GORM's gorm.Open method. (db.go file).
3. Application defines data models that map to tables in the MySQL database. (model.go file)
4. GORM automatically migrates the database schema based on the model definitions.
5. Execute any CRUD operations: "Create, Read, Update, Delete".
6. When all operations are complete, the GO application cloeses the database connection.
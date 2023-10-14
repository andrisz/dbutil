Universal connector/SQL query placeholder updater for MySQL, PostgreSQL and Oracle.

Build project with tags mysql, postgres or oracle. 

Connect to datbase with dbutil.Connect(\<host:[port]\>, \<user\>, \<password\>, \<dbname\>)

Use '?' for placeholders like MySQL and c call dbutil.SetPlaceholders() to translate it to native format.

Escape '?' in text with '??'. 

Example:

	db, err := dbutil.Connect("localhost", "user", "password", "dbname")

	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database: %s", err))
	}
	defer db.Close()

	rows, err := db.Query(dbutil.SetPlaceholders(`select count(*) from table_a where field_b > ?`), 1000)
	if err != nil {
		panic(fmt.Sprintf("Cannot execute query: %s", err))
	}


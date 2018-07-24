This is my version of https://gophercises.com/exercises/task.

For all CLI related operations I'm using COBRA library and for DB using bolt db.

I added tasks_test.go for db package and cli_test.go for cli. For db package code covrage goes above 90% but for cli it much lower.

Also for testing CLI I'm using exec.Command() calls.

## This is my version of https://gophercises.com/exercises/task.


For all CLI related operations I'm using COBRA library and for DB using bolt db.

I added tasks_test.go for db package and cli_test.go for cli. For db package code covrage goes above 90% but for cli it much lower.

Also for testing CLI I'm using exec.Command() calls. As an alternative to COBRA and bolt db I tried pure command line arguments and basic 

file I/O provided by golang. Creating CLI functionality on our own is relatively easy. For db opeartions like add task or list task are 

also easy but it gets difficult updating existing records using fseek like functions. So later I abandoned that approach and switched back 

to COBRA and bolt. 


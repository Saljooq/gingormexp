# Backend Database

## Prerequisites
- docker installed
- docker on path

## Instructions to run
```bash
❯ cd db
❯ sudo docker compose up
[+] Running 1/0
 ⠿ Container db-postgres-1  Created                                                                                       0.0s
Attaching to db-postgres-1
db-postgres-1  | 
db-postgres-1  | PostgreSQL Database directory appears to contain a database; Skipping initialization
db-postgres-1  | 
db-postgres-1  | 2023-03-14 05:10:54.574 UTC [1] LOG:  starting PostgreSQL 15.2 (Debian 15.2-1.pgdg110+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 10.2.1-6) 10.2.1 20210110, 64-bit
db-postgres-1  | 2023-03-14 05:10:54.574 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
db-postgres-1  | 2023-03-14 05:10:54.574 UTC [1] LOG:  listening on IPv6 address "::", port 5432
db-postgres-1  | 2023-03-14 05:10:54.578 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
db-postgres-1  | 2023-03-14 05:10:54.581 UTC [29] LOG:  database system was shut down at 2023-03-14 05:08:55 UTC
db-postgres-1  | 2023-03-14 05:10:54.586 UTC [1] LOG:  database system is ready to accept connections

```


## Instructions to stop
Press Ctrl + C to stop the process first
```bash
^CGracefully stopping... (press Ctrl+C again to force)
Aborting on container exit...
[+] Running 1/1
 ⠿ Container db-postgres-1  Stopped                                                                                       0.2s
canceled

❯ docker compose down
[+] Running 2/2
 ⠿ Container db-postgres-1  Removed                                                                                       0.0s
 ⠿ Network db_default       Removed                                                                                       0.2s


❯ sudo rm  -rfd postgres-data
❯ mkdir postgres-data
```


### Instructions to check the running db
```bash
❯ psql --host=0.0.0.0 --port=5432 --user=postgres
Password for user postgres: 
psql (15.2)
Type "help" for help.

postgres=# \dt
         List of relations
 Schema |  Name  | Type  |  Owner   
--------+--------+-------+----------
 public | mdpage | table | postgres
 public | person | table | postgres
(2 rows)

postgres=# select * from person;
 person_id |     name      | nickname 
-----------+---------------+----------
         1 | Saljooq Altaf | Sal
         2 | Bailey Coan   | Bay
(2 rows)


```


Plans for the future:

1. Research on arrays (avoid one:many relations if possible)
2. Research on how to hash the name-spaces for people
3. Work on a signup page
4. Better Errors/messages/animations (for loading for instance)

Long term
----
- Chat feature
- actual portfolio :)
- look into cool image generators for cool image for main portfolio page
- build (maybe a gprc) so that we can build the web-page when saving
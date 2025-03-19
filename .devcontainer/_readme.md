### Check if the container is running
```docker ps -f name=cushon-local-db```

### Connecting to the local db 
```mysql -h 127.0.0.1 -P 3306 -u root -plocaldev cushondb```

### Check migration history:
```migrate -path database/migrations -database "mysql://root:localdev@tcp(localhost:3306)/cushondb" version```

### Nuke the local database:
```bash
  # From the root of the project \
  make local-db-dn \
  docker volume rm $(docker volume ls -q -f name=mysql_data) \
  make run
```


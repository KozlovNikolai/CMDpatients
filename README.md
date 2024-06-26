docker run --name db-cmd-pg-13.3 -p 35433:5432 -e POSTGRES_USER=dbuser -e POSTGRES_PASSWORD=dbpass -e POSTGRES_DB=db_cmd -d postgres:13.3
docker start db-cmd-pg-13.3

/home/user/go/src/url-shortener/configs/local.yaml
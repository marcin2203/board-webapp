FROM postgres:12 as db
COPY init.sh /docker-entrypoint-initdb.d
COPY start.sql start.sql
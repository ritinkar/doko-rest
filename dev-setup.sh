docker run \
-e POSTGRES_PASSWORD=pass1234 \
-v ~/postgres-data:/var/lib/postgresql/data \
-p 5432:5432 \
-d postgres


psql \
-h localhost \
-U postgres \
-p 5432 \
-d doko
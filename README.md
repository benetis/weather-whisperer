# weather-whisperer

# Disclaimer

I am making this project for the purpose of getting familiar with Temporal.
However, feel free to use it as you wish.

## Meteo.lt API
- Uses https://api.meteo.lt API
- Note: meteo.lt API has 180 requests per minute limit and 20k requests per day limit
  - This project does not handle rate limiting

#### Run temporal server

Instructions from https://github.com/temporalio/docker-compose

```bash
git clone https://github.com/temporalio/docker-compose.git temporal-docker-compose
cd  temporal-docker-compose
docker-compose up
```

Create DB. Tables will be auto migrated
```bash
./scripts/initiliase_db.sh
```

How to start temporal:

```bash
1. Make sure temporal service is running by running `docker-compose up`
2. Run worker `go run cmd/worker.go`
3. Run client `go run cmd/client.go`
```

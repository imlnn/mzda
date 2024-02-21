docker-compose down && docker build . -t mzda && docker-compose --env-file .env.prod up -d

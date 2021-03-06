git checkout -f master && git pull origin master

docker image build -t jewerly-api:0.1 -f ./deploy/Dockerfile .

if [ "$(docker ps -q -f name=jewerly-api)" ]; then
    if [ ! "$(docker ps -aq -f status=exited -f name=jewerly-api)" ]; then
        docker rm $(docker stop jewerly-api)
    fi
fi

docker run --env-file ../.env.prod -v /root/jewerly-shop/api/logs/prod:/logs -d --restart always --publish 8000:8000 --name jewerly-api --link=jewerly-db:db jewerly-api:0.1
git checkout -f develop && git pull origin develop

docker image build -t jewerly-api:0.1 -f ./deploy/Dockerfile .

if [ "$(docker ps -q -f name=jewerly-api-stage)" ]; then
    if [ ! "$(docker ps -aq -f status=exited -f name=jewerly-api-stage)" ]; then
        docker rm $(docker stop jewerly-api-stage)
    fi
fi

docker run --env-file ../.env.stage -v /root/jewerly-shop/api/logs/stage:/logs -d --restart always --publish 8001:8000 --name jewerly-api-stage --link=jewerly-db:db jewerly-api:0.1
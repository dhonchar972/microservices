cp config/Dockerfile $1/Dockerfile
cd $1

./gradlew clean build

docker build . -t dhonchar/$1:$2
docker push dhonchar/$1:$2

rm Dockerfile

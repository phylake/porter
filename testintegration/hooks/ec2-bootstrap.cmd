#!/bin/bash

echo "DEV_MODE $DEV_MODE" 1>&2

if [[ $PORTER_ENVIRONMENT = 'CustomVPC' ]]; then

cat <<'EOF'
SECRET=$(porter_get_secrets)
[[ $SECRET = 'hi' ]] || exit 1
EOF

fi

if [[ $TCPDUMP = '1' ]]; then

cat <<'EOF'
cat <<'DOCKER_POST_RUN' > /usr/bin/porter_docker_post_run
#!/bin/bash

HOST_PORT=$1
if [[ -z "$HOST_PORT" ]]; then
    exit 2
fi
IMAGE_NAME=tcpdump-$HOST_PORT

cat <<'DOCKERFILE' | docker build -t $IMAGE_NAME -
FROM ubuntu:16.04

RUN apt-get update
RUN apt-get install -y tcpdump

CMD tcpdump -w /host/tcpdump.$PORT -n -s 100 -i lo port $PORT
DOCKERFILE

docker run -d \
-v /home/ec2-user:/host \
--net=host \
--privileged \
-e PORT=$HOST_PORT \
$IMAGE_NAME
done

DOCKER_POST_RUN
chmod 755 /usr/bin/porter_docker_post_run
EOF

fi

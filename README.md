# grpc-latency

1) `git clone https://github.com/eladshoval/grpc-latency.git`
2) `cd grpc-latency`
3) `docker build -t pokemon .`
4) Run server -
   1) `docker run -it -e WORK_MODE='server' -e LISTEN_PORT='6666' -p 6666:6666 pokemon:latest`
5) Run client -
   1) `docker run -it -e WORK_MODE='client' -e TARGET_HOST='<YOUR-HOSTNAME>' -e TARGET_PORT='6666' pokemon:latest`
   2) Example:
      1) `docker run -it -e WORK_MODE='client' -e TARGET_HOST='ec2-1-1-1-1.ap-southeast-1.compute.amazonaws.com' -e TARGET_PORT='6666' pokemon:latest`

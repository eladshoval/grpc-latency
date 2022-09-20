# grpc-latency

1) `git clone https://github.com/eladshoval/grpc-latency.git`
2) `cd grpc-latency`
3) `docker build -t pokemon .`
4) Run the **server** -
   * `docker run -it -e WORK_MODE='server' -e LISTEN_PORT='6666' -p 6666:6666 pokemon:latest`
5) Run the **client** -
   * `docker run -it -e WORK_MODE='client' -e TARGET_HOST='<SERVER-HOSTNAME>' -e TARGET_PORT='6666' pokemon:latest` <br/><br/>    
     * Example: `docker run -it -e WORK_MODE='client' -e TARGET_HOST='ec2-1-1-1-1.ap-southeast-1.compute.amazonaws.com' -e TARGET_PORT='6666' pokemon:latest`

## 1) Question : How to check logs from my docker container and how to filter only last 200 lines from container logs ?
Ans) command : **docker container logs <container_name>  --tail 200**

## 2) Question : what will happen to container logs if you restart the container, is it lost?
Ans) No, we will not lost the logs if you start,stop, restart. but if you delete the container or terminate the container the date will lost, due to emphmeral or stateless nature (we can using centralized logging)

## 3) Question : You encounter a docker image with size 2.7GB. Will this be a cause of concern ? If yes how would you tackle this ?
Ans) we can use multistage docker file , dextro images to reduce the size of docker images
---
. Bigger docker image would result in longer build time

. Docker image download error or API rate limit issue

· Application will also become bulkier

Solution :

Smaller Image Base (Alpine Images)

. Multi-Stage Builds Feature when building docker image

. Remove package binaries after installing and don't install packages that isn't required
---

## 4) What is the difernces between docker images and docker layers?
Ans) docker images is image of container with all requirement to run the container, where as docker layers are building blocks or steps to build docker images

## 5) what is differences between CMD and ENTRYPOINT  in docker?
it is the best practices to use both CMD and entrypoint. and we can write the scripts in entrypoints and passes to CMD **[Entrypoints scripts]**
---

Only CMD

FROM ubuntu:latest
/CMD ["echo", "This is CMD"]

Only
ENTRYPOINT

FROM ubuntu:latest
ENTRYPOINT ["echo", "This is EntryPoint"]
-7

ENTRYPOINT
and CMD

FROM ubuntu:latest
ENTRYPOINT ["echo"]
CMD ["This is EntryPoint"]

Both run during docker container runtime
Either one of them are suggested to be used as best practice
Using both of them in the same dockerfile is also possible

---
## 6)  Question : Given 2 different docker file below . Which one do you think will have lesser size when built ?

# syntax=docker/dockerfile:1
FROM node:12-alpine
RUN apk add -- no-cache python2 g++ make
WORKDIR /appCOPY . .RUN yarn install --
productionCMD ["node", "src/index.js"]

# syntax=docker/dockerfile:1
FROM node
RUN apk add -- no-cache python2 g++ make
WORKDIR /appCOPY . .RUN yarn install --
productionCMD ["node", "src/index.js"]

Ans) first one due to **Alpine image**


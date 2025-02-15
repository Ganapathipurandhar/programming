## 1) Question : Can we have a Jenkins agent which is a docker container and runour test inside this docker container ?
ANS) Yes, we can implement docker container and runour test inside this docker container.

  This is used for isolation of the pipeline steps

. Here the code is expected to run inside a docker container which can use a custom test image if required

  This is the best way to perform testing without having every jenkins agent needing to have packages installed
~~~
Jenkinsfile (Declarative Pipeline)
pipeline {
agent {
docker { image 'node:16.13.1-alpine'}
}
stages {
stage('Test') {
steps {
sh 'node -- version'
}
}
}
~~~
Source : https://www.jenkins.io/doc/pipeline/tour/agents/

to use any kind docker container of basically we add it as a part of agent , we can call any kind of image from dockerhub or custom built images, whatever the stages memtion in the pipeline that we can run inside the container, it is best way for multi-node jenkins setup, completely test and isolate it.   

## 2) Question : Have you use Jenkins in multi node setup ? If yes, could you explain how to add a new slave/followed to master ?
jenkins --> manage jenkins--> nodes--> add node with details.

## 3) Question : Currently you have a jenkins job , but we don't have any kind of notifications on job success or failures. We use slack for our internal communications.How would you enable a notification service in this case ?

Ans)
    · We can leverage jenkins plugin here.
    · Jenkins - Slack plugin can be installed (We would need a slack webhook URL ) Pipeline job
      slackSend color: "failure", message: "Pipeline has failed. Kindly take a look"

## 4) Question : What is a multi branch pipeline ?
Ans) 
The Multibranch Pipeline project type enables you to
implement different Jenkinsfiles for different branches
of the same project.

master> git repo ->Jenkins --> production

staging> git repo ->Jenkins --> staging

dev/testing --> test/dev branches --> git repo -->Jenkins --> production

for this [we need to have a proper branching strategy]

## 5) Question : What is a Blue-Green Deployment ?
Ans) blue-green deployment is meture way of deployment when it comes to large user traffic so it cricial to make sure, any new deployment don't causes the deployment, it has a two identical prod enveronment, one env hold's pervious version and other env holds current version,(uses route53, f5 etc) route the traffic to new version and it make easy to revert back easyly




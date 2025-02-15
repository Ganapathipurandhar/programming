## 1) Question : Do you use helm ? Which version do you use currently and do you see benefit of using helm ?
Ans) 
 Yes, we use helm

Helm version is 3.18 for all new application

. Helm is basically packaging all your kubernetes files into a single chart format which will help in
Deployment
Rollback
Upgrade
Prod. Staging. Dev setup

## 2) Question : why helm using K8S?
Helm helps you manage Kubernetes applications - Helm Charts help you define, install, and upgrade even the most complex Kubernetes application.

It becomes complex when you have a lot of yaml to handle (deployment, back-up,service,serviceaccount etc)

Charts are easy to create, version, share, and publish 

## 3) Question : Which helm repository do you use today to store/access helm charts ?
It's a common place where helm charts can be stored, which can be accessed via our deployment.

This also maintains chart version also which shall allow us to do easy rollback, upgrade etc.

Cloudsmith
JFrog Artifactory
AWS S3
Google cloud storage

Open Source : [text](https://artifacthub.io/packages/)

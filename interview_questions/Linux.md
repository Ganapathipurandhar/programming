## 1) your EC2 instances running out of disk space. what action you will take to mitigate the issue?
Ans)  EC2 disk space means we are speaking of EBS volume

We will first check if its a root volume or other volume
/root - OS
/application - for our application

If its root volume then i will first try to check logs and clear some space if not the
instance might shut down also

If its application volume , then I will use EBS feature to take snapshot and increase disk
space for the EC2 instance

--> **first, need to check how many volumes are attached to the EC2(root,application volume) then effected volume(root or application), due to the log, error this logs generated,If its application volume , then I will use EBS feature to take snapshot and increase disk
space for the EC2 instance**

## 2) what is bastion host or gateway Server ? what is the role of it?
Ans) bastion host is using the manage access to an internal or private subnet/network to external network . it is also called as jump box or jump server or gateway server (if you three EC2 intances, you can monitor all the three instances, you restrict the access according the groups), also you can completely the cut access by destroying the bastion, if need you can create by automation

## 3) Mutiple EC2 instances in ASG is getting terminated and this is causing downtime on the application. EC2 pricing , quota all looks good. how would you start debugging this issue?
Ans) if EC2 instances unhealthly then only it will terminated -->(it may causes due to 100% CPU ultility , disk space , memory full)
1) for 100% ultitity  --> check **top** command which threads is using more CPU --> reach developer to inform the causes 
2) disk space(EBS may 100% full) --> check the volumes (/var/logs , /root , /tmp , /etc etc) this may causes the issues
3) no memory left --> free -mt (swap memory) if it is full it may causes the problem

## 4) Question : Create a linux script that will push certain logs to S3 automatically.Explain only high level design which is enough. Share what all steps you would do to achieve this above script.
Bonus : Also, The script should run at particular time.
Ans) For high level design(bigger picture) --> using **IAM roles or access cred**  to get access to s3 for ec2 instances the write a script to get the logs and push the logs using **AWSCLI**  cammands. To run at particular time we can use cron jobs

## 5) Question : why is logging is improtant ? what is centralised logging and what tools helps to achieve centralized logging?
Ans) logging(collection logs) is important to track the behavious application or vm. but in real time production our application is running on scalable manner(instances will automatical created and deleted so we lost the data but we can use centralised logging like S3, cloudwatch, eleaticsreach, spluck, datadog etc to store the logs in outside the instances)
 

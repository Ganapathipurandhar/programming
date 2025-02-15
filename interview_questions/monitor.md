## 1) explain differnet way that promethues get the metrics?
ans) prometheus collect the metrics in two way --> pull base approch , push base approch

1) pull base : prometheus scrape the application to get the logs [prometheua---(pull by scrape)--->applications(java, python etc)]
2) push base : created a custom-metrics and push the metrics to prometheus by using gatewayapi [prometheus<--- gataway<---custom-metrics]

both approch are validate but according to application we use either of them

## 2) Question : Your development team needs your help to monitor the API endpoint. Which HTTP response would you monitor and when will you trigger the alert ?

. First we have to understand that all application will have API endpoints

. Example /status or /home etc

. Any endpoint you hit (get call) then you will get different response

· These response are called HTTP response code
  Informational responses (100-199)
  Successful responses (200-299)
  Redirection messages (300-399)
  Client error responses (400-499)
  Server error responses (500-599)

If HTTP response code is 400 or 500 range then we can trigger the alert

/home-->HTTP response code-->Monitoringcode/service

--> i will write a script using python or shell scripting and get the response by (get call) in the endpoint /status etc, if the response is 400+ or 500's then it tregger the alerts

## 3) what are the some way to can set the alerting?

Usually alerting will be a common process flow that is defined by DevOps engineers

It's not feasible to have multiplying alerting methodology in the same org

Alerting notification can be done via Email, mobile phone call, Slack massage etc

Open source ways :

· Prometheus --> Alertmanager

· Cloudwatch --> SNS

. Nagios --> Alerting

## 4) Question: Please explain how you create your dashboard for monitoring services in your current project ?
Ans) As we know monitoring and dashboarding are the two main oberverbility parameters what we following the devops engineering , dashboard help us to present the metrics to developes, testers and ourself's,because i am using the k8s projects i am mainly using grafana for dashboarding.[for AWS --> we use cloudwatch , other ELK Stack, datadog etc]


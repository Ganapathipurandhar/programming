## 1) Question : Explain which python module will you use to make a simple API testing code.The code should just check if a API endpoint is working or not

Ans) You need to worry if you don't know python .Explain the logic here in any other programing language you are aware of

Here, I will focus on python and the module is **requests module** in python
Part of sample code :
~~~
status_code = requests.get(website).status_code
~~~

Status code (200)- API endpoint is good Any other value then we have an issue

## 2) Question : explain in simple team what is the package in python?
Ans) package is collection modules, each modules hold the differnet function which we help us meet requried feature.

package-->modules-->functions

## 3) Question : A dev team wishes to go with a NoSQL DB, But what is this NoSQL DB and do we have something in AWS that can get can used as a NoSQL DB
Ans)  basically NOSQL is ideal for which there is no predefine structure defination before the table creation [means all the incoming data changes will be accepted it will be best for **ecommernce application**]

SQL:  it a structure database we need define schema before creating tables[it don't accept other column otherthen defined --> **banking,etc**]

## 4) Question : We have a on-premises Linux server which needs some monitoring enabled. Being a DevOps engineer only you have access to this server. How would you setup a basic monitoring script on this on-premises server ?
Ans) 

· Scripting is a important tool/knowledge that DevOps engineer must have

. It is not that you will use it on day-to-day bases but its required for quick
testing/automation and development.

. Based on the interview question asked here : Only quick scripting can help because
integrating a on-premises server to our central monitoring tools can take too long

. With this question they are testing your quick scripting and little bit of coding
knowledge also.

. Don't worry if you don't know python, Take this opportunity to learn with me and I will
explain my best so you understand this code.

~~~

import psutil

# CPU times
print(psutil.cpu_times())

# Avg CPU load
print(psutil.getloadavg())

# Memory
print(psutil.virtual_memory())

# Swap memory
print(psutil.swap_memory())

# Disk usage
print(psutil.disk_usage('/'))

# Disk IO utilization
print(psutil.disk_io_counters(perdisk=False))

# Temp
print(psutil.sensors_temperatures())
~~~

## 5) Question : Are you supporting any application ? If yes what is the application server you are using ?
DevOps engineer in many cases support many applications in a organization.

If you are supporting a project
server you are using or implemented in the team.

Java :
Tomcat
JBoss V
Weblogic -License

NodeJS:
IIS

you are supporting such teams then it's good to know what kind of application

Python :
Django


## 6) Question : Which build tool are you aware of and what are the build tools used in your project ?
. Build tools are critical when it comes to packing the application code and shipping it

We can ship just raw code files to servers/clients

Based on different programming language we have different build tools
Java
- Maven
- Gradle

Nodejs
- npm

Python
- Pybuilder

This question can come as a follow up after jenkins questions


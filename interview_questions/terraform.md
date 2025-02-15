## 1) Question : When using terraform to create a RDS, How to save the DB username and password securely ?
Ans) 
DB username, DB Password are considered secrets [store in vault like hashicorp, KMS, AZURE vault]

. Best practice is we should never commit these information in github repo

· But do we save our secrets then ?
Because without this we can't run terraform code

We can use terraform - Vault integration

. Mention vault provider in the provider section

. Then reference the secrets in vault via data block in terraform

## 2) Question : Your security team has insisted that certain packages (security) related has to be present on all the instances for a defense project. How would you come up with a solution for such a design and ensure these packages will be present on all EC2 instances that shall be used.
Ans) we can use like **hashicorp Packer, Custom AMI images** 

---
. Hashicorp Packer is the solutionk

. Packer is an open source tool that enables you to create identical machine images for multiple platforms
from a single source template.

. A common use case is creating "golden images" that teams across an organization can use in cloud
infrastructure.

JSON/yaml template-->Packer-->OSImage
---


## 3) Question : What is remote-exec in terraform and when will you use it ?

The remote-exec provisioner invokes a script on a remote resource after it is created.

This can be used to run a configuration management tool, bootstrap into a cluster, etc.

To invoke a local process, see the local-exec provisioner instead.

The remote-exec provisioner requires a connection and supports both ssh and winrm.

## 4) Question : How to validate the variable during terraform plan time ? Example you take a variable called AMI . We need to make sure its of a proper format.
Ans we can use the validation block to validate the variable is proper are not, we can use custom error-message for the uses
~~~
variable "ami" {
= string
description = "EC2 instance AMI ID"
= "ami-90394321875931"

validation {
condition= length(var.ami) > 4 && substr(var.ami, 0, 4) == "ami-"
 error_message = "Enter a proper AMI ID."
}
}
~~~
## 5) Question : What does terraform state lock really mean ? Do we have a practical use of it ?
. If supported by your backend, Terraform will lock your state for all operations that could write state.

. This prevents others from acquiring the lock and potentially corrupting your state.

State locking happens automatically on all operations that could write state.

. You won't see any message that it is happening. If state locking fails, Terraform will not continue.

. You can disable state locking for most commands with the -lock flag but it is not recommended.

. If acquiring the lock is taking longer than expected, Terraform will output a status message.

· If Terraform doesn't output a message, state locking is still occurring if your backend supports it.

## 6) Question : When running terraform apply, The process gets hung. Now if you cancel this it mightresult in corrupt terraform state How would you debug this issue ?
- This is not a common issue
- But this can occur if we are not careful with terraform apply, especially if you are doing it too often.
- If the this process get's hung, We can use **terraform state pull** command to pull the state to local system.



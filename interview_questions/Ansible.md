## 1) Question : Let's say you have 2 servers with different ports and username then how can ansible handle different machines for applying and running the same runbook.

We can achieve this by using below hostfile format

[webservers]
10.4.20.90 ansible_port=4000 ansible_user=roger
39.12.3.23 ansible_port=8001 ansible_user=luffy

when we run the different ports and username, then in the hostfile we can specific the port, username for specfic ip address 

## 2) Question : In a ansible playbook, We have a section which could fail on certain nodes. But we don't want the playbook to stop or exist because of this. Is it possible to ignore this if this part of playbook fails ?
Yes we can achieve it using **ingore_errors** sub arguments, if you have error in the middle of the playbook execute it fails it all will contoious to run the playbook
~~~
- name: Do not count this as a failure
  ansible.builtin.command: /bin/false
  **ignore_errors: yes**
~~~
## 3) Question : How can we handle secrets in ansible ? Let's say we have playbook which needs to login to premises server with login name and password.

Ans) · ansible-vault create vault.yml

     · echo 'data to be encrypted' > encrypt_file.txt

     . ansible-vault encrypt encrypt_file.txt

Ansible[have screts] --> create vaults[ansible vault encrpt the screts]--> when need it decrypt the data

in ansible itself we have an command like ansible-vault to store screts in encrypted and also external integration like hashicrop-vault,KMS etc can used to scalability


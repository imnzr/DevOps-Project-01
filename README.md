# Deploy Your Code On a Docker Container Using Jenkins on AWS

![diagram-export-4-22-2025-7_09_24-PM](https://github.com/user-attachments/assets/44669afe-7ed3-46cd-b890-be4ec4b88dde)

Description 

## Agenda
1. Setup jenkins
2. Setup and configure git
3. Integration github with jenkins
4. Setup docker host
5. Integrate docker with jenkins
6. Automate the build and deploy process using jenkins
7. test the deployment

## Prequisites
1. Amazon Web Services (AWS)
2. SonarQube
3. Github account with the source code

## STEP 1 : Setup Jenkins Server on AWS EC2
1. Setup the linux instance
2. Install Java
3. Install Jenkins
4. Configure and Start Jenkins
5. Access web UI on PORT 8080

Login ke halaman Amazon Management Console, buka EC2 Dashboard, dan klik Launch Instance:
![Screenshot from 2025-04-22 19-20-46](https://github.com/user-attachments/assets/8702d2a9-4229-422e-accf-04ddc4cb8414)

Setelah meng-klik launch instance, berikan nama untuk server jenkins:
![Screenshot from 2025-04-22 19-31-18](https://github.com/user-attachments/assets/ad91c197-bcc9-4c93-a925-c1d1753be85a)

Selanjutnya memilih sistem operasi yang akan digunakan untuk server jenkins, disini saya memakai OS Ubuntu 22.04:
![Screenshot from 2025-04-22 19-39-41](https://github.com/user-attachments/assets/128d55a2-7fb7-41be-9441-144bade883d9)

Pilih instance type yang akan digunakan pada server jenkins:
![image](https://github.com/user-attachments/assets/07b7ff07-83b2-4e44-b21e-b77134ef4e09)

Selanjutnya membuat key pairs untuk mengakses server jenkins di terminal, buat key-pair baru:
![Screenshot from 2025-04-22 19-43-46](https://github.com/user-attachments/assets/cb999d63-c730-453f-bf8e-17c5a06e30ad)

Berikan nama untuk key-pair lalu pilih key pair 'RSA' dan file format nya '.pem'
![Screenshot from 2025-04-22 19-44-51](https://github.com/user-attachments/assets/c4418687-1df5-4c41-b23b-c5aeb78358b2)

Selanjutnya membuat storage, setelah itu klik launch instance:
![image](https://github.com/user-attachments/assets/dd64f31c-a3f0-43e7-97b7-1e8821fe7da3)








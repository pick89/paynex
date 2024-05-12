# paynex

Paynex is way to learn Golang in this project there are API, Database.

Stack: 
Golang 
Docker 
Postgres

#### Step1:

Set up your .env following this template 

#.env file

#Database Environment Variables

POSTGRES_USER={*POSTGRES_USER*}
POSTGRES_PASSWORD={*POSTGRE_PASSWORD*}
POSTGRES_DB={*POSTGRES_DB*}
POSTGRES_HOST={*POSTGRES_HOST*}
POSTGRES_PORT={*POSTGRES_PORT*}

#pgAdmin Environment Variables

PGADMIN_DEFAULT_EMAIL={*PGADMIN_DEFAULT_EMAIL*}
PGADMIN_DEFAULT_PASSWORD={*PGADMIN_DEFAULT_PASSWORD*}

#Server Configuration

SERVER_PORT={*server_port*}


#### Step2: Docker Compose 

docker-compose up 

- check the status of docker using command: **docker-compose ps**
- Ensure the database is up by going to **http://localhost:5050/**
The log should be indicate everythig is **UP**

Inspect the docker contain the Postgres for retrive the IP adrress for connect the data ase 
# STAND UP API

## Running Locally

- To run this project you will need to have Docker installed and running on your machine
- Use the .env_sample to create a new .env file
- Open your terminal and run the command `make docker-up` to setup the postgres database and run the required migrations in docker container
- In a different terminal, run the command `make start-server` to start the server 
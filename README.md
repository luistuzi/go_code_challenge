# go_code_challenge
Repository created in order to store the challenge itself and it's changes

## TO STAR THE APP THROUGH DOCKER

All configurations are being setted on Dockerfile, if any need to be changed, don't forget to change and save the file

Endpois are being documented on swagger

execute the comands:

# To perform the build
docker compose build --no-cache

# To run the whole image containing the app and the mysql database for testing
docker compose up -d

# To check if both images are running
docker compose ps

# To test the endpoints

http://localhost:8080/api/v1/device/swagger/index.html 

# To stop the docker
docker compose down -v




## TO START THE APP LOCALLY

# To ensure all packages are in

go mod tidy

# To run the app

go run /cmd/api/main.go

# To test the endpoints

http://localhost:8080/api/v1/device/swagger/index.html 



## PROBLEMS

Using direct and not encrypted passwords on properties file instead of using secrets
Not using completely profiles for properties by env
No implementation of JWT auth for endpoints
No implementation of any use o cert validation
No implementation of a atomic control of database transactions
Perhaps a use of a framework?

## Improvements

Implementation of a spring profile like properties management 
Implementation of a proper JWT auth
Implementation of a @transactional like control of db transactions



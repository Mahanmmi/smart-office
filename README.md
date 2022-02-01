# fuzzy-lamp
An IoT system for managing offices and rooms door, light, etc.

## Used Languages & Technologies & Protocols
- Golang
- CPP
- MQTT
- HTTP
- Docker
- PostgreSQL
- NodeMCU ESP8266 Board
- Postman

## Architecture
![image](https://user-images.githubusercontent.com/44570354/152060473-85385d9c-d07a-4883-92b8-3e6948932b35.png)

## Run the project

- `docker-compose up`
- To run main server: `cd main-server && go run .`
- To run local-server: `cd local-server && go run .`

To run postman tests:
- Import postman collection from `main-server/fuzzy-lamp.postman_collection.json`
- Run `Office/register` -> `admin/register` -> `admin/login` -> `admin/user register` requests in this exact order to import 
required data (an office, an admin and a user. **SERVER WILL NOT WORK AS INTENDED IF YOU DON'T DO THIS**

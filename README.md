# fuzzy-lamp

To run the project:

- `docker-compose up`
- To run main server: `cd main-server && go run .`
- To run local-server: `cd local-server && go run .`

To run postman tests:
- Import postman collection from `main-server/fuzzy-lamp.postman_collection.json`
- Run `Office/register` -> `admin/register` -> `admin/login` -> `admin/user register` requests in this exact order to import 
required data (an office, an admin and a user. **SERVER WILL NOT WORK AS INTENDED IF YOU DON'T DO THIS**
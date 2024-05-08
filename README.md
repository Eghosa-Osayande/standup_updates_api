# STAND UP API

## Running Locally

- To run this project you will need to have Docker installed and running on your machine
- Use the .env_sample to create a new `.env` file in the root directory
- Open your terminal and run the command `make docker-up` to setup the postgres database and run the required migrations in docker container
- During the database migration, a super-admin employee record is created in the employees table
- The super-admin password MUST be set as an environment variable through the .env file.
- In a different terminal, run the command `make start-server` to start the server

## API Documentation

This repository includes a postman collection for testing purposes.

### Employee
#### Login Employee
Returns the access token

**path**: `/v1/employees/login`
**method**: `POST`
**authorization**: `none`
##### Body
```json
{
    "name":"admin", // required
    "password":"password" // required
}
```
##### Sample Response
```json
{
    "name":"admin", // required
    "password":"password" // required
}
```

#### Create Employee
Only employees with `admin` roles can create new employees

**path**: `/v1/employees/create`
**method**: `POST`
**authorization**: `Bearer Token`
##### Body
```json
{
    "name": "John Doe", // required
    "password": "12345", // required
    "role": "employee",// required, "employee" or "admin"
}
```
##### Sample Response
```json
{
    "id": "3bc5665a-0d4a-11ef-9fb0-dfda52db21c3",
    "name": "John Doe",
    "password": "$2a$10$dUtRnwpZzHGHq8.Z3MPWF.l0gAS6.I7nej1C2ABoxq1xHltrD8.6i",
    "created_at": "2024-05-08T14:49:55.326464+01:00",
    "role": "employee"
}
```

#### Fetch All Employees 
Only employees with `admin` roles can see all employees.

**path**: `/v1/employees/all`
**method**: `GET`
**authorization**: `Bearer Token`
##### Body
```json
{
    "page":1, // required
    "per_page":20 // required
}
```
##### Sample Response
```json
{
    "status": true,
    "message": "success",
    "data": [
        {
            "id": "a62f2cca-0d44-11ef-af5f-8bc53e8a9c66",
            "name": "admin",
            "password": "$2a$10$d0VQ7cPV.BEoRoyH7wGHhuqsY6WoG6bFizVbckfH8HrMI8GRIS9rW",
            "created_at": "2024-05-08T15:09:56.727435+01:00",
            "role": "admin"
        },
        {
            "id": "3bc5665a-0d4a-11ef-9fb0-dfda52db21c3",
            "name": "John Doe",
            "password": "$2a$10$dUtRnwpZzHGHq8.Z3MPWF.l0gAS6.I7nej1C2ABoxq1xHltrD8.6i",
            "created_at": "2024-05-08T14:49:55.326464+01:00",
            "role": "employee"
        }
    ],
    "cursor": {
        "next_page": 2,
        "per_page": 20
    }
}
```

### Sprint
#### Create Sprint 
Only employees with `admin` roles can create sprints.

**path**: `/v1/sprints/create`
**method**: `POST`
**authorization**: `Bearer Token`
##### Body
```json
{
    "name": "Sprint 1", // required
    "standup_start_time": "13:00", // required
    "standup_end_time": "13:09" // required
    // standup_start_time and standup_end_time are given
    // as the hour and minute in UTC (GMT)
    // i.e; HH:MM
    
}
```
##### Sample Response
```json
{
    "id": "8f2469b6-0e05-11ef-bc83-070503e341bb",
    "name": "Sprint 1",
    "created_at": "2024-05-09T13:10:51.092417+01:00",
    "standup_start_time": "13:00 UTC",
    "standup_end_time": "13:09 UTC"
}
```

#### Fetch All Sprints 

**path**: `/v1/sprints/all`
**method**: `GET`
**authorization**: `Bearer Token`
##### Body
```json
{
    "page":1,
    "per_page":20
}
```
##### Sample Response
```json
{
    "status": true,
    "message": "success",
    "data": [
        {
            "id": "c634ede2-0e03-11ef-aeef-97c3964a02ee",
            "name": "Sprint 1",
            "created_at": "2024-05-09T12:58:04.47782+01:00",
            "standup_start_time": "12:40 UTC",
            "standup_end_time": "12:50 UTC"
        },
    ],
    "cursor": {
        "next_page": 2,
        "per_page": 20
    }
}
```

### Updates
#### Submit Update 

**path**: `/v1/updates/create`
**method**: `POST`
**authorization**: `Bearer Token`
##### Body
```json
{
    "employee_id":"a62f2cca-0d44-11ef-af5f-8bc53e8a9c66", // required
    "sprint_id":"8f2469b6-0e05-11ef-bc83-070503e341bb", // required
    "yesterday":"Nothing", // required
    "today":"Nothing", // required
    "blocked_by":[], // required, array of string ids
    "breakaway":"Nothing", // required
    "tasks":[] // required, array of string ids
}
```
##### Sample Response
```json
{
    "status": true,
    "message": "success",
    "data": {
        "id": "9e4f809c-0e05-11ef-bc83-67531e5b6061",
        "employee_id": "a62f2cca-0d44-11ef-af5f-8bc53e8a9c66",
        "sprint_id": "8f2469b6-0e05-11ef-bc83-070503e341bb",
        "created_at": "2024-05-09T13:11:16.540322+01:00",
        "yesterday": "Nothing",
        "today": "Nothing",
        "blocked_by": [],
        "breakaway": "Nothing",
        "check_in_time": "13:11 UTC",
        "status": "after standup"
    }
}
```

#### Fetch All Updates 

**path**: `/v1/updates/all`
**method**: `GET`
**authorization**: `Bearer Token`
##### Body
```json
{
    "page":1, // required
    "per_page":20, // required
    "employee_id":"a62f2cca-0d44-11ef-af5f-8bc53e8a9c66", // optional
    "sprint_id":"75e9c7d8-0e01-11ef-a4bd-e794c3f16063", // optional
    "start_date":"2024-05-08", // optional
    "end_date":"2024-05-09" // optional
    // start_date and end_date represent the year, month and day
    // ie; year-month-day
}
```
##### Sample Response
```json
{
    "status": true,
    "message": "success",
    "data": [
        {
            "id": "7e0b5abc-0e01-11ef-a4bd-8348477b0c82",
            "employee_id": "a62f2cca-0d44-11ef-af5f-8bc53e8a9c66",
            "employee_name": "admin",
            "sprint_id": "75e9c7d8-0e01-11ef-a4bd-e794c3f16063",
            "created_at": "2024-05-09T12:41:44.419646+01:00",
            "yesterday": "Nothing",
            "today": "Nothing",
            "blocked_by": [],
            "breakaway": "Nothing",
            "status": "within standup"
        }
    ],
    "cursor": {
        "next_page": 2,
        "per_page": 20
    }
}
```
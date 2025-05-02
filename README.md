# Healthcare Appointment System

## ✅ Requirements Checklist

### Core Requirements
- [x] RESTful API with JSON input/output
- [x] Patient management (CRUD operations)
- [x] Appointment management (CRUD operations)
- [x] Relationship between patients and appointments
- [x] GORM integration with SQLite
- [x] Auto-migrate schema on application start
- [x] Proper error handling with HTTP status codes
- [x] Modular project structure
- [x] Startup behavior without manual setup

### Bonus Features
- [x] Input validation for fields
- [x] Request logging middleware
- [x] Pagination for patient and appointment lists
- [x] Docker containerization
- [x] Postman collection (available [here](#postman-collection))

A backend API for managing patients and their appointments, built with Go, Gin, and GORM.

## Postman Collection
You can find the Postman collection for testing the API here:  
[![Run in Postman](https://run.pstmn.io/button.svg)](https://tournament-9964.postman.co/workspace/My-Workspace~f8ca06f3-afdb-4100-8dff-f35443c006c0/collection/36796614-62c833df-b4f5-4b2c-a7ab-121754a18bb5?action=share&creator=36796614)

## Setup Instructions

1. **Prerequisites**:
   - Go 1.24 installed
   - SQLite (in-memory database used)
   - Docker (optional)

2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the Application**:
   ```bash
   go run cmd/api/main.go
   ```

4. **Run with Docker**:
   Build the Docker image:
   ```bash
   docker build -t healthcare-api .
   ```
   
   Run the container:
   ```bash
   docker run -p 8080:8080 --name healthcare-api healthcare-api
   ```

5. **Environment**:
   The application will create a SQLite database file `healthcare.db` in the project root.

## API Endpoints

### Patients

- `GET /patients` - Get all patients (paginated)
- `POST /patients` - Create a new patient
- `GET /patients/:id` - Get a specific patient by ID
- `GET /patients/:id/appointments` - Get all appointments for a specific patient

### Appointments

- `GET /appointments` - Get all appointments (paginated)
- `POST /appointments` - Create a new appointment

## Validation

### Patient Validation
- `email`: Must be a valid email format
- `gender`: Must be one of ["male", "female", "other"]
- `dateOfBirth`: Must be a valid date in the past

### Appointment Validation
- `date`: Must be a valid future date
- `patientID`: Must reference an existing patient

## Example Requests and Responses

### Create a Patient
**Request**:
```bash
POST /patients
Content-Type: application/json

{
    "firstName": "John",
    "lastName": "Doe",
    "gender": "male",
    "dateOfBirth": "1990-01-01T00:00:00Z",
    "phoneNumber": "1234567890",
    "email": "john.doe@example.com"
}
```

**Response**:
```json
{
    "id": 1,
    "createdAt": "2023-10-01T12:00:00Z",
    "updatedAt": "2023-10-01T12:00:00Z",
    "firstName": "John",
    "lastName": "Doe",
    "gender": "male",
    "dateOfBirth": "1990-01-01T00:00:00Z",
    "phoneNumber": "1234567890",
    "email": "john.doe@example.com"
}
```

### Create an Appointment
**Request**:
```bash
POST /appointments
Content-Type: application/json

{
    "patient_id": 1,
    "reason": "Annual checkup",
    "date": "2024-01-15T10:00:00Z"
}
```

**Response**:
```json
{
    "id": 1,
    "createdAt": "2023-10-01T12:00:00Z",
    "updatedAt": "2023-10-01T12:00:00Z",
    "reason": "Annual checkup",
    "date": "2024-01-15T10:00:00Z",
    "patientID": 1
}
```

### Get Patient by ID
**Request**:
```bash
GET /patients/1
```

**Response**:
```json
{
    "id": 1,
    "createdAt": "2023-10-01T12:00:00Z",
    "updatedAt": "2023-10-01T12:00:00Z",
    "firstName": "John",
    "lastName": "Doe",
    "gender": "male",
    "dateOfBirth": "1990-01-01T00:00:00Z",
    "phoneNumber": "1234567890",
    "email": "john.doe@example.com"
}
```

### Get Appointments for a Patient
**Request**:
```bash
GET /patients/1/appointments
```

**Response**:
```json
[
    {
        "id": 1,
        "createdAt": "2023-10-01T12:00:00Z",
        "updatedAt": "2023-10-01T12:00:00Z",
        "reason": "Annual checkup",
        "date": "2024-01-15T10:00:00Z",
        "patientID": 1
    }
]
```

## Pagination
Both `GET /patients` and `GET /appointments` support pagination via query parameters:
- `limit`: Number of items per page (default: 10)
- `offset`: Starting position (default: 0)

Example: `GET /patients?limit=5&offset=10`
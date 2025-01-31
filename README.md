# Project Search API - GoSolve

This project demonstrates a proposed search gateway, structured with Handler, Controller, Service and Repository to allows search for a value, retriveing it index.

## Quick Start

Run the backend and frontend using Docker Compose:

```bash
docker-compose up --build -d
```

## Makefile Commands

| Command	    | Description
|---------------|-------------------------------------------------------|
| make up	    |  🚀 Start backend and frontend (Docker Compose)       |
| make down	    | 🛑 Stop and remove all containers                     |
| make restart	| 🔄 Restart backend and frontend                       |
| make logs	    | 📜 Show logs for all services                         |
| make test	    | 🧪 Run backend tests                                  |
| make coverage	| 📊 Run tests & generate coverage report               |
|---------------|-------------------------------------------------------|


### Key Features
- **Backend**: `http://localhost:8080`
- **KFrontend**: `http://localhost:5173`
- **Swagger API Documentation:**: [Swagger UI](http://localhost:8080/swagger/index.html)
---

## Detailed Flow

1. User submits a number via the frontend (React)
2. Frontend sends a request to http://localhost:8080/search/{value}
3. Backend processes the request:
    - Checks if the number exists in the data, by performing a binary search
    - If found, returns the index and value.
    - If not found, returns the closest value within a 10% margin.
4. Frontend displays results (or error messages).

## Frontend Notes

  - I used [Vite](https://vite.dev/guide/) for faster and development experience (vite is new for me)
  - I chose Material-UI for its simplicity and ease of use (I have used for other personal proejcts).


---

## Design/Business Decisions and Considerations

- **ZAP Logger:** I created two Zap logger functions:
    - utils/zap.go:NewLogger – A simple wrapper for the Zap logger.
    - utils/zap.go:NewCustomLogger – Allows customization of the logger level (info, warn, error) based on the .env - 
    - Future Improvement:
        - Ideally, the logging configuration should be moved from .env to a YAML or JSON file for better flexibility and maintainability.

- **Server Port Configuration:**
    - The server port is currently loaded from environment variables (SERVER_PORT).
    - This allows quick changes via .env without modifying the code.
    - Future Improvement:
        - SAME as defined in ZAP Logger section

- **CFF and FX from Uber Golang**:
  - We can explore CFF (a concurrency toolkit for Go) and Uber Golang's FX for dependency injection and modularization, e.g., main.go. 

- **Mutext**:
  - For the file data acess I deced to use RWMutex in data repository

- **Project Architecture**:
  - We decided to follow a layered architecture with separation of concerns, ensuring that each layer has a single responsibility, making the system testable, maintainable, and scalable.

## API Documentation

The API follows RESTful principles. Below is an example of how it works:

### Rxample Request

```bash
GET /search/30
```

### Example Response

```bash
{
  "index": 5,
  "value": 30,
  "message": "Closest match found"
}
```

### Error Response

```bsh
{
  "message": "No closest match found",
  "value": -1
}
```

### Coverage

The main backend files are covered with unit tests.

```bash
gosolve/backend/internal/api/handlers.go:24:			    NewHandler		      100.0%
gosolve/backend/internal/api/handlers.go:40:			    SearchHandler		  84.2%
gosolve/backend/internal/api/handlers.go:75:			    handleAPIError		  66.7%
gosolve/backend/internal/api/router.go:13:			        SetupRouter		      0.0%
gosolve/backend/internal/controller/controller.go:19:       New		        	  0.0%
gosolve/backend/internal/controller/controller.go:23:	    newController		  0.0%
gosolve/backend/internal/controller/controller.go:34:	    SearchValue	    	  0.0%
gosolve/backend/internal/errors/errors.go:13:			    Error		    	  100.0%
gosolve/backend/internal/errors/errors.go:18:			    New		              100.0%
gosolve/backend/internal/mocks/mock_repository.go:19:		GetData			      100.0%
gosolve/backend/internal/mocks/mock_repository.go:25:		FindClosestMatchIndex 100.0%
gosolve/backend/internal/mocks/mock_repository.go:31:		LoadData		      0.0%
gosolve/backend/internal/mocks/mock_repository.go:42:		SearchValue		      100.0%
gosolve/backend/internal/repository/data_repository.go:29:	NewDataRepository	  0.0%
gosolve/backend/internal/repository/data_repository.go:50:	GetData			      100.0%
gosolve/backend/internal/repository/data_repository.go:58:	loadData	          90.0%
gosolve/backend/internal/repository/data_repository.go:91:	FindClosestMatchIndex 81.8%
gosolve/backend/internal/services/search_service.go:19:		NewSearchService	  100.0%
gosolve/backend/internal/services/search_service.go:32:		FindClosest		      100.0%
gosolve/backend/internal/utils/utils.go:11:			        Abs		              100.0%
gosolve/backend/internal/utils/utils.go:19:			        EncodeResponse		  100.0%
gosolve/backend/internal/utils/zap.go:8:			        NewLogger		      0.0%
```

## Error Codes

The following table describes the error codes used in the project:

    | Error Name                 | Error Message                                 | HTTP Status Code         |
    |----------------------------|-----------------------------------------------|--------------------------|
    | `ErrInvalidRequest`        | Invalid request format                        | 400 Bad Request          |
    | `ErrInvalidRequestValue`   | Invalid request value                         | 404 Not Found            |
    | `ErrNoClosestMatchFound`   | No closest match found                        | 404 Not Found            |
    | `ErrDataNotLoaded`         | Data is empty                                 | 500 Internal Server Error|
    | `ErrInvalidIndexFromRepo`  | Invalid index from repository                 | 500 Internal Server Error|
    | `ErrInputFileNotSet`       | Input file not set                            | 500 Internal Server Error|
    | `ErrFailedToInitiateRepo`  | Failed to initiate repository                 | 500 Internal Server Error|
    | `ErrProcessingData`        | Error processing da                           | 401 Unauthorized         |
    | `ErrInvalidValueInFile`    | Invalid value in file                         | 500 Internal Server Error|
    | `ErrInternalServer`        | Internal server error                         | 500 Internal Server Error|

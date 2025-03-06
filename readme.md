# Shrinkify - The Mighty URL Shrinker 🌀

## Description
Ever felt your URLs are too long and just plain ugly? Worry no more! **Shrinkify** is here to save the day. This Go-powered URL shortener takes your monstrous links and turns them into sleek, tiny, shareable beauties. Plus, it runs smoothly inside a Docker container, so you can deploy it effortlessly like a boss. 🚀

## Prerequisites
Make sure you have the following installed:
- Go (latest version recommended)
- Docker & Docker Compose
- PostgreSQL (or use Docker for the database)

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/ripu2/URL-shorter.git
   cd URL-shorter
   ```
2. Copy the `.env.example` file to `.env` and configure your environment variables.

3. Build and run the application using Docker:
   ```sh
   docker-compose up --build
   ```

## Running the Application Locally (Without Docker)
1. Install dependencies:
   ```sh
   go mod tidy
   ```
2. Start the application:
   ```sh
   go run main.go
   ```

## Docker Setup
This project includes a `Dockerfile` and `docker-compose.yml` for easy setup.
- To build the Docker image manually:
  ```sh
  docker build -t shrinkify .
  ```
- To start the container:
  ```sh
  docker run --env-file .env -p 8080:8080 shrinkify
  ```

## Database Setup
If using Docker, PostgreSQL will be automatically set up via `docker-compose.yml`.
Otherwise, make sure PostgreSQL is running locally and configured properly.

## API Documentation
You can find the Postman collection for testing the API [here](postman_collection.json). Import it into Postman to explore available endpoints and test Shrinkify like a pro! 🚀

## Example `.env` File
Create a `.env` file in the root directory with the following format:

```
PORT=8080
POSTGRES_USER=your_db_user
POSTGRES_PASSWORD=your_db_password
POSTGRES_DB=your_db_name
POSTGRES_PORT=5432
POSTGRES_HOST=db
```

Modify these values as per your environment setup.

## Contributing
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes and commit (`git commit -m "Added new feature"`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a Pull Request.

## License
This project is licensed under the [MIT License](LICENSE).

## Contact
For any queries or issues, feel free to reach out to the maintainers at [your-email@example.com].


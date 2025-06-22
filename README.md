![nexproject_logo](https://github.com/user-attachments/assets/6f01ed33-778b-4b45-a359-e28fc7d29866)

# nexproject API
nexproject is a platform dedicated to advancing Sustainable Development Goal 8 (SDG 8) by supporting inclusive economic growth and promoting decent work for all. Designed for Small and Medium Enterprises (SMEs)

## About SDG 8
**SDG 8** is a global initiative that aims to promote sustained, inclusive, and sustainable economic growth, full and productive employment, and decent work for all. By facilitating job creation and connecting SMEs with diverse talent, nexproject contributes to a fairer job market where economic opportunities are accessible to everyone.

## Getting Started
### Prerequisites
- **Go**: Ensure the latest version of Go is installed.
- **Gin Gonic**: A web framework for building the API in Golang.
- **GORM**: An ORM library for Golang used for database management.
- **Database**: PostgreSQL for data storage.

### Installation
1. **Clone the repository**:
    ```bash
    git clone https://github.com/gnlehc/nexproject-api
    cd nexproject-api
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

3. **Configure environment variables**:
   - Create a `.env` file in the project root directory.
   - Add your nexproject API Key, database connection details, and any other necessary configurations.
   ```dotenv
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name

### Running the Application
1. **Start the server**:
    ```bash
    go run main.go
    ```

2. **Access the API**:
   The API can be accessed locally at `http://localhost:8080`.
   

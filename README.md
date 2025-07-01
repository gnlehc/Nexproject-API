<img src="https://drive.google.com/uc?export=view&id=1OBNMfwxtol7bCXhroFOKVem7VMu2TuE9" alt="NexProject Preview" width="700"/>

# Nexproject
Nexproject adalah platform pertama di Indonesia yang menghubungkan mahasiswa dan UMKM - memberikan ruang bagi mahasiswa untuk membangun portofolio, sekaligus membantu UMKM memenuhi kebutuhan kreatif dan digital mereka.
---

## ⚙️ Tech Stack

- **Frontend**: Vue.js + Vite
- **Backend**: Go (Gin framework)
- **Database**: PostgreSQL
- **Deployment**: Docker
---
  
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
   

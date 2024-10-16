# Full Stack Todo Application

This project is a full-stack Todo application built using Angular for the frontend and Golang for the backend. It provides a simple and intuitive interface for users to manage their tasks. The app supports CRUD operations, task completion, and task deletion.

## Features

- **Add Tasks**: Users can add new tasks to the list.
- **Edit Tasks**: Users can edit existing tasks directly from the list.
- **Complete Tasks**: Toggle the completion status of tasks.
- **Delete Tasks**: Remove tasks from the list with a confirmation prompt.
- **Responsive UI**: The interface is designed to work well on both desktop and mobile devices.

## Tech Stack

- **Frontend**: Angular 17, Angular Material, TypeScript, HTML, CSS
- **Backend**: Golang, RESTful API
- **Database**: SQLite (for simplicity in this project, but can be replaced by another DBMS)
- **Deployment**: Deployed on Heroku (or any preferred cloud platform)

## Screenshots

![image](https://github.com/user-attachments/assets/8d4e326e-557a-4a4e-b9fd-87b45f9fc3f4)

## Setup and Installation

### Prerequisites

- Node.js and npm installed on your machine.
- Go installed on your machine.
- SQLite installed (or other DBMS if you plan to modify the backend for another database).

### Frontend Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/your-repo-name.git
    cd your-repo-name
    ```

2. Navigate to the frontend folder:
    ```bash
    cd frontend
    ```

3. Install Angular dependencies:
    ```bash
    npm install
    ```

4. Run the Angular application:
    ```bash
    ng serve
    ```

   The application will be running on `http://localhost:4200/`.

### Backend Setup

1. Navigate to the backend folder:
    ```bash
    cd backend
    ```

2. Run the Golang server:
    ```bash
    go run main.go
    ```

   The backend API will be running on `http://localhost:8080/`.

### Database Setup

1. SQLite is used for this application. If needed, initialize the database file:
    ```bash
    sqlite3 todo.db
    ```

2. Run the SQL migrations provided in the `migrations/` folder.

### Environment Variables

Create a `.env` file in the root of the backend with the following content:

```env
PORT=8080
DB_PATH=./todo.db

# Media Storage Server

A simple media storage server for storing various types of media. Can be used for images and videos.

## Table of Contents

-   [Introduction](#introduction)
-   [Features](#features)
-   [Getting Started](#getting-started)
    -   [Prerequisites](#prerequisites)
    -   [Installation](#installation)
    -   [Database](#database)
-   [Usage](#usage)
    -   [Creating Folders]()
    -   [Getting Folders]()
    -   [Editing Folders]()
    -   [Deleting Folders]()
-   [Contributing](#contributing)
-   [License](#license)

## Introduction

This project provides a basic template for building a Go media storage server. It includes CRUD functionality for both folders and images. The images are stored within a folder which creates a more structured storage system. URLs are generated for use displaying the media stored.

## Features

- Image & video storage.
- Folders for sorting images.
- Meta data storage in a database.
- Generate a URL to display media stored.

## Getting Started

Follow these steps to get the project up and running on your local machine.

### Prerequisites

-   Go (1.16 or higher)
-   PostgreSQL database
-   Git (optional)

### Installation

1. Clone the repository (if you haven't already):

```bash
    git clone https://github.com/Azpect3120/MediaStorageServer.git
```

2. Set up your PostgreSQL database and configure the connection details in the `.env` file:

```.env
  # This url can found in the dashboard of most PSQL hosts or can be constructed using the required pieces
  db_url=your-connection-url-here
```

3. Install dependencies:

```bash
  go mod tidy
```

4. Build and run the server:

```bash
  go build -o ./bin/server cmd/mediaStorageServer/main.go
  ./bin/server
```

Your authentication server should now be running on `http://localhost:3000`.

### Database

Once the server is up and running you will need to connect to a PostgreSQL database. If you would like the code to work out of the box, you may copy the database schema provided below.

```sql
  CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

  CREATE TABLE IF NOT EXISTS folders (
      id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      name VARCHAR(32) UNIQUE,
      createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

  CREATE TABLE IF NOT EXISTS images (
      id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      folderId UUID REFERENCES folders(id) ON DELETE CASCADE,
      name TEXT,
      size BIGINT,
      type TEXT,
      uploadedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );
```

## Usage







## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the project.
2. Create a new branch for your feature or bug fix.
3. Make your changes.
4. Test your changes thoroughly.
5. Create a pull request.

## License

The project is licensed under username the **MIT License**

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
    -   [Creating Folders](#creating-folders)
    -   [Getting Folders](#getting-folders)
    -   [Editing Folders]()
    -   [Deleting Folders]()
    -   [Uploading Images]()
    -   [Displaying Images]()
    -   [Deleting Images]()
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

### <a id="creating-folders"></a>Creating Folders

Folders are created to allow users to store images in a organized format. All you need to do is send a `POST` request to the `/folders` endpoint.

NOTE: The name provided must be unique. Response will return an error if the name is invalid.

```json
  {
    "name": "folder-name-here"
  }
```

Ex. Response

```json
  {
    "folder": {
      "ID": "generated-id-here",
      "Name": "folder-name-here",
      "CreatedAt": "timestamp"
    },
    "status": 201
  }
```


### <a id="getting-folders"></a>Getting Folders

Folders store very little data on their own, but by getting a folder using its `ID` you can view a list of each image stored in the folder. To do this, send a `GET` request to the `/folders/<folder_id>` endpoint. 

NOTE: If a folder is not found, PSQL will return an error with a `400` error code.


```bash
  http://localhost:3000/folders/<folder_id>
```

Ex. Response

```json
  {
    "folder": {
      "ID": "generated-id-here",
      "Name": "folder-name-here",
      "CreatedAt": "timestamp"
    },
    "images": [
      {
        "ID": "image-id",
        "FolderId": "generated-id-here",
        "Name": "image-name.jpg",
        "Size": 940184,
        "Format": "image/jpeg",
        "UploadedAt": "timestamp",
        "Path": "uploads/generated-id-here/image-id.jpg"
      }
    ],
    "status": 200
  }
```






## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the project.
2. Create a new branch for your feature or bug fix.
3. Make your changes.
4. Test your changes thoroughly.
5. Create a pull request.

## License

The project is licensed under username the **MIT License**

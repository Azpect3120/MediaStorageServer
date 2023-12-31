# Media Storage Server

A simple media storage server for storing images and videos of various formats.

## Table of Contents

-   [Introduction](#introduction)
-   [Version](#version)
-   [Features](#features)
-   [Getting Started](#getting-started)
    -   [Prerequisites](#prerequisites)
    -   [Installation](#installation)
    -   [Database](#database)
-   [Usage](#usage)
    - [Folders](#creating-folders)
        -   [Creating Folders](#creating-folders)
        -   [Getting Folders](#getting-folders)
        -   [Editing Folders](#editing-folders)
        -   [Deleting Folders](#deleting-folders)
    - [Media](#uploading-media)
        -   [Uploading Media](#uploading-media)
        -   [Getting Media](#getting-media)
        -   [Displaying Media](#displaying-media)
        -   [Deleting Media](#deleting-media)
    -   [Generating Reports](#generating-reports)
-   [Contributing](#contributing)
-   [License](#license)

## <a id="introduction"></a> Introduction

This project provides a basic template for building a Go media storage server. It includes CRUD functionality for both folders and images. The images are stored within a folder which creates a more structured storage system. URLs are generated for use in displaying the media stored. The `/images` endpoints can be used to store both image and video media.

## <a id="version"></a> Version 1.0.1

All endpoints will be preceded by the version. Excluding the `/uploads` endpoint which will remain constant throughout the versions.

Acceptable version endpoints are...
- /v1/

View [change log](https://github.com/Azpect3120/MediaStorageServer/blob/v1.0.0/changelog.md) here

## <a id="features"></a> Features

- Image & video storage.
- Folders for sorting images.
- Metadata storage in a database.
- Generate a URL to display the media stored.
- Generate reports and have them sent to your email.
- Cache for quick media and folder retrieval.
- Duplicate images are stored only once on the disk.
- Supported media formats: [JPEG, PNG, WEBP, BMP, TIFF, ICO, MP4]

## <a id="getting-started"></a> Getting Started

Follow these steps to get the project up and running on your local machine.

### <a id="prerequisites"></a> Prerequisites

-   Go (1.16 or higher)
-   PostgreSQL database
-   Git (optional)

### <a id="installation"></a> Installation

1. Clone the repository (if you haven't already):

```bash
    git clone https://github.com/Azpect3120/MediaStorageServer.git
```

2. Set up your PostgreSQL database and configure the connection details in the `.env` file:

```.env
  # This URL can be found in the dashboard of most PSQL hosts,
  # or it can be constructed using the required pieces
  db_url=your-connection-url-here

  # For more information visit this link:
  # https://www.hostinger.com/tutorials/how-to-use-free-google-smtp-server
  smtp_email=your-email-here
  smtp_password=your-email-app-password-here
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

Your media storage server should now be running on `http://localhost:3000`.

The server's port can be changed in the `cmd/mediaStorageServer/main.go` file:

```go
  err := server.Run("NEW PORT HERE")
```

### <a id="database"></a> Database

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

## <a id="usage"></a> Usage

### <a id="creating-folders"></a>Creating Folders

To create a folder to begin storing images, send a `POST` request to the `/folders` endpoint.

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

Folder meta data and images can be viewed by sending a `GET` request to the `/folders/<folder_id>` endpoint.

NOTE: If a folder is not found, a PSQL error will be returned with a `400` error code.

```bash
  GET http://localhost:3000/<version>/folders/<folder_id>
```

Ex. Response

```json
  {
    "folder": {
      "ID": "generated-id-here",
      "Name": "folder-name-here",
      "CreatedAt": "timestamp"
    },
    "status": 200
  }
```

### <a id="editing-folders"></a> Editing Folders

Folder name can be updated by sending a `PUT` request to the `/folders/<folder_id>` endpoint.

NOTE: Folder names must be valid, if a name is provided that is not valid, it will be cleaned as best as possible. If it cannot be cleaned then an error will be thrown.

```json
  {
    "name": "newFolderName"
  }
```

Ex. Response

```json
  {
    "folder": {
      "ID": "folderIDHere",
      "Name": "newFoldername",
      "CreatedAt": "timestamp"
    },
    "status": 200
  }
```

### <a id="deleting-folders"></a> Deleting Folders

A folder can be deleted by sending a `DELETE` request to the `/folders/<folder_id>` endpoint. Nothing is returned from this request unless an error is encountered.

NOTE: Images will be deleted when its parent folder is deleted.

```bash
  DELETE http://localhost:3000/<version>/folders/<folder_id>
```

### <a id="uploading-media"></a> Uploading Media

Images can be uploaded through forms on a web application. Each "flavor" or framework will do it slightly differently the only information that must remain constant is that the image must be uploaded from a form with the name `media_upload`. Send a `POST` request to the `/images/<folder_id>` endpoint with the form data. The endpoint response will be stored in the `data` variables in the following examples.

HTML Example

```html
    <input type="file" id="fileInput" />
    <button onclick="uploadFile()"> Upload </button>

    <script>
      function uploadFile() {
        const fileInput = document.getElementById("fileInput");
        const file = fileInput.files[0];

        const formData = new FormData();
        formData.append("media_upload", file);

        fetch("http://localhost:3000/<version>/images/<folder_id>", {
            method: "POST",
            body: formData,
        })
        .then(res => res.json())
        .then(data => console.log(data));
      }
    </script>
```

React Example

```javascript
  import React, { useState } from "react";

  function FileUploadComponent () {
    const [file, setFile] = useState(null);

    const handleFileChange = (event) => {
      setFile(event.target.files[0]);
    };
    
    uploadFile = () => {
      if (!file) {
        console.error("No file selected.");
        return;
      }

      const formData = new FormData();
      formData.append("media_upload", file);

      fetch("http://localhost:3000/<version>/images/<folder_id>", {
          method: "POST",
          body: formData,
      })
      .then(res => res.json())
      .then(data => console.log(data));
  };

    return (
      <>
        <input type="file" onChange={handleFileChange} />
        <button onClick={uploadFile}> Upload </button>
      </>
    )
  }

  export default FileUploadComponent;
```

Ex. Response

```json
  {
    "image": {
      "ID": "image-id",
      "FolderId": "generated-id-here",
      "Name": "image-name.jpg",
      "Size": 940184,
      "Format": "image/jpeg",
      "UploadedAt": "timestamp",
      "Path": "uploads/generated-id-here/image-id.jpg"
    }
  }
```

### <a id="getting-media"></a> Getting Media

An images metadata can be viewed by sending a `GET` request to the `/images/<image_id>` endpoint. The `path` property that is returned can be used to display the image.

```bash
  GET http://localhost:3000/<version>/images/<image_id>
```

Ex. Response

```json
{
  "image": {
    "ID": "image-id",
    "FolderId": "generated-id-here",
    "Name": "image-name.jpg",
    "Size": 940184,
    "Format": "image/jpeg",
    "UploadedAt": "timestamp",
    "Path": "uploads/generated-id-here/image-id.jpg"
  },
  "status": 200
}
```

Another way you can get retrieve media is through their parent folders. Sending a GET request to the `/folders/:id/images` endpoint will return a list of images in the folder.

```bash
  GET http://localhost:3000/<version>/folders/<folder_id>/images?limit=<num_per_page>&page=<page_num>
```

Ex. Response

```json
{
    "status": 200,
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
    ]
}

```

### <a id="displaying-media"></a> Displaying Media

In web applications, the images can be displayed using their path, which can be found in the image metadata. Simply use the path as the `src` attribute in any web application.

```html
  <img src="http://localhost:3000/uploads/<folder_id>/<image_id>" alt="Media Storage Image">
```

### <a id="deleting-media"></a> Deleting Media

An image can be deleted by sending a `DELETE` request to the `/images/<image_id>` endpoint. Nothing is returned from this request unless an error is encountered.

```bash
  DELETE http://localhost:3000/<version>/images/<image_id>
```

### <a id="generating-reports"></a> Generating Reports

A folder report can be generated and sent to your email by sending a `GET` request to the `/reports/:id/:email` endpoint.

```bash
  GET http://localhost:3000/<version>/reports/<folder_id>/<your_email>
```

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the project.
2. Create a new branch for your feature or bug fix.
3. Make your changes.
4. Test your changes thoroughly.
5. Create a pull request.

## License

The project is licensed under the **MIT License**

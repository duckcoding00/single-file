# Go Image File I/O Handler

This repository is a mini-project focused on handling file I/O operations for image files using the Go standard library. The project provides basic functionality to upload images, retrieve a list of stored images, and serve images from local storage. The goal is to improve skills in managing file I/O in Go.

## Features
- Upload image files.
- Get a list of stored images.
- Retrieve and serve an image.
- Uses only Go's standard library and gorilla/mux.
- Stores images locally in the repository directory.

## Requirements
- Go (latest stable version)

## Installation & Setup
```sh
# Clone the repository
git clone https://github.com/duckcoding00/single-file.git
cd single-file

# Run the application
go run main.go
```

## API Endpoints
### 1. Upload Image
- **Endpoint:** `POST /upload`
- **Description:** Uploads an image file and saves it locally.
- **Request:** Multipart form-data with an image file.
- **Response:**
  ```json
  {
    "message": "CREATED",
    "data": "filepath"
  }
  ```

### 2. Get List of Images
- **Endpoint:** `GET /images`
- **Description:** Retrieves a list of stored image filenames.
- **Response:**
  ```json
  {
    "message": "CREATED",
    "data": [
      "image1",
      "image2",
    ]
  }
  ```

### 3. Get Image
- **Endpoint:** `GET /images/{filename}`
- **Description:** Serves the requested image.
- **Response:** Returns the image file as binary data.

## Project Purpose
This project is created to practice and improve skills in handling file I/O operations using Go's standard library. It focuses on working with local storage without relying on external packages.

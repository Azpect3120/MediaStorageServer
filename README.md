# Outline


Files will be stored locally on the server using its own hardware.

`Clouds` will be created to store the images more securely. Users can only perform actions on the images if they have the `cloud_id` provided when creating a cloud

Metadata will be stored in a PSQL database
- File name
- User information
- Time stamp
- Image Dimensions
- File size
- ID


Endpoints:
- `POST /cloud` : Create a new cloud
- `GET /cloud` : Get cloud metadata
- `POST /upload/{cloud_id}` : To upload images
- `GET /images/{cloud_id}/{id}` : Retrieve a specific image
- `PUT /images/{cloud_id}/{id}` : Update image metadata
- `DELETE /images/{cloud_id}/{id}` : Delete an image and its metadata

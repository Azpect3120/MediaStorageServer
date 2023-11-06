# Outline


Files will be stored locally on the server using its own hardware.

`Folders` will be created to store the images more securely. Users can only perform actions on the images if they have the `folder_id` provided when creating a folder

Metadata (stored in psql db)
- File name
- User information
- Time stamp
- Image Dimensions
- File size
- ID


Folder Routes:
- `POST /folders` : Create a new folder.
- `GET /folders/{folder_id}` : Get folder metadata.
- `PUT /folders/{folder_id}` : Update a folders metadata.
- `DELETE /folders/{folder_id}` : Delete a folder and it's images.


Image Routes:
- `POST /images/{folder_id}` : Upload an image to a specific folder.
- `GET /images/{image_id}` : Retrieve a specific image and its meta data.
- `PUT /images/{image_id}` : Update image metadata.
- `DELETE /images/{image_id}` : Delete an image and its metadata.


Schema:

```sql
CREATE TABLE IF NOT EXISTS folders (
    id UUID PRIMARY KEY DEFAULT generate_uuid_v4(),
    name VARCHAR(32) UNIQUE,
);
```

```sql
CREATE TABLE IF NOT EXISTS images (
    id UUID PRIMARY KEY DEFAULT generate_uuid_v4(),
    cloudID UUID REFERENCES folders(id),

);
```

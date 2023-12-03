import React, { useState } from "react";

function FileUploadComponent() {
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

        fetch("<host_url>/v1/images/<folder_id>", {
            method: "POST",
            body: formData,
        })
            .then((res) => res.json())
            .then((data) => console.log(data));
    };

    return (
        <>
            <input type="file" onChange={handleFileChange} />
            <button onClick={uploadFile}> Upload </button>
        </>
    );
}

export default FileUploadComponent;

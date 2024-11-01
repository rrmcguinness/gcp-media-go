// Copyright 2024 Google, LLC
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     https://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Button, Container, Typography } from "@mui/material";
import { useState } from "react";
import { FileUploader } from "react-drag-drop-files";
import "./FileUpload.css"
import axios from "axios";

const SupportedFileTypes = ["mp4"]

const FileUpload = () => {
  const [files, setFiles] = useState<File[]>([]);
  
  const handleChange = (file: any) => {
    setFiles(file);
  }
  const onDrop = (file: any) => {
    console.log('drop', file);
  };
  const onSelect = (file: any) => {
    console.log('test', file);
  };

  const onTypeError = (err = 1) => console.log(err);
  const onSizeError = (err = 1) => console.log(err);
  console.log(files);

  const submitData = () => {
    if (!files) {
      return
    }
    const form = new FormData()
    for (const file of files) {
      form.append("files", file)
    }
    axios.post("http://localhost:8080/api/v1/uploads", form, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    }).then(r => {
      setFiles([])
      console.log(r)
    }).catch(e => {
      console.log(e)
    })
  }
  
  
  return(
    <Container>
      <Typography variant={'h5'}>Upload a video file</Typography>
      <FileUploader
        classes="upload-files"
        fileOrFiles={files}
        onTypeError={onTypeError}
        handleChange={handleChange}
        name="file"
        types={SupportedFileTypes}
        onSizeError={onSizeError}
        onDrop={onDrop}
        onSelect={onSelect}
        label="Upload file here"
        dropMessageStyle={{backgroundColor: 'red'}}
        multiple
      />
      <Button variant="contained" onClick={submitData}>Submit</Button>
    </Container>
  )
}

export default FileUpload;
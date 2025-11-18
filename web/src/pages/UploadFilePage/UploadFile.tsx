import UploadFileComp from '../../components/UploadFileComponent/UploadFileComp.jsx';
import { useEffect, useState } from 'react'
import { useParams, useNavigate } from 'react-router-dom';
import Cookies from 'js-cookie';
import convertBytesToString from '../../utils/convertBytesToString.js';
import axiosInstance from '../../utils/axiosInstance.js';
import type { AxiosError } from 'axios';

const UploadFile = () => {
  const params = useParams();
  const navigate = useNavigate();
  const [path, setPath] = useState<string>(params.path ?? "./");
  const [file, setFile] = useState<File | undefined>();
  const [res, setRes] = useState<string>("");
  const [err, setErr] = useState<string>("");
  const [uploading, setUploading] = useState<boolean>(false);
  const [uploadProgress, setUploadProgress] = useState<number>(0);

  useEffect(() => {
    document.title = "Upload - folderhost"
    if (!Cookies.get("ip") && !Cookies.get("token")) {
      navigate("/login");
    }
  }, [])

  useEffect(() => {
    setRes("");
    setErr("");
    setUploadProgress(0);
    setUploading(false)
  }, [file])

  async function uploadFile() {
    if (file === undefined) {
      alert("Please select a file to upload!")
      return
    }
    const chunkSize: number = 5 * 1024 * 1024; // 5 MB
    const totalChunks: number = Math.ceil(file.size / chunkSize)
    const progressIndex: number = 100 / totalChunks;
    const fileID: string = Date.now().toString();
    const fileName: string = file.name;

    setRes("");
    setErr("");
    setUploadProgress(0);

    try {
      for (let i = 0; i < totalChunks; i++) {
        const chunk = file.slice(i * chunkSize, (i + 1) * chunkSize);
        const formData = new FormData();
        formData.append('file', chunk);
        formData.append('chunkIndex', i.toString());
        formData.append('totalChunks', totalChunks.toString());
        formData.append('fileID', fileID);
        formData.append('fileName', file.name);

        const response = await axiosInstance.post(`/upload?path=${path.slice(1)}`, formData);


        if (response.data.response && !response.data.uploaded) {
          setUploadProgress((prev) => prev + progressIndex)
          setRes(`Uploading ${convertBytesToString(chunkSize * i)} / ${convertBytesToString(file.size)}`)
        }

        if (response.data.uploaded) {
          setRes(response.data.response);
          setUploadProgress(100)
        }
      }
    } catch (error) {
      const err = error as AxiosError<{ err?: string }>;
      console.error(err);
      if (err.response?.data?.err) {
        setErr(err.response?.data?.err || "Upload failed");
      }
    } finally {
      setUploading(false);
    }
  }

  return (
    <div>
      <UploadFileComp
        setFile={setFile}
        uploadFile={uploadFile}
        response={res}
        error={err}
        uploadProgress={uploadProgress}
        path={path}
        uploading={uploading}
        setUploading={setUploading}
      />
    </div>
  )
}

export default UploadFile
import ExplorerPage from './pages/ExplorerPage/Explorer';
import Home from './pages/HomePage/Home';
import Login from './pages/LoginPage/Login';
import "./global.css";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import CodeEditor from './pages/CodeEditorPage/CodeEditor';
import UploadFile from './pages/UploadFilePage/UploadFile';
import NoPage from './pages/NoPage';
import Recovery from './pages/Recovery/Recovery';
import Users from './pages/Users/Users';
import Logs from './pages/Logs/Logs';
import Services from './pages/Services/Services';
import NewUser from './pages/NewUser/NewUser';
import EditUser from './pages/EditUser/EditUser';
import Default from './components/templates/Default';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/">
          <Route index element={<Home />} />
          <Route path="login" element={<Login />} />
          <Route path="explorer" >
            <Route index element={<Default><NoPage /></Default>} />
            <Route path=':path' element={<Default><ExplorerPage /></Default>} />
          </Route>
          <Route path="recovery" >
            <Route index element={<Default><Recovery /></Default>} />
          </Route>
          <Route path="users" >
            <Route index element={<Default><Users /></Default>} />
            <Route path=':username' element={<Default><EditUser /></Default>} />
            <Route path='new' element={<Default><NewUser /></Default>} />
          </Route>
          <Route path="logs" >
            <Route index element={<Default><Logs /></Default>} />
          </Route>
          <Route path="services" >
            <Route index element={<Default><Services /></Default>} />
          </Route>
          <Route path="editor" >
            <Route index element={<NoPage />} />
            <Route path=':path' element={<CodeEditor />} />
          </Route>
          <Route path="upload" >
            <Route index element={<NoPage />} />
            <Route path=':path' element={<UploadFile />} />
          </Route>
          <Route path="*" element={<NoPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;

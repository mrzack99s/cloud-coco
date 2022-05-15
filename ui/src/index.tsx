import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import "primereact/resources/themes/fluent-light/theme.css";
import "primereact/resources/primereact.min.css";
import "primeicons/primeicons.css";
import 'primeflex/primeflex.css'
import '@mdi/font/css/materialdesignicons.css'
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { CookiesProvider } from 'react-cookie';
import Portal from './views/portal';
import ApiProvider from './utils/api_connector';
import Directories from './views/directories';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <CookiesProvider>
      <ApiProvider options={{
        baseUrl: "http://localhost:8000/api"
      }}>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Portal />} />
            <Route path="/directories" element={<Directories />} />
          </Routes>
        </BrowserRouter>
      </ApiProvider>
    </CookiesProvider>

  </React.StrictMode >
);


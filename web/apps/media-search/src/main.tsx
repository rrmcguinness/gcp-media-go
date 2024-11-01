import React from 'react';
import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import './index.css'

import {createBrowserRouter, Navigate, RouterProvider} from 'react-router-dom';
import FileUpload from './pages/FileUpload.tsx';


const Search = React.lazy(() => import('./pages/Search'));

const ErrorBoundary = () => {
  return (<Navigate to="/"/>)
}

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        index: true,
        element: <Search />,
        errorElement: <ErrorBoundary />,
      },
      {
        path: "/uploads",
        element: <FileUpload />,
        errorElement: <ErrorBoundary />,
      }
    ]
  }
])

createRoot(document.getElementById('root')!).render(<RouterProvider router={router}/>)
import React from 'react';
import {createRoot} from 'react-dom/client'
import App from './App.tsx'
import './index.css'

import {createBrowserRouter, Navigate, RouterProvider} from 'react-router-dom';

const Search = React.lazy(() => import('./pages/Search'));
const FileUpload = React.lazy(() => import('./pages/FileUpload'));
const Dashboard = React.lazy(() => import('./pages/Dashboard'))

const ErrorBoundary = () => {
    return (<Navigate to="/"/>)
}

const router = createBrowserRouter([
    {
        path: "/",
        element: <App/>,
        children: [
            {
                index: true,
                element: <Search/>,
                errorElement: <ErrorBoundary/>,
            },
            {
                path: "/uploads",
                element: <FileUpload/>,
                errorElement: <ErrorBoundary/>,
            },
            {
                path: "/dashboard",
                element: <Dashboard/>,
                errorElement: <ErrorBoundary/>,
            }
        ]
    }
])

createRoot(document.getElementById('root')!).render(<RouterProvider router={router}/>)
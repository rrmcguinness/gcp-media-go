import "./App.css";
import {Box, Container, createTheme, CssBaseline, ThemeOptions, ThemeProvider, Typography} from "@mui/material";
import {LoadingIcon} from "./components/LoadingIcon";
import React from "react";
import {Outlet} from "react-router-dom";
import Footer from "./components/Footer";
import TopNav from "./components/TopNav";

const themeOptions: ThemeOptions = {
    palette: {
        mode: 'dark',
        primary: {
            main: '#1565c0',
        },
        secondary: {
            main: '#37474f',
        },
        error: {
            main: '#b71c1c',
        },
        warning: {
            main: '#f4511e',
        },
    },
};

const theme = createTheme(themeOptions);

function App() {
    return (
        <ThemeProvider theme={theme}>
            <TopNav/>
            <Typography variant="h3" sx={{ml: 2, mt: 2, fontFamily: 'Google Sans', fontWeight: 800, color: '#4285F4'}}>Me<span
                style={{color: '#FBBC04'}}>d</span>ia <span style={{color: '#DB4437'}}>S</span>ea<span
                style={{color: '#0F9D58'}}>r</span>ch</Typography>
            <Box
                sx={{
                    position: "relative",
                    display: "flex",
                    flexDirection: "column",
                    minHeight: "100vh",
                    minWidth: "100vw"

                }}
            >
                <CssBaseline/>
                <Container
                    component="main"
                    sx={{mt: 3, pb: "3.5em", mb: 2}}
                    maxWidth="xl"
                >
                    <React.Suspense fallback={<LoadingIcon/>}>
                        <Outlet/>
                    </React.Suspense>
                </Container>
                <Footer/>
            </Box>
        </ThemeProvider>
    );
}

export default App;

import "./App.css";
import { Box, Container, createTheme, CssBaseline, ThemeOptions, ThemeProvider } from "@mui/material";
import { LoadingIcon } from "./components/LoadingIcon";
import React from "react";
import { Outlet } from "react-router-dom";

const themeOptions: ThemeOptions = {
  palette: {
      mode: 'light',
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
      <Box
        sx={{
          position: "relative",
          display: "flex",
          flexDirection: "column",
          minHeight: "100vh",
        }}
      >
        <CssBaseline />
        <Container
          component="main"
          sx={{ mt: 3, pb: "3.5em", mb: 2, overflow: "auto" }}
          maxWidth="xl"
        >
          <React.Suspense fallback={<LoadingIcon />}>
            <Outlet />
          </React.Suspense>
        </Container>
      </Box>
    </ThemeProvider>
  );
}

export default App;

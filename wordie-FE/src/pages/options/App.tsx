import { SnackbarProvider } from 'notistack';
import { StrictMode } from "react";
import { AuthProvider } from "react-auth-kit";
import { RouterProvider } from "react-router-dom";
import { router } from "./router";

export const App = () => {
  return <AuthProvider authType="cookie" authName="auth" cookieDomain="localhost" cookieSecure={false}>
    <StrictMode>
      <SnackbarProvider>
        <RouterProvider router={router} />
      </SnackbarProvider>
    </StrictMode>
  </AuthProvider>
}

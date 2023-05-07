import SignIn from "@src/others/signin";
import SignUp from "@src/others/signup";
import { RequireAuth } from "react-auth-kit";
import { Route, createHashRouter, createRoutesFromElements } from "react-router-dom";
import Options from "./Options";

export const router = createHashRouter(
  createRoutesFromElements(
    <>
      <Route path="/" element={
        <RequireAuth loginPath="/login">
          <Options />
        </RequireAuth>
      }>
      </Route>
      <Route path="/login" element={<SignIn></SignIn>} />
      <Route path="/signup" element={<SignUp></SignUp>} />
    </>
  )
)

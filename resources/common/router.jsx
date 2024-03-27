import { createBrowserRouter } from "react-router-dom";
import Dashboard from "../core/Dashboard";
import UserList from "../users/UserList";

export default createBrowserRouter([
    { path: '', element: <Dashboard/> },
    { path: 'users', element: <UserList/> },
]);

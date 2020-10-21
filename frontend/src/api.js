import Axios from "axios";
import { SERVER_URL } from "./config";

const api = Axios.create({
    baseURL: `${SERVER_URL}/api/v1`,
    headers: {
      "Content-Type": "application/json"
    }
  });

api.login = ({ access_token, user }) => {
    localStorage.setItem("token", access_token);
    localStorage.setItem("user", JSON.stringify(user));
};

api.setUser = (user) => {
    localStorage.setItem("user", JSON.stringify(user));
};
  
api.logout = () => {
    localStorage.removeItem("user");
    localStorage.removeItem("token");
    api.defaults.headers.common["Authorization"] = undefined;
};

api.setToken = () => {
    api.defaults.headers.put["Authorization"] = `Bearer ${localStorage.token}`;
}

api.loggedIn = () => {
    return localStorage["user"]
};

api.setUserInterests = ({interests}) => {
    let user = JSON.parse(localStorage.user)
    user.interests = interests;
    localStorage.setItem("user", JSON.stringify(user));
}

api.getUsername = () => {
    const json = api.userJSON()
    return json && json.username ? json.username : "G"
}

api.userPhoto = () => {
    const json = api.userJSON()
    return !!json && !!json.photo ? json.photo : null
}

api.userJSON = () => {
    return localStorage.user ? JSON.parse(localStorage.user) : null
}

api.getBookId = () => {
    const json = api.userJSON()
    return json ? json.id : null
}

export default api;
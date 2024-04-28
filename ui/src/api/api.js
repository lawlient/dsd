import axios from "axios";



var HOSTNAME = import.meta.env.VITE_BASE_URL
var BASEPATH = import.meta.env.BASE_URL
axios.defaults.baseURL = "https://jovan.vip.cpolar.cn/dsd"
axios.defaults.headers["Authorization"] = localStorage.getItem("__token__")

axios.interceptors.response.use((response) => {
    return response;
}, (error) => { // Anything except 2XX goes to here
    const status = error.response?.status || 500;
    if (status === 401) {
    } else {
        return Promise.reject(error); // Delegate error to calling side
    }
});


export function getdotpng(request) {
    return axios.post("/", request)
}
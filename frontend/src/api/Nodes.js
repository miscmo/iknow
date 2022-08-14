import axios from "axios";

export function SearchNodes(parameter) {
   return axios({
        method: "POST",
        url: "/api/v1/searchNodes",
        data: parameter,
        headers: {"content-type": "text/plain"}
   })
}

export function AddNodes(parameter) {
    return axios({
        method: "POST",
        url: "/api/v1/addNodes",
        data: parameter,
        headers: {"content-type": "text/plain"}
    })
}
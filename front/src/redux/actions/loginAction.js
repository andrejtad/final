import axios from 'axios';
import { loginTypes } from '../types/loginTypes'

export const setToken = payload => ({
    type: loginTypes.SET_TOKEN,
    payload
})

export const deleteToken = () => ({
    type: loginTypes.DEL_TOKEN,
})

export const loginUser = (username, password) => dispatch => {
    const data = JSON.stringify({ username, password })
    const url = '/api/auth/sign-in'

    axios.post(url, data)
        .then(res => {
            if (res.data && res.data.token) {
                localStorage.setItem("token", res.data.token)
                dispatch(setToken(res.data.token))
            }
        })
        .catch(err => console.log(err))
}

export const registerUser = (name, username, password) => dispatch => {
    const data = JSON.stringify({ name, username, password })
    const url = '/api/auth/sign-up'

    axios.post(url, data)
        .then(res => console.log(res))
        .catch(err => console.log(err))
}
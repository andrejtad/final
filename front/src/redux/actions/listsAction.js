import axios from 'axios'
import { listsTypes } from '../types/listsTypes'

export const setList = payload => ({
    type: listsTypes.SET_LIST,
    payload
})

export const setItems = (data, id) => ({
    type: listsTypes.SET_ITEM,
    data,
    id
})

export const setSpecs = (data, id) => ({
    type: listsTypes.SET_SPEC,
    data,
    id
})

export const setLinkType = payload => ({
    type: listsTypes.SET_LINK_TYPE,
    payload
})

export const setLink = payload => ({
    type: listsTypes.SET_LINK,
    payload
})

export const getList = () => (dispatch, getState) => {
    const token = getState().login.token
    const url = '/api/dataowners'
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };

    axios.get(url, config)
        .then(res => {
            if (res.data && res.data.data) {
                dispatch(setList(res.data.data))
            }
        })
        .catch(err => console.log(err))
}

export const addList = (title, description) => (dispatch, getState) => {
    const token = getState().login.token

    const url = '/api/dataowners'
    const data = JSON.stringify({ title, description })
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };

    axios.post(url, data, config)
        .then(res => {
            if (res.status === 200) {
                dispatch(getList())
            }
        })
        .catch(err => console.log(err))
}

export const removeList = id => (dispatch, getState) => {
    const token = getState().login.token

    const url = `/api/dataowners/${id}`
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };

    axios.delete(url, config)
        .then(res => {
            if (res.status === 200) {
                dispatch(getList())
            }
        })
        .catch(err => console.log(err))
}

export const getDatasets = id => (dispatch, getState) => {
    const token = getState().login.token
    const url = `/api/dataowners/${id}/datasets`
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };

    axios.get(url, config)
        .then(res => {
            if (res.data && res.data.data) {
                dispatch(setItems(res.data.data, id))
            }
        })
        .catch(err => console.log(err))
}

export const addItem = (title, description, id) => (dispatch, getState) => {
    const token = getState().login.token

    const url = `/api/dataowners/${id}/datasets`
    const data = JSON.stringify({ title, description })
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };

    axios.post(url, data, config)
        .then(res => {
            if (res.status === 200) {
                dispatch(getDatasets(id))
            }
        })
        .catch(err => console.log(err))
}

export const getSpecs = id => (dispatch, getState) => {
    const token = getState().login.token
    const url = `/api/datasets/${id}/specifications`
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };

    axios.get(url, config)
        .then(res => {
            if (res.data && res.data.data) {
                dispatch(setSpecs(res.data.data, id))
            }
        })
        .catch(err => console.log(err))
}

export const getLinkType = () => (dispatch, getState) => {
    const token = getState().login.token
    const url = `/api/linktype`
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };

    axios.get(url, config)
        .then(res => {
            if (res.data && res.data.data) {
                dispatch(setLinkType(res.data.data))
            }
        })
        .catch(err => console.log(err))
}

export const getLink = id => (dispatch, getState) => {
    const token = getState().login.token
    const url = `/api/link/${id}/specification`
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };

    axios.get(url, config)
        .then(res => {
            if (res.data && res.data.data) {
                dispatch(setLink(res.data.data))
            }
        })
        .catch(err => console.log(err))
}


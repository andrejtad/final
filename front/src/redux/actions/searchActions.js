import { searchTypes } from '../types/searchTypes'

export const setTag = payload => ({
    type: searchTypes.SET_TAG,
    payload
})

export const deleteTag = payload => ({
    type: searchTypes.DELETE_TAG,
    payload
})
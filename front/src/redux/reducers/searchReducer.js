import { searchTypes } from '../types/searchTypes'

const initialState = {
    tags: []
}

export const searchReducer = (state = initialState, action) => {
    switch (action.type) {
        case searchTypes.SET_TAG:
            return {
                ...state,
                tags: [...state.tags, action.payload]
            };
        case searchTypes.DELETE_TAG:
            return { ...state, tags: [] };
        default:
            return state;
    }
};
import { loginTypes } from '../types/loginTypes'

const initialState = {
    token: null,
}

export const loginReducer = (state = initialState, action) => {
    switch (action.type) {
        case loginTypes.SET_TOKEN:
            return { ...state, token: action.payload };
        case loginTypes.DEL_TOKEN:
            return { ...state, token: null };
        default:
            return state;
    }
};
import { listsTypes } from '../types/listsTypes'

const initialState = {
    dataowners: [],
    datasets: [],
    specifications: [],
    linkTypes: [],
    links: []
}

export const listsReducer = (state = initialState, action) => {
    switch (action.type) {
        case listsTypes.SET_LIST:
            return { ...state, dataowners: action.payload };
        case listsTypes.SET_ITEM:
            return {
                ...state,
                datasets: {
                    ...state.datasets,
                    [action.id]: action.data
                }
            };
        case listsTypes.SET_SPEC:
            return {
                ...state,
                specifications: {
                    ...state.specifications,
                    [action.id]: action.data
                }
            };
        case listsTypes.SET_LINK_TYPE:
            return { ...state, linkTypes: action.payload };
        case listsTypes.SET_LINK:
            return { ...state, links: [...state.links, action.payload]};
        default:
            return state;
    }
};

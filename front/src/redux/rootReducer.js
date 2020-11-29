import { combineReducers } from 'redux'
import { loginReducer } from './reducers/loginReducer'
import { listsReducer } from './reducers/listsReducer'
import { searchReducer } from './reducers/searchReducer'

export const rootReducer = combineReducers({
    search: searchReducer,
    login: loginReducer,
    lists: listsReducer,
})
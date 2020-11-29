import React from 'react'
import { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useHistory } from 'react-router-dom';
import { setToken } from '../../redux/actions/loginAction'
import { Search } from '../../components/Search/Search'
import './Main.scss'

export function Main() {
    const { token } = useSelector(state => state.login)
    const dispatch = useDispatch()
    const history = useHistory()

    useEffect(() => {
        if (!token) dispatch(setToken(localStorage.getItem("token")))
    }, [])

    if (!token) history.push('/login')

    return (
        <div className='main'>
            <div className='container main__search'>
                <Search/>
            </div>
        </div>
    )
}
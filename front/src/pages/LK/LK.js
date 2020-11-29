import React from 'react'
import { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useHistory } from 'react-router-dom';
import { setToken } from '../../redux/actions/loginAction'
import { Search } from '../../components/Search/Search'
import { Ministry } from '../../components/Minitstry/Ministry'
import './LK.scss'

export function LK() {
    const { token } = useSelector(state => state.login)
    const dispatch = useDispatch()
    const history = useHistory()

    useEffect(() => {
        if (!token) dispatch(setToken(localStorage.getItem("token")))
    }, [])

    if (!token) history.push('/login')

    return (
        <div className='lk'>
            <Search/>

            <div className='container'>
                <Ministry/>
            </div>
        </div>
    )
}
import React from 'react'
import { Link, useHistory } from 'react-router-dom'

import { useDispatch } from 'react-redux'
import { deleteToken } from '../../../redux/actions/loginAction'
import Logo from '../../../assets/img/logo.png'
import './Header.scss'

export function Header() {
    const dispatch = useDispatch()
    const history = useHistory()

    const handleClickLogout = () => {
        dispatch(deleteToken())
        history.push('/login')
        localStorage.removeItem('token')
    }

    return (
        <header className='header'>
            <div className='container header_pos'>
                <Link to={'/'}>
                    <div className='header__logo'>
                        <img src={Logo} alt=""/>
                        <div className='header__left'>
                            Система предоставления государственных данных
                        </div>
                    </div>
                </Link>

                <div className='header__right'>
                    <div
                        onClick={handleClickLogout}
                        className="header__right__logout"
                    />
                </div>
            </div>
        </header>
    )
}

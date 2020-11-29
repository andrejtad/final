import React from 'react'
import Book from '../../assets/img/book.png'
import Diagram from '../../assets/img/diagram.png'
import './Sidebar.scss'
import { Link } from 'react-router-dom'

export function Sidebar() {
    return (
        <div className='sidebar'>
            <Link to={'/lk'}><img src={Book} alt=""/></Link>
            <Link to={'/diagram'}><img src={Diagram} alt=""/></Link>
        </div>
    )
}
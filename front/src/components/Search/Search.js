import React, { useState } from 'react'
import { Button, Form, Toast } from 'react-bootstrap'
import { useDispatch, useSelector } from 'react-redux'
import { useHistory } from 'react-router-dom';
import { deleteTag, setTag } from '../../redux/actions/searchActions'
import groupIcons from '../../assets/img/group-icons-0.png'
import './Search.scss'

export function Search() {
    const dispatch = useDispatch()
    const history = useHistory()
    const { tags } = useSelector(state => state.search)
    const [value, setValue] = useState('');

    const isMainPage = window.location.pathname === '/'

    const handleChange = e => setValue(e.target.value)

    const search = () => {
        if (!value) return

        setValue('')
        dispatch(setTag(value))
        history.push('/diagram')
    }

    const handleDelete = () => {
        dispatch(deleteTag())
    }

    return (
        <div className={`search ${isMainPage ? 'search_main' : ''}`}>
            <Form.Control onChange={handleChange} value={value} placeholder="Поисковый запрос: "/>
            <div className="search__bottom">
                <div className='search__bottom__tags'/>
                <Chips tags={tags} handleDelete={handleDelete}/>
                <Button onClick={search} variant="primary search__bottom__button">Найти</Button>
                {isMainPage && <img src={groupIcons} alt="" className='search__bottom__icons'/>}
            </div>
        </div>
    )
}

function Chips({tags, handleDelete}) {
    return tags.map((tag, i) =>
        <Toast key={i} className='mr-2' onClose={handleDelete}>
            <Toast.Header>
                <strong className="mr-auto">{tag}</strong>
            </Toast.Header>
        </Toast>
    )
}
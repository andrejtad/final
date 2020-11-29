import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { addList, getList } from '../../redux/actions/listsAction'
import { Accordion } from '../Accordion/Accordion'
import { Alert, Button, Form } from 'react-bootstrap'
import './Ministry.scss'

export function Ministry() {
    const dispatch = useDispatch()
    const { dataowners } = useSelector(state => state.lists)
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [isAddFormMin, setIsAddFormMin] = useState(false);

    useEffect(() => {
        if (dataowners.length === 0) {
            dispatch(getList())
        }
    }, [])

    const handleChangeTitle = e => setTitle(e.target.value)
    const handleChangeDescription = e => setDescription(e.target.value)

    const addFormMin = () => setIsAddFormMin(prev => !prev)

    const addMinistry = () => {
        if (!title & !description) return addFormMin()
        setTitle('')
        setDescription('')
        addFormMin()
        dispatch(addList(title, description))
    };

    return (
        <div className='ministry'>
            <Alert variant={'dark'}>
                <h4 className='d-flex justify-content-between align-items-center mb-0'>
                    <div>Министерства</div>
                    {!isAddFormMin && <Button onClick={addFormMin} variant="secondary">Добавить министерство</Button>}
                </h4>
            </Alert>

            {isAddFormMin &&
                <>
                    <div className='ministry__form'>
                        <Form.Label>Название министерства</Form.Label>
                        <Form.Control onChange={handleChangeTitle} value={title} placeholder="Введите название министерства"/>
                        <Form.Label style={{marginTop: 15}}>Описание министерства</Form.Label>
                        <Form.Control onChange={handleChangeDescription} value={description}
                                      placeholder="Введите описание министертсва"/>
                        <Button onClick={addMinistry} variant="outline-primary mt-3">Сохранить</Button>
                    </div>
                </>
            }

            <div className='ministry__list'>{dataowners.length > 0 && Accordion(dataowners.sort((a, b) => b.id - a.id))}</div>
        </div>
    )
}

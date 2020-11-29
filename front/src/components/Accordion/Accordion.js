import React, { Fragment, useState } from 'react';
import { Collapse } from 'react-collapse';
import './Accordion.scss'
import { useDispatch, useSelector } from 'react-redux'
import { addItem, getDatasets, removeList } from '../../redux/actions/listsAction'
import { ItemList } from '../ItemList/ItemList'
import { Button, Form } from 'react-bootstrap'

const Ministry = min => {
    const { datasets } = useSelector(state => state.lists)
    const dispatch = useDispatch()
    const [isAddFormData, setIsAddFormData] = useState(false);
    const [isOpen, setIsOpen] = useState(false);
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');

    const handleClick = () => {
        setIsOpen(prev => !prev);
        dispatch(getDatasets(min.id))
    }

    const handleChangeTitle = e => setTitle(e.target.value)
    const handleChangeDescription = e => setDescription(e.target.value)

    const addFormData = () => setIsAddFormData(prev => !prev)

    const handleClickRemove = () => {
        dispatch(removeList(min.id))
    }

    const addItems = () => {
        if (!title && !description) return addFormData()
        setTitle('')
        setDescription('')
        addFormData()
        dispatch(addItem(title, description, min.id))
    };

    return (
        <>
            <div className={`accordion__category ${isOpen ? 'accordion__category_active' : ''}`} onClick={handleClick}>
                <span className='accordion__category__title'>{min.title}</span>
                <span onClick={handleClickRemove}>X</span>
            </div>

            <Collapse isOpened={isOpen}>
                <div className='accordion__category__desc'>{min.description}</div>
                {isOpen && <hr/>}

                <h5 className='d-flex justify-content-between align-items-center'>
                    <div>Набор данных</div>
                    {!isAddFormData && <Button onClick={addFormData} variant="outline-secondary">Добавить набор данных</Button>}
                </h5>

                {isAddFormData &&
                <>
                    <div className='ministry__form mt-3'>
                        <Form.Label>Название набора данных</Form.Label>
                        <Form.Control onChange={handleChangeTitle} value={title} placeholder="Введите название набора данных"/>
                        <Form.Label style={{marginTop: 15}}>Набор данных</Form.Label>
                        <Form.Control onChange={handleChangeDescription} value={description}
                                      placeholder="Введите описание набора данных"/>
                        <Button onClick={addItems} variant="outline-primary mt-3">Сохранить</Button>
                    </div>
                </>
                }

                <ItemList list={datasets[min.id]}/>
            </Collapse>
        </>
    );
}

export const Accordion = props =>
    <div className='accordion'>
        {props.map(item => <Ministry key={item.id} {...item}/>)}
    </div>;


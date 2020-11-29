import React, { useState } from 'react'
import { useHistory } from 'react-router-dom'
import { useDispatch } from 'react-redux'
import { addOrder } from '../../redux/actions/orderAction'
import './AddOrder.scss'

export function AddOrder() {
    const history = useHistory()
    const dispatch = useDispatch()
    const [order, setOrder] = useState({
        title: '',
        description: '',
        price: null,
    })
    const [steps, setSteps] = useState([{
        number: 0,
        title: '',
        description: '',
        deadline: null
    }])

    const hanldeClickAddStep = () => {
        setSteps(prev => ([
                ...prev,
                {
                    number: prev.length,
                    stepTitle: '',
                    stepDescription: '',
                    deadline: null
                }
            ]
        ))
    }

    const handleOnChange = e => {
        const name = e.target.name
        const value = e.target.value

        setOrder(prev => ({...prev, [name]: value }))
    }

    const handleOnChangeStep = (e, number) => {
        const name = e.target.name
        const value = e.target.value
        setSteps(prev => prev.map(step =>
            step.number === number
                ? {...step, [name]: value}
                : step
        ))
    }

    const handleSubmit = () => {
        const orderWithSteps = {
            order,
            steps
        }
        dispatch(addOrder(orderWithSteps))
        history.push('/lk/customer')
    }
    const Step = ({number, title, description, deadline}, i) => (
        <div key={i} className='d-flex'>
            <div className='steps__step__number'>{number + 1}.</div>
            <div className='steps__step__desc'>
                <div>Название</div>
                <input className='input' type='text' name='stepTitle' value={steps[number].stepTitle} onChange={e => handleOnChangeStep(e, number)} placeholder='Название'/>
                <div>Описание</div>
                <textarea className='input textarea' name='stepDescription' value={steps[number].stepDescription} onChange={e => handleOnChangeStep(e, number)} placeholder='Описание'/>
                <div>Срок</div>
                <input  className='input' type="text" name='deadline' value={steps[number].deadline} onChange={e => handleOnChangeStep(e, number)} placeholder='Цена'/>
            </div>
        </div>
    )

    return (
        <div className='add-order'>
            <p className='add-order__title'>Добавление заказа</p>
            <div className='add-order__form'>
                <input className='input' name='title' type='text' value={order.title} onChange={handleOnChange} placeholder='Название'/>
                <textarea className='input textarea' value={order.description} onChange={handleOnChange} name='description' placeholder='Описание'/>
                <input className='input' type="number" value={order.price} onChange={handleOnChange} name='number' placeholder='Цена'/>
            </div>

            <p className='add-order__title'>Этапы</p>

            <div className='steps'>
                {!!steps.length && steps.map(Step)}
                <div className='steps__add button__white width-200 mb-3' onClick={hanldeClickAddStep}>Добавить этап</div>
                <div className='steps__add button__white width-200' onClick={handleSubmit} style={{color: '#FFFFFF'}}>Разместить заказ</div>
            </div>
        </div>
    )
}
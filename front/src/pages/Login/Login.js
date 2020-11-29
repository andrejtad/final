import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link, useHistory } from 'react-router-dom';
import { Form, Button } from 'react-bootstrap'

import { loginUser, registerUser } from '../../redux/actions/loginAction';
import logo from '../../assets/img/logo-login.png';
import './Login.scss';

export function Login() {
    const dispatch = useDispatch();
    const { token } = useSelector(state => state.login);
    const history = useHistory();
    const [name, setName] = useState('');
    const [login, setLogin] = useState('');
    const [pass, setPass] = useState('');
    const isLogin = window.location.pathname === '/login';

    useEffect(() => {
        setLogin('');
        setPass('');
    }, [isLogin])

    const handleChangeName = e => setName(e.target.value);
    const handleChangeLogin = e => setLogin(e.target.value);
    const handleChangePass = e => setPass(e.target.value);

    const handleSubmit = () => {
        if (isLogin) {
            dispatch(loginUser(login, pass));
        } else {
            dispatch(registerUser(name, login, pass))
            history.push('/login')
        }

        setName('');
        setLogin('');
        setPass('');
    };

    const data = {
        entry: isLogin ? 'Войти в систему' : 'Регистрация',
        smallText: isLogin ? 'Для входа введите логин и пароль' : 'Для регистрации введите имя, логин и пароль',
        submit: isLogin ? 'Войти' : 'Зарегистрироваться',
        linkName: isLogin ? 'Зарегистрироваться' : 'Войти',
        link: isLogin ? '/sing-up' : '/login'
    }

    if (token) history.push('/')

    return (
        <div className='login'>
            <div className='login__window login__window_pos'>
                <img src={logo} alt=''/>

                <h5 className='mt-4 mb-0'>{data.entry}</h5>
                <div className='login__small-text mb-3'>{data.smallText}</div>

                {!isLogin && <Form.Group>
                    <Form.Control onChange={handleChangeName} value={name} type="email" placeholder="Введи ваше имя" />
                </Form.Group>}

                <Form.Group>
                    <Form.Control onChange={handleChangeLogin} value={login} type="email" placeholder="Введи логин" />
                </Form.Group>

                <Form.Group >
                    <Form.Control onChange={handleChangePass} value={pass} type="password" placeholder="Введите пароль" />
                </Form.Group>

                <Form.Group className='d-flex justify-content-between align-items-center'>
                    <Button onClick={handleSubmit} variant="primary" type="submit" className='button button_blue'>
                        {data.submit}
                    </Button>
                    <Link to={data.link}>{data.linkName}</Link>
                </Form.Group>
            </div>
        </div>
    );
}
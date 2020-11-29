import React from 'react'

import { Header } from './Header/Header'
import { Footer } from './Footer/Footer'
import { Sidebar } from '../Sidebar/Sidebar'

export const Layout = props => (
    <div style={{display: 'flex'}}>
        <Sidebar/>
        <div style={{width: '100%'}}>
            <Header />
            {props.children}
            {/*<Footer />*/}
        </div>
    </div>
)

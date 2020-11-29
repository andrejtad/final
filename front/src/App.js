import { createBrowserHistory } from 'history';
import React, { useEffect } from 'react'
import { Router, Route, Switch, Redirect } from 'react-router-dom';

import { Layout } from './components/Layout';
import { Main } from './pages/Main/Main';
import { Login } from './pages/Login/Login';
import './App.scss'
import { DiagramPage } from './pages/Diagram/Diagram'
import { LK } from './pages/LK/LK'

const history = createBrowserHistory();

const withLayout = Component => props => (
    <Layout>
        <Component props={props} />
    </Layout>
);


function App() {
    return (
        <Router history={history}>
            <Switch>
                <Route path={'/'} exact component={Main}/>
                <Route path={'/login'} component={Login}/>
                <Route path={'/sing-up'} component={Login}/>
                <Route path={'/lk'} component={withLayout(LK)}/>
                <Route path={'/diagram'} component={withLayout(DiagramPage)}/>
            </Switch>
        </Router>
    );
}

export default App;

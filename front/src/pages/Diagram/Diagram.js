import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useHistory } from 'react-router-dom';

import { Search } from '../../components/Search/Search'
import { getDatasets, getLink, getLinkType, getList, getSpecs } from '../../redux/actions/listsAction'
import { setToken } from '../../redux/actions/loginAction'
import './Diagram.scss'
import { Canvas } from '../../components/Canvas/Canvas'

export function DiagramPage() {
    const dispatch = useDispatch()
    const history = useHistory()
    const isN = window.location.search
    const { dataowners, datasets, specifications, linkTypes, links } = useSelector(state => state.lists)
    const { token } = useSelector(state => state.login)

    const handleChange = () => history.push('/diagram?p=2')

    useEffect(() => {
        window.addEventListener('click', handleChange)
        return () => {
            window.removeEventListener('click', handleChange)
        }
    })

    useEffect(() => {
        if (!token) dispatch(setToken(localStorage.getItem("token")))
    }, [])

    useEffect(() => {
        if (dataowners.length === 0 && token) {
            if (links.length === 0) dispatch(getList())
            //if (linkTypes.length === 0)dispatch(getLinkType())
        }
    }, [token])

    useEffect(() => {
        if (dataowners.length > 0 && token) {
            dataowners.sort((a, b) => a.id - b.id).forEach(d => dispatch(getDatasets(d.id)))
        }
    }, [dataowners])

    useEffect(() => {
        if (dataowners.length in datasets &&  Object.values(datasets).length > 0 && token) {
            Object.values(datasets).forEach(dataset => {
                dataset.forEach(d => dispatch(getSpecs(d.id)))
            })
        }
    }, [datasets])

    useEffect(() => {
        if (
            Object.keys(datasets).length === dataowners.length
            && Object.keys(specifications).length === Object.values(datasets).flat().length
            && links.length <= 11
            && Object.keys(datasets).length > 0
            && links.length === 0
            && token
        ) {
            Object.values(specifications).forEach(specification => {
                specification.forEach(s => dispatch(getLink(s.id)))
            })
        }
    }, [specifications])

    return (
        <div className='diagram'>
            <Search/>
            {
                (Object.keys(datasets).length === dataowners.length
                && Object.keys(specifications).length === Object.values(datasets).flat().length
                && Object.keys(datasets).length > 0) && (
                    <Canvas
                        dataowners={dataowners.sort((a, b) => a.id - b.id)}
                        datasets={datasets}
                        specifications={specifications}
                        linkTypes={linkTypes}
                        links={links}
                        isN={isN}
                    />
                )
            }
        </div>
    )
}
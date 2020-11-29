import React, { useEffect } from 'react'
import createEngine, { DefaultNodeModel, DiagramModel } from '@projectstorm/react-diagrams'
import { CanvasWidget } from '@projectstorm/react-canvas-core';
import './Canvas.scss'

export function Canvas({dataowners, datasets, specifications, isN, linkTypes, links}) {
    const engine = createEngine();

    const getNodeProps = (dataset) =>  {

        const node = new DefaultNodeModel({
            name: dataset.title,
            color: 'rgb(0,192,255)'
        })

        node.setPosition(dataset.id % 2 ? 200 : 800, 130 * dataset.id)

        let ports = []
        specifications[dataset.id].forEach(specification => {
            ports = [...ports, (dataset.id % 2)
                ? node.addOutPort(specification.title)
                : node.addInPort(specification.title)
            ]
        })

        return node
    };

    const nodes = () => {
        let nodeArray = []
        dataowners.forEach(dataowner => {
            datasets[dataowner.id].forEach(dataset => {
                nodeArray = [
                    ...nodeArray,
                    getNodeProps(dataset)
                ]
            })
        })

        return nodeArray
    }



// node 1
    const node1 = new DefaultNodeModel({
        name: 'ПФР - Зарегистрированные в ПФР лица',
        color: 'rgb(0,192,255)',
    });
    node1.setPosition(150, isN ? 150 : 100);
    let port11 = node1.addOutPort('Фамилия');
    let port12 = node1.addOutPort('Имя');
    let port13 = node1.addOutPort('Отчество');
    let port14 = node1.addOutPort('Дата рождения');
    let port15 = node1.addOutPort('СНИЛС');
    let port16 = node1.addOutPort('Регистрационный номер');
    let port17 = node1.addOutPort('ИНН');
    let port18 = node1.addOutPort('Паспорт');

// node 2
    const node2 = new DefaultNodeModel({
        name: 'ФНС - ЕГРЮЛ',
        color: 'rgb(0,192,255)',
    });

    node2.setPosition(isN ? 450 : 600, isN ? 700 : 100);
    let port21 = node2.addInPort('Название');
    let port22 = node2.addInPort('ОГРН');
    let port23 = node2.addInPort('ИНН');

// node 3
    const node3 = new DefaultNodeModel({
        name: 'ФНС - Лица на налоговом учете',
        color: 'rgb(0,192,255)',
    });
    node3.setPosition(850, isN ? 150 : 100);
    let port31 = node3.addInPort('ИНН');
    let port32 = node3.addInPort('Фамилия');
    let port33 = node3.addInPort('Имя');
    let port34 = node3.addInPort('Отчество');
    let port35 = node3.addInPort('Дата рождения');
    let port36 = node3.addInPort('Паспорт');

// node 4
    const node4 = new DefaultNodeModel({
        name: 'ФНС - База ИНН',
        color: 'rgb(0,192,255)',
    });
    node4.setPosition(isN ? 850 : 950, 400);
    let port41 = node4.addInPort('ИНН');
    let port42 = node4.addInPort('Дата предоставления');

// node 5
    const node5 = new DefaultNodeModel({
        name: 'ПФР - Снилс',
        color: 'rgb(0,192,255)',
    });
    node5.setPosition(isN ? 600 : 300, isN ? 50 : 400);
    let port51 = node5.addOutPort('Номер');


// node 6
    const node6 = new DefaultNodeModel({
        name: 'ФОМС - Полисы ОМС',
        color: 'rgb(0,192,255)',
    });
    node6.setPosition(150, isN ? 650 : 550);
    let port61 = node6.addOutPort('Фамилия');
    let port62 = node6.addOutPort('Имя');
    let port63 = node6.addOutPort('Отчество');
    let port64 = node6.addOutPort('Дата рождения');
    let port65 = node6.addOutPort('Паспорт');
    let port66 = node6.addOutPort('Страховая организация');
    let port67 = node6.addOutPort('Дата постановки на учет');

// node 7
    const node7 = new DefaultNodeModel({
        name: 'ФОМС - Страховые организации',
        color: 'rgb(0,192,255)',
    });
    node7.setPosition(150, isN ? 450 : 800);
    let port71 = node7.addOutPort('Название');
    let port72 = node7.addOutPort('ИНН');
    let port73 = node7.addOutPort('ОГРН');


// node 8
    const node8 = new DefaultNodeModel({
        name: 'МВД - Паспортные данные',
        color: 'rgb(0,192,255)',
    });
    node8.setPosition(650, isN ? 550 : 650);
    let port81 = node8.addInPort('Номер паспорта');
    let port82 = node8.addInPort('Фамилия');
    let port83 = node8.addInPort('Имя');
    let port84 = node8.addInPort('Отчество');
    let port85 = node8.addInPort('Дата рождения');
    let port86 = node8.addInPort('Место рождения');
    let port87 = node8.addInPort('Дата выдачи');
    let port88 = node8.addInPort('Кем выдан');
    let port89 = node8.addInPort('Код подразделения');

// node 9
    const node9 = new DefaultNodeModel({
        name: 'МВД - Регистрация по месту пребывания',
        color: 'rgb(0,192,255)',
    });
    node9.setPosition(950, isN ? 550 : 650);
    let port91 = node9.addInPort('Номер паспорта');
    let port92 = node9.addInPort('Дата регистрации');
    let port93 = node9.addInPort('Место регистрации');
    let port94 = node9.addInPort('Орган регистрации');

// node 10
    const node10 = new DefaultNodeModel({
        name: 'МВД - Органы внутренних дел',
        color: 'rgb(0,192,255)',
    });
    node10.setPosition(950, isN ? 750 : 850);
    let port101 = node10.addInPort('Название');
    let port102 = node10.addInPort('Территория обслуживания');
    let port103 = node10.addInPort('Адрес');
    let port104 = node10.addInPort('Телефон');




// link them and add a label to the link
    const link11 = port15.link(port51);
    const link12 = port17.link(port31);
    const link13 = port18.link(port81);
    link13.setColor('#0bb7c1');

    const link21 = port22.link(port73);
    link21.setColor('#0bb7c1');
    const link22 = port23.link(port41);

    const link31 = port31.link(port41);
    const link32 = port36.link(port81);
    link32.setColor('#0bb7c1');

    const link41 = port41.link(port72);
    link41.setColor('#0bb7c1');

    const link61 = port66.link(port71);
    link61.setColor('#FF0000');

    const link81 = port81.link(port91);
    const link82 = port88.link(port101);

    const link91 = port94.link(port101);

    const nodess = [
        node1, node2, node3, node4, node5, node6, node7, node8, node9, node10,
        link11, link12, link13, link21, link22, link31, link32, link41, link61, link81, link82, link91
    ]

    const model = new DiagramModel();

    model.addAll(...nodess);
    engine.setModel(model);

    return <CanvasWidget engine={engine} className='canvas' />
}
import React from 'react'

export function ItemList({list = []}) {
    return list.map(Item)
}

function Item(item) {
    return (
        <div key={item.id}>{item.title}</div>
    )
}
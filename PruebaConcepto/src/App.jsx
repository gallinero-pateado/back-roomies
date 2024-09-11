import React from 'react';
import './App.css'
import {Roomies} from './RoomiesCard'
import persona1 from './img/Persona1.jpg'
import persona2 from './img/Persona2.jpg'
import persona3 from './img/Persona3.jpg'
import persona4 from './img/Persona4.jpg'
import persona5 from './img/Persona5.jpg'

const infoo= "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Illo at quaerat eos iure tempora, ni"

const usuarios = [
    {
        userName: "Valeria Henriquez",
        info: infoo,
        ubicacion: "Santiago centro",
        isFav: false,
        perfilImg: persona2
    },

    {
        userName: "John Marston",
        info: infoo,
        ubicacion: "Ñuñoa",
        isFav: true,
        perfilImg: persona3
    },
    
    {
        userName: "Alberto Hurtado",
        info: infoo,
        ubicacion: "La cisterna",
        isFav: false,
        perfilImg: persona4
    },

    {
        userName: "Carolina Rojas",
        info: infoo,
        ubicacion: "Macul",
        isFav: true,
        perfilImg: persona5
    },

    

    {
        userName: "Arthur Morgan",
        info: infoo,
        ubicacion: "La florida",
        isFav: true,
        perfilImg: persona1
    },    
]

export function App(){

    return (
        <React.Fragment>
            <div className='tittle'>
                <h1>Busca tu roomie</h1>
                <h2>Utem Link</h2></div>
            <section className='App'>
                
                    {usuarios.map((usuarios, index) =>(
                        <Roomies
                        key={index}
                        userName={ usuarios.userName}
                        info = {usuarios.info}
                        ubicacion={usuarios.ubicacion}
                        isFav={usuarios.isFav}
                        perfilImg={usuarios.perfilImg}
                        />
                    ))}
            </section>
        </React.Fragment>
    )
}


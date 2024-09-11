import { useState } from "react"
export function Roomies({userName,info ,ubicacion, perfilImg}){

    const [isFav, setIsFav] = useState(false);
    const botonFav = isFav ? 'RoomieCard-button is-fav': 'RoomieCard-button'

    const handleClick = () => {
        setIsFav(!isFav)
    }

    return(
       <article className='RoomieCard'>
        <header className='RoomieCard-header'>
            <img className = "RoomieCard-img" src={perfilImg} alt={'${userName} perfilImg'} />
            <div className='RoomieCard-info'>
                <strong>{userName}</strong>
                <p>{info}</p>
                <span className='RoomieCard-infoUserName'>Ubicacion: {ubicacion}</span>
            </div>
        </header>

        <aside>
            <button className={botonFav} onClick={handleClick}>Favorito</button>
            <button className='RoomieCard-button'>Contactar</button>
        </aside>

       </article>
    )

}


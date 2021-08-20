import {useState} from 'react';

const James = () => {
    const [carName, setCarName] = useState("")
    
    fetch('/products')
    .then(response => response.json())
    .then(data => {
      console.log(data)
      setCarName(data.name)
    });

    return <h1>My name is... {carName}</h1>
}

export default James
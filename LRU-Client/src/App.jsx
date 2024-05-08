import { useEffect, useState } from 'react'
import axios from 'axios';
import ListView from './components/list';
import './App.css'

function App() {
  const [caches, setCaches] = useState([])

  axios.defaults.baseURL = 'http://localhost:3004/api'  

  const getCaches = async () => {
    try {
      const response = await axios.get(`/GetKey`);
      if (response.status === 200) {
        setCaches(response.data);
      }
    } catch (error) {
      console.error('Error fetching cache data:', error.message);
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      await getCaches();
    };
    fetchData();
  }, []); 

  console.log(caches);

  return (
    <>
      {
        caches.length > 0 ? (
            <ListView cache={caches} trigger={getCaches} setCaches={setCaches} />
        ) : (
          <h1>No Data..</h1>
        )
      }
    </>
  );
}

export default App;

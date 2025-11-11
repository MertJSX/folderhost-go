import { useNavigate } from 'react-router-dom'
import Cookies from 'js-cookie';
import React from 'react';

const Home: React.FC = () => {
  const navigate = useNavigate();

  React.useEffect(() => {
    if (Cookies.get("token")) {
      navigate(`/explorer/${encodeURIComponent("./")}`)
    } else {
      navigate("/login")
    }
  }, [])

  return (
    <>
    </>
  )
}

export default Home
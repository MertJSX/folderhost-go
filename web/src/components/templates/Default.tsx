import Header from "../Header/Header"

interface DefaultProps {
  children?: React.ReactNode;
}

const Default: React.FC<DefaultProps> = ({ children }) => {
  return (
    <div>
        <Header />
        {children}
    </div>
  )
}

export default Default
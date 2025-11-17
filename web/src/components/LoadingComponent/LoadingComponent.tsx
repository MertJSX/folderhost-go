import loadingGIF from "../../assets/loading.gif"

const LoadingComponent = () => {
    return (
        <div className='flex items-center justify-center'>
            Loading
            <img src={loadingGIF} width={40} height={40} className='select-none' alt='' />
        </div>
    )
}

export default LoadingComponent
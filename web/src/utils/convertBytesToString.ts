const convertBytesToString = function(bytes: number) {
  const sizes = ["Bytes", "KB", "MB", "GB", "TB"]

  if (bytes == 0) {
    return "N/A"
  }

  const i: number = Math.floor(Math.log(bytes) / Math.log(1024))

  if (i == 0) {
    return bytes + " " + sizes[i]
  }

  return (bytes / Math.pow(1024, i)).toFixed(1) + " " + sizes[i]
}

export default convertBytesToString
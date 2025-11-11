const convertToBytes = (sizeString: string): number => {
    const sizes = ["Bytes", "KB", "MB", "GB", "TB"] as const;

    const [value, unit] = sizeString.split(" ");

    const numericValue = parseFloat(value);
    const index = sizes.indexOf(unit as typeof sizes[number]);

    if (index === -1 || isNaN(numericValue)) {
        return 0;
    }

    return numericValue * Math.pow(1024, index);
}

export default convertToBytes;
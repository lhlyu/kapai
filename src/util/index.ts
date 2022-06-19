

export const Encode = (str: string):string => {
    return encodeURIComponent(str)
}


console.log(Encode('{"id":"123","name":"123","roomId":"123"}'))
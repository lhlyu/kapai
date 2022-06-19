import CryptoJS from 'crypto-js';

export const base64Encode = (str: string):string => {
    const wordArray = CryptoJS.enc.Utf8.parse(str)
    return CryptoJS.enc.Base64.stringify(wordArray)
}

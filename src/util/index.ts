const ALPHABET: string = '123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ'
const ALPHABET_MAP: Record<string, number> = {}
const BASE = 58
for (var i = 0; i < ALPHABET.length; i++) {
    ALPHABET_MAP[ALPHABET.charAt(i)] = i
}

function ToUTF8(str: string): Array<number> {
    var result = new Array()

    var k = 0
    for (var i = 0; i < str.length; i++) {
        var j = encodeURI(str[i])
        if (j.length == 1) {
            result[k++] = j.charCodeAt(0)
        } else {
            var bytes = j.split('%')
            for (var l = 1; l < bytes.length; l++) {
                result[k++] = parseInt('0x' + bytes[l])
            }
        }
    }
    return result
}

export const base58Encode = (str: string) => {
    const buffer = ToUTF8(str)
    if (buffer.length === 0) {
        return ''
    }
    let i,
        j,
        digits = [0]
    for (i = 0; i < buffer.length; i++) {
        for (j = 0; j < digits.length; j++) {
            digits[j] <<= 8
        }
        digits[0] += buffer[i]
        let carry = 0
        for (j = 0; j < digits.length; ++j) {
            digits[j] += carry
            carry = (digits[j] / BASE) | 0
            digits[j] %= BASE
        }
        while (carry) {
            digits.push(carry % BASE)
            carry = (carry / BASE) | 0
        }
    }
    for (i = 0; buffer[i] === 0 && i < buffer.length - 1; i++) {
        digits.push(0)
    }
    return digits
        .reverse()
        .map(function (digit) {
            return ALPHABET[digit]
        })
        .join('')
}

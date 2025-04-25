package cipher

import (
	"strings"
	"math/rand"
)

var letter2Index = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
	"i": 8,
	"j": 9,
	"k": 10,
	"l": 11,
	"m": 12,
	"n": 13,
	"o": 14,
	"p": 15,
	"q": 16,
	"r": 17,
	"s": 18,
	"t": 19,
	"u": 20,
	"v": 21,
	"w": 22,
	"x": 23,
	"y": 24,
	"z": 25,
}

var index2Letter = map[int]string{0: "a", 1: "b", 2: "c", 3: "d", 4: "e", 5: "f", 6: "g", 7: "h", 8: "i", 9: "j", 10: "k", 11: "l", 12: "m", 13: "n", 14: "o", 15: "p", 16: "q", 17: "r", 18: "s", 19: "t", 20: "u", 21: "v", 22: "w", 23: "x", 24: "y", 25: "z"}    


// Define the shift and vigenere types here.
type shift struct {
	shift int
}

type vigenere struct {
	key string
}

func NewCaesar() Cipher {
    return NewShift(3)
}

func NewShift(distance int) Cipher {     
	if distance == 0 {         
		return nil     
	} else if (-25 <= distance && distance <= -1) || (1 <= distance && distance <= 25) {
		return shift{shift: distance}     
	} else {         
		return nil     
	}     
}

func (c shift) Encode(input string) string {
    var result strings.Builder
    input = strings.ToLower(input)
    
    for _, r := range input {
        if r >= 'a' && r <= 'z' {
            shifted := (int(r-'a') + c.shift) % 26
            if shifted < 0 {
                shifted += 26
            }
            result.WriteByte(byte(shifted + 'a'))
        }
    }
    
    return result.String()
}

func (c shift) Decode(input string) string {
    return shift{shift: -c.shift}.Encode(input)
}

func NewVigenere(key string) Cipher {
    // If key is empty, generate a random key
    if key == "" {
        return nil
    }
    
	if strings.ToLower(key) != key {
		return nil
	} 
	
    key = strings.ToLower(key)
    
    // Check if key contains only lowercase letters
    allA := true
    for _, r := range key {
        if r < 'a' || r > 'z' {
            return nil // Invalid character
        }
        if r != 'a' {
            allA = false
        }
    }
    
    // Key cannot be all 'a's
    if allA {
        return nil
    }
    
    return vigenere{key: key}
}

// Helper function to generate a random vigenere cipher with at least 100 characters
func generateRandomVigenere() Cipher {
    const letters = "abcdefghijklmnopqrstuvwxyz"
    keyLength := 100
    
    key := make([]byte, keyLength)
    for i := 0; i < keyLength; i++ {
        // Generate a random index from 0-25
        randomIndex := rand.Intn(len(letters))
        key[i] = letters[randomIndex]
    }
    
    // Ensure at least one character is not 'a'
    key[0] = 'b'
    
    return vigenere{key: string(key)}
}

func (v vigenere) Encode(input string) string {
    var result strings.Builder
    input = strings.ToLower(input)
    keyIndex := 0
    
    for _, r := range input {
        if r >= 'a' && r <= 'z' {
            shift := int(v.key[keyIndex%len(v.key)] - 'a')
            shifted := (int(r-'a') + shift) % 26
            result.WriteByte(byte(shifted + 'a'))            
            keyIndex++
        }
    }
    
    return result.String()
}

func (v vigenere) Decode(input string) string {
    var result strings.Builder
    keyIndex := 0
    
    for _, r := range input {
        if r >= 'a' && r <= 'z' {
            shift := int(v.key[keyIndex%len(v.key)] - 'a')
            
            shifted := (int(r-'a') - shift) % 26
            if shifted < 0 {
                shifted += 26
            }
            result.WriteByte(byte(shifted + 'a'))
            
            keyIndex++
        }
    }
    
    return result.String()
}
package bencode

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
)

func Decode(reader *bufio.Reader) (interface{}, error) {
	ch, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	switch ch {

	// integer
	case 'i':
		var buffer []byte
		for {
			ch, err = reader.ReadByte()
			if err != nil {
				return nil, err
			}

			// if we stumble upon `e`, dict complete, and we return
			if ch == 'e' {
				value, err := strconv.ParseInt(string(buffer), 10, 64)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("invalid integer %s", string(buffer)))
				}
				return value, nil
			}
			buffer = append(buffer, ch)
		}

	// list
	case 'l':
		var listHolder []interface{}
		for {
			ch, err = reader.ReadByte()
			if err != nil {
				return nil, err
			}

			// if we stumble upon `e`, list complete, and we return
			if ch == 'e' {
				return listHolder, nil
			}

			// read the key
			reader.UnreadByte()
			data, err := Decode(reader)
			if err != nil {
				return nil, err
			}

			// put key and value in dictionary
			listHolder = append(listHolder, data)
		}

	// dictioanry
	case 'd':
		dictHolder := map[string]interface{}{}
		for {
			ch, err = reader.ReadByte()
			if err != nil {
				return nil, err
			}

			// if we stumble upon `e`, dict complete, and we return
			if ch == 'e' {
				return dictHolder, nil
			}

			// read the key
			reader.UnreadByte()
			data, err := Decode(reader)
			if err != nil {
				return nil, err
			}

			// key has to be a string, if not then error
			key, ok := data.(string)
			if !ok {
				return nil, errors.New("key of the dictionary is not string")
			}

			// read the value
			value, err := Decode(reader)
			if err != nil {
				return nil, err
			}

			// put key and value in dictionary
			dictHolder[key] = value
		}

	// string
	default:
		reader.UnreadByte()

		var lengthBuf []byte
		for {
			ch, err := reader.ReadByte()
			if err != nil {
				return nil, err
			}
			if ch == ':' {
				break
			}
			lengthBuf = append(lengthBuf, ch)
		}

		length, err := strconv.Atoi(string(lengthBuf))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("invaid integer %s", string(lengthBuf)))
		}

		var strBuf []byte
		for i := 0; i < length; i++ {
			ch, err := reader.ReadByte()
			if err != nil {
				return nil, err
			}
			strBuf = append(strBuf, ch)
		}

		return string(strBuf), nil
	}
}
